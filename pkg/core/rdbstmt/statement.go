package rdbstmt

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type Statement interface {
	Name() string
	FullName() string
	Args() []interface{}
	String() string
	ShortString() string
}

type Replacer func(*BaseStatement)

func NewStatement(args []interface{}) Statement {
	var copied = make([]interface{}, len(args))
	copy(copied, args)

	b := &BaseStatement{
		args: copied,
	}
	b.replacer = getReplacer(b.Name())
	return b
}

type BaseStatement struct {
	args     []interface{}
	replacer Replacer
}

func (b *BaseStatement) Name() string {
	if len(b.args) == 0 {
		return ""
	}

	return strings.ToLower(b.stringArg(0))
}

func (b *BaseStatement) FullName() string {
	switch name := b.Name(); name {
	case "cluster", "command":
		if len(b.args) == 1 {
			return name
		}
		if s2, ok := b.args[1].(string); ok {
			return name + " " + s2
		}
		return name
	default:
		return name
	}
}

func (b *BaseStatement) Args() []interface{} {
	return b.args
}

func (b *BaseStatement) stringArg(pos int) string {
	if pos < 0 || pos >= len(b.args) {
		return ""
	}
	s, _ := b.args[pos].(string)
	return s
}

func (b *BaseStatement) ShortString() string {
	var copied []interface{}

	if nil != b.replacer {
		copied = make([]interface{}, len(b.args))
		copy(copied, b.args)
		b.replacer(b)
	}

	str := b.String()
	if copied != nil {
		b.args = copied
	}
	return str
}

func (b *BaseStatement) String() string {
	p := make([]byte, 0, 64)

	for i, arg := range b.Args() {
		if i > 0 {
			p = append(p, ' ')
		}
		p = appendArg(p, arg)
	}
	return string(p)
}

func appendRune(b []byte, r rune) []byte {
	if r < utf8.RuneSelf {
		switch c := byte(r); c {
		case '\n':
			return append(b, "\\n"...)
		case '\r':
			return append(b, "\\r"...)
		default:
			return append(b, c)
		}
	}

	l := len(b)
	b = append(b, make([]byte, utf8.UTFMax)...)
	n := utf8.EncodeRune(b[l:l+utf8.UTFMax], r)
	b = b[:l+n]

	return b
}

func appendUTF8String(b []byte, s string) []byte {
	for _, r := range s {
		b = appendRune(b, r)
	}
	return b
}

func appendArg(b []byte, v interface{}) []byte {
	switch v := v.(type) {
	case nil:
		return append(b, "<nil>"...)
	case string:
		return appendUTF8String(b, v)
	case []byte:
		return appendUTF8String(b, string(v))
	case int:
		return strconv.AppendInt(b, int64(v), 10)
	case int8:
		return strconv.AppendInt(b, int64(v), 10)
	case int16:
		return strconv.AppendInt(b, int64(v), 10)
	case int32:
		return strconv.AppendInt(b, int64(v), 10)
	case int64:
		return strconv.AppendInt(b, v, 10)
	case uint:
		return strconv.AppendUint(b, uint64(v), 10)
	case uint8:
		return strconv.AppendUint(b, uint64(v), 10)
	case uint16:
		return strconv.AppendUint(b, uint64(v), 10)
	case uint32:
		return strconv.AppendUint(b, uint64(v), 10)
	case uint64:
		return strconv.AppendUint(b, v, 10)
	case float32:
		return strconv.AppendFloat(b, float64(v), 'f', -1, 64)
	case float64:
		return strconv.AppendFloat(b, v, 'f', -1, 64)
	case bool:
		if v {
			return append(b, "true"...)
		}
		return append(b, "false"...)
	case time.Time:
		return v.AppendFormat(b, time.RFC3339Nano)
	default:
		return append(b, fmt.Sprint(v)...)
	}
}
