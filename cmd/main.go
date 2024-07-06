package main

import (
	"gin-Vue/dao"
	"gin-Vue/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 加载配置
	dao.InitDB()
	// 路由加载
	r = routers.NewRouter(r)
	_ = r.Run(":1061")
}
