package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//import "laya-go/main/cmd"

//func main() {
//	//c := cmd.Cmd{}
//	//c.Run()
//	t := time.Now()
//	var wg sync.WaitGroup
//	f, _ := os.OpenFile("a.txt", os.O_WRONLY|os.O_APPEND, 0644)
//	content := "每行只有3个数字\r"
//	var i int64
//	for i = 1; i <= 10000000; i++ {
//		wg.Add(1)
//		go func(f *os.File, c string) {
//			_, _ = f.Write([]byte(c))
//			defer wg.Add(-1)
//		}(f, content)
//	}
//
//	wg.Wait()
//	defer f.Close()
//	elapsed := time.Since(t)
//	fmt.Println("app run time", elapsed)
//}
func main() {
	r := gin.Default()
	r.Use(response())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})

	r.GET("/status", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.Set("message", "hello, world")
	})

	r.Run()
}

func response() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Written() {
			return
		}

		params := c.Keys
		if len(params) == 0 {
			return
		}
		c.JSON(http.StatusOK, params)
	}
}