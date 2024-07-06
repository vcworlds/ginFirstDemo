package routers

import (
	"gin-Vue/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("api/auth/login", controller.Login)
	return r
}
