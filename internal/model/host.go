package model

// Host 主机配置表
type Host struct {
	ID     string `json:"id"`
	Name   string `json:"name"`   //主机名称
	Group  string `json:"group`   //主机分组
	IP     string `json:"ip"`     //主机ip
	Port   int    `json:"port"`   //主机端口
	User   string `json:"user"`   //登录用户名
	Passwd string `json:"passwd"` //登录密码
	Key    string `json:"key"`    //登录私钥
}
