package routes

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"hmdp/controllers"
	"log"
	"net/http"
)

var (
	g errgroup.Group
)

func setupRouter() http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())
	blog := router.Group("/blog")
	{
		blog.GET("/hot", controllers.QueryHotBlog)
	}
	shopType := router.Group("/shop-type")
	{
		shopType.GET("/list", controllers.QueryTypeList)
	}
	user := router.Group("/user")
	{
		user.POST("/code", controllers.SendCode)
		user.POST("/login", controllers.Login)
		user.GET("/me", controllers.Me)
	}
	shop := router.Group("/shop")
	{
		shop.GET("/:id", controllers.QueryShopById)
		shop.GET("/of/type", controllers.QueryShopByType)
	}
	voucher := router.Group("/voucher")
	{
		voucher.GET("/list/:shopId", controllers.QueryVoucherOfShop)
		voucher.POST("/", controllers.AddVoucher)
		voucher.POST("/seckill", controllers.AddSeckillVoucher)
	}
	voucherOrder := router.Group("/voucher-order")
	{
		voucherOrder.POST("/seckill/:id", controllers.SeckillVoucher)
	}

	return router
}

// RunServer 启动服务器
func RunServer() {
	server01 := &http.Server{
		Addr:    ":8081",
		Handler: setupRouter(),
	}
	server02 := &http.Server{
		Addr:    ":8082",
		Handler: setupRouter(),
	}
	g.Go(func() error {
		log.Printf("启动8081端口")
		return server01.ListenAndServe()
	})
	g.Go(func() error {
		log.Printf("启动8082端口")
		return server02.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
