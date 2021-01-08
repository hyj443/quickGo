package cache

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

// RedisClient redis client实例
var RedisClient *redis.Client

// Redis 初始化redis连接
func Redis() {
	// REDIS_DB环境变量 解析成十进制的数字 ParseUint类似于ParseInt，只是它是无符号的
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)

	// 连接到redis服务器的client
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"), // host:port address.
		Password: os.Getenv("REDIS_PW"),
		DB:       int(db), // 0-10数据库用哪个
	})

	// 测试是否成功连接
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}

	RedisClient = client
}
