package model

// User 用户表
type User struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"` //用户名
	Password string `json:"passwod"`   //密码
	Salt     string `json:"salt"`      //密码加密盐
}
