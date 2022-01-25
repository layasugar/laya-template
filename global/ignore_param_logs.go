package global

import "strings"

// 不需要打印入参和出参的路由
// 不需要打印入参和出参的前缀
// 不需要打印入参和出参的后缀
type logParams struct {
	NoLogParams       map[string]string
	NoLogParamsPrefix []string
	NoLogParamsSuffix []string
}

// 不想打印的路由分组
var noLogParamsRules logParams

// SetNoLogParams 设置不需要打印的路由
func SetNoLogParams(path ...string) {
	for _, v := range path {
		noLogParamsRules.NoLogParams[v] = v
	}
}

// SetNoLogParamsPrefix 设置不需要打印入参和出参的路由前缀
func SetNoLogParamsPrefix(path ...string) {
	for _, v := range path {
		noLogParamsRules.NoLogParamsPrefix = append(noLogParamsRules.NoLogParamsPrefix, v)
	}
}

// SetNoLogParamsSuffix 设置不需要打印的入参和出参的路由后缀
func SetNoLogParamsSuffix(path ...string) {
	for _, v := range path {
		noLogParamsRules.NoLogParamsSuffix = append(noLogParamsRules.NoLogParamsSuffix, v)
	}
}

// CheckNoLogParams 判断是否需要打印入参出参日志, 不需要打印返回true
func CheckNoLogParams(origin string) bool {
	if len(noLogParamsRules.NoLogParams) > 0 {
		if _, ok := noLogParamsRules.NoLogParams[origin]; ok {
			return true
		}
	}

	if len(noLogParamsRules.NoLogParamsPrefix) > 0 {
		for _, v := range noLogParamsRules.NoLogParamsPrefix {
			if strings.HasPrefix(origin, v) {
				return true
			}
		}
	}

	if len(noLogParamsRules.NoLogParamsSuffix) > 0 {
		for _, v := range noLogParamsRules.NoLogParamsSuffix {
			if strings.HasSuffix(origin, v) {
				return true
			}
		}
	}

	return false
}
