package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var EmptyData struct{}

type ResponseData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(ctx *gin.Context, c MyCode) {
	rd := &ResponseData{
		Success: false,
		Message: c.Msg(),
		Data:    EmptyData,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code MyCode, errMsg string) {
	rd := &ResponseData{
		Success: false,
		Message: errMsg,
		Data:    EmptyData,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Success: true,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccessWithMsg(ctx *gin.Context, data interface{}, msg string) {
	rd := &ResponseData{
		Success: true,
		Message: msg,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
