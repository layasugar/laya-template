package response

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"strings"
)

//取出头部language
func GetLang(Lang string) string {
	rs := []rune(Lang)
	result := string(rs[0:2])
	return result
}

//获取对应语言信息
func GetLanguage(parameter string, language string) string {
	result := strings.Split(parameter, "|")
	res := make([]interface{}, len(result))
	for i, v := range result {
		res[i] = v
	}
	var message string
	var EntireLang = make(map[string]string)
	config.Get(result[0]).Scan(&EntireLang)
	for k, v := range EntireLang {
		if language == k {
			message = v
		}
	}

	slice4 := res[1:]
	message = fmt.Sprintf(message, slice4...)

	return message
}
