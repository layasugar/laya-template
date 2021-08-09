package controllers

import (
	"github.com/gin-gonic/gin"
)

// test
func (ctrl *BaseCtrl) Test(c *gin.Context) {
	var body map[string]interface{}
	_ = c.ShouldBindJSON(&body)

	//// 生产数据
	//partition, offset, err := dao.Kafka.SendMsg("layatest", "1111111111111")
	//if err != nil {
	//	log.Print(err.Error())
	//} else {
	//	log.Printf("Message partion: %d, Message offset: %d.", partition, offset)
	//}
	////钉钉推送
	//var alarmData = &glogs.AlarmData{
	//	Title:       "我是一个快乐的机器人",
	//	Description: "快乐的机器人",
	//	Content: map[string]interface{}{
	//		"time": time.Now(),
	//		"haha": "流弊机器人",
	//	},
	//}
	//glogs.SendDing(alarmData)
	//fmt.Println(body)
	ctrl.Suc(c, body, "success")
}
