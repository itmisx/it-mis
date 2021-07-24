package model

// HostLog 主机操作日志
type HostLog struct {
	ID         string `json:"id"`
	Log        string `json:"log"`         //日志内容
	CreateTime int64  `json:"create_time"` //创建时间
}
