// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TbBlogColumns holds the columns for the "tb_blog" table.
	TbBlogColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "shop_id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "title", Type: field.TypeString},
		{Name: "images", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "liked", Type: field.TypeInt},
		{Name: "comments", Type: field.TypeInt},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// TbBlogTable holds the schema information for the "tb_blog" table.
	TbBlogTable = &schema.Table{
		Name:       "tb_blog",
		Columns:    TbBlogColumns,
		PrimaryKey: []*schema.Column{TbBlogColumns[0]},
	}
	// TbSeckillVoucherColumns holds the columns for the "tb_seckill_voucher" table.
	TbSeckillVoucherColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "stock", Type: field.TypeUint64},
		{Name: "begin_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "voucher_id", Type: field.TypeUint64},
	}
	// TbSeckillVoucherTable holds the schema information for the "tb_seckill_voucher" table.
	TbSeckillVoucherTable = &schema.Table{
		Name:       "tb_seckill_voucher",
		Columns:    TbSeckillVoucherColumns,
		PrimaryKey: []*schema.Column{TbSeckillVoucherColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tb_seckill_voucher_tb_voucher_getMore",
				Columns:    []*schema.Column{TbSeckillVoucherColumns[6]},
				RefColumns: []*schema.Column{TbVoucherColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TbShopColumns holds the columns for the "tb_shop" table.
	TbShopColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type_id", Type: field.TypeUint64},
		{Name: "images", Type: field.TypeString},
		{Name: "area", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "x", Type: field.TypeFloat64},
		{Name: "y", Type: field.TypeFloat64},
		{Name: "avg_price", Type: field.TypeUint64},
		{Name: "sold", Type: field.TypeUint64},
		{Name: "comments", Type: field.TypeUint64},
		{Name: "score", Type: field.TypeInt8},
		{Name: "open_hours", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// TbShopTable holds the schema information for the "tb_shop" table.
	TbShopTable = &schema.Table{
		Name:       "tb_shop",
		Columns:    TbShopColumns,
		PrimaryKey: []*schema.Column{TbShopColumns[0]},
	}
	// TbShopTypeColumns holds the columns for the "tb_shop_type" table.
	TbShopTypeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "icon", Type: field.TypeString},
		{Name: "sort", Type: field.TypeInt},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// TbShopTypeTable holds the schema information for the "tb_shop_type" table.
	TbShopTypeTable = &schema.Table{
		Name:       "tb_shop_type",
		Columns:    TbShopTypeColumns,
		PrimaryKey: []*schema.Column{TbShopTypeColumns[0]},
	}
	// TbUserColumns holds the columns for the "tb_user" table.
	TbUserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "phone", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "nick_name", Type: field.TypeString},
		{Name: "icon", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// TbUserTable holds the schema information for the "tb_user" table.
	TbUserTable = &schema.Table{
		Name:       "tb_user",
		Columns:    TbUserColumns,
		PrimaryKey: []*schema.Column{TbUserColumns[0]},
	}
	// TbVoucherColumns holds the columns for the "tb_voucher" table.
	TbVoucherColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "shop_id", Type: field.TypeUint64},
		{Name: "title", Type: field.TypeString},
		{Name: "sub_title", Type: field.TypeString},
		{Name: "rules", Type: field.TypeString},
		{Name: "pay_value", Type: field.TypeUint64},
		{Name: "actual_value", Type: field.TypeInt64},
		{Name: "type", Type: field.TypeInt8},
		{Name: "status", Type: field.TypeInt8},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// TbVoucherTable holds the schema information for the "tb_voucher" table.
	TbVoucherTable = &schema.Table{
		Name:       "tb_voucher",
		Columns:    TbVoucherColumns,
		PrimaryKey: []*schema.Column{TbVoucherColumns[0]},
	}
	// TbVoucherOrderColumns holds the columns for the "tb_voucher_order" table.
	TbVoucherOrderColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeUint64},
		{Name: "voucher_id", Type: field.TypeUint64},
		{Name: "pay_type", Type: field.TypeInt8, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Nullable: true},
		{Name: "pay_time", Type: field.TypeTime, Nullable: true},
		{Name: "use_time", Type: field.TypeTime, Nullable: true},
		{Name: "refund_time", Type: field.TypeTime, Nullable: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// TbVoucherOrderTable holds the schema information for the "tb_voucher_order" table.
	TbVoucherOrderTable = &schema.Table{
		Name:       "tb_voucher_order",
		Columns:    TbVoucherOrderColumns,
		PrimaryKey: []*schema.Column{TbVoucherOrderColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TbBlogTable,
		TbSeckillVoucherTable,
		TbShopTable,
		TbShopTypeTable,
		TbUserTable,
		TbVoucherTable,
		TbVoucherOrderTable,
	}
)

func init() {
	TbBlogTable.Annotation = &entsql.Annotation{
		Table: "tb_blog",
	}
	TbSeckillVoucherTable.ForeignKeys[0].RefTable = TbVoucherTable
	TbSeckillVoucherTable.Annotation = &entsql.Annotation{
		Table: "tb_seckill_voucher",
	}
	TbShopTable.Annotation = &entsql.Annotation{
		Table: "tb_shop",
	}
	TbShopTypeTable.Annotation = &entsql.Annotation{
		Table: "tb_shop_type",
	}
	TbUserTable.Annotation = &entsql.Annotation{
		Table: "tb_user",
	}
	TbVoucherTable.Annotation = &entsql.Annotation{
		Table: "tb_voucher",
	}
	TbVoucherOrderTable.Annotation = &entsql.Annotation{
		Table: "tb_voucher_order",
	}
}
