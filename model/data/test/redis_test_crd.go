package test

import (
	"log"
	"time"

	"github.com/layasugar/laya"

	"github.com/layasugar/laya-template/model/dao"
	"github.com/layasugar/laya-template/model/dao/rdb"
)

const redisTestPrefix = "test:prefix:"

// RedisTestCurd 测试curd
func RedisTestCurd(ctx *laya.Context) {
	redisKey := redisTestPrefix + "laya-template"
	data := rdb.User{
		ID:       1,
		Username: "laya",
		Nickname: "layasugar",
		Avatar:   "https://layasugar.cn",
		Mobile:   "12345678910",
		Status:   1,
	}

	err := dao.Rdb().Set(ctx, redisKey, data.String(), 100*time.Second).Err()
	if err != nil {
		panic(err)
	}
	log.Print("redis set success")

	res, err := dao.Rdb().Get(ctx, redisKey).Result()
	if err != nil {
		panic(err)
	}
	log.Printf("redis 获取结果: %v", res)

	err = dao.Rdb().Del(ctx, redisKey).Err()
	if err != nil {
		panic(err)
	}
	log.Print("redis del success")
}
