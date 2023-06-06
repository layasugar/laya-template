package db

import (
	"context"
	"fmt"
	"time"

	"github.com/layasugar/laya"
	"github.com/layasugar/laya/core/constants"
	l "github.com/layasugar/laya/core/logger"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

const (
	sqlTitle = "mysql"
)

func Default(level logger.LogLevel) logger.Interface {
	var config = logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      level,
		Colorful:      true,
	}
	var (
		infoStr      = "{\"line\": \"%s\", \"level\": \"[info]\", \"msg\": \"%s\"}"
		warnStr      = "{\"line\": \"%s\", \"level\": \"[warn]\", \"msg\": \"%s\"}"
		errStr       = "{\"line\": \"%s\", \"level\": \"[error]\", \"msg\": \"%s\"}"
		traceStr     = "{\"line\": \"%s\", \"耗时\": \"%.3fms\", \"rows\": \"%v\" ,\"sql\": \"%s\"}"
		traceWarnStr = "{\"line\": \"%s\", \"错误\": \"%s\", \"耗时\": \"%.3fms\", \"rows\": \"%v\", \"sql\": \"%s\"}"
		traceErrStr  = "{\"line\": \"%s\", \"slow\": \"%s\", \"耗时\": \"%.3fms\", \"rows\": \"%v\", \"sql\": \"%s\"}"
	)

	return &gormLogger{
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

type gormLogger struct {
	logger.Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode logger mode
func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		errInfo := fmt.Sprintf(msg, data...)
		gormWriter(ctx, levelInfo, fmt.Sprintf(l.infoStr, utils.FileWithLineNum(), errInfo))
	}
}

// Warn print warn messages
func (l *gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		errInfo := fmt.Sprintf(msg, data...)
		gormWriter(ctx, levelWarn, fmt.Sprintf(l.infoStr, utils.FileWithLineNum(), errInfo))
	}
}

// Error print error messages
func (l *gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		errInfo := fmt.Sprintf(msg, data...)
		gormWriter(ctx, levelError, fmt.Sprintf(l.infoStr, utils.FileWithLineNum(), errInfo))
	}
}

// Trace print sql message
func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= logger.Error:
			sql, rows := fc()
			if rows == -1 {
				msg := fmt.Sprintf(l.traceErrStr, utils.FileWithLineNum(), err.Error(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
				gormWriter(ctx, levelError, msg)
			} else {
				msg := fmt.Sprintf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
				gormWriter(ctx, levelError, msg)
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			if rows == -1 {
				msg := fmt.Sprintf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
				gormWriter(ctx, levelWarn, msg)
			} else {
				msg := fmt.Sprintf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
				gormWriter(ctx, levelWarn, msg)
			}
		case l.LogLevel >= logger.Info:
			sql, rows := fc()
			if rows == -1 {
				msg := fmt.Sprintf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
				gormWriter(ctx, levelInfo, msg)
			} else {
				msg := fmt.Sprintf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
				gormWriter(ctx, levelInfo, msg)
			}
		}
	}
}

func gormWriter(ctx context.Context, level, msg string) {
	var logId string
	if webCtx, okInterface := ctx.(*laya.Context); okInterface {
		logId = webCtx.LogId()
	}

	var fields = map[string]interface{}{
		constants.X_REQUESTID: logId,
		constants.LOGGERTITLE: sqlTitle,
	}

	switch level {
	case levelInfo:
		l.GetSugar().WithFields(fields).Info(msg)
	case levelWarn:
		l.GetSugar().WithFields(fields).Warn(msg)
	case levelError:
		l.GetSugar().WithFields(fields).Error(msg)
	}
}
