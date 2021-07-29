package global

import (
	"github.com/layasugar/laya-go/models/dao"
	"github.com/layasugar/laya/gconf"
	"github.com/layasugar/laya/genv"
	"log"
	"sync"
	"time"
)

var lock sync.Mutex

var changeTime time.Time

// 配置文件变化默认回调
func ConfChangeHandler(we *gconf.WatcherEvent) error {
	lock.Lock()
	defer lock.Unlock()
	now := time.Now()
	// 50毫秒以内的事件是不处理得
	if !changeTime.IsZero() {
		if now.Before(changeTime.Add(50 * time.Millisecond)) {
			return nil
		}
	}
	changeTime = now
	log.Print(we)
	switch we.Type {
	case gconf.WatcherEventChange:
		log.Printf("[file_watcher] file changed, begin reload config")
		// 重载配置
		err := gconf.InitConfig(genv.ConfigPath)
		if err != nil {
			return nil
		}
		log.Printf("[file_watcher] config reload success")
		// 重载连接池
		go dao.Init()
	case gconf.WatcherEventCreate:
	case gconf.WatcherEventDelete:
	}
	return nil
}
