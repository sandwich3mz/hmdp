package voucher_service

import (
	"context"
	"hmdp/model"
)

type Repository interface {
	QueryVoucherOfShop(ctx context.Context, shopId uint64) (res []*model.Voucher)
	AddSeckillVoucher(ctx context.Context, voucher model.Voucher) bool
	SeckillVoucher(ctx context.Context, voucherId uint64) (int64, error)
}
