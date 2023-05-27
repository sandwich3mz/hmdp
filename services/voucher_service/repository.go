package voucher_service

import (
	"context"
	"hmdp/model"
)

type Repository interface {
	QueryVoucherOfShop(ctx context.Context, shopId uint64) (res []*model.Voucher)
	AddSeckillVoucher(ctx context.Context, Voucher model.Voucher) bool
}
