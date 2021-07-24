package system

import (
	"github.com/gin-gonic/gin"
)

type AccessAuth struct {
}

//  SignIn 登录
func (au AccessAuth) SignIn(c *gin.Context) {
	// 普通登录
	// 开启安全认证的登录
}

// SignOut 登出
func (au AccessAuth) SignOut(c *gin.Context) {
	//销毁token
}

// SignStatus 登录状态
func (au AccessAuth) SignStatus(c *gin.Context) {
	//token验证
	//token+实名认证（邮箱）
}
