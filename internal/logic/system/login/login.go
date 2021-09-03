package login

import (
	"it-mis/init/db"
	"it-mis/internal/logic/system/member"
	"it-mis/internal/model"
	"it-mis/internal/pkg/encrypt"
	"it-mis/internal/pkg/errorx"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// LoginParams 登录参数
type LoginParams struct {
	User        string `json:"user"`
	Passwd      string `json:"passwd"`
	SecureCode  string `json:"secure_code"`
	SecureToken string `json:"secure_token"`
}

// LoginResult 登录结果
type LoginResult struct {
	UserToken string          `json:"user_token"`
	UserInfo  member.UserInfo `json:"user_info"`
}

type Login struct{}

// ParamsValidate 参数校验
func (l Login) validateParams(c *gin.Context, scene string) (userp *model.User, err error) {
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
	return &user, nil
}

// Login 登录
// return token & userInfo
func (l Login) Login(c *gin.Context) (userInfo *member.UserInfo, err error) {
	// 参数校验
	userp, err := l.validateParams(c, "")
	if err != nil {
		return nil, err
	}
	// 获取用户信息
	userInfo, _ = member.Member{}.Info(userp.ID)
	// 返回用户信息
	return userInfo, nil
}

// Logout 登出
func (l Login) Logout(c *gin.Context) error {
	userToken, _ := c.Cookie("user_token")
	// 删除userToken记录
	db.New().Model(&model.UserToken{}).
		Where("token = ?", userToken).
		Delete(&model.UserToken{})
	// 重置user_token cookie
	c.SetCookie("user_token", "", 0, "/", CookieDomain(c), false, true)
	return nil
}
