// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hmdp/ent/blog"
	"hmdp/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogUpdate is the builder for updating Blog entities.
type BlogUpdate struct {
	config
	hooks    []Hook
	mutation *BlogMutation
}

// Where appends a list predicates to the BlogUpdate builder.
func (bu *BlogUpdate) Where(ps ...predicate.Blog) *BlogUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetShopId sets the "shopId" field.
func (bu *BlogUpdate) SetShopId(i int64) *BlogUpdate {
	bu.mutation.ResetShopId()
	bu.mutation.SetShopId(i)
	return bu
}

// AddShopId adds i to the "shopId" field.
func (bu *BlogUpdate) AddShopId(i int64) *BlogUpdate {
	bu.mutation.AddShopId(i)
	return bu
}

// SetUserId sets the "userId" field.
func (bu *BlogUpdate) SetUserId(i int64) *BlogUpdate {
	bu.mutation.ResetUserId()
	bu.mutation.SetUserId(i)
	return bu
}

// AddUserId adds i to the "userId" field.
func (bu *BlogUpdate) AddUserId(i int64) *BlogUpdate {
	bu.mutation.AddUserId(i)
	return bu
}

// SetTitle sets the "title" field.
func (bu *BlogUpdate) SetTitle(s string) *BlogUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetImages sets the "images" field.
func (bu *BlogUpdate) SetImages(s string) *BlogUpdate {
	bu.mutation.SetImages(s)
	return bu
}

// SetContent sets the "content" field.
func (bu *BlogUpdate) SetContent(s string) *BlogUpdate {
	bu.mutation.SetContent(s)
	return bu
}

// SetLiked sets the "liked" field.
func (bu *BlogUpdate) SetLiked(i int) *BlogUpdate {
	bu.mutation.ResetLiked()
	bu.mutation.SetLiked(i)
	return bu
}

// AddLiked adds i to the "liked" field.
func (bu *BlogUpdate) AddLiked(i int) *BlogUpdate {
	bu.mutation.AddLiked(i)
	return bu
}

// SetComments sets the "comments" field.
func (bu *BlogUpdate) SetComments(i int) *BlogUpdate {
	bu.mutation.ResetComments()
	bu.mutation.SetComments(i)
	return bu
}

// AddComments adds i to the "comments" field.
func (bu *BlogUpdate) AddComments(i int) *BlogUpdate {
	bu.mutation.AddComments(i)
	return bu
}

// SetCreateTime sets the "createTime" field.
func (bu *BlogUpdate) SetCreateTime(t time.Time) *BlogUpdate {
	bu.mutation.SetCreateTime(t)
	return bu
}

// SetUpdateTime sets the "updateTime" field.
func (bu *BlogUpdate) SetUpdateTime(t time.Time) *BlogUpdate {
	bu.mutation.SetUpdateTime(t)
	return bu
}

// Mutation returns the BlogMutation object of the builder.
func (bu *BlogUpdate) Mutation() *BlogMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlogUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlogUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlogUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BlogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(blog.Table, blog.Columns, sqlgraph.NewFieldSpec(blog.FieldID, field.TypeInt64))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.ShopId(); ok {
		_spec.SetField(blog.FieldShopId, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.AddedShopId(); ok {
		_spec.AddField(blog.FieldShopId, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.UserId(); ok {
		_spec.SetField(blog.FieldUserId, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.AddedUserId(); ok {
		_spec.AddField(blog.FieldUserId, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.SetField(blog.FieldTitle, field.TypeString, value)
	}
	if value, ok := bu.mutation.Images(); ok {
		_spec.SetField(blog.FieldImages, field.TypeString, value)
	}
	if value, ok := bu.mutation.Content(); ok {
		_spec.SetField(blog.FieldContent, field.TypeString, value)
	}
	if value, ok := bu.mutation.Liked(); ok {
		_spec.SetField(blog.FieldLiked, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedLiked(); ok {
		_spec.AddField(blog.FieldLiked, field.TypeInt, value)
	}
	if value, ok := bu.mutation.Comments(); ok {
		_spec.SetField(blog.FieldComments, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedComments(); ok {
		_spec.AddField(blog.FieldComments, field.TypeInt, value)
	}
	if value, ok := bu.mutation.CreateTime(); ok {
		_spec.SetField(blog.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := bu.mutation.UpdateTime(); ok {
		_spec.SetField(blog.FieldUpdateTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BlogUpdateOne is the builder for updating a single Blog entity.
type BlogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlogMutation
}

// SetShopId sets the "shopId" field.
func (buo *BlogUpdateOne) SetShopId(i int64) *BlogUpdateOne {
	buo.mutation.ResetShopId()
	buo.mutation.SetShopId(i)
	return buo
}

// AddShopId adds i to the "shopId" field.
func (buo *BlogUpdateOne) AddShopId(i int64) *BlogUpdateOne {
	buo.mutation.AddShopId(i)
	return buo
}

// SetUserId sets the "userId" field.
func (buo *BlogUpdateOne) SetUserId(i int64) *BlogUpdateOne {
	buo.mutation.ResetUserId()
	buo.mutation.SetUserId(i)
	return buo
}

// AddUserId adds i to the "userId" field.
func (buo *BlogUpdateOne) AddUserId(i int64) *BlogUpdateOne {
	buo.mutation.AddUserId(i)
	return buo
}

// SetTitle sets the "title" field.
func (buo *BlogUpdateOne) SetTitle(s string) *BlogUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetImages sets the "images" field.
func (buo *BlogUpdateOne) SetImages(s string) *BlogUpdateOne {
	buo.mutation.SetImages(s)
	return buo
}

// SetContent sets the "content" field.
func (buo *BlogUpdateOne) SetContent(s string) *BlogUpdateOne {
	buo.mutation.SetContent(s)
	return buo
}

// SetLiked sets the "liked" field.
func (buo *BlogUpdateOne) SetLiked(i int) *BlogUpdateOne {
	buo.mutation.ResetLiked()
	buo.mutation.SetLiked(i)
	return buo
}

// AddLiked adds i to the "liked" field.
func (buo *BlogUpdateOne) AddLiked(i int) *BlogUpdateOne {
	buo.mutation.AddLiked(i)
	return buo
}

// SetComments sets the "comments" field.
func (buo *BlogUpdateOne) SetComments(i int) *BlogUpdateOne {
	buo.mutation.ResetComments()
	buo.mutation.SetComments(i)
	return buo
}

// AddComments adds i to the "comments" field.
func (buo *BlogUpdateOne) AddComments(i int) *BlogUpdateOne {
	buo.mutation.AddComments(i)
	return buo
}

// SetCreateTime sets the "createTime" field.
func (buo *BlogUpdateOne) SetCreateTime(t time.Time) *BlogUpdateOne {
	buo.mutation.SetCreateTime(t)
	return buo
}

// SetUpdateTime sets the "updateTime" field.
func (buo *BlogUpdateOne) SetUpdateTime(t time.Time) *BlogUpdateOne {
	buo.mutation.SetUpdateTime(t)
	return buo
}

// Mutation returns the BlogMutation object of the builder.
func (buo *BlogUpdateOne) Mutation() *BlogMutation {
	return buo.mutation
}

// Where appends a list predicates to the BlogUpdate builder.
func (buo *BlogUpdateOne) Where(ps ...predicate.Blog) *BlogUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlogUpdateOne) Select(field string, fields ...string) *BlogUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Blog entity.
func (buo *BlogUpdateOne) Save(ctx context.Context) (*Blog, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlogUpdateOne) SaveX(ctx context.Context) *Blog {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlogUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlogUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BlogUpdateOne) sqlSave(ctx context.Context) (_node *Blog, err error) {
	_spec := sqlgraph.NewUpdateSpec(blog.Table, blog.Columns, sqlgraph.NewFieldSpec(blog.FieldID, field.TypeInt64))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Blog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blog.FieldID)
		for _, f := range fields {
			if !blog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.ShopId(); ok {
		_spec.SetField(blog.FieldShopId, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.AddedShopId(); ok {
		_spec.AddField(blog.FieldShopId, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.UserId(); ok {
		_spec.SetField(blog.FieldUserId, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.AddedUserId(); ok {
		_spec.AddField(blog.FieldUserId, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.SetField(blog.FieldTitle, field.TypeString, value)
	}
	if value, ok := buo.mutation.Images(); ok {
		_spec.SetField(blog.FieldImages, field.TypeString, value)
	}
	if value, ok := buo.mutation.Content(); ok {
		_spec.SetField(blog.FieldContent, field.TypeString, value)
	}
	if value, ok := buo.mutation.Liked(); ok {
		_spec.SetField(blog.FieldLiked, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedLiked(); ok {
		_spec.AddField(blog.FieldLiked, field.TypeInt, value)
	}
	if value, ok := buo.mutation.Comments(); ok {
		_spec.SetField(blog.FieldComments, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedComments(); ok {
		_spec.AddField(blog.FieldComments, field.TypeInt, value)
	}
	if value, ok := buo.mutation.CreateTime(); ok {
		_spec.SetField(blog.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := buo.mutation.UpdateTime(); ok {
		_spec.SetField(blog.FieldUpdateTime, field.TypeTime, value)
	}
	_node = &Blog{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
