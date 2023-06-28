package alarm

import "fmt"

type Alarm interface {
	Push(title string, content string, data map[string]interface{})
}

func NewContext() Alarm {
	return getAlarm()
}

type DefaultContext struct{}

func (ctx *DefaultContext) Push(title string, content string, data map[string]interface{}) {
	fmt.Printf("Alarm info, title: %s\r\ncontext: %s\r\ndata: %v\r\n", title, content, data)
}
