package dao

import (
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

var Mem *cache.Cache

func InitMemory() {
	Mem = cache.New(0, 1000*time.Minute)
	log.Printf("[store_memory] open success")
}
