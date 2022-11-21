package admin

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/global/page"
	"github.com/layasugar/laya-template/models/page/admin"
)

func (ctrl *controller) GetUserList(ctx *laya.WebContext) {
	var param admin.GetUserListReq
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctrl.Fail(ctx, err)
		return
	}

	if param.Page == 0 {
		param.Page = page.DefaultPage
	}
	if param.PageSize == 0 {
		param.PageSize = page.DefaultPageSize
	}

	resp, err := admin.GetUserList(ctx, &param)
	if err != nil {
		ctx.InfoF("获取用户列表, err: %s", err.Error())
		ctrl.Fail(ctx, err)
		return
	}
	ctrl.Suc(ctx, resp)
}
