package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"hmdp/ent"
	"hmdp/ent/voucher"
	"hmdp/model"
	"hmdp/services/voucher_service"
	"log"
	"time"
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
			eachVoucher.BeginTime = getMore.BeginTime
			eachVoucher.EndTime = getMore.EndTime
		} else {
			eachVoucher.BeginTime = time.Time{}
			forever, _ := time.Parse("2006-01-02", "2099-01-01")
			eachVoucher.EndTime = forever
		}

		res = append(res, &eachVoucher)
	}
	return res
}

func (p *pgImpl) AddSeckillVoucher(ctx context.Context, voucher model.Voucher) bool {
	tx, err := p.dbClient.Tx(ctx)
	if err != nil {
		logrus.Println("starting a transaction: %w", err)
		return false
	}
	res, err := tx.Voucher.Create().
		SetShopID(voucher.ShopId).
		SetActualValue(voucher.ActualValue).
		SetPayValue(voucher.PayValue).
		SetRules(voucher.Rules).
		SetStatus(voucher.Status).
		SetTitle(voucher.Title).
		SetSubTitle(voucher.SubTitle).
		SetType(voucher.Type).
		Save(context.Background())
	if err != nil {

		log.Printf("Failed to set Voucher: %v\n", err)
		err := rollback(tx, fmt.Errorf("failed creating the group: %w", err))
		if err != nil {
			log.Printf("Failed to rollback: %v\n", err)
		}
		return false
	}
	_, err = tx.SeckillVoucher.Create().
		SetVoucherID(res.ID).
		SetStock(voucher.Stock).
		SetCreateTime(res.CreateTime).
		SetUpdateTime(res.CreateTime).
		SetBeginTime(voucher.BeginTime).
		SetEndTime(voucher.EndTime).
		Save(context.Background())
	if err != nil {
		log.Printf("Failed to set skillVoucher: %v\n", err)
		err := rollback(tx, fmt.Errorf("failed creating the group: %w", err))
		if err != nil {
			log.Printf("Failed to rollback: %v\n", err)
		}
		return false
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("Failed to commit transactions: %v\n", err)
		return false
	}
	return true
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
