package errno

// RspError
type RspError struct {
	Code uint32
	Msg  string
}

func (re *RspError) Error() string {
	return re.Msg
}

func Err(code uint32, msg string) (err error) {
	err = &RspError{
		Code: code,
		Msg:  msg,
	}
	return err
}

// Render
func (re *RspError) Render() (code uint32, msg string) {
	return re.Code, re.Msg
}
