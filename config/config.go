package config

import "github.com/layasugar/laya/gcnf"

func GetZkAddr() string {
	return gcnf.GetString("extra.auto_metrics")
}
