// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"hmdp/ent/voucher"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Voucher is the model entity for the Voucher schema.
type Voucher struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID uint64 `json:"id,omitempty"`
	// 商铺id
	ShopID uint64 `json:"shop_id,omitempty"`
	// 代金券标题
	Title string `json:"title,omitempty"`
	// 副标题
	SubTitle string `json:"sub_title,omitempty"`
	// 使用规则
	Rules string `json:"rules,omitempty"`
	// 支付金额
	PayValue uint64 `json:"pay_value,omitempty"`
	// 抵扣金额
	ActualValue int64 `json:"actual_value,omitempty"`
	// 优惠券类型
	Type int8 `json:"type,omitempty"`
	// 优惠券类型
	Status int8 `json:"status,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VoucherQuery when eager-loading is set.
	Edges        VoucherEdges `json:"edges"`
	selectValues sql.SelectValues
}

// VoucherEdges holds the relations/edges for other nodes in the graph.
type VoucherEdges struct {
	// GetMore holds the value of the getMore edge.
	GetMore []*SeckillVoucher `json:"getMore,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GetMoreOrErr returns the GetMore value or an error if the edge
// was not loaded in eager-loading.
func (e VoucherEdges) GetMoreOrErr() ([]*SeckillVoucher, error) {
	if e.loadedTypes[0] {
		return e.GetMore, nil
	}
	return nil, &NotLoadedError{edge: "getMore"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Voucher) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case voucher.FieldID, voucher.FieldShopID, voucher.FieldPayValue, voucher.FieldActualValue, voucher.FieldType, voucher.FieldStatus:
			values[i] = new(sql.NullInt64)
		case voucher.FieldTitle, voucher.FieldSubTitle, voucher.FieldRules:
			values[i] = new(sql.NullString)
		case voucher.FieldCreateTime, voucher.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Voucher fields.
func (v *Voucher) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case voucher.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = uint64(value.Int64)
		case voucher.FieldShopID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field shop_id", values[i])
			} else if value.Valid {
				v.ShopID = uint64(value.Int64)
			}
		case voucher.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				v.Title = value.String
			}
		case voucher.FieldSubTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sub_title", values[i])
			} else if value.Valid {
				v.SubTitle = value.String
			}
		case voucher.FieldRules:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rules", values[i])
			} else if value.Valid {
				v.Rules = value.String
			}
		case voucher.FieldPayValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pay_value", values[i])
			} else if value.Valid {
				v.PayValue = uint64(value.Int64)
			}
		case voucher.FieldActualValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field actual_value", values[i])
			} else if value.Valid {
				v.ActualValue = value.Int64
			}
		case voucher.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				v.Type = int8(value.Int64)
			}
		case voucher.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				v.Status = int8(value.Int64)
			}
		case voucher.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				v.CreateTime = value.Time
			}
		case voucher.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				v.UpdateTime = value.Time
			}
		default:
			v.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Voucher.
// This includes values selected through modifiers, order, etc.
func (v *Voucher) Value(name string) (ent.Value, error) {
	return v.selectValues.Get(name)
}

// QueryGetMore queries the "getMore" edge of the Voucher entity.
func (v *Voucher) QueryGetMore() *SeckillVoucherQuery {
	return NewVoucherClient(v.config).QueryGetMore(v)
}

// Update returns a builder for updating this Voucher.
// Note that you need to call Voucher.Unwrap() before calling this method if this Voucher
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Voucher) Update() *VoucherUpdateOne {
	return NewVoucherClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Voucher entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Voucher) Unwrap() *Voucher {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Voucher is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Voucher) String() string {
	var builder strings.Builder
	builder.WriteString("Voucher(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("shop_id=")
	builder.WriteString(fmt.Sprintf("%v", v.ShopID))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(v.Title)
	builder.WriteString(", ")
	builder.WriteString("sub_title=")
	builder.WriteString(v.SubTitle)
	builder.WriteString(", ")
	builder.WriteString("rules=")
	builder.WriteString(v.Rules)
	builder.WriteString(", ")
	builder.WriteString("pay_value=")
	builder.WriteString(fmt.Sprintf("%v", v.PayValue))
	builder.WriteString(", ")
	builder.WriteString("actual_value=")
	builder.WriteString(fmt.Sprintf("%v", v.ActualValue))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", v.Type))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", v.Status))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(v.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(v.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Vouchers is a parsable slice of Voucher.
type Vouchers []*Voucher
