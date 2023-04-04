// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/contrib/entgql/internal/todo/ent/category"
	"entgo.io/contrib/entgql/internal/todo/ent/schema/schematype"
	"entgo.io/contrib/entgql/internal/todo/ent/todo"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryCreate is the builder for creating a Category entity.
type CategoryCreate struct {
	config
	mutation *CategoryMutation
	hooks    []Hook
}

// SetText sets the "text" field.
func (cc *CategoryCreate) SetText(s string) *CategoryCreate {
	cc.mutation.SetText(s)
	return cc
}

// SetStatus sets the "status" field.
func (cc *CategoryCreate) SetStatus(c category.Status) *CategoryCreate {
	cc.mutation.SetStatus(c)
	return cc
}

// SetConfig sets the "config" field.
func (cc *CategoryCreate) SetConfig(sc *schematype.CategoryConfig) *CategoryCreate {
	cc.mutation.SetConfig(sc)
	return cc
}

// SetDuration sets the "duration" field.
func (cc *CategoryCreate) SetDuration(t time.Duration) *CategoryCreate {
	cc.mutation.SetDuration(t)
	return cc
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableDuration(t *time.Duration) *CategoryCreate {
	if t != nil {
		cc.SetDuration(*t)
	}
	return cc
}

// SetCount sets the "count" field.
func (cc *CategoryCreate) SetCount(u uint64) *CategoryCreate {
	cc.mutation.SetCount(u)
	return cc
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableCount(u *uint64) *CategoryCreate {
	if u != nil {
		cc.SetCount(*u)
	}
	return cc
}

// SetStrings sets the "strings" field.
func (cc *CategoryCreate) SetStrings(s []string) *CategoryCreate {
	cc.mutation.SetStrings(s)
	return cc
}

// SetID sets the "id" field.
func (cc *CategoryCreate) SetID(i int) *CategoryCreate {
	cc.mutation.SetID(i)
	return cc
}

// AddTodoIDs adds the "todos" edge to the Todo entity by IDs.
func (cc *CategoryCreate) AddTodoIDs(ids ...int) *CategoryCreate {
	cc.mutation.AddTodoIDs(ids...)
	return cc
}

// AddTodos adds the "todos" edges to the Todo entity.
func (cc *CategoryCreate) AddTodos(t ...*Todo) *CategoryCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cc.AddTodoIDs(ids...)
}

// AddSubCategoryIDs adds the "sub_categories" edge to the Category entity by IDs.
func (cc *CategoryCreate) AddSubCategoryIDs(ids ...int) *CategoryCreate {
	cc.mutation.AddSubCategoryIDs(ids...)
	return cc
}

// AddSubCategories adds the "sub_categories" edges to the Category entity.
func (cc *CategoryCreate) AddSubCategories(c ...*Category) *CategoryCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddSubCategoryIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cc *CategoryCreate) Mutation() *CategoryMutation {
	return cc.mutation
}

// Save creates the Category in the database.
func (cc *CategoryCreate) Save(ctx context.Context) (*Category, error) {
	return withHooks[*Category, CategoryMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CategoryCreate) SaveX(ctx context.Context) *Category {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CategoryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CategoryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CategoryCreate) check() error {
	if _, ok := cc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Category.text"`)}
	}
	if v, ok := cc.mutation.Text(); ok {
		if err := category.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Category.text": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Category.status"`)}
	}
	if v, ok := cc.mutation.Status(); ok {
		if err := category.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Category.status": %w`, err)}
		}
	}
	return nil
}

func (cc *CategoryCreate) sqlSave(ctx context.Context) (*Category, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CategoryCreate) createSpec() (*Category, *sqlgraph.CreateSpec) {
	var (
		_node = &Category{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(category.Table, sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Text(); ok {
		_spec.SetField(category.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(category.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := cc.mutation.Config(); ok {
		_spec.SetField(category.FieldConfig, field.TypeOther, value)
		_node.Config = value
	}
	if value, ok := cc.mutation.Duration(); ok {
		_spec.SetField(category.FieldDuration, field.TypeInt64, value)
		_node.Duration = value
	}
	if value, ok := cc.mutation.Count(); ok {
		_spec.SetField(category.FieldCount, field.TypeUint64, value)
		_node.Count = value
	}
	if value, ok := cc.mutation.Strings(); ok {
		_spec.SetField(category.FieldStrings, field.TypeJSON, value)
		_node.Strings = value
	}
	if nodes := cc.mutation.TodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.TodosTable,
			Columns: []string{category.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.SubCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.SubCategoriesTable,
			Columns: category.SubCategoriesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CategoryCreateBulk is the builder for creating many Category entities in bulk.
type CategoryCreateBulk struct {
	config
	builders []*CategoryCreate
}

// Save creates the Category entities in the database.
func (ccb *CategoryCreateBulk) Save(ctx context.Context) ([]*Category, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Category, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CategoryCreateBulk) SaveX(ctx context.Context) []*Category {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CategoryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
