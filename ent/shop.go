// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"hmdp/ent/shop"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Shop is the model entity for the Shop schema.
type Shop struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID int64 `json:"id,omitempty"`
	// 商铺名称
	Name string `json:"name,omitempty"`
	// 商铺类型的id
	TypeID uint64 `json:"type_id,omitempty"`
	// 商铺图片
	Images string `json:"images,omitempty"`
	// 商圈
	Area string `json:"area,omitempty"`
	// 地址
	Address string `json:"address,omitempty"`
	// 经度
	X float64 `json:"x,omitempty"`
	// 维度
	Y float64 `json:"y,omitempty"`
	// 均价
	AvgPrice uint64 `json:"avg_price,omitempty"`
	// 销量
	Sold uint64 `json:"sold,omitempty"`
	// 评论数量
	Comments uint64 `json:"comments,omitempty"`
	// 评分，1~5分，乘10保存，避免小数
	Score int8 `json:"score,omitempty"`
	// 更新时间
	OpenHours string `json:"open_hours,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"createTime,omitempty"`
	// 更新时间
	UpdateTime   time.Time `json:"updateTime,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Shop) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case shop.FieldX, shop.FieldY:
			values[i] = new(sql.NullFloat64)
		case shop.FieldID, shop.FieldTypeID, shop.FieldAvgPrice, shop.FieldSold, shop.FieldComments, shop.FieldScore:
			values[i] = new(sql.NullInt64)
		case shop.FieldName, shop.FieldImages, shop.FieldArea, shop.FieldAddress, shop.FieldOpenHours:
			values[i] = new(sql.NullString)
		case shop.FieldCreateTime, shop.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Shop fields.
func (s *Shop) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shop.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int64(value.Int64)
		case shop.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case shop.FieldTypeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type_id", values[i])
			} else if value.Valid {
				s.TypeID = uint64(value.Int64)
			}
		case shop.FieldImages:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field images", values[i])
			} else if value.Valid {
				s.Images = value.String
			}
		case shop.FieldArea:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field area", values[i])
			} else if value.Valid {
				s.Area = value.String
			}
		case shop.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				s.Address = value.String
			}
		case shop.FieldX:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field x", values[i])
			} else if value.Valid {
				s.X = value.Float64
			}
		case shop.FieldY:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field y", values[i])
			} else if value.Valid {
				s.Y = value.Float64
			}
		case shop.FieldAvgPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field avg_price", values[i])
			} else if value.Valid {
				s.AvgPrice = uint64(value.Int64)
			}
		case shop.FieldSold:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sold", values[i])
			} else if value.Valid {
				s.Sold = uint64(value.Int64)
			}
		case shop.FieldComments:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field comments", values[i])
			} else if value.Valid {
				s.Comments = uint64(value.Int64)
			}
		case shop.FieldScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				s.Score = int8(value.Int64)
			}
		case shop.FieldOpenHours:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field open_hours", values[i])
			} else if value.Valid {
				s.OpenHours = value.String
			}
		case shop.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createTime", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case shop.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updateTime", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Shop.
// This includes values selected through modifiers, order, etc.
func (s *Shop) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Shop.
// Note that you need to call Shop.Unwrap() before calling this method if this Shop
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Shop) Update() *ShopUpdateOne {
	return NewShopClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Shop entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Shop) Unwrap() *Shop {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Shop is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Shop) String() string {
	var builder strings.Builder
	builder.WriteString("Shop(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("type_id=")
	builder.WriteString(fmt.Sprintf("%v", s.TypeID))
	builder.WriteString(", ")
	builder.WriteString("images=")
	builder.WriteString(s.Images)
	builder.WriteString(", ")
	builder.WriteString("area=")
	builder.WriteString(s.Area)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(s.Address)
	builder.WriteString(", ")
	builder.WriteString("x=")
	builder.WriteString(fmt.Sprintf("%v", s.X))
	builder.WriteString(", ")
	builder.WriteString("y=")
	builder.WriteString(fmt.Sprintf("%v", s.Y))
	builder.WriteString(", ")
	builder.WriteString("avg_price=")
	builder.WriteString(fmt.Sprintf("%v", s.AvgPrice))
	builder.WriteString(", ")
	builder.WriteString("sold=")
	builder.WriteString(fmt.Sprintf("%v", s.Sold))
	builder.WriteString(", ")
	builder.WriteString("comments=")
	builder.WriteString(fmt.Sprintf("%v", s.Comments))
	builder.WriteString(", ")
	builder.WriteString("score=")
	builder.WriteString(fmt.Sprintf("%v", s.Score))
	builder.WriteString(", ")
	builder.WriteString("open_hours=")
	builder.WriteString(s.OpenHours)
	builder.WriteString(", ")
	builder.WriteString("createTime=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updateTime=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Shops is a parsable slice of Shop.
type Shops []*Shop
