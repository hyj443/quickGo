package conf

import (
	"os"
	"quickGo/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var SignKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Init 初始化各种配置项
func Init() {
	godotenv.Load()

	if os.Getenv("GIN_MODE") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	if os.Getenv("RIM") == "use" {
		model.DataBase(os.Getenv("MYSQL_DSN"))

		

	}

}
