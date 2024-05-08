package converter

import (
	"errors"
	"reflect"
)

type RawConverter struct{}

// Pack the data package
func (*RawConverter) Pack(data interface{}) ([]byte, error) {
	switch t := data.(type) {
	case string:
		return []byte(t), nil
	case []byte:
		return t, nil
	default:
		return nil, errors.New("raw pack error: unknown raw data type")
	}
}

// UnPack the data package
func (*RawConverter) UnPack(data interface{}, res interface{}) error {
	rtype := reflect.TypeOf(res)
	rvalue := reflect.ValueOf(res)
	if rtype.Kind() != reflect.Ptr || rvalue.IsNil() {
		return errors.New("raw unpack error: can't unpack raw converter, converter type must by pointer and not nil")
	}
	if rvalue.Elem().Kind() != reflect.Slice || !rvalue.Elem().CanSet() {
		return errors.New("raw unpack error: can't unpack raw converter, converter result must be []byte, and can set")
	}
	rvalue.Elem().Set(reflect.ValueOf(data))
	return nil
}
