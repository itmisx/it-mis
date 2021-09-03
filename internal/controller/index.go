package controller

import (
	"it-mis/internal/logic/system/login"
	"it-mis/internal/logic/system/member"
	"it-mis/internal/pkg/errorx"
	"it-mis/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

// GetSecureToken 获取安全token
func (i IndexController) GetSecureToken(c *gin.Context) {
	// data, err := login.GetSecureToken(c)
}

// GetSecureCode 获取安全码
func (i IndexController) GetSecureCode(c *gin.Context) {
}

// 获取csrf_token
func (i IndexController) CSRFToken(c *gin.Context) {
	response.JSON(c, gin.H{
		"csrf_token": login.SetCSRFToken(c),
	}, nil)
}

// Login 登录
func (i IndexController) Login(c *gin.Context) {
	// 登录，并获取登录用户的信息，及登录凭证
	user, err := login.Login{}.Login(c)
	//  错误返回
	if err != nil {
		response.JSON(c, nil, err)
		return
	}
	// 设置userToken
	login.SetUserToken(c, user.ID)
	// 获取用户信息
	userInfo, err := member.Member{}.Info(user.ID)
	if err != nil {
		response.JSON(c, nil, err)
		return
	}
	// 返回结果
	response.JSON(c, userInfo, nil)
}

// LoginStatus 登录状态
func (i IndexController) LoginStatus(c *gin.Context) {
	// 获取user_token
	userToken, _ := c.Cookie("user_token")
	if userToken == "" {
		response.JSON(c, nil, errorx.New("", 200000))
		return
	}
	// 获取user_token_info
	tokenInfo := login.GetUserTokenInfo(userToken)
	if tokenInfo == nil {
		response.JSON(c, nil, errorx.New("", 200000))
		return
	}
	userInfo, err := member.Member{}.Info(tokenInfo.UserID)
	// 查询用户信息
	if err == nil {
		response.JSON(c, userInfo, nil)
		return
	}
	// 重新登录
	response.JSON(c, nil, errorx.New("", 200000))
}

// Logout 登出
func (i IndexController) Logout(c *gin.Context) {
	err := login.Login{}.Logout(c)
	response.JSON(c, nil, err)
}
