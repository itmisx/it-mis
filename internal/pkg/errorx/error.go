package errorx

import (
	"it-mis/internal/pkg/debugx"
	"strconv"
)

// 自定义Error，并实现Error()方法
type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Error 打印错误信息
func (e Error) Error() string {
	return e.Msg
}

// 实例化Error
func New(msg string, code ...int) Error {
	var (
		e     Error
		print = debugx.Print{
			Skip: 2,
		}
	)
	e.Msg = msg
	if len(code) > 0 {
		e.Code = code[0]
	} else {
		e.Code = -1
	}
	//打印错误信息
	print.Print(msg+"("+strconv.Itoa(e.Code)+")", "error")
	// 返回错误
	return e
}
