package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

//import "laya-go/main/cmd"

func main() {
	//c := cmd.Cmd{}
	//c.Run()
	t := time.Now()
	var wg sync.WaitGroup
	f, _ := os.OpenFile("a.txt", os.O_WRONLY|os.O_APPEND, 0644)
	content := "每行只有3个数字\r"
	var i int64
	for i = 1; i <= 10000000; i++ {
		wg.Add(1)
		go func(f *os.File, c string) {
			_, _ = f.Write([]byte(c))
			defer wg.Add(-1)
		}(f, content)
	}

	wg.Wait()
	defer f.Close()
	elapsed := time.Since(t)
	fmt.Println("app run time", elapsed)
}
