package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type SeckillVoucher struct {
	ent.Schema
}

func (SeckillVoucher) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("主键"),
		field.Uint64("voucher_id").Comment("优惠券id"),
		field.Uint64("stock").Comment("库存"),
		field.Time("begin_time").Comment("生效时间").StructTag(`json:"beginTime,omitempty"`),
		field.Time("end_time").Comment("失效时间").StructTag(`json:"endTime,omitempty"`),
		field.Time("create_time").Comment("创建时间"),
		field.Time("update_time").Comment("更新时间"),
	}
}

func (SeckillVoucher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("getForm", Voucher.Type).
			Ref("getMore").
			Field("voucher_id").
			Unique().
			Required(),
	}
}

func (SeckillVoucher) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "tb_seckill_voucher",
		},
	}
}
