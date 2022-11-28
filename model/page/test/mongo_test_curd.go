package test

import (
	"log"

	"github.com/layasugar/laya"

	"github.com/layasugar/laya-template/model/data/test"
)

func mongoTestCurd(ctx *laya.Context) {
	mid, err := test.MongoUserCreate(ctx)
	if err != nil {
		panic(err)
	}
	log.Print("mongo 数据增加成功, _id=" + mid)

	err = test.MongoUserUpdate(ctx, mid)
	if err != nil {
		panic(err)
	}
	log.Print("mongo 数据修改成功")

	res, err := test.MongoUserSelect(ctx, mid)
	if err != nil {
		panic(err)
	}
	log.Printf("mongo 数据查询成功, %v", res)

	err = test.MongoUserDel(ctx, mid)
	if err != nil {
		panic(err)
	}
	log.Print("mongo 数据删除成功")
}
