package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// Md5 md5
func Md5(s string) string {
	m := md5.Sum([]byte(s))
	return hex.EncodeToString(m[:])
}

// GenerateLogId 获取logId
func GenerateLogId() string {
	s := uuid.NewV4().String()
	return Md5(s)
}

// GetString 只能是map和slice
func GetString(d interface{}) string {
	bytesD, err := CJson.Marshal(d)
	if err != nil {
		return fmt.Sprintf("%v", d)
	} else {
		return string(bytesD)
	}
}
