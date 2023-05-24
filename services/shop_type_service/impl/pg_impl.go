package impl

import (
	"context"
	"hmdp/ent"
	"hmdp/ent/shoptype"
	"hmdp/services/shop_type_service"
)

type pgImpl struct {
	dbClient *ent.Client
}

func NewPgImpl(dbClient *ent.Client) shop_type_service.Repository {
	return &pgImpl{
		dbClient: dbClient,
	}
}

// QueryTypeList 查询商店分类
func (p *pgImpl) QueryTypeList(ctx context.Context) (res []*ent.ShopType, err error) {
	res, err = p.dbClient.ShopType.
		Query().
		Order(ent.Asc(shoptype.FieldSort)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
