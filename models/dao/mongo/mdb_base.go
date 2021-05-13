// Package mongodb 封装 mongo 数据库调用的一些列方法，包括数据库
package mdb

//// Mdb Extends DBConnInfo
//type Mdb struct {
//	DBConnInfo
//	ctx context.RequestContext
//}
//
//func GetConn(cluster string, ctx ...context.RequestContext) *Mdb {
//	db := &Mdb{
//		GetConnInfo(cluster),
//		nil,
//	}
//	if len(ctx) == 1 {
//		db.ctx = ctx[0]
//	}
//	return db
//}
//
//func (db *Mdb) Aggregate(col string, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
//	ctx := db.initCalContext("Aggregate")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	cur, err := db.DB.Collection(col).Aggregate(ct.Background(), pipeline, opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return cur, err
//}
//
//func (db *Mdb) BulkWrite(col string, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
//	ctx := db.initCalContext("BulkWrite")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	bwres, err := db.DB.Collection(col).BulkWrite(ct.Background(), models, opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return bwres, err
//}
//
//func (db *Mdb) Clone(col string, opts ...*options.CollectionOptions) (*mongo.Collection, error) {
//	return db.DB.Collection(col).Clone(opts...)
//}
//
//func (db *Mdb) Count(col string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
//	ctx := db.initCalContext("Count")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	count, err := db.DB.Collection(col).CountDocuments(ct.Background(), filter, opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return count, err
//}
//
//func (db *Mdb) CountDocuments(col string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
//	ctx := db.initCalContext("CountDocuments")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	count, err := db.DB.Collection(col).CountDocuments(ct.Background(), filter, opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return count, err
//}
//
//func (db *Mdb) Database(col string) *mongo.Database { return db.DB.Collection(col).Database() }
//
//func (db *Mdb) DeleteMany(col string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
//	ctx := db.initCalContext("DeleteMany")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	dmres, err := db.DB.Collection(col).DeleteMany(ct.Background(), filter, opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return dmres, err
//}
//
//func (db *Mdb) DeleteOne(col string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
//	ctx := db.initCalContext("DeleteOne")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	dor, err := db.DB.Collection(col).DeleteOne(ct.Background(), filter, opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return dor, err
//}
//
//func (db *Mdb) Distinct(col string, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
//	ctx := db.initCalContext("DeleteOne")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	distinct, err := db.DB.Collection(col).Distinct(ct.Background(), fieldName, filter, opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return distinct, err
//}
//
//func (db *Mdb) Drop(col string) error {
//	ctx := db.initCalContext("Drop")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	err := db.DB.Collection(col).Drop(ct.Background())
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return err
//}
//
//func (db *Mdb) EstimatedDocumentCount(col string, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
//	ctx := db.initCalContext("EstimatedDocumentCount")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	count, err := db.DB.Collection(col).EstimatedDocumentCount(ct.Background(), opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return count, err
//}
//
//func (db *Mdb) Find(col string, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
//	ctx := db.initCalContext("Find")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	cur, err := db.DB.Collection(col).Find(ct.Background(), filter, opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return cur, err
//}
//
//func (db *Mdb) FindOne(col string, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
//	ctx := db.initCalContext("FindOne")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	r := db.DB.Collection(col).FindOne(ct.Background(), filter, opts...)
//	ctx.TimeStatisStop("talk")
//	return r
//}
//
//func (db *Mdb) FindOneAndDelete(col string, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
//	ctx := db.initCalContext("FindOneAndDelete")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	r:= db.DB.Collection(col).FindOneAndDelete(ct.Background(), filter, opts...)
//	ctx.TimeStatisStop("talk")
//	return r
//}
//
//func (db *Mdb) FindOneAndReplace(col string, filter, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult {
//	ctx := db.initCalContext("FindOneAndReplace")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	r:= db.DB.Collection(col).FindOneAndReplace(ct.Background(), filter, replacement, opts...)
//	ctx.TimeStatisStop("talk")
//	return  r
//}
//
//func (db *Mdb) FindOneAndUpdate(col string, filter, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
//	ctx := db.initCalContext("FindOneAndUpdate")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	r:= db.DB.Collection(col).FindOneAndUpdate(ct.Background(), filter, update, opts...)
//	ctx.TimeStatisStop("talk")
//	return  r
//}
//
//func (db *Mdb) Indexes(col string) mongo.IndexView { return db.DB.Collection(col).Indexes() }
//
//func (db *Mdb) InsertMany(col string, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
//	ctx := db.initCalContext("InsertMany")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	insmres, err := db.DB.Collection(col).InsertMany(ct.Background(), documents, opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return insmres, err
//}
//
//func (db *Mdb) InsertOne(col string, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
//	ctx := db.initCalContext("InsertOne")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	insores, err := db.DB.Collection(col).InsertOne(ct.Background(), document, opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return insores, err
//}
//
//func (db *Mdb) Name(col string) string { return db.DB.Collection(col).Name() }
//
//func (db *Mdb) ReplaceOne(col string, filter, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
//	ctx := db.initCalContext("ReplaceOne")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	repres, err := db.DB.Collection(col).ReplaceOne(ct.Background(), filter, replacement, opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return repres, err
//}
//
//func (db *Mdb) UpdateMany(col string, filter, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
//	ctx := db.initCalContext("UpdateMany")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	umres, err := db.DB.Collection(col).UpdateMany(ct.Background(), filter, replacement, opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return umres, err
//}
//
//func (db *Mdb) UpdateOne(col string, filter, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
//	ctx := db.initCalContext("UpdateOne")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	uores, err := db.DB.Collection(col).UpdateOne(ct.Background(), filter, replacement, opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return uores, err
//}
//
//func (db *Mdb) Watch(col string, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
//	ctx := db.initCalContext("Watch")
//	defer ctx.FlushLog()
//	ctx.TimeStatisStart("cost")
//	defer ctx.TimeStatisStop("cost")
//	ctx.TimeStatisStart("talk")
//	cs, err := db.DB.Collection(col).Watch(ct.Background(), pipeline, opts...)
//	ctx.TimeStatisStop("talk")
//	if err != nil {
//		ctx.CurRecord().Error = err
//	}
//	return cs, err
//}
//
//func (db *Mdb) Collection(col string) *mongo.Collection {
//	return db.DB.Collection(col)
//}
//
//
//// Checker ...
//func (db *Mdb) Checker(lastIDORAffected int64, err error) error {
//	if lastIDORAffected == 0 {
//		return errors.New("InsertLastID OR RowsAffected is 0")
//	}
//	return err
//}
//
//func initLogID(ctx *context.Context) string {
//	var logID string
//	if ctx.ReqContext != nil {
//		logID = ctx.ReqContext.GetLogID()
//	}
//	if logID == "" {
//		logID = produce.NewLogID()
//	}
//	return logID
//}
//
//func (db *Mdb) initCalContext(method string) *context.Context {
//	ctx := context.NewContext()
//	ctx.Caller = "Mongodb"
//	ctx.Method = strings.ToLower(method)
//	ctx.ReqContext = db.ctx
//	ctx.ServiceName = db.ClusterName
//	ctx.LogID = initLogID(ctx)
//	return ctx
//}
