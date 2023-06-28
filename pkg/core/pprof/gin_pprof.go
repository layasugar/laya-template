package pprof

import (
	"net/http/pprof"
	"strings"

	"github.com/gin-gonic/gin"
)

// Wrap adds several routes from package `net/http/pprofx` to *gin.Engine object
func Wrap(router *gin.Engine) {
	WrapGroup(&router.RouterGroup)
}

// Wrapper make sure we are backward compatible
var Wrapper = Wrap

// WrapGroup adds several routes from package `net/http/pprofx` to *gin.RouterGroup object
func WrapGroup(router *gin.RouterGroup) {
	routers := []struct {
		Method  string
		Path    string
		Handler gin.HandlerFunc
	}{
		{"GET", "/debug/pprofx/", IndexHandler()},
		{"GET", "/debug/pprofx/heap", HeapHandler()},
		{"GET", "/debug/pprofx/goroutine", GoroutineHandler()},
		{"GET", "/debug/pprofx/allocs", AllocsHandler()},
		{"GET", "/debug/pprofx/block", BlockHandler()},
		{"GET", "/debug/pprofx/threadcreate", ThreadCreateHandler()},
		{"GET", "/debug/pprofx/cmdline", CmdlineHandler()},
		{"GET", "/debug/pprofx/profile", ProfileHandler()},
		{"GET", "/debug/pprofx/symbol", SymbolHandler()},
		{"POST", "/debug/pprofx/symbol", SymbolHandler()},
		{"GET", "/debug/pprofx/tracex", TraceHandler()},
		{"GET", "/debug/pprofx/mutex", MutexHandler()},
	}

	basePath := strings.TrimSuffix(router.BasePath(), "/")
	var prefix string

	switch {
	case basePath == "":
		prefix = ""
	case strings.HasSuffix(basePath, "/debug"):
		prefix = "/debug"
	case strings.HasSuffix(basePath, "/debug/pprofx"):
		prefix = "/debug/pprofx"
	}

	for _, r := range routers {
		router.Handle(r.Method, strings.TrimPrefix(r.Path, prefix), r.Handler)
	}
}

// IndexHandler will pass the call from /debug/pprofx to pprofx
func IndexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Index(ctx.Writer, ctx.Request)
	}
}

// HeapHandler will pass the call from /debug/pprofx/heap to pprofx
func HeapHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("heap").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// GoroutineHandler will pass the call from /debug/pprofx/goroutine to pprofx
func GoroutineHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("goroutine").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// AllocsHandler will pass the call from /debug/pprofx/allocs to pprofx
func AllocsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("allocs").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// BlockHandler will pass the call from /debug/pprofx/block to pprofx
func BlockHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("block").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// ThreadCreateHandler will pass the call from /debug/pprofx/threadcreate to pprofx
func ThreadCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("threadcreate").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// CmdlineHandler will pass the call from /debug/pprofx/cmdline to pprofx
func CmdlineHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Cmdline(ctx.Writer, ctx.Request)
	}
}

// ProfileHandler will pass the call from /debug/pprofx/profile to pprofx
func ProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Profile(ctx.Writer, ctx.Request)
	}
}

// SymbolHandler will pass the call from /debug/pprofx/symbol to pprofx
func SymbolHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Symbol(ctx.Writer, ctx.Request)
	}
}

// TraceHandler will pass the call from /debug/pprofx/tracex to pprofx
func TraceHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Trace(ctx.Writer, ctx.Request)
	}
}

// MutexHandler will pass the call from /debug/pprofx/mutex to pprofx
func MutexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("mutex").ServeHTTP(ctx.Writer, ctx.Request)
	}
}
