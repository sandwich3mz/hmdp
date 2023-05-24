package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type ShopType struct {
	ent.Schema
}

// Fields of the ShopType.
func (ShopType) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("主键").Unique().Positive(),
		field.String("name").Comment("类型名称"),
		field.String("icon").Comment("图标"),
		field.Int("sort").Comment("顺序"),
		field.Time("createTime").Comment("创建时间"),
		field.Time("updateTime").Comment("更新时间"),
	}
}

// Edges of the ShopType.
func (ShopType) Edges() []ent.Edge {
	return nil
}

func (ShopType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tb_shop_type"},
	}
}
