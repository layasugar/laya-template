package db

import (
	"context"

	"github.com/layasugar/laya"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

const (
	contextKey = "gorm:context"
	spanKey    = "span:key"
	tSpanName  = "mysql:"
)

func Wrap(ctx *laya.Context, dbName ...string) *gorm.DB {
	var db *gorm.DB
	if len(dbName) > 0 {
		db = getGormDB(dbName[0])
	} else {
		db = getGormDB(defaultDbName)
	}

	return db.Set(contextKey, ctx)
}

func registerCallbacks(db *gorm.DB) {
	prefix := db.Dialector.Name() + ":"

	_ = db.Callback().Create().Before("gorm:begin_transaction").Register("aotgorm_before_create", newBefore(prefix+"create"))
	_ = db.Callback().Create().After("gorm:commit_or_rollback_transaction").Register("otgorm_after_create", newAfter())

	_ = db.Callback().Update().Before("gorm:begin_transaction").Register("otgorm_before_update", newBefore(prefix+"update"))
	_ = db.Callback().Update().After("gorm:commit_or_rollback_transaction").Register("otgorm_after_update", newAfter())

	_ = db.Callback().Query().Before("gorm:query").Register("otgorm_before_query", newBefore(prefix+"query"))
	_ = db.Callback().Query().After("gorm:after_query").Register("otgorm_after_query", newAfter())

	_ = db.Callback().Delete().Before("gorm:begin_transaction").Register("otgorm_before_delete", newBefore(prefix+"delete"))
	_ = db.Callback().Delete().After("gorm:commit_or_rollback_transaction").Register("otgorm_after_delete", newAfter())

	_ = db.Callback().Row().Before("gorm:row").Register("otgorm_before_row", newBefore(prefix+"row"))
	_ = db.Callback().Row().After("gorm:row").Register("otgorm_after_row", newAfter())

	_ = db.Callback().Raw().Before("gorm:raw").Register("otgorm_before_raw", newBefore(prefix+"raw"))
	_ = db.Callback().Raw().After("gorm:raw").Register("otgorm_after_raw", newAfter())
}

func newBefore(name string) func(*gorm.DB) {
	return func(db *gorm.DB) {
		if v, ok := db.Get(contextKey); ok {
			switch ctx := v.(type) {
			case *laya.Context:
				_, span := ctx.Start(context.TODO(), tSpanName+name)
				if nil != span {
					db.Set(spanKey, span)
				}
			}
		}
	}
}

func newAfter() func(*gorm.DB) {
	return func(db *gorm.DB) {
		if spanx, ok := db.Get(spanKey); ok {
			switch span := spanx.(type) {
			case trace.Span:
				span.End()
			}
		}
	}
}
