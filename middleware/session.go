package middleware

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// gin-contrib/sessions提供了不同的存储引擎，比如redis cookie，如果是开发成大型项目就用redis，如果不需要，用cookie就行

var store sessions.Store

// SessionRedis 返回session中间件。先创建以redis作为存储引擎的store，再返回一个中间件，这个中间件会
// 创建一个名为gin-session的session对象，挂载到Context
func SessionRedis(secret string) gin.HandlerFunc {
	// 创建基于redis的存储引擎，传入用于加密的密钥
	store, _ = redis.NewStore(10, "tcp", os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PW"), []byte(secret))

	store.Options(sessions.Options{
		MaxAge:   7 * 24 * 60 * 60, //seconds
		Path:     "/",
		HttpOnly: true,
	})

	// 返回session中间件，参数gin-session代表是session的名字
	return sessions.Sessions("gin-session", store)
	// 	func Sessions(name string, store Store) gin.HandlerFunc {
	// 	return func(c *gin.Context) {
	// 		s := &session{name, c.Request, store, nil, false, c.Writer}
	// 		c.Set(DefaultKey, s)
	// 		defer context.Clear(c.Request)
	// 		c.Next()
	// 	}
	// }
}

// SessionCookie 中间件，初始化Session，用cookie作为存储引擎
func SessionCookie(secret string) gin.HandlerFunc {
	// 创建基于cookie的存储引擎，传入了用于加密的密钥
	store = cookie.NewStore([]byte(secret))

	store.Options(sessions.Options{
		MaxAge:   7 * 24 * 60 * 60, //seconds
		Path:     "/",
		HttpOnly: true,
	})

	// 返回session中间件，参数gin-session代表是cookie的名字
	return sessions.Sessions("gin-session", store)
}
