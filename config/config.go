package config

import "github.com/layasugar/laya-template/pkg/gcnf"

func GetZkAddr() string {
	return gcnf.GetString("extra.auto_metrics")
}
