package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// test
func (ctrl *BaseCtrl) Test(c *gin.Context) {
	var body map[string]interface{}
	_ = c.ShouldBindJSON(&body)
	fmt.Println(body)
	ctrl.Suc(c, body, "success")
}
