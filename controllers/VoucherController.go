package controllers

import (
	"github.com/gin-gonic/gin"
	"hmdp/common"
	"hmdp/services"
	"strconv"
)

type Voucher struct {
	Id           uint64
	Shop_id      uint64
	Title        string
	Sub_title    string
	Rules        string
	Pay_value    uint64
	Actual_value int64
	Type         int8
	status       int8
}

func AddSeckillVoucher(ctx *gin.Context) {

}

func QueryVoucherOfShop(ctx *gin.Context) {
	shopId, _ := strconv.Atoi(ctx.Param("shopId"))
	res := services.VoucherRepository.QueryVoucherOfShop(ctx, uint64(shopId))
	common.ResponseSuccess(ctx, res)
}
