package worker

import (
	"github.com/robfig/cron/v3"
	"log"
)

func (ctrl *controller) Top10() {
	c := cron.New()
	_, _ = c.AddFunc("@midnight", Top10Timer)
	_, _ = c.AddFunc("@every 250s", Top10Timer)
	_, _ = c.AddFunc("0 23 * * *", Top10Timer)
	c.Start()
	log.Println("来了老弟")
}

func Top10Timer() {
	log.Println("来了老弟")
}
