package middlewares

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/global"
	"github.com/layasugar/laya/genv"
	"github.com/layasugar/laya/glogs"
	"github.com/layasugar/laya/gutils"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"strings"
)

// SetHeader implements the gin.handlerFunc
func SetHeader(ctx *laya.WebContext) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	requestID := ctx.GetHeader(glogs.RequestIDName)
	if requestID == "" {
		ctx.Request.Header.Set(glogs.RequestIDName, uuid.NewV4().String())
	}
	ctx.Next()
}

func SetTrace(ctx *laya.WebContext) {
	if !gutils.InSliceString(ctx.Request.RequestURI, gutils.IgnoreRoutes) {
		span := glogs.StartSpanR(ctx.Request, ctx.Request.RequestURI)
		if span != nil {
			span.Tag(glogs.RequestIDName, ctx.GetHeader(glogs.RequestIDName))
			_ = glogs.Inject(context.WithValue(context.Background(), glogs.GetSpanContextKey(), span.Context()), ctx.Request)
			ctx.Next()
			span.Finish()
		}
	}
}

// LogInParams 记录框架出入参
func LogInParams(ctx *laya.WebContext) {
	if genv.ParamLog() {
		if !global.CheckNoLogParams(ctx.Request.RequestURI) {
			requestData, _ := ctx.GetRawData()
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestData))
			ct := ctx.GetHeader("Content-Type")
			sct := strings.Split(ct, ";")
			switch sct[0] {
			case "application/json":
				var in map[string]interface{}
				_ = json.NewDecoder(bytes.NewBuffer(requestData)).Decode(&in)
				inJson, _ := json.Marshal(&in)
				ctx.Infof("入参", string(inJson), glogs.String("header", ctx.Request.Header))
			case "application/x-www-form-urlencoded", "multipart/form-data":
				ctx.Infof("入参", string(requestData), glogs.String("header", ctx.Request.Header))
			default:
				ctx.Infof("入参", string(requestData), glogs.String("header", ctx.Request.Header))
			}
		}
	}

	ctx.Next()
}
