package datetime

import (
	"strings"
	"time"
)

var (
	// 时区，如8对应东八区
	default_timezone interface{}
)

// SetDefaultTimeZone 设置默认时区
func SetDefaultTimeZone(timezone int) {
	default_timezone = timezone
}

// Date 时间日期格式化显示
// format 格式化  Y-m-d H:i:s
// timestamp 时间戳，-1为当前时间戳
// timezone 时区，nil为系统默认时区
func Date(format string, timestamp ...int64) string {
	var ts = time.Now().Unix()
	if len(timestamp) > 0 && timestamp[0] > 0 {
		ts = timestamp[0]
	}
	// 全部转为小写
	format = strings.ToLower(format)
	// 格式化替换为go内置的格式化字符串
	format = strings.ReplaceAll(format, "y", "2006")
	format = strings.ReplaceAll(format, "m", "1")
	format = strings.ReplaceAll(format, "d", "02")
	format = strings.ReplaceAll(format, "h", "15")
	format = strings.ReplaceAll(format, "i", "04")
	format = strings.ReplaceAll(format, "s", "05")
	// 时间戳转time.Time
	tm := time.Unix(ts, 0)
	// 指定时区
	if timezone, ok := default_timezone.(int); ok {
		var cstZone = time.FixedZone("CST", timezone*3600)
		return tm.In(cstZone).Format(format)
	}
	// 默认本地时区
	return tm.Format(format)
}
