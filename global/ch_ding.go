package global

import "log"

var DingCh = make(chan *AlarmData, 10)

func InitDingCh() {
	go func() {
		for {
			d := <-DingCh
			err := SendAlarm(d)
			if err != nil {
				log.Printf("发送钉钉失败：err=%s", err.Error())
			}
		}
	}()
}
