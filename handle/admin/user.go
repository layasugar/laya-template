package admin

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/global/pagination"
	"github.com/layasugar/laya-template/model/page/admin"
)

func (ctrl *controller) GetUserList(ctx *laya.Context) {
	var param admin.GetUserListReq
	err := ctx.Gin().ShouldBindJSON(&param)
	if err != nil {
		ctrl.Fail(ctx, err)
		return
	}

	if param.Page == 0 {
		param.Page = pagination.DefaultPage
	}
	if param.PageSize == 0 {
		param.PageSize = pagination.DefaultPageSize
	}

	resp, err := admin.GetUserList(ctx, &param)
	if err != nil {
		ctx.Info("获取用户列表, err: %s", err.Error())
		ctrl.Fail(ctx, err)
		return
	}
	ctrl.Suc(ctx, resp)
}
