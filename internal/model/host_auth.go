package model

// HostAuth 主机授权
type HostAuth struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`  // 用户id
	HostID  string `json:"host_id"`  // 主机id
	GroupID string `json:"group_id"` // 主机组id
}
