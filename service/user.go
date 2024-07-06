package service

import (
	"gin-Vue/dao"
	"gin-Vue/models"
	"gin-Vue/pkg/e"
	"gin-Vue/serialize"
	"net/http"
)

type RegisterService struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
}

func (service *RegisterService) Register() serialize.Response {
	print(service.Phone)
	if len(service.Phone) < 11 {
		code := e.Error
		return serialize.Response{
			Status: code,
			Msg:    e.GetMgsg(code),
			Data:   nil,
			Error:  "手机号格式不正确",
		}
	}
	if len(service.Password) < 6 {
		code := e.Error
		return serialize.Response{
			Status: code,
			Msg:    e.GetMgsg(code),
			Data:   nil,
			Error:  "密码长度不能少于6",
		}
	}
	exit := dao.PhoneIsExit(service.Phone)
	if exit {
		return serialize.Response{
			Status: http.StatusOK,
			Msg:    "手机号已被注册",
			Data:   nil,
			Error:  "手机号已被注册",
		}
	}
	var userinfo = &models.User{
		Name:     service.Name,
		Password: service.Password,
		Phone:    service.Phone,
	}
	dao.DB.Create(&userinfo)
	code := e.Success
	return serialize.Response{
		Status: code,
		Msg:    "注册成功",
		Data:   userinfo,
		Error:  "",
	}

}
