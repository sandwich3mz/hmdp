package controllers

import (
	"github.com/gin-gonic/gin"
	"hmdp/common"
	"hmdp/services"
	"log"
	"strconv"
)

func QueryShopById(ctx *gin.Context) {
	id := ctx.Param("id")
	shop, err := services.ShopRepository.QueryShopById(ctx, id)
	if err != nil {
		log.Printf("Failed to query: %v", err)
	}
	common.ResponseSuccess(ctx, shop)
}

func QueryShopByType(ctx *gin.Context) {
	typeId, _ := strconv.Atoi(ctx.Query("typeId"))
	current, _ := strconv.Atoi(ctx.Query("current"))
	if current == 0 {
		current = 1
	}
	x, _ := strconv.Atoi(ctx.Query("x"))
	y, _ := strconv.Atoi(ctx.Query("y"))
	res := services.ShopRepository.QueryShopByType(ctx, int64(typeId), int64(current), float64(x), float64(y))
	common.ResponseSuccess(ctx, res)
}
