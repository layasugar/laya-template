package admin

import (
	"github.com/layasugar/laya/glogs"
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/data"
	"time"
)

type UserParam struct {
	Id int64 `json:"id" binding:"required"`
}

func GetUserList(ctx *laya.WebContext, param *UserParam) (interface{}, error) {
	span1 := glogs.StartSpan("GetUserList")
	time.Sleep(time.Second)
	glogs.StopSpan(span1)

	span2 := glogs.StartSpanP(span1.Context(), "haha")
	time.Sleep(100 * time.Microsecond)
	glogs.StopSpan(span2)

	span3 := glogs.StartSpanP(span1.Context(), "heihei")
	time.Sleep(200 * time.Microsecond)
	glogs.StopSpan(span3)

	users, err := data.GetUserListByZone(ctx, param.Id)
	return users, err
}
