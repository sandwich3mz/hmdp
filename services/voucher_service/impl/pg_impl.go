package impl

import (
	"context"
	"github.com/sirupsen/logrus"
	"hmdp/ent"
	"hmdp/ent/voucher"
	"hmdp/services/voucher_service"
)

type pgImpl struct {
	dbClient *ent.Client
}

func NewPgImpl(dbClient *ent.Client) voucher_service.Repository {
	return &pgImpl{
		dbClient: dbClient,
	}
}

func (p *pgImpl) QueryVoucherOfShop(ctx context.Context, shopId uint64) (res []*ent.Voucher) {

	//var VoucherOfShop *[]model.Voucher
	res, err := p.dbClient.Debug().Voucher.
		Query().
		//Where(func(s *sql.Selector) {
		//	t := sql.Table(seckillvoucher.Table)
		//	s.LeftJoin(t).On(s.C(voucher.FieldID), t.C(seckillvoucher.FieldID))
		//}).
		Where(voucher.And(
			voucher.ShopIDEQ(shopId),
			voucher.StatusEQ(1),
		)).
		WithGetMore().
		All(context.Background())

	if err != nil {
		logrus.Println(err)
	}
	//if res != nil {
	//	if err := json.Unmarshal(res, VoucherOfShop); err != nil {
	//		log.Printf("Failed to unmarshal VoucherOfShop: %v", err)
	//	}
	//}
	return res
}
