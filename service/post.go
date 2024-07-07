package service

import (
	"gin-Vue/dao"
	"gin-Vue/models"
	"gin-Vue/pkg/e"
	"gin-Vue/serialize"
	"net/http"
	"time"
)

type PostService struct {
	Id         uint   `json:"id" form:"id"`
	Content    string `json:"content" form:"content"`
	Title      string `json:"title" form:"title"`
	CategoryId uint   `json:"categoryId" form:"categoryId"`
	PostImg    string `json:"postImg" form:"postImg"`
}

func (service PostService) CreatePost(user models.User) serialize.Response {
	var postController = &models.Post{
		Title:      service.Title,
		Content:    service.Content,
		UserId:     user.ID,
		CategoryId: service.CategoryId,
		PostImg:    service.PostImg,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}
	if err := dao.DB.Create(&postController).Error; err != nil {
		code := http.StatusUnprocessableEntity
		return serialize.Response{
			Status: code,
			Msg:    e.GetMgsg(code),
			Data:   nil,
			Error:  "数据创建失败",
		}
	}
	return serialize.Response{
		Status: http.StatusOK,
		Msg:    "创建成功",
		Data:   postController,
		Error:  "",
	}
}

func (service PostService) UpdatePost(user models.User, postId string) serialize.Response {
	updatePost := models.Post{
		Content:    service.Content,
		Title:      service.Title,
		CategoryId: service.CategoryId,
		PostImg:    service.PostImg,
		UpdateAt:   time.Now(),
	}
	str, post, exits := dao.UpdatePost(postId, user.ID, updatePost)
	if !exits {
		code := http.StatusUnprocessableEntity
		return serialize.Response{
			Status: code,
			Msg:    "",
			Data:   nil,
			Error:  str,
		}
	}
	return serialize.Response{
		Status: http.StatusOK,
		Msg:    "更新成功",
		Data:   post,
		Error:  "",
	}
}
