package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/layatips/laya-go/models/data"
	"github.com/layatips/laya/glogs"
	"time"
)

type UserParam struct {
	Id int64 `json:"id" binding:"required"`
}

func GetUserList(c *gin.Context, param *UserParam) (interface{}, error) {
	span1 := glogs.StartSpan("GetUserList")
	time.Sleep(time.Second)
	glogs.StopSpan(span1)

	span2 := glogs.StartSpanP(span1.Context(), "haha")
	time.Sleep(100 * time.Microsecond)
	glogs.StopSpan(span2)

	span3 := glogs.StartSpanP(span1.Context(), "heihei")
	time.Sleep(200 * time.Microsecond)
	glogs.StopSpan(span3)

	users, err := data.GetUserListByZone(c, param.Id)
	return users, err
}
