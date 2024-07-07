package controller

import (
	"gin-Vue/dao"
	"gin-Vue/models"
	"gin-Vue/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
}

func NewCategoryController() ICategoryController {
	return CategoryController{}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var category models.Category
	err := ctx.ShouldBind(&category)
	if err != nil {
		response.Error(ctx, "获取数据失败")
		return
	}
	if category.Name == "" {
		response.Error(ctx, "分类名称必填")
		return
	}

	dao.DB.Create(&category)
	if category.ID == 0 {
		response.Error(ctx, "该分类已被创建")
		return
	}
	response.Success(ctx, category, "创建成功")

}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 尝试将URL参数中的id转换为整数
	categoryId, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		// 如果转换失败，返回错误信息
		response.Error(ctx, "获取分类id失败，请刷新")
		return
	}

	// 直接删除指定ID的分类记录
	if err := dao.DB.Debug().Delete(&models.Category{}, categoryId).Error; err != nil {
		// 如果删除失败，返回错误信息
		response.Error(ctx, "删除失败")
		return
	}
	// 如果删除成功，返回成功信息
	response.Success(ctx, nil, "删除成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var category models.Category
	err := ctx.ShouldBind(&category)
	if err != nil {
		response.Error(ctx, "获取数据失败")
		return
	}
	if category.Name == "" {
		response.Error(ctx, "更新名称必填")
		return
	}
	categoryId, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		response.Error(ctx, "获取分类id失败")
		return
	}
	var updateCategory models.Category
	dao.DB.Take(&updateCategory, categoryId).Update("name", category.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "更新成功")

}

func (c CategoryController) Show(ctx *gin.Context) {
	categoryID, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		response.Error(ctx, "获取分类id失败请刷新")
		return
	}
	var category models.Category
	dao.DB.Take(&category, categoryID)
	response.Success(ctx, gin.H{"category": category}, "获取成功")
}
