// Code generated by ent, DO NOT EDIT.

package seckillvoucher

import (
	"hmdp/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldLTE(FieldID, id))
}

// VoucherID applies equality check predicate on the "voucher_id" field. It's identical to VoucherIDEQ.
func VoucherID(v uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldEQ(FieldVoucherID, v))
}

// Stock applies equality check predicate on the "stock" field. It's identical to StockEQ.
func Stock(v uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldEQ(FieldStock, v))
}

// BeginTime applies equality check predicate on the "begin_time" field. It's identical to BeginTimeEQ.
func BeginTime(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldEQ(FieldBeginTime, v))
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldEQ(FieldEndTime, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldEQ(FieldUpdateTime, v))
}

// VoucherIDEQ applies the EQ predicate on the "voucher_id" field.
func VoucherIDEQ(v uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldEQ(FieldVoucherID, v))
}

// VoucherIDNEQ applies the NEQ predicate on the "voucher_id" field.
func VoucherIDNEQ(v uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldNEQ(FieldVoucherID, v))
}

// VoucherIDIn applies the In predicate on the "voucher_id" field.
func VoucherIDIn(vs ...uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldIn(FieldVoucherID, vs...))
}

// VoucherIDNotIn applies the NotIn predicate on the "voucher_id" field.
func VoucherIDNotIn(vs ...uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldNotIn(FieldVoucherID, vs...))
}

// StockEQ applies the EQ predicate on the "stock" field.
func StockEQ(v uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldEQ(FieldStock, v))
}

// StockNEQ applies the NEQ predicate on the "stock" field.
func StockNEQ(v uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldNEQ(FieldStock, v))
}

// StockIn applies the In predicate on the "stock" field.
func StockIn(vs ...uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldIn(FieldStock, vs...))
}

// StockNotIn applies the NotIn predicate on the "stock" field.
func StockNotIn(vs ...uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldNotIn(FieldStock, vs...))
}

// StockGT applies the GT predicate on the "stock" field.
func StockGT(v uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldGT(FieldStock, v))
}

// StockGTE applies the GTE predicate on the "stock" field.
func StockGTE(v uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldGTE(FieldStock, v))
}

// StockLT applies the LT predicate on the "stock" field.
func StockLT(v uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldLT(FieldStock, v))
}

// StockLTE applies the LTE predicate on the "stock" field.
func StockLTE(v uint64) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldLTE(FieldStock, v))
}

// BeginTimeEQ applies the EQ predicate on the "begin_time" field.
func BeginTimeEQ(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldEQ(FieldBeginTime, v))
}

// BeginTimeNEQ applies the NEQ predicate on the "begin_time" field.
func BeginTimeNEQ(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldNEQ(FieldBeginTime, v))
}

// BeginTimeIn applies the In predicate on the "begin_time" field.
func BeginTimeIn(vs ...time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldIn(FieldBeginTime, vs...))
}

// BeginTimeNotIn applies the NotIn predicate on the "begin_time" field.
func BeginTimeNotIn(vs ...time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldNotIn(FieldBeginTime, vs...))
}

// BeginTimeGT applies the GT predicate on the "begin_time" field.
func BeginTimeGT(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldGT(FieldBeginTime, v))
}

// BeginTimeGTE applies the GTE predicate on the "begin_time" field.
func BeginTimeGTE(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldGTE(FieldBeginTime, v))
}

// BeginTimeLT applies the LT predicate on the "begin_time" field.
func BeginTimeLT(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldLT(FieldBeginTime, v))
}

// BeginTimeLTE applies the LTE predicate on the "begin_time" field.
func BeginTimeLTE(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldLTE(FieldBeginTime, v))
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldEQ(FieldEndTime, v))
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldNEQ(FieldEndTime, v))
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldIn(FieldEndTime, vs...))
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldNotIn(FieldEndTime, vs...))
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldGT(FieldEndTime, v))
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldGTE(FieldEndTime, v))
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldLT(FieldEndTime, v))
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldLTE(FieldEndTime, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(sql.FieldLTE(FieldUpdateTime, v))
}

// HasGetForm applies the HasEdge predicate on the "getForm" edge.
func HasGetForm() predicate.SeckillVoucher {
	return predicate.SeckillVoucher(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GetFormTable, GetFormColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGetFormWith applies the HasEdge predicate on the "getForm" edge with a given conditions (other predicates).
func HasGetFormWith(preds ...predicate.Voucher) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(func(s *sql.Selector) {
		step := newGetFormStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SeckillVoucher) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SeckillVoucher) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SeckillVoucher) predicate.SeckillVoucher {
	return predicate.SeckillVoucher(func(s *sql.Selector) {
		p(s.Not())
	})
}