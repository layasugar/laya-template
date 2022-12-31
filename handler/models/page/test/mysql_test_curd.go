package test

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/handle/model/data/test"
)

func mysqlTestCurd(ctx *laya.Context) {
	if err := test.MysqlUserCreate(ctx, test.User{
		Username: "layasugar",
		Nickname: "laya",
		Avatar:   "https://layasugar.cn",
		Password: "123456",
		Salt:     "aaa",
		Mobile:   "12345678910",
		Status:   1,
	}); err != nil {
		panic(err)
	}
	ctx.Info("用户创建成功")

	if err := test.MysqlUserUpdate(ctx, "12345678910", "layasugar"); err != nil {
		panic(err)
	}
	ctx.Info("用户更新成功")

	userInfo, err := test.MysqlUserSelect(ctx, "12345678910")
	if err != nil {
		panic(err)
	}
	ctx.Info("%v", userInfo)

	if err = test.MysqlUserDel(ctx, "12345678910"); err != nil {
		panic(err)
	}
	ctx.Info("用户删除成功")
}
