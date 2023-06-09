// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hmdp/ent/predicate"
	"hmdp/ent/voucherorder"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VoucherOrderUpdate is the builder for updating VoucherOrder entities.
type VoucherOrderUpdate struct {
	config
	hooks    []Hook
	mutation *VoucherOrderMutation
}

// Where appends a list predicates to the VoucherOrderUpdate builder.
func (vou *VoucherOrderUpdate) Where(ps ...predicate.VoucherOrder) *VoucherOrderUpdate {
	vou.mutation.Where(ps...)
	return vou
}

// SetUserID sets the "user_id" field.
func (vou *VoucherOrderUpdate) SetUserID(u uint64) *VoucherOrderUpdate {
	vou.mutation.ResetUserID()
	vou.mutation.SetUserID(u)
	return vou
}

// AddUserID adds u to the "user_id" field.
func (vou *VoucherOrderUpdate) AddUserID(u int64) *VoucherOrderUpdate {
	vou.mutation.AddUserID(u)
	return vou
}

// SetVoucherID sets the "voucher_id" field.
func (vou *VoucherOrderUpdate) SetVoucherID(u uint64) *VoucherOrderUpdate {
	vou.mutation.ResetVoucherID()
	vou.mutation.SetVoucherID(u)
	return vou
}

// AddVoucherID adds u to the "voucher_id" field.
func (vou *VoucherOrderUpdate) AddVoucherID(u int64) *VoucherOrderUpdate {
	vou.mutation.AddVoucherID(u)
	return vou
}

// SetPayType sets the "pay_type" field.
func (vou *VoucherOrderUpdate) SetPayType(i int8) *VoucherOrderUpdate {
	vou.mutation.ResetPayType()
	vou.mutation.SetPayType(i)
	return vou
}

// SetNillablePayType sets the "pay_type" field if the given value is not nil.
func (vou *VoucherOrderUpdate) SetNillablePayType(i *int8) *VoucherOrderUpdate {
	if i != nil {
		vou.SetPayType(*i)
	}
	return vou
}

// AddPayType adds i to the "pay_type" field.
func (vou *VoucherOrderUpdate) AddPayType(i int8) *VoucherOrderUpdate {
	vou.mutation.AddPayType(i)
	return vou
}

// ClearPayType clears the value of the "pay_type" field.
func (vou *VoucherOrderUpdate) ClearPayType() *VoucherOrderUpdate {
	vou.mutation.ClearPayType()
	return vou
}

// SetStatus sets the "status" field.
func (vou *VoucherOrderUpdate) SetStatus(i int8) *VoucherOrderUpdate {
	vou.mutation.ResetStatus()
	vou.mutation.SetStatus(i)
	return vou
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (vou *VoucherOrderUpdate) SetNillableStatus(i *int8) *VoucherOrderUpdate {
	if i != nil {
		vou.SetStatus(*i)
	}
	return vou
}

// AddStatus adds i to the "status" field.
func (vou *VoucherOrderUpdate) AddStatus(i int8) *VoucherOrderUpdate {
	vou.mutation.AddStatus(i)
	return vou
}

// ClearStatus clears the value of the "status" field.
func (vou *VoucherOrderUpdate) ClearStatus() *VoucherOrderUpdate {
	vou.mutation.ClearStatus()
	return vou
}

// SetPayTime sets the "pay_time" field.
func (vou *VoucherOrderUpdate) SetPayTime(t time.Time) *VoucherOrderUpdate {
	vou.mutation.SetPayTime(t)
	return vou
}

// SetNillablePayTime sets the "pay_time" field if the given value is not nil.
func (vou *VoucherOrderUpdate) SetNillablePayTime(t *time.Time) *VoucherOrderUpdate {
	if t != nil {
		vou.SetPayTime(*t)
	}
	return vou
}

// ClearPayTime clears the value of the "pay_time" field.
func (vou *VoucherOrderUpdate) ClearPayTime() *VoucherOrderUpdate {
	vou.mutation.ClearPayTime()
	return vou
}

// SetUseTime sets the "use_time" field.
func (vou *VoucherOrderUpdate) SetUseTime(t time.Time) *VoucherOrderUpdate {
	vou.mutation.SetUseTime(t)
	return vou
}

// SetNillableUseTime sets the "use_time" field if the given value is not nil.
func (vou *VoucherOrderUpdate) SetNillableUseTime(t *time.Time) *VoucherOrderUpdate {
	if t != nil {
		vou.SetUseTime(*t)
	}
	return vou
}

// ClearUseTime clears the value of the "use_time" field.
func (vou *VoucherOrderUpdate) ClearUseTime() *VoucherOrderUpdate {
	vou.mutation.ClearUseTime()
	return vou
}

// SetRefundTime sets the "refund_time" field.
func (vou *VoucherOrderUpdate) SetRefundTime(t time.Time) *VoucherOrderUpdate {
	vou.mutation.SetRefundTime(t)
	return vou
}

// SetNillableRefundTime sets the "refund_time" field if the given value is not nil.
func (vou *VoucherOrderUpdate) SetNillableRefundTime(t *time.Time) *VoucherOrderUpdate {
	if t != nil {
		vou.SetRefundTime(*t)
	}
	return vou
}

// ClearRefundTime clears the value of the "refund_time" field.
func (vou *VoucherOrderUpdate) ClearRefundTime() *VoucherOrderUpdate {
	vou.mutation.ClearRefundTime()
	return vou
}

// SetCreateTime sets the "create_time" field.
func (vou *VoucherOrderUpdate) SetCreateTime(t time.Time) *VoucherOrderUpdate {
	vou.mutation.SetCreateTime(t)
	return vou
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (vou *VoucherOrderUpdate) SetNillableCreateTime(t *time.Time) *VoucherOrderUpdate {
	if t != nil {
		vou.SetCreateTime(*t)
	}
	return vou
}

// SetUpdateTime sets the "update_time" field.
func (vou *VoucherOrderUpdate) SetUpdateTime(t time.Time) *VoucherOrderUpdate {
	vou.mutation.SetUpdateTime(t)
	return vou
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (vou *VoucherOrderUpdate) SetNillableUpdateTime(t *time.Time) *VoucherOrderUpdate {
	if t != nil {
		vou.SetUpdateTime(*t)
	}
	return vou
}

// Mutation returns the VoucherOrderMutation object of the builder.
func (vou *VoucherOrderUpdate) Mutation() *VoucherOrderMutation {
	return vou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vou *VoucherOrderUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, vou.sqlSave, vou.mutation, vou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vou *VoucherOrderUpdate) SaveX(ctx context.Context) int {
	affected, err := vou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vou *VoucherOrderUpdate) Exec(ctx context.Context) error {
	_, err := vou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vou *VoucherOrderUpdate) ExecX(ctx context.Context) {
	if err := vou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vou *VoucherOrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(voucherorder.Table, voucherorder.Columns, sqlgraph.NewFieldSpec(voucherorder.FieldID, field.TypeInt64))
	if ps := vou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vou.mutation.UserID(); ok {
		_spec.SetField(voucherorder.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := vou.mutation.AddedUserID(); ok {
		_spec.AddField(voucherorder.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := vou.mutation.VoucherID(); ok {
		_spec.SetField(voucherorder.FieldVoucherID, field.TypeUint64, value)
	}
	if value, ok := vou.mutation.AddedVoucherID(); ok {
		_spec.AddField(voucherorder.FieldVoucherID, field.TypeUint64, value)
	}
	if value, ok := vou.mutation.PayType(); ok {
		_spec.SetField(voucherorder.FieldPayType, field.TypeInt8, value)
	}
	if value, ok := vou.mutation.AddedPayType(); ok {
		_spec.AddField(voucherorder.FieldPayType, field.TypeInt8, value)
	}
	if vou.mutation.PayTypeCleared() {
		_spec.ClearField(voucherorder.FieldPayType, field.TypeInt8)
	}
	if value, ok := vou.mutation.Status(); ok {
		_spec.SetField(voucherorder.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := vou.mutation.AddedStatus(); ok {
		_spec.AddField(voucherorder.FieldStatus, field.TypeInt8, value)
	}
	if vou.mutation.StatusCleared() {
		_spec.ClearField(voucherorder.FieldStatus, field.TypeInt8)
	}
	if value, ok := vou.mutation.PayTime(); ok {
		_spec.SetField(voucherorder.FieldPayTime, field.TypeTime, value)
	}
	if vou.mutation.PayTimeCleared() {
		_spec.ClearField(voucherorder.FieldPayTime, field.TypeTime)
	}
	if value, ok := vou.mutation.UseTime(); ok {
		_spec.SetField(voucherorder.FieldUseTime, field.TypeTime, value)
	}
	if vou.mutation.UseTimeCleared() {
		_spec.ClearField(voucherorder.FieldUseTime, field.TypeTime)
	}
	if value, ok := vou.mutation.RefundTime(); ok {
		_spec.SetField(voucherorder.FieldRefundTime, field.TypeTime, value)
	}
	if vou.mutation.RefundTimeCleared() {
		_spec.ClearField(voucherorder.FieldRefundTime, field.TypeTime)
	}
	if value, ok := vou.mutation.CreateTime(); ok {
		_spec.SetField(voucherorder.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := vou.mutation.UpdateTime(); ok {
		_spec.SetField(voucherorder.FieldUpdateTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{voucherorder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vou.mutation.done = true
	return n, nil
}

// VoucherOrderUpdateOne is the builder for updating a single VoucherOrder entity.
type VoucherOrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VoucherOrderMutation
}

// SetUserID sets the "user_id" field.
func (vouo *VoucherOrderUpdateOne) SetUserID(u uint64) *VoucherOrderUpdateOne {
	vouo.mutation.ResetUserID()
	vouo.mutation.SetUserID(u)
	return vouo
}

// AddUserID adds u to the "user_id" field.
func (vouo *VoucherOrderUpdateOne) AddUserID(u int64) *VoucherOrderUpdateOne {
	vouo.mutation.AddUserID(u)
	return vouo
}

// SetVoucherID sets the "voucher_id" field.
func (vouo *VoucherOrderUpdateOne) SetVoucherID(u uint64) *VoucherOrderUpdateOne {
	vouo.mutation.ResetVoucherID()
	vouo.mutation.SetVoucherID(u)
	return vouo
}

// AddVoucherID adds u to the "voucher_id" field.
func (vouo *VoucherOrderUpdateOne) AddVoucherID(u int64) *VoucherOrderUpdateOne {
	vouo.mutation.AddVoucherID(u)
	return vouo
}

// SetPayType sets the "pay_type" field.
func (vouo *VoucherOrderUpdateOne) SetPayType(i int8) *VoucherOrderUpdateOne {
	vouo.mutation.ResetPayType()
	vouo.mutation.SetPayType(i)
	return vouo
}

// SetNillablePayType sets the "pay_type" field if the given value is not nil.
func (vouo *VoucherOrderUpdateOne) SetNillablePayType(i *int8) *VoucherOrderUpdateOne {
	if i != nil {
		vouo.SetPayType(*i)
	}
	return vouo
}

// AddPayType adds i to the "pay_type" field.
func (vouo *VoucherOrderUpdateOne) AddPayType(i int8) *VoucherOrderUpdateOne {
	vouo.mutation.AddPayType(i)
	return vouo
}

// ClearPayType clears the value of the "pay_type" field.
func (vouo *VoucherOrderUpdateOne) ClearPayType() *VoucherOrderUpdateOne {
	vouo.mutation.ClearPayType()
	return vouo
}

// SetStatus sets the "status" field.
func (vouo *VoucherOrderUpdateOne) SetStatus(i int8) *VoucherOrderUpdateOne {
	vouo.mutation.ResetStatus()
	vouo.mutation.SetStatus(i)
	return vouo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (vouo *VoucherOrderUpdateOne) SetNillableStatus(i *int8) *VoucherOrderUpdateOne {
	if i != nil {
		vouo.SetStatus(*i)
	}
	return vouo
}

// AddStatus adds i to the "status" field.
func (vouo *VoucherOrderUpdateOne) AddStatus(i int8) *VoucherOrderUpdateOne {
	vouo.mutation.AddStatus(i)
	return vouo
}

// ClearStatus clears the value of the "status" field.
func (vouo *VoucherOrderUpdateOne) ClearStatus() *VoucherOrderUpdateOne {
	vouo.mutation.ClearStatus()
	return vouo
}

// SetPayTime sets the "pay_time" field.
func (vouo *VoucherOrderUpdateOne) SetPayTime(t time.Time) *VoucherOrderUpdateOne {
	vouo.mutation.SetPayTime(t)
	return vouo
}

// SetNillablePayTime sets the "pay_time" field if the given value is not nil.
func (vouo *VoucherOrderUpdateOne) SetNillablePayTime(t *time.Time) *VoucherOrderUpdateOne {
	if t != nil {
		vouo.SetPayTime(*t)
	}
	return vouo
}

// ClearPayTime clears the value of the "pay_time" field.
func (vouo *VoucherOrderUpdateOne) ClearPayTime() *VoucherOrderUpdateOne {
	vouo.mutation.ClearPayTime()
	return vouo
}

// SetUseTime sets the "use_time" field.
func (vouo *VoucherOrderUpdateOne) SetUseTime(t time.Time) *VoucherOrderUpdateOne {
	vouo.mutation.SetUseTime(t)
	return vouo
}

// SetNillableUseTime sets the "use_time" field if the given value is not nil.
func (vouo *VoucherOrderUpdateOne) SetNillableUseTime(t *time.Time) *VoucherOrderUpdateOne {
	if t != nil {
		vouo.SetUseTime(*t)
	}
	return vouo
}

// ClearUseTime clears the value of the "use_time" field.
func (vouo *VoucherOrderUpdateOne) ClearUseTime() *VoucherOrderUpdateOne {
	vouo.mutation.ClearUseTime()
	return vouo
}

// SetRefundTime sets the "refund_time" field.
func (vouo *VoucherOrderUpdateOne) SetRefundTime(t time.Time) *VoucherOrderUpdateOne {
	vouo.mutation.SetRefundTime(t)
	return vouo
}

// SetNillableRefundTime sets the "refund_time" field if the given value is not nil.
func (vouo *VoucherOrderUpdateOne) SetNillableRefundTime(t *time.Time) *VoucherOrderUpdateOne {
	if t != nil {
		vouo.SetRefundTime(*t)
	}
	return vouo
}

// ClearRefundTime clears the value of the "refund_time" field.
func (vouo *VoucherOrderUpdateOne) ClearRefundTime() *VoucherOrderUpdateOne {
	vouo.mutation.ClearRefundTime()
	return vouo
}

// SetCreateTime sets the "create_time" field.
func (vouo *VoucherOrderUpdateOne) SetCreateTime(t time.Time) *VoucherOrderUpdateOne {
	vouo.mutation.SetCreateTime(t)
	return vouo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (vouo *VoucherOrderUpdateOne) SetNillableCreateTime(t *time.Time) *VoucherOrderUpdateOne {
	if t != nil {
		vouo.SetCreateTime(*t)
	}
	return vouo
}

// SetUpdateTime sets the "update_time" field.
func (vouo *VoucherOrderUpdateOne) SetUpdateTime(t time.Time) *VoucherOrderUpdateOne {
	vouo.mutation.SetUpdateTime(t)
	return vouo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (vouo *VoucherOrderUpdateOne) SetNillableUpdateTime(t *time.Time) *VoucherOrderUpdateOne {
	if t != nil {
		vouo.SetUpdateTime(*t)
	}
	return vouo
}

// Mutation returns the VoucherOrderMutation object of the builder.
func (vouo *VoucherOrderUpdateOne) Mutation() *VoucherOrderMutation {
	return vouo.mutation
}

// Where appends a list predicates to the VoucherOrderUpdate builder.
func (vouo *VoucherOrderUpdateOne) Where(ps ...predicate.VoucherOrder) *VoucherOrderUpdateOne {
	vouo.mutation.Where(ps...)
	return vouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vouo *VoucherOrderUpdateOne) Select(field string, fields ...string) *VoucherOrderUpdateOne {
	vouo.fields = append([]string{field}, fields...)
	return vouo
}

// Save executes the query and returns the updated VoucherOrder entity.
func (vouo *VoucherOrderUpdateOne) Save(ctx context.Context) (*VoucherOrder, error) {
	return withHooks(ctx, vouo.sqlSave, vouo.mutation, vouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vouo *VoucherOrderUpdateOne) SaveX(ctx context.Context) *VoucherOrder {
	node, err := vouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vouo *VoucherOrderUpdateOne) Exec(ctx context.Context) error {
	_, err := vouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vouo *VoucherOrderUpdateOne) ExecX(ctx context.Context) {
	if err := vouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vouo *VoucherOrderUpdateOne) sqlSave(ctx context.Context) (_node *VoucherOrder, err error) {
	_spec := sqlgraph.NewUpdateSpec(voucherorder.Table, voucherorder.Columns, sqlgraph.NewFieldSpec(voucherorder.FieldID, field.TypeInt64))
	id, ok := vouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "VoucherOrder.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, voucherorder.FieldID)
		for _, f := range fields {
			if !voucherorder.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != voucherorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vouo.mutation.UserID(); ok {
		_spec.SetField(voucherorder.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := vouo.mutation.AddedUserID(); ok {
		_spec.AddField(voucherorder.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := vouo.mutation.VoucherID(); ok {
		_spec.SetField(voucherorder.FieldVoucherID, field.TypeUint64, value)
	}
	if value, ok := vouo.mutation.AddedVoucherID(); ok {
		_spec.AddField(voucherorder.FieldVoucherID, field.TypeUint64, value)
	}
	if value, ok := vouo.mutation.PayType(); ok {
		_spec.SetField(voucherorder.FieldPayType, field.TypeInt8, value)
	}
	if value, ok := vouo.mutation.AddedPayType(); ok {
		_spec.AddField(voucherorder.FieldPayType, field.TypeInt8, value)
	}
	if vouo.mutation.PayTypeCleared() {
		_spec.ClearField(voucherorder.FieldPayType, field.TypeInt8)
	}
	if value, ok := vouo.mutation.Status(); ok {
		_spec.SetField(voucherorder.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := vouo.mutation.AddedStatus(); ok {
		_spec.AddField(voucherorder.FieldStatus, field.TypeInt8, value)
	}
	if vouo.mutation.StatusCleared() {
		_spec.ClearField(voucherorder.FieldStatus, field.TypeInt8)
	}
	if value, ok := vouo.mutation.PayTime(); ok {
		_spec.SetField(voucherorder.FieldPayTime, field.TypeTime, value)
	}
	if vouo.mutation.PayTimeCleared() {
		_spec.ClearField(voucherorder.FieldPayTime, field.TypeTime)
	}
	if value, ok := vouo.mutation.UseTime(); ok {
		_spec.SetField(voucherorder.FieldUseTime, field.TypeTime, value)
	}
	if vouo.mutation.UseTimeCleared() {
		_spec.ClearField(voucherorder.FieldUseTime, field.TypeTime)
	}
	if value, ok := vouo.mutation.RefundTime(); ok {
		_spec.SetField(voucherorder.FieldRefundTime, field.TypeTime, value)
	}
	if vouo.mutation.RefundTimeCleared() {
		_spec.ClearField(voucherorder.FieldRefundTime, field.TypeTime)
	}
	if value, ok := vouo.mutation.CreateTime(); ok {
		_spec.SetField(voucherorder.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := vouo.mutation.UpdateTime(); ok {
		_spec.SetField(voucherorder.FieldUpdateTime, field.TypeTime, value)
	}
	_node = &VoucherOrder{config: vouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{voucherorder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vouo.mutation.done = true
	return _node, nil
}
