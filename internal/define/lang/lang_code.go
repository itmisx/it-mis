package lang

// LangCode 主要用于错误码
var LangCode = map[string]map[interface{}]interface{}{
	"zh-cn": {
		// 100000 系统错误
		100000: "系统错误",
		// 2***** 认证错误
		// 基本登录
		200000: "请先登录",
		200001: "用户名或密码错误",
		200002: "账号已被禁用",
		// 300000 业务错误
		// 301xxx 用户
		301001: "用户不存在",
	},
}
