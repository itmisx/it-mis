package lang

// LangTemplate  主要用于消息模板的翻译
var LangTemplate = map[string]map[interface{}]interface{}{
	"zh-cn": {
		"hello{name}": "你好{name}",
	},
	"en-us": {
		"hello{name}": "hello{name}",
	},
}
