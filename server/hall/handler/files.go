package handler
//
//import (
//	"crypto/rand"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"laya-go/ship"
//	"net/http"
//	"os"
//	"path/filepath"
//	"strconv"
//	"strings"
//	"time"
//)
//
//type res struct {
//	ImgUrl string
//}
//
//const uploadPath = "files"
//
//func Upload(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := ship.GetLang(lang)
//	file, _ := c.FormFile("File")
//	if file == nil {
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//	fileType := c.PostForm("Type")
//	token := c.GetHeader("Token")
//	var uid string
//	if fileType != "3" {
//		uid, _ = ship.Redis.HGet("user:uid", token).Result()
//		if uid == "" {
//			c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//			return
//		}
//	} else {
//		uid = "0"
//	}
//	// initialize filepath
//	newPath := getNewPath(fileType, uid, file.Filename)
//
//	imgUrl := newPath
//	if fileType == "1" {
//		imgUrl += "?t=" + strconv.FormatInt(time.Now().Unix(), 10)
//	}
//	err := c.SaveUploadedFile(file, newPath)
//	if err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//	}
//
//	c.JSON(http.StatusOK, ship.GetMessage("Success", 1, language, res{ImgUrl: imgUrl}))
//}
//
//func getNewPath(fileType string, uid string, fileName string) string {
//	// initialize filepath
//	path := uploadPath
//	kv := strings.Split(fileName, ".")
//	mimeType := kv[len(kv)-1]
//	newName := randToken(12)
//	newPath := filepath.Join(path, newName+"."+mimeType)
//	switch fileType {
//	case "1":
//		path += "/" + uid + "/avatar"
//		_ = os.MkdirAll(path, 777)
//		newPath = filepath.Join(path, uid+"."+mimeType)
//	case "2":
//		path += "/" + uid + "/recharge"
//		_ = os.MkdirAll(path, 777)
//		newPath = filepath.Join(path, newName+"."+mimeType)
//	case "3":
//		path += "/admin"
//		_ = os.MkdirAll(path, 777)
//		newPath = filepath.Join(path, newName+"."+mimeType)
//	}
//	return newPath
//}
//
//func randToken(len int) string {
//	b := make([]byte, len)
//	_, _ = rand.Read(b)
//	return fmt.Sprintf("%x", b)
//}
