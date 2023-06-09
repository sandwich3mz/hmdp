// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hmdp/ent/seckillvoucher"
	"hmdp/ent/voucher"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VoucherCreate is the builder for creating a Voucher entity.
type VoucherCreate struct {
	config
	mutation *VoucherMutation
	hooks    []Hook
}

// SetShopID sets the "shop_id" field.
func (vc *VoucherCreate) SetShopID(u uint64) *VoucherCreate {
	vc.mutation.SetShopID(u)
	return vc
}

// SetTitle sets the "title" field.
func (vc *VoucherCreate) SetTitle(s string) *VoucherCreate {
	vc.mutation.SetTitle(s)
	return vc
}

// SetSubTitle sets the "sub_title" field.
func (vc *VoucherCreate) SetSubTitle(s string) *VoucherCreate {
	vc.mutation.SetSubTitle(s)
	return vc
}

// SetRules sets the "rules" field.
func (vc *VoucherCreate) SetRules(s string) *VoucherCreate {
	vc.mutation.SetRules(s)
	return vc
}

// SetPayValue sets the "pay_value" field.
func (vc *VoucherCreate) SetPayValue(u uint64) *VoucherCreate {
	vc.mutation.SetPayValue(u)
	return vc
}

// SetActualValue sets the "actual_value" field.
func (vc *VoucherCreate) SetActualValue(i int64) *VoucherCreate {
	vc.mutation.SetActualValue(i)
	return vc
}

// SetType sets the "type" field.
func (vc *VoucherCreate) SetType(i int8) *VoucherCreate {
	vc.mutation.SetType(i)
	return vc
}

// SetStatus sets the "status" field.
func (vc *VoucherCreate) SetStatus(i int8) *VoucherCreate {
	vc.mutation.SetStatus(i)
	return vc
}

// SetCreateTime sets the "create_time" field.
func (vc *VoucherCreate) SetCreateTime(t time.Time) *VoucherCreate {
	vc.mutation.SetCreateTime(t)
	return vc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (vc *VoucherCreate) SetNillableCreateTime(t *time.Time) *VoucherCreate {
	if t != nil {
		vc.SetCreateTime(*t)
	}
	return vc
}

// SetUpdateTime sets the "update_time" field.
func (vc *VoucherCreate) SetUpdateTime(t time.Time) *VoucherCreate {
	vc.mutation.SetUpdateTime(t)
	return vc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (vc *VoucherCreate) SetNillableUpdateTime(t *time.Time) *VoucherCreate {
	if t != nil {
		vc.SetUpdateTime(*t)
	}
	return vc
}

// SetID sets the "id" field.
func (vc *VoucherCreate) SetID(u uint64) *VoucherCreate {
	vc.mutation.SetID(u)
	return vc
}

// AddGetMoreIDs adds the "getMore" edge to the SeckillVoucher entity by IDs.
func (vc *VoucherCreate) AddGetMoreIDs(ids ...uint64) *VoucherCreate {
	vc.mutation.AddGetMoreIDs(ids...)
	return vc
}

// AddGetMore adds the "getMore" edges to the SeckillVoucher entity.
func (vc *VoucherCreate) AddGetMore(s ...*SeckillVoucher) *VoucherCreate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return vc.AddGetMoreIDs(ids...)
}

// Mutation returns the VoucherMutation object of the builder.
func (vc *VoucherCreate) Mutation() *VoucherMutation {
	return vc.mutation
}

// Save creates the Voucher in the database.
func (vc *VoucherCreate) Save(ctx context.Context) (*Voucher, error) {
	vc.defaults()
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VoucherCreate) SaveX(ctx context.Context) *Voucher {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VoucherCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VoucherCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VoucherCreate) defaults() {
	if _, ok := vc.mutation.CreateTime(); !ok {
		v := voucher.DefaultCreateTime
		vc.mutation.SetCreateTime(v)
	}
	if _, ok := vc.mutation.UpdateTime(); !ok {
		v := voucher.DefaultUpdateTime
		vc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VoucherCreate) check() error {
	if _, ok := vc.mutation.ShopID(); !ok {
		return &ValidationError{Name: "shop_id", err: errors.New(`ent: missing required field "Voucher.shop_id"`)}
	}
	if _, ok := vc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Voucher.title"`)}
	}
	if _, ok := vc.mutation.SubTitle(); !ok {
		return &ValidationError{Name: "sub_title", err: errors.New(`ent: missing required field "Voucher.sub_title"`)}
	}
	if _, ok := vc.mutation.Rules(); !ok {
		return &ValidationError{Name: "rules", err: errors.New(`ent: missing required field "Voucher.rules"`)}
	}
	if _, ok := vc.mutation.PayValue(); !ok {
		return &ValidationError{Name: "pay_value", err: errors.New(`ent: missing required field "Voucher.pay_value"`)}
	}
	if _, ok := vc.mutation.ActualValue(); !ok {
		return &ValidationError{Name: "actual_value", err: errors.New(`ent: missing required field "Voucher.actual_value"`)}
	}
	if _, ok := vc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Voucher.type"`)}
	}
	if _, ok := vc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Voucher.status"`)}
	}
	if _, ok := vc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Voucher.create_time"`)}
	}
	if _, ok := vc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Voucher.update_time"`)}
	}
	return nil
}

func (vc *VoucherCreate) sqlSave(ctx context.Context) (*Voucher, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VoucherCreate) createSpec() (*Voucher, *sqlgraph.CreateSpec) {
	var (
		_node = &Voucher{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(voucher.Table, sqlgraph.NewFieldSpec(voucher.FieldID, field.TypeUint64))
	)
	if id, ok := vc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vc.mutation.ShopID(); ok {
		_spec.SetField(voucher.FieldShopID, field.TypeUint64, value)
		_node.ShopID = value
	}
	if value, ok := vc.mutation.Title(); ok {
		_spec.SetField(voucher.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := vc.mutation.SubTitle(); ok {
		_spec.SetField(voucher.FieldSubTitle, field.TypeString, value)
		_node.SubTitle = value
	}
	if value, ok := vc.mutation.Rules(); ok {
		_spec.SetField(voucher.FieldRules, field.TypeString, value)
		_node.Rules = value
	}
	if value, ok := vc.mutation.PayValue(); ok {
		_spec.SetField(voucher.FieldPayValue, field.TypeUint64, value)
		_node.PayValue = value
	}
	if value, ok := vc.mutation.ActualValue(); ok {
		_spec.SetField(voucher.FieldActualValue, field.TypeInt64, value)
		_node.ActualValue = value
	}
	if value, ok := vc.mutation.GetType(); ok {
		_spec.SetField(voucher.FieldType, field.TypeInt8, value)
		_node.Type = value
	}
	if value, ok := vc.mutation.Status(); ok {
		_spec.SetField(voucher.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := vc.mutation.CreateTime(); ok {
		_spec.SetField(voucher.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := vc.mutation.UpdateTime(); ok {
		_spec.SetField(voucher.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if nodes := vc.mutation.GetMoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   voucher.GetMoreTable,
			Columns: []string{voucher.GetMoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(seckillvoucher.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VoucherCreateBulk is the builder for creating many Voucher entities in bulk.
type VoucherCreateBulk struct {
	config
	builders []*VoucherCreate
}

// Save creates the Voucher entities in the database.
func (vcb *VoucherCreateBulk) Save(ctx context.Context) ([]*Voucher, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Voucher, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VoucherMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VoucherCreateBulk) SaveX(ctx context.Context) []*Voucher {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VoucherCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VoucherCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
