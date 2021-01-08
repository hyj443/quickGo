package conf

import (
	"fmt"
	"os"
	"quickGo/cache"
	"quickGo/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var SignKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Init 初始化各种配置项
func Init() {
	// 从.env文件读取并设置环境变量
	godotenv.Load(".env")

	if os.Getenv("GIN_MODE") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	// 环境变量RIM为use时，会启动数据库连接，设置为notuse不会打开连接
	if os.Getenv("RIM") == "use" {
		model.DataBase(os.Getenv("MYSQL_DSN"))
		cache.Redis()
	}

	if gin.Mode() == gin.ReleaseMode {
		go func() {
			time.Sleep(time.Second)
			fmt.Println("服务器已经成功启动，当前是Release模式")
		}()
	}
}
