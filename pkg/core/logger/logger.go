// Package logger
// logger: this is extend package, use https://github.com/sirupsen/logrus
package logger

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/layasugar/laya/core/constants"
	"github.com/layasugar/laya/core/rotatelog"
	"github.com/sirupsen/logrus"
)

var sugar *logrus.Logger

var defaultConfig = &Config{
	AppName:       constants.DEFAULT_APPNAME,
	AppMode:       constants.DEFAULT_APPMODE,
	LogType:       constants.DEFAULT_LOGTYPE,
	LogPath:       constants.DEFAULT_LOGPATH,
	ChildPath:     constants.DEFAULT_LOGCHILDPATH,
	RotationSize:  constants.DEFAULT_LOGMAXSIZE,
	RotationCount: constants.DEFAULT_LOGMAXCOUNT,
	RotationTime:  constants.DEFAULT_LOGMAXTIME,
	MaxAge:        constants.DEFAULT_LOGMAXAGE,
}

type Config struct {
	AppName       string        // 应用名
	AppMode       string        // 应用环境
	ChildPath     string        // 日志子路径+文件名
	LogType       string        // 日志类型
	LogPath       string        // 日志主路径
	LogLevel      string        // 日志等级
	RotationSize  int64         // 单个文件大小
	RotationCount uint          // 可以保留的文件个数
	RotationTime  time.Duration // 日志分割的时间
	MaxAge        time.Duration // 日志最大保留的天数
}

func GetSugar() *logrus.Logger {
	if sugar == nil {
		sugar = InitSugar(defaultConfig)
	}
	return sugar
}

func InitSugar(lc *Config) *logrus.Logger {
	level, err := logrus.ParseLevel(lc.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
	logPath := fmt.Sprintf("%s/%s/%s", lc.LogPath, lc.AppName, lc.ChildPath)
	if lc.LogType == constants.DEFAULT_LOGTYPE {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(GetWriter(logPath, lc))
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	}
	log.Printf("[app] logger success")
	sugar = logrus.WithField("app_name", lc.AppName).WithField("app_mode", lc.AppMode).Logger
	return sugar
}

func Debug(logId, template string, args ...interface{}) {
	entry := GetSugar().WithField(constants.X_REQUESTID, logId)
	msg, entry := dealWithArgs(entry, template, args...)
	entry.Debug(msg)
}

func Info(logId, template string, args ...interface{}) {
	entry := GetSugar().WithField(constants.X_REQUESTID, logId)
	msg, entry := dealWithArgs(entry, template, args...)
	entry.Info(msg)
}

func Warn(logId, template string, args ...interface{}) {
	entry := GetSugar().WithField(constants.X_REQUESTID, logId)
	msg, entry := dealWithArgs(entry, template, args...)
	entry.Warn(msg)
}

func Error(logId, template string, args ...interface{}) {
	entry := GetSugar().WithField(constants.X_REQUESTID, logId)
	msg, entry := dealWithArgs(entry, template, args...)
	entry.Error(msg)
}

func dealWithArgs(entry *logrus.Entry, tmp string, args ...interface{}) (msg string, l *logrus.Entry) {
	l = entry
	var tmpArgs []interface{}
	if len(args) > 0 {
		for _, item := range args {
			if nil == item {
				continue
			}
			if fields, ok := item.(logrus.Fields); ok {
				l = l.WithFields(fields)
			} else {
				tmpArgs = append(tmpArgs, item)
			}
		}
	}
	return fmt.Sprintf(tmp, tmpArgs...), l
}

// GetWriter 按天切割按大小切割
// filename 文件名
// RotationSize 每个文件的大小
// MaxAge 文件最大保留天数
// RotationCount 最大保留文件个数
// RotationTime 设置文件分割时间
// RotationCount 设置保留的最大文件数量
func GetWriter(filename string, lc *Config) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 stream-2021-5-20.logger
	// demo.log是指向最新日志的连接
	// 保存7天内的日志，每1小时(整点)分割一第二天志
	var options []rotatelog.Option
	options = append(options,
		rotatelog.WithRotationSize(lc.RotationSize),
		rotatelog.WithRotationCount(lc.RotationCount),
		rotatelog.WithRotationTime(lc.RotationTime),
		rotatelog.WithMaxAge(lc.MaxAge))

	hook, err := rotatelog.New(
		filename,
		options...,
	)

	if err != nil {
		panic(err)
	}
	return hook
}
