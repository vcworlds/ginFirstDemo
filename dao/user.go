package dao

import (
	"gin-Vue/models"
)

func PhoneIsExit(phone string) (exit bool) {
	var user models.User
	DB.Model(models.User{}).Where("phone = ?", phone).Take(&user)
	if user.ID == 0 {
		return false
	}
	return true
}

func GetUserInfo(phone, password string) (*models.User, string) {
	var user *models.User
	DB.Model(&models.User{}).Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		return nil, "查询用户失败"
	}
	if user.Password != password {
		return nil, "密码错误"
	}
	return user, "登陆成功"
}
