package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("主键").Unique().Positive(),
		field.String("phone").Comment("手机号"),
		field.String("password").Comment("密码"),
		field.String("nick_name").Comment("昵称"),
		field.String("icon").Comment("头像"),
		field.Time("create_time").Comment("创建时间"),
		field.Time("update_time").Comment("更新时间"),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tb_user"},
	}
}
