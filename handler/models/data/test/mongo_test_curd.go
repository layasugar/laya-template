package test

import (
	"github.com/layasugar/laya-template/handle/model/dao"
	"github.com/layasugar/laya-template/handle/model/dao/mdb"
	"time"

	"github.com/layasugar/laya"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MongoUserCreate(ctx *laya.Context) (string, error) {
	data := mdb.User{
		ID:        1,
		Username:  "laya",
		Nickname:  "layasugar",
		Avatar:    "https://layasugar.cn",
		Mobile:    "12345678910",
		Status:    1,
		CreatedAt: time.Now(),
	}
	mid, err := dao.Mdb().Database(data.Database()).Collection(data.Collection()).InsertOne(ctx, data)
	if err != nil {
		return "", err
	}
	if a, ok := mid.InsertedID.(primitive.ObjectID); ok {
		return a.Hex(), nil
	}
	return "", nil
}

func MongoUserUpdate(ctx *laya.Context, mid string) error {
	var data mdb.User
	id, _ := primitive.ObjectIDFromHex(mid)
	filter := bson.E{Key: "_id", Value: id}
	update := bson.E{Key: "$set", Value: bson.E{Key: "username", Value: "layasugar"}}
	_, err := dao.Mdb().Database(data.Database()).Collection(data.Collection()).UpdateMany(ctx, filter, update)
	return err
}

func MongoUserSelect(ctx *laya.Context, mid string) (*mdb.User, error) {
	var data mdb.User
	id, _ := primitive.ObjectIDFromHex(mid)
	filter := bson.E{Key: "_id", Value: id}
	res := dao.Mdb().Database(data.Database()).Collection(data.Collection()).FindOne(ctx, filter)
	err := res.Decode(&data)
	return &data, err
}

func MongoUserDel(ctx *laya.Context, mid string) error {
	var data mdb.User
	id, _ := primitive.ObjectIDFromHex(mid)
	filter := bson.E{Key: "_id", Value: id}
	_, err := dao.Mdb().Database(data.Database()).Collection(data.Collection()).DeleteOne(ctx, filter)
	return err
}
