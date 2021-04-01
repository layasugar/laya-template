package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/layatips/laya-go/models/dao"
	"github.com/layatips/laya/gconf"
	"github.com/layatips/laya/glogs"
	"net/http"
)

// test
func (ctrl *BaseCtrl) Test(c *gin.Context) {
	var body map[string]interface{}
	_ = c.ShouldBindJSON(&body)
	fmt.Println(body)
	ctrl.Suc(c, body, "success")
}

type Memories struct {
	M Condition `json:"M"`
}

type Condition struct {
	Count int                      `json:"count"`
	Item  []map[string]interface{} `json:"item"`
}

// memory status
func (ctrl *BaseCtrl) MemoryStatus(c *gin.Context) {
	var resp Memories
	resp.M.Count = dao.Mem.ItemCount()
	dmi := dao.Mem.Items()
	counter := 0
	for k, v := range dmi {
		var temp = map[string]interface{}{}
		temp[k] = v.Object
		resp.M.Item = append(resp.M.Item, temp)
		counter++
		if counter >= 10 {
			break
		}
	}
	ctrl.Suc(c, resp)
}

// version
func (ctrl *BaseCtrl) Version(c *gin.Context) {
	bca := gconf.GetBaseConf()
	res := bca.AppName + " api version: 1.0.0\r\n" + "app_url: " + bca.AppUrl
	_, _ = c.Writer.Write([]byte(res))
	return
}

func (ctrl *BaseCtrl) HealthCheck(c *gin.Context) {
	ctrl.Suc(c, "", "success")
}

func (ctrl *BaseCtrl) ReadyCheck(c *gin.Context) {
	// mysql检测
	mc := gconf.GetDBConf()
	if mc.Open {
		sqlDB, err := dao.DB.DB()
		if err != nil {
			glogs.ErrorF("探针存活检测失败,mysql凉凉")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		err = sqlDB.Ping()
		if err != nil {
			glogs.ErrorF("探针存活检测失败,mysql凉凉")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	// redis检测
	rc := gconf.GetRdbConf()
	if rc.Open {
		pong, err := dao.Rdb.Ping(c).Result()
		if err != nil && err != redis.Nil {
			glogs.ErrorF("探针存活检测失败,redis凉凉")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctrl.Suc(c, pong, "success")
	}
}
