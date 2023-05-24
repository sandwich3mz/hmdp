package controllers

import (
	"github.com/gin-gonic/gin"
	"hmdp/common"
	"hmdp/services"
)

// QueryTypeList 查询商店分类
func QueryTypeList(ctx *gin.Context) {
	list, err := services.ShopTypeRepository.QueryTypeList(ctx)
	if err != nil {
		common.ResponseErrorWithMsg(ctx, common.CodeServerError, err.Error())
		return
	}
	//ctx.JSON(http.StatusOK, gin.H{
	//	"success": "true",
	//	"data":    list,
	//})
	common.ResponseSuccess(ctx, list)
}
