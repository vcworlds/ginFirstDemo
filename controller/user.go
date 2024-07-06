package controller

import (
	"gin-Vue/response"
	"gin-Vue/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx *gin.Context) {
	var registerService = service.RegisterService{}
	err := ctx.ShouldBind(&registerService)
	if err != nil {
		response.Error(ctx, "获取数据失败")
		return
	}
	res := registerService.Register()
	if res.Status != 200 {
		response.Response(ctx, http.StatusUnprocessableEntity, res.Status, res.Data, res.Error)
		return
	}
	response.Response(ctx, http.StatusOK, res.Status, res.Data, res.Msg)
}

func Login(ctx *gin.Context) {
	var LoginService = service.LoginService{}
	err := ctx.ShouldBind(&LoginService)
	if err != nil {
		response.Error(ctx, "获取数据失败")
		return
	}
	res := LoginService.Login()
	if res.Status != 200 {
		response.Response(ctx, http.StatusUnprocessableEntity, res.Status, res.Data, res.Error)
		return
	}
	response.Success(ctx, res.Data, res.Msg)
}
