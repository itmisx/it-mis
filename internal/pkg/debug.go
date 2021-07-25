package pkg

import (
	"fmt"
	"runtime"
	"strconv"
)

//
func PrintError(err error) {
	_, file, line, _ := runtime.Caller(1)
	position := file + ":" + strconv.Itoa(line)
	PrintWithColor(position, "31", "", 0, "1")
	PrintWithColor(err.Error(), "32", "", 2, "1")
}

// printWithColor 带颜色的打印
// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见
func PrintWithColor(msg string, foreground string, background string, newline int, code ...string) {
	format := "\n %c["
	body := []interface{}{}
	for _, c := range code {
		format = format + c + ";"
	}
	format = format + foreground + ";" + background + "m"
	format = format + "%s%c[0m"
	for newline > 0 {
		newline--
		format = format + "\n"
	}
	body = append(body, 0x1B, msg, 0x1B)
	fmt.Printf(format, body...)
}
