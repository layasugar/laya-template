package global

import (
	"github.com/layasugar/laya-go/models/dao"
	"log"
)

func Init() {
	// 先激活消费
	go work()

	// 再初始化消费者
	go dao.Kafka.InitConsumer(DataChan)

	// 生产数据
	partition, offset, err := dao.Kafka.SendMsg("layatest", "asdddddddasdadasd")
	if err != nil {
		log.Print(err.Error())
	} else {
		log.Printf("Message partion: %d, Message offset: %d.", partition, offset)
	}
	return
}

// 消费kafka
func work() {
	for item := range DataChan {
		log.Printf(string(item))
	}
}
