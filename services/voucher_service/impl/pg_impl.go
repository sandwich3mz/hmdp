package impl

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"hmdp/ent"
	"hmdp/ent/voucher"
	"hmdp/model"
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

func (p *pgImpl) QueryVoucherOfShop(ctx context.Context, shopId uint64) (res []*model.Voucher) {

	allVoucher, err := p.dbClient.Voucher.
		Query().
		Where(voucher.And(
			voucher.StatusEQ(1),
			voucher.ShopIDEQ(shopId),
		)).
		WithGetMore().
		All(context.Background())

	if err != nil {
		logrus.Println(err)
	}

	// 遍历 Voucher 结果集
	for i := range allVoucher {
		// 将 Voucher 结构体转换为 model.Voucher 并添加到结果集中
		b, err := json.Marshal(allVoucher[i])
		if err != nil {
			logrus.Println(err)
			continue
		}
		var eachVoucher model.Voucher
		err = json.Unmarshal(b, &eachVoucher)
		if err != nil {
			logrus.Println(err)
			continue
		}
		// 是否存在 GetMore 结构体
		if len(allVoucher[i].Edges.GetMore) > 0 {
			// 取出 GetMore 结构体中的信息
			getMore := allVoucher[i].Edges.GetMore[0]
			eachVoucher.Stock = getMore.Stock
			eachVoucher.BeginTime = &getMore.BeginTime
			eachVoucher.EndTime = &getMore.EndTime
		} else {
			eachVoucher.BeginTime = nil
			eachVoucher.EndTime = nil
		}

		res = append(res, &eachVoucher)
	}
	return res
}
