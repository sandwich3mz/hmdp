package shop_type_service

import (
	"context"
	"hmdp/ent"
)

type Repository interface {
	QueryTypeList(ctx context.Context) (res []*ent.ShopType, err error)
}
