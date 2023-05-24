package voucher_service

import (
	"context"
	"hmdp/ent"
)

type Repository interface {
	QueryVoucherOfShop(ctx context.Context, shopId uint64) (res []*ent.Voucher)
}
