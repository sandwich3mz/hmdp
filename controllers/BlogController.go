package controllers

import (
	"github.com/gin-gonic/gin"
	"hmdp/common"
	"hmdp/services"
	"strconv"
)

// QueryHotBlog 查找最热的博客
func QueryHotBlog(ctx *gin.Context) {
	current, _ := strconv.Atoi(ctx.Query("current"))

	blog, err := services.BlogRepository.QueryHotBlog(ctx, current)
	if err != nil {
		common.ResponseErrorWithMsg(ctx, common.CodeServerError, err.Error())
		return
	}
	//ctx.JSON(http.StatusOK, gin.H{
	//	"success": "true",
	//	"data":    blog,
	//})
	common.ResponseSuccess(ctx, blog)
}
