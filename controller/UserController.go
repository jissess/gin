package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"myGin/model"
	"myGin/response"
	"myGin/service"
)

type UserController struct {
}

func (uc UserController) Register(ctx *gin.Context) {
	username := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	if len(telephone) <= 0 || len(telephone) > 11 {
		response.Failed(ctx, "手机号必须为11位", nil)
		return
	}
	if len(password) < 6 {
		response.Failed(ctx, "密码长度不少于6位", nil)
		return
	}
	if len(username) == 0 {
		username = telephone
	}
	userService := service.UserService{}
	if userService.IsPhoneExists(telephone) {
		response.Failed(ctx, "用户已存在", nil)
		return
	}
	encryptedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Failed(ctx, "加密错误", nil)
		return
	}
	log.Println(username, telephone, password)
	user := model.User{
		UserName:  username,
		Telephone: telephone,
		Password:  string(encryptedPwd),
	}
	_, err = userService.Register(user)
	if err != nil {
		response.Failed(ctx, "注册失败", nil)
		return
	}
	response.Success(ctx, nil, "注册成功")
}
