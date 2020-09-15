package main

import (
	"laya-go/server/worker/handler"
	"laya-go/ship"
)

func main() {
	ship.Init()
	handler.Top10()
}
