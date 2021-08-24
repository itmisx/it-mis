package login

import (
	"it-mis/init/db"
	"it-mis/internal/model"
	"it-mis/internal/pkg/encrypt"
	"it-mis/internal/pkg/errorx"
	"it-mis/internal/pkg/randx"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     int64  `json:"id"`
	User   string `json:"user"`
	Status int    `json:"status"`
}

// LoginParams 登录参数
type LoginParams struct {
	User        string `json:"user"`
	Passwd      string `json:"passwd"`
	SecureCode  string `json:"secure_code"`
	SecureToken string `json:"secure_token"`
}

// LoginResult 登录结果
type LoginResult struct {
	UserToken string   `json:"user_token"`
	UserInfo  UserInfo `json:"user_info"`
}

type Login struct {
}

// ParamsValidate 参数校验
func (l Login) validateParams(c *gin.Context, scene string) (userInfo *UserInfo, err error) {
	var (
		user   model.User
		db     = db.New()
		params LoginParams
	)
	c.ShouldBindBodyWith(&params, binding.JSON)
	// 基本参数验证
	if params.User == "" {
		return nil, errorx.New("用户名不能为空")
	}
	if params.Passwd == "" {
		return nil, errorx.New("密码不能为空")
	}

	//
	userNameExits := model.User{}
	db.Model(&model.User{}).Where(&model.User{
		User: params.User,
	}).Take(&userNameExits)
	if userNameExits.ID == 0 {
		return nil, errorx.New("不存在的用户")
	}

	// 查询用户是否存在
	db.Model(&model.User{}).
		Where(&model.User{
			User:   params.User,
			Passwd: encrypt.Sha256(params.Passwd + userNameExits.Salt),
		}).
		Take(&user)
	if user.ID == 0 {
		return nil, errorx.New("用户或密码有误")
	}
	// 查询用户的状态
	if user.Status != 1 {
		return nil, errorx.New("用户已被删除或禁用")
	}
	// 特殊场景参数
	switch scene {
	case "secure_code":
		if params.SecureCode == "" {
			return nil, errorx.New("安全码不能为空")
		}
	case "secure_token":
		if params.SecureCode == "" {
			return nil, errorx.New("安全码不能为空")
		}
	}
	// 返回用户信息
	return &UserInfo{
		ID:     user.ID,
		User:   user.User,
		Status: user.Status,
	}, nil
}

// GetSecureToken 获取安全token
func (l Login) GetSecureToken(c *gin.Context) (secureToken string, err error) {
	var (
		userInfo = &UserInfo{}
		db       = db.New()
	)

	// 参数验证
	userInfo, err = l.validateParams(c, "secure_token")
	if err != nil {
		return "", err
	}
	// 生成secureToken
	token := ""
	count := db.Model(&model.UserSecureToken{}).
		Create(&model.UserSecureToken{
			UserID:     userInfo.ID,
			Token:      token,
			ExpireTime: time.Now().Unix() + 30*24*3600,
		}).
		RowsAffected
	if count == 1 {
		return token, nil
	}
	return "", errorx.New("生成安全令牌失败")
}

// GetSecureCode 获取安全码
func (l Login) GetSecureCode(c *gin.Context) (secureCode string, err error) {
	// 用户名、密码验证
	var (
		userInfo = &UserInfo{}
		db       = db.New()
	)

	// 参数验证
	userInfo, err = l.validateParams(c, "secure_token")
	if err != nil {
		return "", err
	}
	// 生成secureCode
	code := ""
	count := db.Model(&model.UserCode{}).
		Create(&model.UserCode{
			UserID:     userInfo.ID,
			Code:       code,
			ExpireTime: time.Now().Unix() + 180,
		}).
		RowsAffected
	if count == 1 {
		return code, nil
	}
	return "", errorx.New("生成安全码失败")
}

// Login 登录
// return token & userInfo
func (l Login) Login(c *gin.Context) (userInfo *UserInfo, err error) {
	// 参数校验
	userInfo, err = l.validateParams(c, "")
	if err != nil {
		return nil, err
	}
	// 返回用户信息
	return userInfo, nil
}

// Logout 登出
// 销毁token,secureToken
func (l Login) Logout(c *gin.Context) error {
	// 清理session信息
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	return nil
}

func genUserToken(c *gin.Context, userInfo UserInfo) string {
	token := randx.RandString(64)
	count := db.New().Model(&model.UserToken{}).
		Create(&model.UserToken{
			UserID:     userInfo.ID,
			Token:      token,
			ExpireTime: time.Now().Unix() + 3600*24*30,
		}).RowsAffected
	if count > 0 {
		return token
	}
	return ""
}

func ValidateUserToken(token string) bool {
	var count int64
	db.New().Model(&model.UserToken{}).
		Where("token = ?", token).
		Count(&count)
	if count > 0 {
		return true
	}
	return false
}
