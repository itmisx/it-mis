package debugx

import (
	"fmt"
	"it-mis/internal/pkg/datetime"
	"runtime"
	"strconv"
)

// PrintSuccess 打印信息
// 带行号及颜色
func PrintSuccess(msg interface{}, detail ...interface{}) {
	print(msg, "success", detail...)
}

// PrintWarning 打印信息
// 带行号及颜色
func PrintWarning(msg interface{}, detail ...interface{}) {
	print(msg, "warning", detail...)
}

// PrintInfo 打印信息
// 带行号及颜色
func PrintInfo(msg interface{}, detail ...interface{}) {
	print(msg, "info", detail...)
}

// PrintError 打印错误
// 带行号及颜色
func PrintError(msg interface{}, detail ...interface{}) {
	print(msg, "error", detail...)
}

// print 打印内容
// print_type, error or  info
func print(msg interface{}, printType string, detail ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	time := "[" + datetime.Date("h:i:s", -1) + "]"
	position := file + ":" + strconv.Itoa(line)
	errStr := fmt.Sprintf("%+v", msg)
	foreground := "37"
	switch printType {
	case "error":
		foreground = "31"
	case "success":
		foreground = "32"
	case "warning":
		foreground = "33"
	case "info":
		foreground = "34"
	}
	// 打印时间、位置
	printWithColor(time+position, foreground, "", 1, "1")
	// 打印错误信息
	printWithColor(errStr, foreground, "", 1, "1")
	// 循环打印详细信息
	for _, d := range detail {
		dstr := fmt.Sprint(d)
		printWithColor(dstr, "32", "", 1, "1")
	}
	// 全部信息打印结束，增加一行换行
	printWithColor("", "32", "", 1, "1")
}

// printWithColor 带颜色的打印
///
// 前景 背景 颜色
//
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
//
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见
func printWithColor(msg string, foreground string, background string, newline int, code ...string) {
	format := "%c["
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
