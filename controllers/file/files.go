package file

import (
	"crypto/rand"
	"fmt"
	"github.com/LaYa-op/laya"
	"github.com/LaYa-op/laya/response"
	"github.com/gin-gonic/gin"
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

func (ctrl *controller) Upload(c *gin.Context) {
	file, _ := c.FormFile("File")
	if file == nil {
		c.Set("$.Upload.NoFile.code", response.NoFile)
		return
	}
	fileType := c.PostForm("FileType")
	if fileType == "" {
		c.Set("$.Upload.NoFileType.code", response.NoFileType)
		return
	}

	token := c.GetHeader("Token")
	var uid string
	if fileType != laya.FileTypeAdmin {
		if token == "" {
			c.Set("$.Upload.TokenErr.code", response.TokenErr)
			return
		}
		uid, _ = laya.Redis.HGet("user:uid", token).Result()
		if uid == "" {
			c.Set("$.Upload.TokenErr.code", response.TokenErr)
			return
		}
	} else {
		uid = "0"
	}
	// initialize filepath
	newPath := ctrl.getNewPath(fileType, uid, file.Filename)

	imgUrl := newPath
	if fileType == laya.FileTypeAvatar {
		imgUrl += "?t=" + strconv.FormatInt(time.Now().Unix(), 10)
	}
	err := c.SaveUploadedFile(file, newPath)
	if err != nil {
		c.Set("$.Upload.SaveUploadedFail.code", response.SaveUploadedFail)
		return
	}

	//c.Set("$.Upload.Success.response", response.Response{Code: response.Success, Data: res{ImgUrl: imgUrl}})
	return
}

func (ctrl *controller) getNewPath(fileType string, uid string, fileName string) string {
	// initialize filepath
	path := uploadPath
	kv := strings.Split(fileName, ".")
	mimeType := kv[len(kv)-1]
	newName := randToken(12)
	newPath := filepath.Join(path, newName+"."+mimeType)
	switch fileType {
	case laya.FileTypeAvatar:
		path += "/" + uid + "/avatar"
		_ = os.MkdirAll(path, 777)
		newPath = filepath.Join(path, uid+"."+mimeType)
	case laya.FileTypeUser:
		path += "/" + uid + "/other"
		_ = os.MkdirAll(path, 777)
		newPath = filepath.Join(path, newName+"."+mimeType)
	case laya.FileTypeAdmin:
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
