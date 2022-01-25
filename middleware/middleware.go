package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/layasugar/glogs"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/global"
	"github.com/layasugar/laya/genv"
	"github.com/layasugar/laya/gutils"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"strings"
)

// SetHeader implements the gin.handlerFunc
func SetHeader(c *laya.WebContext) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	requestID := c.GetHeader(glogs.RequestIDName)
	if requestID == "" {
		c.Request.Header.Set(glogs.RequestIDName, uuid.NewV4().String())
	}
	c.Next()
}

func SetTrace(c *laya.WebContext) {
	if !gutils.InSliceString(c.Request.RequestURI, gutils.IgnoreRoutes) {
		span := glogs.StartSpanR(c.Request, c.Request.RequestURI)
		if span != nil {
			span.Tag(glogs.RequestIDName, c.GetHeader(glogs.RequestIDName))
			_ = glogs.Inject(context.WithValue(context.Background(), glogs.GetSpanContextKey(), span.Context()), c.Request)
			c.Next()
			span.Finish()
		}
	}
}

// LogInParams 记录框架出入参
func LogInParams(c *laya.WebContext) {
	if genv.ParamLog() {
		if !global.CheckNoLogParams(c.Request.RequestURI) {
			requestData, _ := c.GetRawData()
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestData))
			ct := c.GetHeader("Content-Type")
			sct := strings.Split(ct, ";")
			switch sct[0] {
			case "application/json":
				var in map[string]interface{}
				_ = json.NewDecoder(bytes.NewBuffer(requestData)).Decode(&in)
				inJson, _ := json.Marshal(&in)
				glogs.InfoF(c.Request, "入参", string(inJson), glogs.String("header", c.Request.Header))
			case "application/x-www-form-urlencoded", "multipart/form-data":
				glogs.InfoF(c.Request, "入参", string(requestData), glogs.String("header", c.Request.Header))
			default:
				glogs.InfoF(c.Request, "入参", string(requestData), glogs.String("header", c.Request.Header))
			}
		}
	}

	c.Next()
}
