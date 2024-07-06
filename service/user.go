package service

import (
	"gin-Vue/dao"
	"gin-Vue/models"
	"gin-Vue/pkg/e"
	"gin-Vue/serialize"
	"gin-Vue/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterService struct {
	Name       string `json:"name" form:"name"`
	Password   string `json:"password" form:"password"`
	RePassword string `json:"re_password" form:"re_password"`
	Phone      string `json:"phone" form:"phone"`
}
type LoginService struct {
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
}

func (service *RegisterService) Register() serialize.Response {
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
	if service.RePassword != service.Password {
		code := e.Error
		return serialize.Response{
			Status: code,
			Msg:    e.GetMgsg(code),
			Data:   nil,
			Error:  "两次输入密码不一致",
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

func (service *LoginService) Login() serialize.Response {
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

	user, str := dao.GetUserInfo(service.Phone, service.Password)
	if user == nil {
		return serialize.Response{
			Status: 420,
			Msg:    "获取用户失败",
			Data:   nil,
			Error:  str,
		}
	}
	tokenString := utils.ReleaseToken(user)
	return serialize.Response{
		Status: 200,
		Msg:    str,
		Data:   gin.H{"userInfo": user, "token": tokenString},
		Error:  "",
	}
}
