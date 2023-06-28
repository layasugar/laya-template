package core

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/layasugar/laya/core/constants"
	"github.com/layasugar/laya/core/metautils"
	"github.com/layasugar/laya/core/pprof"
	"github.com/layasugar/laya/core/util"
	"github.com/layasugar/laya/gcnf"
)

// WebServer 基于http协议的服务
// 这里的实现是基于gin框架, 封装了gin的所有的方法
// gin 的核心是高效路由, 但是gin.Engine和gin.IRouter(s)的高耦合让我们无法复用, gin的作者认为它的路由就是引擎吧
type WebServer struct {
	// 重写所有的路由相关的方法
	*WebRoute
	// 继承引擎本身的其他方法
	*gin.Engine
}

// newWebServer 创建WebServer
func newWebServer(mode string) *WebServer {
	gin.SetMode(mode)

	server := &WebServer{
		Engine: gin.New(),
		WebRoute: &WebRoute{
			root: true,
		},
	}
	server.WebRoute.server = server
	server.WebRoute.RouterGroup = &server.Engine.RouterGroup

	if mode == "debug" {
		pprof.Wrap(server.Engine)
	}

	return server
}

// RouterRegister 路由注册者
type RouterRegister func(*WebServer)
type WebHandlerFunc func(ctx *Context)

// type RouterRegister func(WebRouter)

// Register 注册路由
func (webServer *WebServer) Register(rr RouterRegister) {
	rr(webServer)
}

const (
	defaultReadTimeout  = time.Second * 3
	defaultWriteTimeout = time.Second * 3
)

// RunGrace 实现Server接口
func (webServer *WebServer) RunGrace(addr string, timeouts ...time.Duration) error {
	readTimeout, writeTimeout := defaultReadTimeout, defaultWriteTimeout
	if len(timeouts) > 0 {
		readTimeout = timeouts[0]
		if len(timeouts) > 1 {
			writeTimeout = timeouts[1]
		}
	}

	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Handler:      webServer.Engine,
	}
	return server.ListenAndServe()
}

// Delims 设置模板的分解符
// 重写gin方法
func (webServer *WebServer) Delims(left, right string) *WebServer {
	webServer.Engine.Delims(left, right)
	return webServer
}

// HandleContext re-enter a contextx that has been rewritten.
// This can be done by setting c.Request.URL.Path to your new target.
// Disclaimer: You can loop yourself to death with this, use wisely.
func (webServer *WebServer) HandleContext(wc *Context) {
	webServer.Engine.HandleContext(wc.Gin())
}

// NoRoute adds handlers for NoRoute. It return a 404 code by default.
// 重写gin方法
func (webServer *WebServer) NoRoute(handlers ...WebHandlerFunc) {
	webServer.Engine.NoRoute(decorateWebHandlers(handlers)...)
}

// NoMethod sets the handlers called when...
// 重写gin方法
func (webServer *WebServer) NoMethod(handlers ...WebHandlerFunc) {
	webServer.Engine.NoMethod(decorateWebHandlers(handlers)...)
}

// Use adds middleware to the group, see example code in github.
func (webServer *WebServer) Use(middleware ...WebHandlerFunc) WebRouter {
	return webServer.WebRoute.Use(middleware...)
}

// toGinHandlerFunc 转换为Gin的HandlerFunc
func (hdlr WebHandlerFunc) toGinHandlerFunc() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		name := ginCtx.Request.RequestURI
		md := metautils.NiceMD(ginCtx.Request.Header)
		ctx := NewContext(constants.SERVERGIN, name, md, ginCtx)
		hdlr(ctx)
	}
}

func ginWebHandler(ginHandler gin.HandlerFunc) WebHandlerFunc {
	return func(ctx *Context) {
		ginHandler(ctx.Gin())
	}
}

func decorateWebHandlers(handlers []WebHandlerFunc) []gin.HandlerFunc {
	_handlers := []gin.HandlerFunc{}
	for _, hdlr := range handlers {
		_handlers = append(_handlers, hdlr.toGinHandlerFunc())
	}
	return _handlers
}

var _ WebRouter = &WebRoute{}

// WebRouter interface WebRequest Router
// 它合并了 gin.IRoute 和 gin.IRoutes
type WebRouter interface {
	// Group gin.IRoute.Group
	Group(string, ...WebHandlerFunc) WebRouter

	// Use gin.IRoutes.Use
	Use(...WebHandlerFunc) WebRouter

	Any(pattern string, handlers ...WebHandlerFunc) WebRouter
	GET(pattern string, handlers ...WebHandlerFunc) WebRouter
	POST(pattern string, handlers ...WebHandlerFunc) WebRouter
	DELETE(pattern string, handlers ...WebHandlerFunc) WebRouter
	PATCH(pattern string, handlers ...WebHandlerFunc) WebRouter
	PUT(pattern string, handlers ...WebHandlerFunc) WebRouter
	OPTIONS(pattern string, handlers ...WebHandlerFunc) WebRouter
	HEAD(pattern string, handlers ...WebHandlerFunc) WebRouter

	StaticFile(relativePath, filepath string) WebRouter
	Static(relativePath, root string) WebRouter
	StaticFS(relativePath string, fs http.FileSystem) WebRouter

	// TODO
	// AutoGET(...WebHandlerFunc) WebRouter
	// AutoPOST(...WebHandlerFunc) WebRouter
	// AutoDELETE(...WebHandlerFunc) WebRouter
	// AutoPATCH(...WebHandlerFunc) WebRouter
	// AutoPUT(...WebHandlerFunc) WebRouter
	// AutoOPTIONS(...WebHandlerFunc) WebRouter
	// AutoHEAD(...WebHandlerFunc) WebRouter
	// AutoHandle(string, ...WebHandlerFunc) WebRouter
}

// WebRouterGroup interface
// type WebRouterGroup interface {
// 	WebRouter
// 	Group(string, ...WebHandlerFunc) WebRouterGroup
// }

// WebRoute struct
// 它实现了gin.IRoutes, gin.IRoute
type WebRoute struct {
	RouterGroup *gin.RouterGroup
	server      *WebServer
	root        bool
}

// Group creates a new web router group
func (wrc *WebRoute) Group(pattern string, handlers ...WebHandlerFunc) WebRouter {
	return &WebRoute{
		RouterGroup: wrc.RouterGroup.Group(pattern, decorateWebHandlers(handlers)...),
		server:      wrc.server,
		root:        false,
	}
}

// Use attachs a global middleware to the router. ie. the middleware attached though Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
func (wrc *WebRoute) Use(middleware ...WebHandlerFunc) WebRouter {
	wrc.RouterGroup.Use(decorateWebHandlers(middleware)...)
	return wrc.returnObject()
}

// Any 注册所有的方法
func (wrc *WebRoute) Any(pattern string, handlers ...WebHandlerFunc) WebRouter {
	wrc.RouterGroup.Any(pattern, decorateWebHandlers(handlers)...)
	return wrc.returnObject()
}

// GET 注册GET方法
func (wrc *WebRoute) GET(pattern string, handlers ...WebHandlerFunc) WebRouter {
	wrc.RouterGroup.GET(pattern, decorateWebHandlers(handlers)...)
	return wrc.returnObject()
}

// POST 注册POST方法
func (wrc *WebRoute) POST(pattern string, handlers ...WebHandlerFunc) WebRouter {
	wrc.RouterGroup.POST(pattern, decorateWebHandlers(handlers)...)
	return wrc.returnObject()
}

// DELETE 注册DELETE方法
func (wrc *WebRoute) DELETE(pattern string, handlers ...WebHandlerFunc) WebRouter {
	wrc.RouterGroup.DELETE(pattern, decorateWebHandlers(handlers)...)
	return wrc.returnObject()
}

// PATCH 注册PATCH方法
func (wrc *WebRoute) PATCH(pattern string, handlers ...WebHandlerFunc) WebRouter {
	wrc.RouterGroup.PATCH(pattern, decorateWebHandlers(handlers)...)
	return wrc.returnObject()
}

// PUT 注册PUT方法
func (wrc *WebRoute) PUT(pattern string, handlers ...WebHandlerFunc) WebRouter {
	wrc.RouterGroup.PUT(pattern, decorateWebHandlers(handlers)...)
	return wrc.returnObject()
}

// OPTIONS 注册OPTIONS方法
func (wrc *WebRoute) OPTIONS(pattern string, handlers ...WebHandlerFunc) WebRouter {
	wrc.RouterGroup.OPTIONS(pattern, decorateWebHandlers(handlers)...)
	return wrc.returnObject()
}

// HEAD 注册HEAD方法
func (wrc *WebRoute) HEAD(pattern string, handlers ...WebHandlerFunc) WebRouter {
	wrc.RouterGroup.HEAD(pattern, decorateWebHandlers(handlers)...)
	return wrc.returnObject()
}

// StaticFile 静态文件
func (wrc *WebRoute) StaticFile(relativePath, filepath string) WebRouter {
	wrc.RouterGroup.StaticFile(relativePath, filepath)
	return wrc.returnObject()
}

// Static 静态文件
func (wrc *WebRoute) Static(relativePath, root string) WebRouter {
	wrc.RouterGroup.Static(relativePath, root)
	return wrc.returnObject()
}

// StaticFS 静态文件
func (wrc *WebRoute) StaticFS(relativePath string, fs http.FileSystem) WebRouter {
	wrc.RouterGroup.StaticFS(relativePath, fs)
	return wrc.returnObject()
}

func (wrc *WebRoute) returnObject() WebRouter {
	if wrc.root {
		return wrc.server
	}
	return wrc
}

// 打印出入参数
func webBoundLog(ctx *Context) {
	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: ctx.Gin().Writer}
	ctx.Gin().Writer = w
	requestData, _ := ctx.Gin().GetRawData()
	ctx.Gin().Request.Body = io.NopCloser(bytes.NewBuffer(requestData))
	ctx.Gin().Next()
	if gcnf.CheckLogParams(ctx.Gin().Request.RequestURI) {
		ctx.Info("params_log", ctx.Field("header", util.GetString(ctx.Gin().Request.Header)),
			ctx.Field("path", ctx.Gin().Request.RequestURI), ctx.Field("protocol", constants.PROTOCOLHTTP),
			ctx.Field("inbound", string(requestData)), ctx.Field("outbound", w.body.String()))
	}
	ctx.End(ctx.TopSpan())
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r responseBodyWriter) WriteString(s string) (n int, err error) {
	r.body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}

// defaultWebServerMiddlewares 默认的Http Server中间件
// 其实应该保证TowerLogware 不panic，但是无法保证，多一个recovery来保证业务日志崩溃后依旧有访问日志
var defaultWebServerMiddlewares = []WebHandlerFunc{
	webBoundLog,
	ginWebHandler(gin.Recovery()),
	recovery,
}

// 拦截到错误后处理span, 记录日志, 然后panic
func recovery(ctx *Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.End(ctx.TopSpan())
			ctx.Error("系统错误, err: %v", err)
			panic(err)
		}
	}()
	ctx.Gin().Next()
}
