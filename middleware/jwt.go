package middleware

import (
	"net/http"
	"quickGo/auth"
	"quickGo/cache"
	"quickGo/conf"
	"quickGo/serializer"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// JwtRequired 中间件 需要在header中传token
// 从头部获取token字符串，验证一下是否传了，将它拆分开来，验证格式是否正确，然后调用ParseWithClaims对它进行解析
// 得到token结构体，将Raw字符串挂载到context上，并将对应的用户模型存到context上
func JwtRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken := c.Request.Header.Get("Authorization")

		if userToken == "" {
			c.JSON(http.StatusOK, serializer.Response{
				Code:    serializer.UserNotPermissionError,
				Message: "token不能为传空",
			}.Format())
			c.Abort()
			return
		}

		list := strings.Split(userToken, " ")
		if len(list) != 2 || list[0] == "Bearer" {
			c.JSON(http.StatusOK, serializer.Response{
				Code:    serializer.UserNotPermissionError,
				Message: "token令牌的格式不正确",
			}.Format())
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(list[1], &auth.Jwt{}, func(token *jwt.Token) (interface{}, error) { return conf.SignKey, nil })
		// token过期或非正确处理
		if err != nil || token.Valid != true {
			c.JSON(http.StatusOK, serializer.Response{
				Code:    serializer.UserNotPermissionError,
				Message: "令牌错误！",
			}.Format())
			c.Abort()
			return
		}
		// 判断token是否是在黑名单里面，这个黑名单集合存到redis的集合里，jwt:baned集合里，判断token raw字符串是否在这个集合里
		if res, _ := cache.RedisClient.SIsMember("jwt:baned", token.Raw).Result(); res {
			c.JSON(http.StatusOK, serializer.Response{
				Code:    serializer.UserNotPermissionError,
				Message: "这个用户的token已经注销",
			}.Format())
			c.Abort()
			return
		}
		// token字符串挂载到context，用于注销 添加黑名单
		c.Set("token", token.Raw)

		// 将对应的用户模型的地址存入context
		if jwtStruct, ok := token.Claims.(*auth.Jwt); ok {
			c.Set("user", &jwtStruct.Data)
		}
	}

}
