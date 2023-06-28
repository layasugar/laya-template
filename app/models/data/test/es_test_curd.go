package test

import (
	"encoding/json"
	"github.com/layasugar/laya-template/handle/model/dao"
	"github.com/layasugar/laya-template/handle/model/dao/es"
	"time"
)

func EsUserCreate(ctx *core.Context) (string, error) {
	data := es.User{
		ID:        1,
		Username:  "laya",
		Nickname:  "layasugar",
		Avatar:    "https://layasugar.cn",
		Mobile:    "12345678910",
		Status:    1,
		CreatedAt: time.Now(),
	}

	res, err := dao.Edb().Index().Index("user").BodyJson(data).Do(ctx)
	if err != nil {
		return "", err
	}

	return res.Id, nil
}

func EsUserUpdate(ctx *core.Context, eid string) error {
	var data = map[string]interface{}{
		"username": "layasugar",
	}

	_, err := dao.Edb().Update().Index("user").Id(eid).Upsert(data).Doc(data).Do(ctx)
	return err
}

func EsUserSelect(ctx *core.Context, eid string) (*es.User, error) {
	var u es.User
	res, err := dao.Edb().Get().Index("user").Id(eid).Do(ctx)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res.Source, &u)
	return &u, err
}

func EsUserDel(ctx *core.Context, eid string) error {
	_, err := dao.Edb().Delete().Index("user").Id(eid).Do(ctx)
	return err
}
