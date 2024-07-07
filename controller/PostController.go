package controller

import (
	"gin-Vue/models"
	"gin-Vue/response"
	"gin-Vue/service"
	"github.com/gin-gonic/gin"
)

type IPostController interface {
	RestController
}

type PostController struct {
}

func NewPostController() ICategoryController {
	return PostController{}
}

func (p PostController) Create(ctx *gin.Context) {
	var postService service.PostService
	err := ctx.ShouldBind(&postService)
	if err != nil {
		response.Error(ctx, "数据绑定失败")
		return
	}
	user, exits := ctx.Get("user")
	if !exits {
		response.Error(ctx, "token解析失败")
		return
	}
	res := postService.CreatePost(user.(models.User))
	if res.Status != 200 {
		response.Error(ctx, res.Error)
		return
	}
	response.Success(ctx, res.Data, res.Msg)
}

func (p PostController) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p PostController) Update(ctx *gin.Context) {
	var postService service.PostService
	err := ctx.ShouldBind(&postService)
	if err != nil {
		response.Error(ctx, "获取数据失败")
		return
	}
	user, _ := ctx.Get("user")
	postId, exits := ctx.Params.Get("id")
	if !exits {
		response.Error(ctx, "id获取失败")
		return
	}
	res := postService.UpdatePost(user.(models.User), postId)
	if res.Status != 200 {
		response.Error(ctx, res.Error)
		return
	}
	response.Success(ctx, res.Data, res.Msg)
}

func (p PostController) Show(ctx *gin.Context) {
	postId :=
}
