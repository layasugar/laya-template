package config

import "github.com/layasugar/laya/gcf"

func GetZkAddr() string {
	return gcf.GetString("extra.auto_metrics")
}
