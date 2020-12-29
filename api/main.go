package api

import (
	"net/http"
	"quickGo/model"
	"quickGo/serializer"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// writes the given string into the response body.
	c.String(http.StatusOK, "welcome!")
}

func Ping(c *gin.Context) {
	// serializes the given struct as JSON into the response body. It also sets the Content-Type as "application/json".
	c.JSON(http.StatusOK, serializer.Response{
		Message: "Pong",
	}.Format())
}

// GetCurrentUser 获取当前用户
func GetCurrentUser(c *gin.Context) *model.User {
	// 从context上获取
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	// 没有登录，返回nil
	return nil
}
