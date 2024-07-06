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
