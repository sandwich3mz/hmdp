package routes

import (
	"github.com/gin-gonic/gin"
	"hmdp/config"
	"hmdp/controllers"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

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
	}

	return router
}

// RunServer 启动服务器
func RunServer() {
	ginServer := setupRouter()
	//ginServer.LoadHTMLGlob("templates/*")
	//ginServer.Use(middleware.LoggerMiddleware())
	ginServer.Run(":" + config.Conf.App.Port)
}