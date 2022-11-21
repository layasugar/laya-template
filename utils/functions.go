package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// IsMobile 判断是否是手机号
func IsMobile(s string) bool {
	exp := `^1[3|4|5|6|7|8|9]\d{9}$`
	compile := regexp.MustCompile(exp)
	findString := compile.MatchString(s)
	return findString
}

// RandToken 生成用户token
func RandToken() string {
	rand.Seed(time.Now().UnixNano())
	text := randomStr(rand.Intn(20))
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//	RandomStr 随机生成字符串
func randomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	byt := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, byt[r.Intn(len(byt))])
	}
	return string(result) + strconv.FormatInt(time.Now().UnixNano(), 19)
}

func InSliceUint8(k uint8, s []uint8) bool {
	for _, v := range s {
		if k == v {
			return true
		}
	}
	return false
}
