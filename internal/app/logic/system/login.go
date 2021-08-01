package system

import (
	"errors"
	"it-mis/init/mysql"
	"it-mis/internal/model"
	encrpt "it-mis/internal/pkg/encypt"
)

type PassedLogin struct {
}

// Login 登录
// user 用户名
// passwd 密码
func (PassedLogin) GetUser(user string, passwd string) (*model.User, error) {
	var (
		db           = mysql.NewDB()
		passwdSha256 string
		userExist    *model.User
	)
	// 密码明文sha256加密
	passwdSha256 = encrpt.GetMd5FromString(passwd)
	// 查询用户
	db.Model(&model.User{}).Where(&model.User{
		User:   user,
		Passwd: passwdSha256,
	}).Take(&userExist)
	if userExist.ID <= 0 {
		return nil, errors.New("用户或密码错误")
	}
	// 查询用户的状态
	if userExist.Status != 1 {
		return nil, errors.New("用户已被禁用或锁定")
	}
	return userExist, nil
}

type SecureLogin struct {
}

// Login 登录
// user 用户名
// passwd 密码
// secureToken 安全令牌
func (SecureLogin) Login(user string, passwd string, secureToken string) (*model.User, error) {
	var (
		db        = mysql.NewDB()
		userExist *model.User
	)
	// 验证用户名及密码
	userExist, err := PassedLogin{}.GetUser(user, passwd)
	if err != nil {
		return nil, err
	}
	// 验证安全令牌
	db.Model(&model.UserSecureToken{}).
		Where(&model.UserSecureToken{
			UserID: userExist.ID,
		})
}
