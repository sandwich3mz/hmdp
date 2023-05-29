package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

type VoucherOrder struct {
	ent.Schema
}

func (VoucherOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("主键"),
		field.Uint64("user_id").Comment("下单的用户id").StructTag(`json:"userId"`),
		field.Uint64("voucher_id").Comment("购买的优惠券id").StructTag(`json:"voucherId"`),
		field.Int8("pay_type").Comment("支付方式").StructTag(`json:"payType"`).Optional(),
		field.Int8("status").Comment("订单状态").Optional(),
		field.Time("pay_time").Comment("支付时间").StructTag(`json:"payTime"`).Optional(),
		field.Time("use_time").Comment("核销时间").StructTag(`json:"useTime"`).Optional(),
		field.Time("refund_time").Comment("退款时间").StructTag(`json:"refundTime"`).Optional(),
		field.Time("create_time").Comment("创建时间").Default(time.Now()),
		field.Time("update_time").Comment("更新时间").Default(time.Now()),
	}
}

func (VoucherOrder) Edges() []ent.Edge {
	return nil
}

func (VoucherOrder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tb_voucher_order"},
	}
}
