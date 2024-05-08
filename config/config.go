package config

import "github.com/layasugar/laya-template/utils/gcnf"

func GetZkAddr() string {
	return gcnf.GetString("extra.auto_metrics")
}
