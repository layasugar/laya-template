## gin_pprof

#### use

```go
package main

import (
	"github.com/layasugar/laya-template/pkg/core/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// automatically add routers for net/http/pprofx
	// e.g. /debug/pprofx, /debug/pprofx/heap, etc.
	pprof.Wrap(router)

	// pprofx also plays well with *gin.RouterGroup
	// group := router.Group("/debug/pprofx")
	// pprofx.WrapGroup(group)

	router.Run(":8080")
}
```