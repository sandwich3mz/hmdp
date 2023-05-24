// Code generated by ent, DO NOT EDIT.

package blog

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the blog type in the database.
	Label = "blog"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldShopId holds the string denoting the shopid field in the database.
	FieldShopId = "shop_id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldImages holds the string denoting the images field in the database.
	FieldImages = "images"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldLiked holds the string denoting the liked field in the database.
	FieldLiked = "liked"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// FieldCreateTime holds the string denoting the createtime field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the updatetime field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the blog in the database.
	Table = "tb_blog"
)

// Columns holds all SQL columns for blog fields.
var Columns = []string{
	FieldID,
	FieldShopId,
	FieldUserId,
	FieldTitle,
	FieldImages,
	FieldContent,
	FieldLiked,
	FieldComments,
	FieldCreateTime,
	FieldUpdateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int64) error
)

// OrderOption defines the ordering options for the Blog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByShopId orders the results by the shopId field.
func ByShopId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShopId, opts...).ToFunc()
}

// ByUserId orders the results by the userId field.
func ByUserId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserId, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByImages orders the results by the images field.
func ByImages(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImages, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByLiked orders the results by the liked field.
func ByLiked(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLiked, opts...).ToFunc()
}

// ByComments orders the results by the comments field.
func ByComments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComments, opts...).ToFunc()
}

// ByCreateTime orders the results by the createTime field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the updateTime field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}