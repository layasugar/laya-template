package gcnf

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Get(key string) interface{} {
	return viper.Get(key)
}
func GetBool(key string) bool {
	return viper.GetBool(key)
}
func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}
func GetInt(key string) int {
	return viper.GetInt(key)
}
func GetIntSlice(key string) []int {
	return viper.GetIntSlice(key)
}
func GetString(key string) string {
	return viper.GetString(key)
}
func GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}
func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}
func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}
func GetTime(key string) time.Time {
	return viper.GetTime(key)
}
func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}
func IsSet(key string) bool {
	return viper.IsSet(key)
}
func AllSettings() map[string]interface{} {
	return viper.AllSettings()
}

// OnConfigCharge 注册配置改变后的处理
func OnConfigCharge(run func(in fsnotify.Event)) {
	viper.OnConfigChange(run)
}

// LoadErrMsg 根据code加载提示信息
func LoadErrMsg(code uint32) string {
	key := fmt.Sprintf("err_code.%d", code)
	s := viper.GetString(key)
	return s
}

func GetConfigMap(key string) []map[string]interface{} {
	var configMaps []map[string]interface{}
	b := viper.Get(key)
	switch b.(type) {
	case []interface{}:
		si := b.([]interface{})
		for _, item := range si {
			if sim, ok := item.(map[string]interface{}); ok {
				configMaps = append(configMaps, sim)
			}
		}
	}
	return configMaps
}
