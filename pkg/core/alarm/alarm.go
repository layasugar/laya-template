package alarm

import (
	"github.com/layasugar/laya/gcnf"
)

const (
	DINGDING = "dingding"
	FEISHU   = "feishu"
	WEBHOOK  = "web_hook"
)

type Data struct {
	Level       int                    // 告警等级
	Title       string                 // 报警标题
	Description string                 // 报警描述
	Content     map[string]interface{} // kv数据
}

var alarm Alarm

func getAlarm() Alarm {
	if nil == alarm {
		if gcnf.AlarmType() == "" {
			return &DefaultContext{}
		}

		switch gcnf.AlarmType() {
		case DINGDING:
			alarm = &DingContext{
				robotKey:  gcnf.AlarmKey(),
				robotHost: gcnf.AlarmHost(),
			}
		default:
			return &DefaultContext{}
		}
	}
	return alarm
}
