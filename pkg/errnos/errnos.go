package errnos

import (
	"errors"
	"fmt"
)

const (
	//common
	CodeSuccess    = 0
	CodeServiceErr = 10001
	CodeParamErr   = 10002
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	//common
	Success    = NewErrNo(int64(CodeSuccess), "Success")
	ServiceErr = NewErrNo(int64(CodeServiceErr), "Service is unable to start successfully")
	ParamErr   = NewErrNo(int64(CodeParamErr), "Wrong Parameter has been given")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
