package file

import (
	"github.com/layasugar/laya"
)

//type res struct {
//	ImgUrl string
//}

//const uploadPath = "files"

func (ctrl *controller) Upload(ctx *laya.Context) {
	//file, _ := c.FormFile("File")
	//if file == nil {
	//	c.Set("$.Upload.NoFile.code", response.NoFile)
	//	return
	//}
	//fileType := c.PostForm("FileType")
	//if fileType == "" {
	//	c.Set("$.Upload.NoFileType.code", response.NoFileType)
	//	return
	//}
	//
	//token := c.GetHeader("Token")
	//var uid string
	//
	//if token == "" {
	//	c.Set("$.Upload.TokenErr.code", response.TokenErr)
	//	return
	//}
	//uid, _ = rdb.Dao.HGet(context.Background(), "user:uid", token).Result()
	//if uid == "" {
	//	c.Set("$.Upload.TokenErr.code", response.TokenErr)
	//	return
	//}
	//
	//// initialize filepath
	//newPath := ctrl.getNewPath(fileType, uid, file.Filename)
	//
	//imgUrl := newPath
	//imgUrl += "?t=" + strconv.FormatInt(time.Now().Unix(), 10)
	//
	//err := c.SaveUploadedFile(file, newPath)
	//if err != nil {
	//	c.Set("$.Upload.SaveUploadedFail.code", response.SaveUploadedFail)
	//	return
	//}
	//
	////c.Set("$.Upload.Success.response", response.Response{Code: response.Success, Data: resp{ImgUrl: imgUrl}})
	//return
}

//
//func (ctrl *controller) getNewPath(fileType string, uid string, fileName string) string {
//	// initialize filepath
//	path := uploadPath
//	kv := strings.Split(fileName, ".")
//	mimeType := kv[len(kv)-1]
//	newName := randToken(12)
//	newPath := filepath.Join(path, newName+"."+mimeType)
//	path += "/admin"
//	_ = os.MkdirAll(path, 777)
//	newPath = filepath.Join(path, newName+"."+mimeType)
//	return newPath
//}
//
//func randToken(len int) string {
//	b := make([]byte, len)
//	_, _ = rand.Read(b)
//	return fmt.Sprintf("%x", b)
//}
