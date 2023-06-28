package util

import (
	jsonIter "github.com/json-iterator/go"
)

var (
	// CJson 全局json序列化和反序列化
	CJson = jsonIter.ConfigCompatibleWithStandardLibrary
)
