package shop_service

import (
	"context"
	"hmdp/ent"
)

type Repository interface {
	QueryShopById(ctx context.Context, id string) (*ent.Shop, error)
	QueryShopByType(ctx context.Context, typeId int64, current int64, x float64, y float64) (res []*ent.Shop)
}
