package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// InitSession 初始话session
func InitSession() gin.HandlerFunc {
	sessionType := "cookie"
	switch sessionType {
	case "cookie":
		store := cookie.NewStore([]byte("secret"))
		return sessions.Sessions("session", store)
	case "redis":
		store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
		return sessions.Sessions("session", store)
	}
	return nil
}
