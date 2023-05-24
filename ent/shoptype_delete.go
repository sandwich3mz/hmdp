// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"hmdp/ent/predicate"
	"hmdp/ent/shoptype"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShopTypeDelete is the builder for deleting a ShopType entity.
type ShopTypeDelete struct {
	config
	hooks    []Hook
	mutation *ShopTypeMutation
}

// Where appends a list predicates to the ShopTypeDelete builder.
func (std *ShopTypeDelete) Where(ps ...predicate.ShopType) *ShopTypeDelete {
	std.mutation.Where(ps...)
	return std
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (std *ShopTypeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, std.sqlExec, std.mutation, std.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (std *ShopTypeDelete) ExecX(ctx context.Context) int {
	n, err := std.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (std *ShopTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(shoptype.Table, sqlgraph.NewFieldSpec(shoptype.FieldID, field.TypeInt64))
	if ps := std.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, std.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	std.mutation.done = true
	return affected, err
}

// ShopTypeDeleteOne is the builder for deleting a single ShopType entity.
type ShopTypeDeleteOne struct {
	std *ShopTypeDelete
}

// Where appends a list predicates to the ShopTypeDelete builder.
func (stdo *ShopTypeDeleteOne) Where(ps ...predicate.ShopType) *ShopTypeDeleteOne {
	stdo.std.mutation.Where(ps...)
	return stdo
}

// Exec executes the deletion query.
func (stdo *ShopTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := stdo.std.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{shoptype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (stdo *ShopTypeDeleteOne) ExecX(ctx context.Context) {
	if err := stdo.Exec(ctx); err != nil {
		panic(err)
	}
}
