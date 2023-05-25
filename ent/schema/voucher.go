package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Voucher struct {
	ent.Schema
}

func (Voucher) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("主键"),
		field.Uint64("shop_id").Comment("商铺id").StructTag(`json:"shopId, omitempty"`),
		field.String("title").Comment("代金券标题"),
		field.String("sub_title").Comment("副标题").StructTag(`json:"subTitle,omitempty"`),
		field.String("rules").Comment("使用规则"),
		field.Uint64("pay_value").Comment("支付金额").StructTag(`json:"payValue,omitempty"`),
		field.Int64("actual_value").Comment("抵扣金额").StructTag(`json:"actualValue,omitempty"`),
		field.Int8("type").Comment("优惠券类型"),
		field.Int8("status").Comment("优惠券类型"),
		field.Time("create_time").Comment("创建时间"),
		field.Time("update_time").Comment("更新时间"),
	}
}

func (Voucher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("getMore", SeckillVoucher.Type),
	}
}

func (Voucher) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tb_voucher"},
	}
}
