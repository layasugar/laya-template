package handler

import (
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"laya-go/ship"
	r "laya-go/ship/response"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type res struct {
	ImgUrl string
}

const uploadPath = "files"

func Upload(c *gin.Context) {
	file, _ := c.FormFile("File")
	if file == nil {
		c.Set("$.Upload.NoFile.code", r.NoFile)
		return
	}
	fileType := c.PostForm("FileType")
	if fileType == "" {
		c.Set("$.Upload.NoFileType.code", r.NoFileType)
		return
	}

	token := c.GetHeader("Token")
	var uid string
	if fileType != ship.FileTypeAdmin {
		if token == "" {
			c.Set("$.Upload.TokenErr.code", r.TokenErr)
			return
		}
		uid, _ = ship.Redis.HGet("user:uid", token).Result()
		if uid == "" {
			c.Set("$.Upload.TokenErr.code", r.TokenErr)
			return
		}
	} else {
		uid = "0"
	}
	// initialize filepath
	newPath := getNewPath(fileType, uid, file.Filename)

	imgUrl := newPath
	if fileType == ship.FileTypeAvatar {
		imgUrl += "?t=" + strconv.FormatInt(time.Now().Unix(), 10)
	}
	err := c.SaveUploadedFile(file, newPath)
	if err != nil {
		c.Set("$.Upload.SaveUploadedFail.code", r.SaveUploadedFail)
		return
	}

	c.Set("$.Upload.Success.response", r.Response{Code: r.Success, Data: res{ImgUrl: imgUrl}})
	return
}

func getNewPath(fileType string, uid string, fileName string) string {
	// initialize filepath
	path := uploadPath
	kv := strings.Split(fileName, ".")
	mimeType := kv[len(kv)-1]
	newName := randToken(12)
	newPath := filepath.Join(path, newName+"."+mimeType)
	switch fileType {
	case ship.FileTypeAvatar:
		path += "/" + uid + "/avatar"
		_ = os.MkdirAll(path, 777)
		newPath = filepath.Join(path, uid+"."+mimeType)
	case ship.FileTypeUser:
		path += "/" + uid + "/other"
		_ = os.MkdirAll(path, 777)
		newPath = filepath.Join(path, newName+"."+mimeType)
	case ship.FileTypeAdmin:
		path += "/admin"
		_ = os.MkdirAll(path, 777)
		newPath = filepath.Join(path, newName+"."+mimeType)
	}
	return newPath
}

func randToken(len int) string {
	b := make([]byte, len)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}
