package gcnf

import (
	"strings"
	"time"

	"github.com/layasugar/laya/core/constants"
)

// logParams
// 需要打印入参和出参的路由
// 需要打印入参和出参的前缀
// 需要打印入参和出参的后缀
type logParams struct {
	logParams       []string
	logParamsPrefix []string
	logParamsSuffix []string
}

var logParamsRules = &logParams{}

// LogLevel 返回日志等级
func LogLevel() string {
	if IsSet(constants.KEY_APPLOGGERLEVEL) {
		return GetString(constants.KEY_APPLOGGERLEVEL)
	}
	return constants.DEFAULT_LOGLEVEL
}

// LogPath 返回日志基本路径
func LogPath() string {
	if IsSet(constants.KEY_APPLOGGERPATH) {
		return GetString(constants.KEY_APPLOGGERPATH)
	}
	return constants.DEFAULT_LOGPATH
}

// LogChildPath 返回日志子路径
func LogChildPath() string {
	if IsSet(constants.KEY_APPLOGGERCHILDPATH) {
		return GetString(constants.KEY_APPLOGGERCHILDPATH)
	}
	return constants.DEFAULT_LOGCHILDPATH
}

// LogType 返回日志类型
func LogType() string {
	if IsSet(constants.KEY_APPLOGGERTYPE) {
		return GetString(constants.KEY_APPLOGGERTYPE)
	}
	return constants.DEFAULT_LOGTYPE
}

// LogMaxAge 返回日志默认保留7天
func LogMaxAge() time.Duration {
	if IsSet(constants.KEY_APPLOGGERMAXAGE) {
		return time.Duration(GetInt(constants.KEY_APPLOGGERMAXAGE)) * 24 * time.Hour
	}
	return constants.DEFAULT_LOGMAXAGE
}

// LogMaxTime 返回日志切割的时间
func LogMaxTime() time.Duration {
	if IsSet(constants.KEY_APPLOGGERMAXTIME) {
		return time.Duration(GetInt(constants.KEY_APPLOGGERMAXTIME)) * 24 * time.Hour
	}
	return constants.DEFAULT_LOGMAXTIME
}

// SdkLog 返回日志切割的时间
func SdkLog() bool {
	if IsSet(constants.KEY_APPLOGPARAMSSDK) {
		return GetBool(constants.KEY_APPLOGPARAMSSDK)
	}
	return constants.DEFAULT_BOOLTRUE
}

// LogMaxCount 返回日志默认限制为30个
func LogMaxCount() uint {
	if IsSet(constants.KEY_APPLOGGERMAXCOUNT) {
		return uint(GetInt(constants.KEY_APPLOGGERMAXCOUNT))
	}
	return constants.DEFAULT_LOGMAXCOUNT
}

// LogMaxSize 返回单个日志文件大小
func LogMaxSize() int64 {
	if IsSet(constants.KEY_APPLOGGERMAXSIZE) {
		return int64(GetInt(constants.KEY_APPLOGGERMAXSIZE))
	}
	return constants.DEFAULT_LOGMAXSIZE
}

func loadLogParams() *logParams {
	if IsSet(constants.KEY_APPLOGPARAMSLOGURI) {
		rules := GetStringSlice(constants.KEY_APPLOGPARAMSLOGURI)
		if len(rules) > 0 {
			logParamsRules.logParams = rules
		}
	}

	if IsSet(constants.KEY_APPLOGPARAMSLOGPREFIXURI) {
		rules := GetStringSlice(constants.KEY_APPLOGPARAMSLOGPREFIXURI)
		if len(rules) > 0 {
			logParamsRules.logParamsPrefix = rules
		}
	}

	if IsSet(constants.KEY_APPLOGPARAMSLOGSUFFIXURI) {
		rules := GetStringSlice(constants.KEY_APPLOGPARAMSLOGSUFFIXURI)
		if len(rules) > 0 {
			logParamsRules.logParamsSuffix = rules
		}
	}

	return logParamsRules
}

// CheckLogParams 判断是否需要打印入参出参日志, 需要打印返回true
func CheckLogParams(origin string) bool {
	if len(logParamsRules.logParams) > 0 {
		for _, v := range logParamsRules.logParams {
			if v == origin || v == constants.DEFAULT_ALLOWALLURI {
				return true
			}
		}
	}

	if len(logParamsRules.logParamsPrefix) > 0 {
		for _, v := range logParamsRules.logParamsPrefix {
			if strings.HasPrefix(origin, v) {
				return true
			}
		}
	}

	if len(logParamsRules.logParamsSuffix) > 0 {
		for _, v := range logParamsRules.logParamsSuffix {
			if strings.HasSuffix(origin, v) {
				return true
			}
		}
	}

	return false
}
