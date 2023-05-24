package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct {
	ent.Schema
}

// Fields of the Blog.
func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("主键").Unique().Positive(),
		field.Int64("shopId").Comment("商户id"),
		field.Int64("userId").Comment("用户id"),
		field.String("title").Comment("标题"),
		field.String("images").Comment("探店的照片"),
		field.String("content").Comment("探店的内容"),
		field.Int("liked").Comment("点赞的数量"),
		field.Int("comments").Comment("评论数量"),
		field.Time("createTime").Comment("创建时间"),
		field.Time("updateTime").Comment("更新时间"),
	}
}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge {
	return nil
}

func (Blog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tb_blog"},
	}
}
