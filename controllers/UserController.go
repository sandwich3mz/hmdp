package controllers

import (
	"github.com/gin-gonic/gin"
	"hmdp/common"
	"hmdp/dto"
	"hmdp/global"
	"hmdp/services"
	_ "hmdp/services/user_service/impl"
	"regexp"
)

// SendCode 发送手机验证码
func SendCode(ctx *gin.Context) {
	phone := ctx.Query("phone")

	if ok, _ := regexp.MatchString(`^1([38][0-9]|4[579]|5[0-3,5-9]|6[6]|7[0135678]|9[89])\d{8}$`, phone); !ok { // 手机号格式不正确
		common.ResponseErrorWithMsg(ctx, 200, "手机号格式不正确")
	} else {
		services.UserRepository.SendCode(ctx, phone)
		common.ResponseSuccess(ctx, nil)
	}
}

// Login 登录
func Login(ctx *gin.Context) {
	var loginForm dto.LoginFormDTO
	err := ctx.Bind(&loginForm)
	if err != nil {
		common.ResponseErrorWithMsg(ctx, 200, "提交数据异常")
		return
	}
	res := services.UserRepository.Login(ctx, loginForm)
	common.ResponseSuccess(ctx, res)
}

// Me 我的
func Me(ctx *gin.Context) {
	common.ResponseSuccess(ctx, global.UserDTO)
}
