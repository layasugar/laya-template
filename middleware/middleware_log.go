package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/layatips/laya/gconf"
	"github.com/layatips/laya/glogs"
	"io/ioutil"
	"strings"
)

func LogInParams(c *gin.Context) {
	if c.Request.RequestURI != "/ready" && c.Request.RequestURI != "/healthz" {
		bc := gconf.GetBaseConf()
		if bc.ParamsLog {
			requestData, _ := c.GetRawData()
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestData))
			ct := c.GetHeader("Content-Type")
			sct := strings.Split(ct, ";")
			switch sct[0] {
			case "application/json":
				var in map[string]interface{}
				_ = json.NewDecoder(bytes.NewBuffer(requestData)).Decode(&in)
				inJson, _ := json.Marshal(&in)
				glogs.InfoFR(c, "title=入参打印,path=%s,content=%s", c.Request.RequestURI, inJson)
			case "application/x-www-form-urlencoded", "multipart/form-data":
				glogs.InfoFR(c, "title=入参打印,path=%s,content=%s", c.Request.RequestURI, string(requestData))
			default:
				glogs.InfoFR(c, "title=入参打印,path=%s,content=%s", c.Request.RequestURI, string(requestData))
			}
		}
	}
	c.Next()
}
