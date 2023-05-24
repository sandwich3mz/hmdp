package services

import (
	"hmdp/db"
	"hmdp/services/blog_service"
	blogImpl "hmdp/services/blog_service/impl"
	"hmdp/services/shop_service"
	shopImpl "hmdp/services/shop_service/impl"
	"hmdp/services/shop_type_service"
	shopTypeImpl "hmdp/services/shop_type_service/impl"
	"hmdp/services/user_service"
	userImpl "hmdp/services/user_service/impl"
	"hmdp/services/voucher_service"
	voucherImpl "hmdp/services/voucher_service/impl"
)

var (
	BlogRepository     blog_service.Repository
	ShopTypeRepository shop_type_service.Repository
	UserRepository     user_service.Repository
	ShopRepository     shop_service.Repository
	VoucherRepository  voucher_service.Repository
)

func Init() {
	dbClient := db.InitializeDB()
	BlogRepository = blogImpl.NewPgImpl(dbClient)
	ShopTypeRepository = shopTypeImpl.NewPgImpl(dbClient)
	UserRepository = userImpl.NewPgImpl(dbClient)
	ShopRepository = shopImpl.NewPgImpl(dbClient)
	VoucherRepository = voucherImpl.NewPgImpl(dbClient)
}
