
1. 还没创建数据库quickGo就去连接它 Error 1049: Unknown database 'quickGo'
2. cannot find module providing package github.com/go-redis/redis: github.com/go-redis/redis@v0.2.0: reading https://mirrors.aliyun.com/goproxy/github.com/go-redis/redis/@v/v0.2.0.zip: 404 Not Found
   1. 镜像代理找不到这个包，就会去源头找。当代理返回 404 时就去源头抓取。
   2. 将 GOPROXY 设置为 https://goproxy.cn,direct
   3. go env -w GOPROXY=https://goproxy.cn,direct
3. gin的engine Use的中间件，是gin.HandlerFunc 类型，不是它的指针类型
4. 我return一个东西，但函数定义了忘了写返回的类型，就报错：too many arguments to return   have (gin.HandlerFunc) want ()
5. Unix() 返回int64 一般用于生成时间戳
6. 使用标准包可以出现代码提示，但是使用自己的包或者第三方库无法出现代码提示，你可以查看一下你的配置项。 "go.inferGopath": true,