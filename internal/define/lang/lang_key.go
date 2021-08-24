package lang

// LangKey 主要用于字段翻译
var LangKey = map[string]map[interface{}]interface{}{
	"zh-cn": {
		"author": "smally84",
		"user": map[string]string{
			"name": "姓名",
			"sex":  "性别",
		},
	},
	"en-us": {
		"user": map[string]string{
			"name": "name",
			"sex":  "sex",
		},
	},
}
