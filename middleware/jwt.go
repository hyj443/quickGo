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

// JwtRequired 中间件 
// 因为我把token返回给客户端，我希望下次客户端请求时要在header中传token
// 这个中间件从header获取token，看看是否传了，将它拆分开来，验证格式是否正确，然后解析它
// 得到token结构体，将Raw字符串挂载到context上，并将对应的用户模型存到context上
func JwtRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Authorization这个header获取token
		userToken := c.Request.Header.Get("Authorization")

		// 如果token没传
		if userToken == "" {
			c.JSON(http.StatusOK, serializer.Response{
				Code:    serializer.UserNotPermissionError,
				Message: "token不能为传空",
			}.Format())
			// 必须要带token才能继续执行后续的handler
			c.Abort()
			return
		}

		// 对token的格式进行检查
		list := strings.Split(userToken, " ")
		if len(list) != 2 || list[0] == "Bearer" {
			c.JSON(http.StatusOK, serializer.Response{
				Code:    serializer.UserNotPermissionError,
				Message: "token令牌的格式不正确",
			}.Format())
			c.Abort()
			return
		}

		// 解析用户传过来的token字符串 
		token, err := jwt.ParseWithClaims(list[1], &auth.Jwt{}, func(token *jwt.Token) (interface{}, error) { return conf.SignKey, nil })
		// 如果解析出错，或拿到token结构体但过期了
		if err != nil || token.Valid != true {
			c.JSON(http.StatusOK, serializer.Response{
				Code:    serializer.UserNotPermissionError,
				Message: "token错误",
			}.Format())
			c.Abort()
			return
		}

		// 判断token是否已经注销，已注销名单集合保存在redis，jwt:banned集合里，判断token raw字符串是否在这个集合里
		if res, _ := cache.RedisClient.SIsMember("jwt:banned", token.Raw).Result(); res {
			c.JSON(http.StatusOK, serializer.Response{
				Code:    serializer.UserNotPermissionError,
				Message: "这个用户的token已经注销",
			}.Format())
			c.Abort()
			return
		}

		// 已经校验好的代表用户的token串挂载到context，用于注销
		c.Set("token", token.Raw)

		// 将当前访问的用户模型，挂载到context的user字段
		if jwtStruct, ok := token.Claims.(*auth.Jwt); ok {
			c.Set("user", &jwtStruct.Data)
		}
	}
}
