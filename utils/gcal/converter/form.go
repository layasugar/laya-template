package converter

import (
	"errors"
	"net/url"
	"strings"
)

type FormConverter struct{}

// Pack the data package
func (*FormConverter) Pack(data interface{}) ([]byte, error) {
	switch t := data.(type) {
	case url.Values:
		return []byte(t.Encode()), nil
	case map[string]string:
		form := url.Values{}
		for k, v := range t {
			form.Add(k, v)
		}
		return []byte(strings.TrimSpace(form.Encode())), nil
	}
	return nil, errors.New("form pack error: unknow form data type")
}

// UnPack the data package
func (*FormConverter) UnPack(data interface{}, rsp interface{}) error {
	return errors.New("form unpack error: not support form unpack")
}
