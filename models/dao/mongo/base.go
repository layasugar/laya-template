package dbdao

import (
	"context"
	"fmt"
	"hash/crc32"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBDao struct {
	ClusterName string
	DB          *mongodb.DB
}

//Dao Layer GetConn
func GetConn(cluster string) *mongodb.DB {
	return mongodb.GetConn(cluster)
}

func (d *DBDao) Aggregate(col string, pipeline interface{}, results interface{}, opts ...*options.AggregateOptions) error {
	cur, err := GetConn(d.ClusterName).Aggregate(col, pipeline)
	if err == nil {
		err = cur.All(context.Background(), results)
	}
	if err != nil {
		return err
	}
	return err
}

func (d *DBDao) BulkWrite(col string, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	bwres, err := GetConn(d.ClusterName).BulkWrite(col, models, opts...)
	if err != nil {
		return nil, err
	}
	return bwres, err
}

func (d *DBDao) Clone(col string, opts ...*options.CollectionOptions) (*mongo.Collection, error) {
	return GetConn(d.ClusterName).Clone(col, opts...)
}

func (d *DBDao) Count(col string, filter interface{}, opts ...*options.CountOptions) (int64, error) {

	count, err := GetConn(d.ClusterName).CountDocuments(col, filter, opts...)
	if err != nil {
		return 0, err
	}
	return count, err
}

func (d *DBDao) CountDocuments(col string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	count, err := GetConn(d.ClusterName).CountDocuments(col, filter, opts...)
	if err != nil {
		return 0, err
	}
	return count, err
}

func (d *DBDao) Database(col string) *mongo.Database { return GetConn(d.ClusterName).Database(col) }

func (d *DBDao) DeleteMany(col string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	dmres, err := GetConn(d.ClusterName).DeleteMany(col, filter, opts...)
	if err != nil {
		return nil, err
	}
	return dmres, err
}

func (d *DBDao) DeleteOne(col string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	dor, err := GetConn(d.ClusterName).DeleteOne(col, filter, opts...)
	if err != nil {
		return nil, err
	}
	return dor, err
}

func (d *DBDao) Distinct(col string, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	distinct, err := GetConn(d.ClusterName).Distinct(col, fieldName, filter, opts...)
	if err != nil {
		return nil, err
	}
	return distinct, err
}

func (d *DBDao) Drop(col string) error {
	err := GetConn(d.ClusterName).Drop(col)
	if err != nil {
		return err
	}
	return err
}

func (d *DBDao) EstimatedDocumentCount(col string, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	count, err := GetConn(d.ClusterName).EstimatedDocumentCount(col, opts...)
	if err != nil {
		return 0, err
	}
	return count, err
}

func (d *DBDao) Find(col string, filter interface{}, results interface{}, opts ...*options.FindOptions) error {
	cur, err := GetConn(d.ClusterName).Find(col, filter, opts...)
	if err == nil {
		err = cur.All(context.Background(), results)
	}
	if err != nil {
		return err
	}
	return err
}

func (d *DBDao) FindOne(col string, filter interface{}, result interface{}, opts ...*options.FindOneOptions) error {
	return GetConn(d.ClusterName).FindOne(col, filter, opts...).Decode(result)
}

func (d *DBDao) FindOneAndDelete(col string, filter interface{}, result interface{}, opts ...*options.FindOneAndDeleteOptions) error {
	return GetConn(d.ClusterName).FindOneAndDelete(col, filter, opts...).Decode(result)
}

func (d *DBDao) FindOneAndReplace(col string, filter, replacement interface{}, result interface{}, opts ...*options.FindOneAndReplaceOptions) error {
	return GetConn(d.ClusterName).FindOneAndReplace(col, filter, replacement, opts...).Decode(result)
}

func (d *DBDao) FindOneAndUpdate(col string, filter, update interface{}, result interface{}, opts ...*options.FindOneAndUpdateOptions) error {
	return GetConn(d.ClusterName).FindOneAndUpdate(col, filter, update, opts...).Decode(result)
}

func (d *DBDao) FindOneAndUpsert(col string, filter, update interface{}, result interface{}, opts ...*options.FindOneAndUpdateOptions) error {
	rd := options.After
	optUpsert := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(rd)
	opts = append(opts, optUpsert)

	return GetConn(d.ClusterName).FindOneAndUpdate(col, filter, update, opts...).Decode(result)
}

func (d *DBDao) Indexes(col string) mongo.IndexView { return GetConn(d.ClusterName).Indexes(col) }

func (d *DBDao) InsertMany(col string, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	insmres, err := GetConn(d.ClusterName).InsertMany(col, documents, opts...)
	if err != nil {
		return nil, err
	}
	return insmres, err
}

func (d *DBDao) InsertOne(col string, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	insores, err := GetConn(d.ClusterName).InsertOne(col, document, opts...)
	if err != nil {
		return nil, err
	}
	return insores, err
}

func (d *DBDao) Name(col string) string { return GetConn(d.ClusterName).Name(col) }

func (d *DBDao) ReplaceOne(col string, filter, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	repres, err := GetConn(d.ClusterName).ReplaceOne(col, filter, replacement, opts...)
	if err != nil {
		return nil, err
	}
	return repres, err
}

func (d *DBDao) UpdateMany(col string, filter, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	umres, err := GetConn(d.ClusterName).UpdateMany(col, filter, replacement, opts...)
	if err != nil {
		return nil, err
	}
	return umres, err
}

func (d *DBDao) UpdateOne(col string, filter, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	uores, err := GetConn(d.ClusterName).UpdateOne(col, filter, replacement, opts...)
	if err != nil {
		return nil, err
	}
	return uores, err
}

func (d *DBDao) Upsert(col string, filter, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	optUpsert := options.Update().SetUpsert(true)
	opts = append(opts, optUpsert)
	uores, err := GetConn(d.ClusterName).UpdateOne(col, filter, replacement, opts...)
	if err != nil {
		return nil, err
	}
	return uores, err
}

func (d *DBDao) Watch(col string, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	cs, err := GetConn(d.ClusterName).Watch(col, pipeline, opts...)
	if err != nil {
		return nil, err
	}
	return cs, err
}

// GetTableName 得到分表后的 表名
// baseName 元素表名
// tableCount 表的总个数
// divisionFlag 分表依据
func GetCollectionName(baseName string, tableCount uint32, divisionFlag int64) string {
	crc32 := crc32.ChecksumIEEE([]byte(strconv.FormatInt(divisionFlag, 10)))
	return fmt.Sprintf("%s_%d", baseName, crc32%tableCount)
}
