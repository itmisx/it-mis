package model

// User 用户表
type User struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"passwod"`
}
