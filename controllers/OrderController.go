package controllers

import (
	"github.com/gin-gonic/gin"
	"hmdp/common"
	"hmdp/services"
	"strconv"
)

func SeckillVoucher(ctx *gin.Context) {
	paramId := ctx.Param("id")
	voucherId, _ := strconv.Atoi(paramId)
	orderId, err := services.VoucherRepository.SeckillVoucher(ctx, uint64(voucherId))
	if err != nil {
		common.ResponseErrorWithMsg(ctx, 200, err.Error())
		return
	}
	common.ResponseSuccess(ctx, orderId)
}
