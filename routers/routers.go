package routers

import (
	"gin-Vue/controller"
	"gin-Vue/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/userinfo", middleware.AuthMiddleware(), controller.Userinfo)
	return r
}
