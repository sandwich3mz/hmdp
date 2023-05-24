package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Shop struct {
	ent.Schema
}

// Fields of the ShopType.
func (Shop) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("主键").Unique().Positive(),
		field.String("name").Comment("商铺名称"),
		field.Uint64("type_id").Comment("商铺类型的id"),
		field.String("images").Comment("商铺图片"),
		field.String("area").Comment("商圈"),
		field.String("address").Comment("地址"),
		field.Float("x").Comment("经度"),
		field.Float("y").Comment("维度"),
		field.Uint64("avg_price").Comment("均价"),
		field.Uint64("sold").Comment("销量"),
		field.Uint64("comments").Comment("评论数量"),
		field.Int8("score").Comment("评分，1~5分，乘10保存，避免小数"),
		field.String("open_hours").Comment("更新时间"),
		field.Time("createTime").Comment("创建时间"),
		field.Time("updateTime").Comment("更新时间"),
	}
}

// Edges of the ShopType.
func (Shop) Edges() []ent.Edge {
	return nil
}

func (Shop) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tb_shop"},
	}
}
