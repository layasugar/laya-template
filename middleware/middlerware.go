package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/layasugar/laya/genv"
	"github.com/layasugar/laya/glogs"
	"github.com/layasugar/laya/gutils"
	"io/ioutil"
	"strings"
)

// implements the gin.handlerFunc
func SetHeader(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	requestID := c.GetHeader(glogs.RequestIDName)
	if requestID == "" {
		c.Request.Header.Set(glogs.RequestIDName, uuid.NewV4().String())
	}
	c.Next()
}

func SetTrace(c *gin.Context) {
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

func LogInParams(c *gin.Context) {
	if !gutils.InSliceString(c.Request.RequestURI, gutils.IgnoreRoutes) {
		if genv.ParamLog() {
			requestData, _ := c.GetRawData()
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestData))
			ct := c.GetHeader("Content-Type")
			sct := strings.Split(ct, ";")
			switch sct[0] {
			case "application/json":
				var in map[string]interface{}
				_ = json.NewDecoder(bytes.NewBuffer(requestData)).Decode(&in)
				inJson, _ := json.Marshal(&in)
				glogs.InfoF(c, "入参打印", "path=%s,content=%s", c.Request.RequestURI, inJson)
			case "application/x-www-form-urlencoded", "multipart/form-data":
				glogs.InfoF(c, "入参打印", "title=入参打印,path=%s,content=%s", c.Request.RequestURI, string(requestData))
			default:
				glogs.InfoF(c, "入参打印", "title=入参打印,path=%s,content=%s", c.Request.RequestURI, string(requestData))
			}
		}
	}
	c.Next()
}
