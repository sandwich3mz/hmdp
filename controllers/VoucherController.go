package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"hmdp/common"
	"hmdp/db"
	"hmdp/ent"
	"hmdp/model"
	"hmdp/services"
	"log"
	"strconv"
)

func AddVoucher(ctx *gin.Context) {
	var voucher ent.Voucher
	client := db.InitializeDB()
	err := ctx.BindJSON(&voucher)
	if err != nil {
		log.Printf("Failed to get the voucher struct: %v", err)
		return
	}
	_, err = client.Voucher.Create().
		SetShopID(voucher.ShopID).
		SetActualValue(voucher.ActualValue).
		SetPayValue(voucher.PayValue).
		SetRules(voucher.Rules).
		SetStatus(voucher.Status).
		SetTitle(voucher.Title).
		SetSubTitle(voucher.SubTitle).
		SetType(voucher.Type).
		Save(context.Background())
	if err != nil {
		log.Printf("Failed to save the voucher:%v", err)
		common.ResponseErrorWithMsg(ctx, 200, "Failed to save the voucher")
		return
	}
	common.ResponseSuccess(ctx, voucher.ID)
}

func QueryVoucherOfShop(ctx *gin.Context) {
	shopId, _ := strconv.Atoi(ctx.Param("shopId"))
	res := services.VoucherRepository.QueryVoucherOfShop(ctx, uint64(shopId))
	common.ResponseSuccess(ctx, res)
}

func AddSeckillVoucher(ctx *gin.Context) {
	var voucher model.Voucher
	err := ctx.BindJSON(&voucher)
	if err != nil {
		log.Printf("Failed to get the seckillVoucher struct: %v", err)
		return
	}
	isSuccess := services.VoucherRepository.AddSeckillVoucher(ctx, voucher)
	if isSuccess {
		common.ResponseSuccess(ctx, "set success")
	} else {
		common.ResponseSuccess(ctx, "set fail")
	}
}
