package model

// SystemConfig 系统配置表
type SystemConfig struct {
	ID         string `json:"id"`
	Key        string `json:"key"`
	Value      string `json:"value"`
	EncrypType string `json:"encryp_type"` //关键信息会加密，如采用rsa等
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}
