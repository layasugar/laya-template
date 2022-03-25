package test

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/dao"
	"github.com/layasugar/laya-template/models/dao/mdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func MongoUserCreate(ctx *laya.WebContext) (string, error) {
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

func MongoUserUpdate(ctx *laya.WebContext, mid string) error {
	var data mdb.User
	id, _ := primitive.ObjectIDFromHex(mid)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"username", "layasugar"}}}}
	_, err := dao.Mdb().Database(data.Database()).Collection(data.Collection()).UpdateMany(ctx, filter, update)
	return err
}

func MongoUserSelect(ctx *laya.WebContext, mid string) (*mdb.User, error) {
	var data mdb.User
	id, _ := primitive.ObjectIDFromHex(mid)
	filter := bson.D{{"_id", id}}
	res := dao.Mdb().Database(data.Database()).Collection(data.Collection()).FindOne(ctx, filter)
	err := res.Decode(&data)
	return &data, err
}

func MongoUserDel(ctx *laya.WebContext, mid string) error {
	var data mdb.User
	id, _ := primitive.ObjectIDFromHex(mid)
	filter := bson.D{{"_id", id}}
	_, err := dao.Mdb().Database(data.Database()).Collection(data.Collection()).DeleteOne(ctx, filter)
	return err
}
