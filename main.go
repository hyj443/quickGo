package main

import (
	"quickGo/conf"
	"quickGo/server"
)

func main() {
	// 初始化各种配置
	conf.Init()

	// 路由装载
	r := server.NewRouter()

	// 运行服务，在3000端口
	r.Run(":3000")
}
