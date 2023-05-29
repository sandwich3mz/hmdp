// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hmdp/ent/predicate"
	"hmdp/ent/voucherorder"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VoucherOrderQuery is the builder for querying VoucherOrder entities.
type VoucherOrderQuery struct {
	config
	ctx        *QueryContext
	order      []voucherorder.OrderOption
	inters     []Interceptor
	predicates []predicate.VoucherOrder
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VoucherOrderQuery builder.
func (voq *VoucherOrderQuery) Where(ps ...predicate.VoucherOrder) *VoucherOrderQuery {
	voq.predicates = append(voq.predicates, ps...)
	return voq
}

// Limit the number of records to be returned by this query.
func (voq *VoucherOrderQuery) Limit(limit int) *VoucherOrderQuery {
	voq.ctx.Limit = &limit
	return voq
}

// Offset to start from.
func (voq *VoucherOrderQuery) Offset(offset int) *VoucherOrderQuery {
	voq.ctx.Offset = &offset
	return voq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (voq *VoucherOrderQuery) Unique(unique bool) *VoucherOrderQuery {
	voq.ctx.Unique = &unique
	return voq
}

// Order specifies how the records should be ordered.
func (voq *VoucherOrderQuery) Order(o ...voucherorder.OrderOption) *VoucherOrderQuery {
	voq.order = append(voq.order, o...)
	return voq
}

// First returns the first VoucherOrder entity from the query.
// Returns a *NotFoundError when no VoucherOrder was found.
func (voq *VoucherOrderQuery) First(ctx context.Context) (*VoucherOrder, error) {
	nodes, err := voq.Limit(1).All(setContextOp(ctx, voq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{voucherorder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (voq *VoucherOrderQuery) FirstX(ctx context.Context) *VoucherOrder {
	node, err := voq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VoucherOrder ID from the query.
// Returns a *NotFoundError when no VoucherOrder ID was found.
func (voq *VoucherOrderQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = voq.Limit(1).IDs(setContextOp(ctx, voq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{voucherorder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (voq *VoucherOrderQuery) FirstIDX(ctx context.Context) int64 {
	id, err := voq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VoucherOrder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VoucherOrder entity is found.
// Returns a *NotFoundError when no VoucherOrder entities are found.
func (voq *VoucherOrderQuery) Only(ctx context.Context) (*VoucherOrder, error) {
	nodes, err := voq.Limit(2).All(setContextOp(ctx, voq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{voucherorder.Label}
	default:
		return nil, &NotSingularError{voucherorder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (voq *VoucherOrderQuery) OnlyX(ctx context.Context) *VoucherOrder {
	node, err := voq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VoucherOrder ID in the query.
// Returns a *NotSingularError when more than one VoucherOrder ID is found.
// Returns a *NotFoundError when no entities are found.
func (voq *VoucherOrderQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = voq.Limit(2).IDs(setContextOp(ctx, voq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{voucherorder.Label}
	default:
		err = &NotSingularError{voucherorder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (voq *VoucherOrderQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := voq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VoucherOrders.
func (voq *VoucherOrderQuery) All(ctx context.Context) ([]*VoucherOrder, error) {
	ctx = setContextOp(ctx, voq.ctx, "All")
	if err := voq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VoucherOrder, *VoucherOrderQuery]()
	return withInterceptors[[]*VoucherOrder](ctx, voq, qr, voq.inters)
}

// AllX is like All, but panics if an error occurs.
func (voq *VoucherOrderQuery) AllX(ctx context.Context) []*VoucherOrder {
	nodes, err := voq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VoucherOrder IDs.
func (voq *VoucherOrderQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if voq.ctx.Unique == nil && voq.path != nil {
		voq.Unique(true)
	}
	ctx = setContextOp(ctx, voq.ctx, "IDs")
	if err = voq.Select(voucherorder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (voq *VoucherOrderQuery) IDsX(ctx context.Context) []int64 {
	ids, err := voq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (voq *VoucherOrderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, voq.ctx, "Count")
	if err := voq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, voq, querierCount[*VoucherOrderQuery](), voq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (voq *VoucherOrderQuery) CountX(ctx context.Context) int {
	count, err := voq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (voq *VoucherOrderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, voq.ctx, "Exist")
	switch _, err := voq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (voq *VoucherOrderQuery) ExistX(ctx context.Context) bool {
	exist, err := voq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VoucherOrderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (voq *VoucherOrderQuery) Clone() *VoucherOrderQuery {
	if voq == nil {
		return nil
	}
	return &VoucherOrderQuery{
		config:     voq.config,
		ctx:        voq.ctx.Clone(),
		order:      append([]voucherorder.OrderOption{}, voq.order...),
		inters:     append([]Interceptor{}, voq.inters...),
		predicates: append([]predicate.VoucherOrder{}, voq.predicates...),
		// clone intermediate query.
		sql:  voq.sql.Clone(),
		path: voq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uint64 `json:"userId"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VoucherOrder.Query().
//		GroupBy(voucherorder.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (voq *VoucherOrderQuery) GroupBy(field string, fields ...string) *VoucherOrderGroupBy {
	voq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VoucherOrderGroupBy{build: voq}
	grbuild.flds = &voq.ctx.Fields
	grbuild.label = voucherorder.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uint64 `json:"userId"`
//	}
//
//	client.VoucherOrder.Query().
//		Select(voucherorder.FieldUserID).
//		Scan(ctx, &v)
func (voq *VoucherOrderQuery) Select(fields ...string) *VoucherOrderSelect {
	voq.ctx.Fields = append(voq.ctx.Fields, fields...)
	sbuild := &VoucherOrderSelect{VoucherOrderQuery: voq}
	sbuild.label = voucherorder.Label
	sbuild.flds, sbuild.scan = &voq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VoucherOrderSelect configured with the given aggregations.
func (voq *VoucherOrderQuery) Aggregate(fns ...AggregateFunc) *VoucherOrderSelect {
	return voq.Select().Aggregate(fns...)
}

func (voq *VoucherOrderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range voq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, voq); err != nil {
				return err
			}
		}
	}
	for _, f := range voq.ctx.Fields {
		if !voucherorder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if voq.path != nil {
		prev, err := voq.path(ctx)
		if err != nil {
			return err
		}
		voq.sql = prev
	}
	return nil
}

func (voq *VoucherOrderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VoucherOrder, error) {
	var (
		nodes = []*VoucherOrder{}
		_spec = voq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VoucherOrder).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VoucherOrder{config: voq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, voq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (voq *VoucherOrderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := voq.querySpec()
	_spec.Node.Columns = voq.ctx.Fields
	if len(voq.ctx.Fields) > 0 {
		_spec.Unique = voq.ctx.Unique != nil && *voq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, voq.driver, _spec)
}

func (voq *VoucherOrderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(voucherorder.Table, voucherorder.Columns, sqlgraph.NewFieldSpec(voucherorder.FieldID, field.TypeInt64))
	_spec.From = voq.sql
	if unique := voq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if voq.path != nil {
		_spec.Unique = true
	}
	if fields := voq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, voucherorder.FieldID)
		for i := range fields {
			if fields[i] != voucherorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := voq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := voq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := voq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := voq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (voq *VoucherOrderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(voq.driver.Dialect())
	t1 := builder.Table(voucherorder.Table)
	columns := voq.ctx.Fields
	if len(columns) == 0 {
		columns = voucherorder.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if voq.sql != nil {
		selector = voq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if voq.ctx.Unique != nil && *voq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range voq.predicates {
		p(selector)
	}
	for _, p := range voq.order {
		p(selector)
	}
	if offset := voq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := voq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VoucherOrderGroupBy is the group-by builder for VoucherOrder entities.
type VoucherOrderGroupBy struct {
	selector
	build *VoucherOrderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vogb *VoucherOrderGroupBy) Aggregate(fns ...AggregateFunc) *VoucherOrderGroupBy {
	vogb.fns = append(vogb.fns, fns...)
	return vogb
}

// Scan applies the selector query and scans the result into the given value.
func (vogb *VoucherOrderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vogb.build.ctx, "GroupBy")
	if err := vogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VoucherOrderQuery, *VoucherOrderGroupBy](ctx, vogb.build, vogb, vogb.build.inters, v)
}

func (vogb *VoucherOrderGroupBy) sqlScan(ctx context.Context, root *VoucherOrderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vogb.fns))
	for _, fn := range vogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vogb.flds)+len(vogb.fns))
		for _, f := range *vogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VoucherOrderSelect is the builder for selecting fields of VoucherOrder entities.
type VoucherOrderSelect struct {
	*VoucherOrderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vos *VoucherOrderSelect) Aggregate(fns ...AggregateFunc) *VoucherOrderSelect {
	vos.fns = append(vos.fns, fns...)
	return vos
}

// Scan applies the selector query and scans the result into the given value.
func (vos *VoucherOrderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vos.ctx, "Select")
	if err := vos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VoucherOrderQuery, *VoucherOrderSelect](ctx, vos.VoucherOrderQuery, vos, vos.inters, v)
}

func (vos *VoucherOrderSelect) sqlScan(ctx context.Context, root *VoucherOrderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vos.fns))
	for _, fn := range vos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
