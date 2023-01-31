// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/contrib/entproto/internal/entprototest/ent/skipedgeexample"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SkipEdgeExampleDelete is the builder for deleting a SkipEdgeExample entity.
type SkipEdgeExampleDelete struct {
	config
	hooks    []Hook
	mutation *SkipEdgeExampleMutation
}

// Where appends a list predicates to the SkipEdgeExampleDelete builder.
func (seed *SkipEdgeExampleDelete) Where(ps ...predicate.SkipEdgeExample) *SkipEdgeExampleDelete {
	seed.mutation.Where(ps...)
	return seed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (seed *SkipEdgeExampleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, SkipEdgeExampleMutation](ctx, seed.sqlExec, seed.mutation, seed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (seed *SkipEdgeExampleDelete) ExecX(ctx context.Context) int {
	n, err := seed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (seed *SkipEdgeExampleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: skipedgeexample.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skipedgeexample.FieldID,
			},
		},
	}
	if ps := seed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, seed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	seed.mutation.done = true
	return affected, err
}

// SkipEdgeExampleDeleteOne is the builder for deleting a single SkipEdgeExample entity.
type SkipEdgeExampleDeleteOne struct {
	seed *SkipEdgeExampleDelete
}

// Exec executes the deletion query.
func (seedo *SkipEdgeExampleDeleteOne) Exec(ctx context.Context) error {
	n, err := seedo.seed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{skipedgeexample.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (seedo *SkipEdgeExampleDeleteOne) ExecX(ctx context.Context) {
	seedo.seed.ExecX(ctx)
}
