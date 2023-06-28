package converter

import (
	"bytes"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type JSONConverter struct{}

var jsoner = jsoniter.ConfigCompatibleWithStandardLibrary

// Pack the data package
func (*JSONConverter) Pack(data interface{}) ([]byte, error) {
	switch t := data.(type) {
	case string:
		return []byte(t), nil
	case []byte:
		return t, nil
	default:
		res, err := jsoner.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("json pack error: %s", err.Error())
		}
		return res, nil
	}
}

// UnPack the data package
func (*JSONConverter) UnPack(data interface{}, rsp interface{}) error {
	dec := jsoner.NewDecoder(bytes.NewReader(data.([]byte)))
	dec.UseNumber()
	err := dec.Decode(rsp)
	if err != nil {
		return fmt.Errorf("json unpack error: %s", err.Error())
	}
	return nil
}
