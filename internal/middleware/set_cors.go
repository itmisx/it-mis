package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetCors 设置跨域
//
// 跨域定义，请求分两类：简单请求和非简单请求
// 一、简单请求定义
//（1)请求方法是以下三种方法之一：
//  - HEAD
//  - GET
//  - POST
//（2）HTTP的头信息不超出以下几种字段：
//  - Accept
//  - Accept-Language
//  - Content-Language
//  - Last-Event-ID
//  - Content-Type：只限于三个值application/x-www-form-urlencoded、multipart/form-data、text/plain
// 其他请求，浏览器会发出预检请求Preflighted
func SetCors() gin.HandlerFunc {
	conf := cors.DefaultConfig()
	conf.AllowOriginFunc = func(origin string) bool {
		// TODO config
		return true
	}
	conf.AddAllowMethods("GET", "POST", "DELETE", "PUT", "PATCH")
	conf.AddAllowHeaders("CSRFToken")
	conf.AllowCredentials = true
	return cors.New(conf)
}
