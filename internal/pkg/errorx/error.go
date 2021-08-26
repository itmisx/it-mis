package errorx

import (
	"it-mis/internal/pkg/debugx"
	"strconv"
)

// 自定义Error，并实现Error()方法
type Error struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	ExtraCode *int   `json:"extra_code"`
}

// Error 打印错误信息
func (e Error) Error() string {
	if e.ExtraCode != nil {
		return e.Msg + "-" + strconv.Itoa(*e.ExtraCode)
	}
	return e.Msg
}

// 实例化Error
func New(msg string, code ...int) Error {
	var (
		e  Error
		dp = debugx.Print{
			Skip: 2,
		}
	)
	e.Msg = msg
	if len(code) == 1 {
		e.Code = code[0]
	} else if len(code) == 2 {
		e.Code = code[0]
		e.ExtraCode = &code[1]
	} else {
		e.Code = -1
	}
	//打印错误信息
	dp.Print(msg+"("+strconv.Itoa(e.Code)+")", "error")
	// 返回错误
	return e
}
