package routers

import (
	"gin-Vue/controller"
	"gin-Vue/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/userinfo", middleware.AuthMiddleware(), controller.Userinfo)
	categoryController := controller.NewCategoryController()
	CategoryRouters := r.Group("/category")
	CategoryRouters.POST("", categoryController.Create)
	CategoryRouters.DELETE("/:id", categoryController.Delete)
	CategoryRouters.PUT("/:id", categoryController.Update)
	CategoryRouters.GET("/:id", categoryController.Show)

	postCategoryRouters := r.Group("/post")
	postController := controller.NewPostController()
	postCategoryRouters.Use(middleware.AuthMiddleware())
	postCategoryRouters.POST("", middleware.AuthMiddleware(), postController.Create)
	postCategoryRouters.DELETE("/:id", postController.Delete)
	postCategoryRouters.PUT("/:id", postController.Update)
	postCategoryRouters.GET("/:id", postController.Show)
	return r
}
