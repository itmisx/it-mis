package controller

import (
	"it-mis/internal/logic/system/login"
	"it-mis/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

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
	userInfo, err := login.Login{}.Login(c)
	//  错误返回
	if err != nil {
		response.JSON(c, nil, err)
		return
	}
	// 设置session信息
	login.SetSession(c, "user_id", userInfo.ID)
	// 设置userToken
	login.SetUserToken(c, *userInfo)
	// 返回结果
	response.JSON(c, userInfo, err)
}

// Logout 登出
func (i IndexController) Logout(c *gin.Context) {
	err := login.Login{}.Logout(c)
	response.JSON(c, nil, err)
}
