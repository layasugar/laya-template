// Package converter 提供了一系列对请求的数据序列化和响应的数据格式化方法
package converter

// Converter 对请求的数据序列化和响应的数据格式化
type Converter interface {
	Pack(interface{}) ([]byte, error)
	UnPack(interface{}, interface{}) error
}

var (
	_ Converter = &FormConverter{}
	_ Converter = &JSONConverter{}
	_ Converter = &RawConverter{}
)

// ConverterType alias
type ConverterType string

const JSON ConverterType = "json"
const FORM ConverterType = "form"
const RAW ConverterType = "raw"

var converterPool = map[ConverterType]Converter{
	FORM: &FormConverter{},
	JSON: &JSONConverter{},
	RAW:  &RawConverter{},
}

// GetConverter 获取Converter
func GetConverter(converterName ConverterType) (c Converter, ok bool) {
	c, ok = converterPool[converterName]
	return
}

// RegisterConverter 注册Converter
func RegisterConverter(typ ConverterType, converter Converter) {
	converterPool[typ] = converter
}
