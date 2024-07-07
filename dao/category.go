package dao

import "gin-Vue/models"

func CategoryIsExits(categoryId uint) bool {
	var category models.Category
	if err := DB.Model(&models.Category{}).Where("id = ?", categoryId).First(category).Error; err != nil {
		return false
	}
	return true
}
