package dao

import (
	"gin-Vue/models"
	"time"
)

func UpdatePost(id string, userid uint, UpdatePost models.Post) (string, *models.Post, bool) {
	var post = &models.Post{}
	if err := DB.Where("id = ?", id).First(&post).Error; err != nil {
		return "获取文章失败", nil, false
	}
	//判断文章作者是否user
	if post.UserId != userid {
		return "你没有该权限", nil, false
	}
	// 更新文章内容
	post.Content = UpdatePost.Content
	post.Title = UpdatePost.Title
	post.CategoryId = UpdatePost.CategoryId
	post.PostImg = UpdatePost.PostImg
	post.UpdateAt = time.Now()
	if err := DB.Save(&post).Error; err != nil {
		return "更新文章失败", nil, false
	}
	return "修改成功", post, true
}

func GetPost(id string) (*models.Post, bool, error) {
	var post models.Post
	if err := DB.Take(&post, id).Error; err != nil {
		return nil, false, err
	}
	return &post, true, nil
}

func DeletePost(id string, userId uint) (bool, error, string) {
	var post models.Post
	if err := DB.Take(&post, id).Error; err != nil {
		return false, err, "查询数据失败"
	}
	if post.UserId != userId {
		return false, nil, "你没有该权限"
	}
	if err := DB.Delete(&post).Error; err != nil {
		return false, err, ""
	}
	return true, nil, "删除成功"
}
