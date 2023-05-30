package impl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"hmdp/ent"
	"hmdp/ent/seckillvoucher"
	"hmdp/ent/voucher"
	"hmdp/ent/voucherorder"
	"hmdp/global"
	"hmdp/model"
	"hmdp/services/voucher_service"
	"hmdp/tools"
	"log"
	"sync"
	"time"
)

type pgImpl struct {
	dbClient *ent.Client
}

var lock sync.Mutex

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

func (p *pgImpl) SeckillVoucher(ctx context.Context, voucherId uint64) (int64, error) {
	// 查询优惠券
	targetVoucher, err := p.dbClient.SeckillVoucher.Query().
		Where(seckillvoucher.VoucherIDEQ(voucherId)).
		Only(context.Background())
	if err != nil {
		log.Printf("Failed to find the targetVoucher: %v", err)
		return -1, err
	}
	// 判断秒杀是否开始
	if !time.Now().After(targetVoucher.BeginTime) {
		log.Printf("The Lightning Deal hasn't started yet")
		err = errors.New("the Lightning Deal hasn't started yet")
		return -1, err
	}

	// 判断秒杀是否结束
	if !time.Now().Before(targetVoucher.EndTime) {
		log.Printf("The Lightning Deal is over")
		err = errors.New("the Lightning Deal is over")
		return -1, err
	}

	// 判断库存是否充足
	if targetVoucher.Stock <= 0 {
		log.Printf("There is not enough inventory")
		err = errors.New("there is not enough inventory")
		return -1, err
	}

	lock.Lock()
	defer lock.Unlock()

	orderId, err := p.createVoucherOrder(ctx, voucherId)
	return orderId, err
}

func (p *pgImpl) createVoucherOrder(ctx context.Context, voucherId uint64) (int64, error) {
	// 一人一单
	userId := global.UserDTO.ID
	tx, err := p.dbClient.Tx(context.Background())
	if err != nil {
		return -1, err
	}
	// 查询是否下过单
	count, err := p.dbClient.VoucherOrder.Query().
		Where(voucherorder.UserIDEQ(userId)).
		Count(context.Background())
	if err != nil {
		return -1, err
	} else if count > 0 {
		err = errors.New("You've already placed an order\n")
		log.Printf("You've already placed an order\n")
		return -1, err
	}

	// 扣减库存
	isSuccess, err := tx.SeckillVoucher.Update().
		Where(seckillvoucher.And(
			seckillvoucher.VoucherID(voucherId),
			seckillvoucher.StockGT(0), // 确保库存 > 0
		)).
		AddStock(-1).
		Save(context.Background())
	if err != nil || isSuccess < 1 {
		log.Printf("Failed to update the stock: %v", err)
		return -1, err
	}

	// 创建订单
	orderId := tools.NextId("order")
	_, err = tx.VoucherOrder.Create().
		SetID(orderId).
		SetUserID(userId).
		SetVoucherID(voucherId).
		Save(context.Background())
	if err != nil {
		log.Printf("Failed to create vouche order: %v", err)
		return -1, err
	}
	return orderId, nil
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
