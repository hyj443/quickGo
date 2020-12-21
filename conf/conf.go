package conf

import (
	"os"
	"quickGo/model"
	"quickGo/cache"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var SignKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Init 初始化各种配置项
func Init() {
	// 从.env文件读取并设置环境变量
	godotenv.Load()

	if os.Getenv("GIN_MODE") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	// 环境变量设置为use时，会启动数据库连接，设置为notuse不会打开连接
	if os.Getenv("RIM") == "use" {
		model.DataBase(os.Getenv("MYSQL_DSN"))
		cache.Redis()
	}

}
