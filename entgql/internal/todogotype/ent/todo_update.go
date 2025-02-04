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

	"entgo.io/contrib/entgql/internal/todo/ent/schema/customstruct"
	"entgo.io/contrib/entgql/internal/todogotype/ent/predicate"
	"entgo.io/contrib/entgql/internal/todogotype/ent/todo"
	"entgo.io/contrib/entgql/internal/todogotype/ent/verysecret"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// TodoUpdate is the builder for updating Todo entities.
type TodoUpdate struct {
	config
	hooks    []Hook
	mutation *TodoMutation
}

// Where appends a list predicates to the TodoUpdate builder.
func (tu *TodoUpdate) Where(ps ...predicate.Todo) *TodoUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetStatus sets the "status" field.
func (tu *TodoUpdate) SetStatus(t todo.Status) *TodoUpdate {
	tu.mutation.SetStatus(t)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableStatus(t *todo.Status) *TodoUpdate {
	if t != nil {
		tu.SetStatus(*t)
	}
	return tu
}

// SetPriority sets the "priority" field.
func (tu *TodoUpdate) SetPriority(i int) *TodoUpdate {
	tu.mutation.ResetPriority()
	tu.mutation.SetPriority(i)
	return tu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tu *TodoUpdate) SetNillablePriority(i *int) *TodoUpdate {
	if i != nil {
		tu.SetPriority(*i)
	}
	return tu
}

// AddPriority adds i to the "priority" field.
func (tu *TodoUpdate) AddPriority(i int) *TodoUpdate {
	tu.mutation.AddPriority(i)
	return tu
}

// SetText sets the "text" field.
func (tu *TodoUpdate) SetText(s string) *TodoUpdate {
	tu.mutation.SetText(s)
	return tu
}

// SetNillableText sets the "text" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableText(s *string) *TodoUpdate {
	if s != nil {
		tu.SetText(*s)
	}
	return tu
}

// SetBlob sets the "blob" field.
func (tu *TodoUpdate) SetBlob(b []byte) *TodoUpdate {
	tu.mutation.SetBlob(b)
	return tu
}

// ClearBlob clears the value of the "blob" field.
func (tu *TodoUpdate) ClearBlob() *TodoUpdate {
	tu.mutation.ClearBlob()
	return tu
}

// SetInit sets the "init" field.
func (tu *TodoUpdate) SetInit(m map[string]interface{}) *TodoUpdate {
	tu.mutation.SetInit(m)
	return tu
}

// ClearInit clears the value of the "init" field.
func (tu *TodoUpdate) ClearInit() *TodoUpdate {
	tu.mutation.ClearInit()
	return tu
}

// SetCustom sets the "custom" field.
func (tu *TodoUpdate) SetCustom(c []customstruct.Custom) *TodoUpdate {
	tu.mutation.SetCustom(c)
	return tu
}

// AppendCustom appends c to the "custom" field.
func (tu *TodoUpdate) AppendCustom(c []customstruct.Custom) *TodoUpdate {
	tu.mutation.AppendCustom(c)
	return tu
}

// ClearCustom clears the value of the "custom" field.
func (tu *TodoUpdate) ClearCustom() *TodoUpdate {
	tu.mutation.ClearCustom()
	return tu
}

// SetCustomp sets the "customp" field.
func (tu *TodoUpdate) SetCustomp(c []*customstruct.Custom) *TodoUpdate {
	tu.mutation.SetCustomp(c)
	return tu
}

// AppendCustomp appends c to the "customp" field.
func (tu *TodoUpdate) AppendCustomp(c []*customstruct.Custom) *TodoUpdate {
	tu.mutation.AppendCustomp(c)
	return tu
}

// ClearCustomp clears the value of the "customp" field.
func (tu *TodoUpdate) ClearCustomp() *TodoUpdate {
	tu.mutation.ClearCustomp()
	return tu
}

// SetValue sets the "value" field.
func (tu *TodoUpdate) SetValue(i int) *TodoUpdate {
	tu.mutation.ResetValue()
	tu.mutation.SetValue(i)
	return tu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableValue(i *int) *TodoUpdate {
	if i != nil {
		tu.SetValue(*i)
	}
	return tu
}

// AddValue adds i to the "value" field.
func (tu *TodoUpdate) AddValue(i int) *TodoUpdate {
	tu.mutation.AddValue(i)
	return tu
}

// SetParentID sets the "parent" edge to the Todo entity by ID.
func (tu *TodoUpdate) SetParentID(id string) *TodoUpdate {
	tu.mutation.SetParentID(id)
	return tu
}

// SetNillableParentID sets the "parent" edge to the Todo entity by ID if the given value is not nil.
func (tu *TodoUpdate) SetNillableParentID(id *string) *TodoUpdate {
	if id != nil {
		tu = tu.SetParentID(*id)
	}
	return tu
}

// SetParent sets the "parent" edge to the Todo entity.
func (tu *TodoUpdate) SetParent(t *Todo) *TodoUpdate {
	return tu.SetParentID(t.ID)
}

// AddChildIDs adds the "children" edge to the Todo entity by IDs.
func (tu *TodoUpdate) AddChildIDs(ids ...string) *TodoUpdate {
	tu.mutation.AddChildIDs(ids...)
	return tu
}

// AddChildren adds the "children" edges to the Todo entity.
func (tu *TodoUpdate) AddChildren(t ...*Todo) *TodoUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddChildIDs(ids...)
}

// SetSecretID sets the "secret" edge to the VerySecret entity by ID.
func (tu *TodoUpdate) SetSecretID(id string) *TodoUpdate {
	tu.mutation.SetSecretID(id)
	return tu
}

// SetNillableSecretID sets the "secret" edge to the VerySecret entity by ID if the given value is not nil.
func (tu *TodoUpdate) SetNillableSecretID(id *string) *TodoUpdate {
	if id != nil {
		tu = tu.SetSecretID(*id)
	}
	return tu
}

// SetSecret sets the "secret" edge to the VerySecret entity.
func (tu *TodoUpdate) SetSecret(v *VerySecret) *TodoUpdate {
	return tu.SetSecretID(v.ID)
}

// Mutation returns the TodoMutation object of the builder.
func (tu *TodoUpdate) Mutation() *TodoMutation {
	return tu.mutation
}

// ClearParent clears the "parent" edge to the Todo entity.
func (tu *TodoUpdate) ClearParent() *TodoUpdate {
	tu.mutation.ClearParent()
	return tu
}

// ClearChildren clears all "children" edges to the Todo entity.
func (tu *TodoUpdate) ClearChildren() *TodoUpdate {
	tu.mutation.ClearChildren()
	return tu
}

// RemoveChildIDs removes the "children" edge to Todo entities by IDs.
func (tu *TodoUpdate) RemoveChildIDs(ids ...string) *TodoUpdate {
	tu.mutation.RemoveChildIDs(ids...)
	return tu
}

// RemoveChildren removes "children" edges to Todo entities.
func (tu *TodoUpdate) RemoveChildren(t ...*Todo) *TodoUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveChildIDs(ids...)
}

// ClearSecret clears the "secret" edge to the VerySecret entity.
func (tu *TodoUpdate) ClearSecret() *TodoUpdate {
	tu.mutation.ClearSecret()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TodoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TodoUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TodoUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TodoUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TodoUpdate) check() error {
	if v, ok := tu.mutation.Status(); ok {
		if err := todo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Todo.status": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Text(); ok {
		if err := todo.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Todo.text": %w`, err)}
		}
	}
	return nil
}

func (tu *TodoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(todo.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.Priority(); ok {
		_spec.SetField(todo.FieldPriority, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedPriority(); ok {
		_spec.AddField(todo.FieldPriority, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Text(); ok {
		_spec.SetField(todo.FieldText, field.TypeString, value)
	}
	if value, ok := tu.mutation.Blob(); ok {
		_spec.SetField(todo.FieldBlob, field.TypeBytes, value)
	}
	if tu.mutation.BlobCleared() {
		_spec.ClearField(todo.FieldBlob, field.TypeBytes)
	}
	if value, ok := tu.mutation.Init(); ok {
		_spec.SetField(todo.FieldInit, field.TypeJSON, value)
	}
	if tu.mutation.InitCleared() {
		_spec.ClearField(todo.FieldInit, field.TypeJSON)
	}
	if value, ok := tu.mutation.Custom(); ok {
		_spec.SetField(todo.FieldCustom, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.AppendedCustom(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, todo.FieldCustom, value)
		})
	}
	if tu.mutation.CustomCleared() {
		_spec.ClearField(todo.FieldCustom, field.TypeJSON)
	}
	if value, ok := tu.mutation.Customp(); ok {
		_spec.SetField(todo.FieldCustomp, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.AppendedCustomp(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, todo.FieldCustomp, value)
		})
	}
	if tu.mutation.CustompCleared() {
		_spec.ClearField(todo.FieldCustomp, field.TypeJSON)
	}
	if value, ok := tu.mutation.Value(); ok {
		_spec.SetField(todo.FieldValue, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedValue(); ok {
		_spec.AddField(todo.FieldValue, field.TypeInt, value)
	}
	if tu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.ParentTable,
			Columns: []string{todo.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.ParentTable,
			Columns: []string{todo.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !tu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.SecretCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.SecretTable,
			Columns: []string{todo.SecretColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(verysecret.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.SecretIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.SecretTable,
			Columns: []string{todo.SecretColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(verysecret.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TodoUpdateOne is the builder for updating a single Todo entity.
type TodoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TodoMutation
}

// SetStatus sets the "status" field.
func (tuo *TodoUpdateOne) SetStatus(t todo.Status) *TodoUpdateOne {
	tuo.mutation.SetStatus(t)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableStatus(t *todo.Status) *TodoUpdateOne {
	if t != nil {
		tuo.SetStatus(*t)
	}
	return tuo
}

// SetPriority sets the "priority" field.
func (tuo *TodoUpdateOne) SetPriority(i int) *TodoUpdateOne {
	tuo.mutation.ResetPriority()
	tuo.mutation.SetPriority(i)
	return tuo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillablePriority(i *int) *TodoUpdateOne {
	if i != nil {
		tuo.SetPriority(*i)
	}
	return tuo
}

// AddPriority adds i to the "priority" field.
func (tuo *TodoUpdateOne) AddPriority(i int) *TodoUpdateOne {
	tuo.mutation.AddPriority(i)
	return tuo
}

// SetText sets the "text" field.
func (tuo *TodoUpdateOne) SetText(s string) *TodoUpdateOne {
	tuo.mutation.SetText(s)
	return tuo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableText(s *string) *TodoUpdateOne {
	if s != nil {
		tuo.SetText(*s)
	}
	return tuo
}

// SetBlob sets the "blob" field.
func (tuo *TodoUpdateOne) SetBlob(b []byte) *TodoUpdateOne {
	tuo.mutation.SetBlob(b)
	return tuo
}

// ClearBlob clears the value of the "blob" field.
func (tuo *TodoUpdateOne) ClearBlob() *TodoUpdateOne {
	tuo.mutation.ClearBlob()
	return tuo
}

// SetInit sets the "init" field.
func (tuo *TodoUpdateOne) SetInit(m map[string]interface{}) *TodoUpdateOne {
	tuo.mutation.SetInit(m)
	return tuo
}

// ClearInit clears the value of the "init" field.
func (tuo *TodoUpdateOne) ClearInit() *TodoUpdateOne {
	tuo.mutation.ClearInit()
	return tuo
}

// SetCustom sets the "custom" field.
func (tuo *TodoUpdateOne) SetCustom(c []customstruct.Custom) *TodoUpdateOne {
	tuo.mutation.SetCustom(c)
	return tuo
}

// AppendCustom appends c to the "custom" field.
func (tuo *TodoUpdateOne) AppendCustom(c []customstruct.Custom) *TodoUpdateOne {
	tuo.mutation.AppendCustom(c)
	return tuo
}

// ClearCustom clears the value of the "custom" field.
func (tuo *TodoUpdateOne) ClearCustom() *TodoUpdateOne {
	tuo.mutation.ClearCustom()
	return tuo
}

// SetCustomp sets the "customp" field.
func (tuo *TodoUpdateOne) SetCustomp(c []*customstruct.Custom) *TodoUpdateOne {
	tuo.mutation.SetCustomp(c)
	return tuo
}

// AppendCustomp appends c to the "customp" field.
func (tuo *TodoUpdateOne) AppendCustomp(c []*customstruct.Custom) *TodoUpdateOne {
	tuo.mutation.AppendCustomp(c)
	return tuo
}

// ClearCustomp clears the value of the "customp" field.
func (tuo *TodoUpdateOne) ClearCustomp() *TodoUpdateOne {
	tuo.mutation.ClearCustomp()
	return tuo
}

// SetValue sets the "value" field.
func (tuo *TodoUpdateOne) SetValue(i int) *TodoUpdateOne {
	tuo.mutation.ResetValue()
	tuo.mutation.SetValue(i)
	return tuo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableValue(i *int) *TodoUpdateOne {
	if i != nil {
		tuo.SetValue(*i)
	}
	return tuo
}

// AddValue adds i to the "value" field.
func (tuo *TodoUpdateOne) AddValue(i int) *TodoUpdateOne {
	tuo.mutation.AddValue(i)
	return tuo
}

// SetParentID sets the "parent" edge to the Todo entity by ID.
func (tuo *TodoUpdateOne) SetParentID(id string) *TodoUpdateOne {
	tuo.mutation.SetParentID(id)
	return tuo
}

// SetNillableParentID sets the "parent" edge to the Todo entity by ID if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableParentID(id *string) *TodoUpdateOne {
	if id != nil {
		tuo = tuo.SetParentID(*id)
	}
	return tuo
}

// SetParent sets the "parent" edge to the Todo entity.
func (tuo *TodoUpdateOne) SetParent(t *Todo) *TodoUpdateOne {
	return tuo.SetParentID(t.ID)
}

// AddChildIDs adds the "children" edge to the Todo entity by IDs.
func (tuo *TodoUpdateOne) AddChildIDs(ids ...string) *TodoUpdateOne {
	tuo.mutation.AddChildIDs(ids...)
	return tuo
}

// AddChildren adds the "children" edges to the Todo entity.
func (tuo *TodoUpdateOne) AddChildren(t ...*Todo) *TodoUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddChildIDs(ids...)
}

// SetSecretID sets the "secret" edge to the VerySecret entity by ID.
func (tuo *TodoUpdateOne) SetSecretID(id string) *TodoUpdateOne {
	tuo.mutation.SetSecretID(id)
	return tuo
}

// SetNillableSecretID sets the "secret" edge to the VerySecret entity by ID if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableSecretID(id *string) *TodoUpdateOne {
	if id != nil {
		tuo = tuo.SetSecretID(*id)
	}
	return tuo
}

// SetSecret sets the "secret" edge to the VerySecret entity.
func (tuo *TodoUpdateOne) SetSecret(v *VerySecret) *TodoUpdateOne {
	return tuo.SetSecretID(v.ID)
}

// Mutation returns the TodoMutation object of the builder.
func (tuo *TodoUpdateOne) Mutation() *TodoMutation {
	return tuo.mutation
}

// ClearParent clears the "parent" edge to the Todo entity.
func (tuo *TodoUpdateOne) ClearParent() *TodoUpdateOne {
	tuo.mutation.ClearParent()
	return tuo
}

// ClearChildren clears all "children" edges to the Todo entity.
func (tuo *TodoUpdateOne) ClearChildren() *TodoUpdateOne {
	tuo.mutation.ClearChildren()
	return tuo
}

// RemoveChildIDs removes the "children" edge to Todo entities by IDs.
func (tuo *TodoUpdateOne) RemoveChildIDs(ids ...string) *TodoUpdateOne {
	tuo.mutation.RemoveChildIDs(ids...)
	return tuo
}

// RemoveChildren removes "children" edges to Todo entities.
func (tuo *TodoUpdateOne) RemoveChildren(t ...*Todo) *TodoUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveChildIDs(ids...)
}

// ClearSecret clears the "secret" edge to the VerySecret entity.
func (tuo *TodoUpdateOne) ClearSecret() *TodoUpdateOne {
	tuo.mutation.ClearSecret()
	return tuo
}

// Where appends a list predicates to the TodoUpdate builder.
func (tuo *TodoUpdateOne) Where(ps ...predicate.Todo) *TodoUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TodoUpdateOne) Select(field string, fields ...string) *TodoUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Todo entity.
func (tuo *TodoUpdateOne) Save(ctx context.Context) (*Todo, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TodoUpdateOne) SaveX(ctx context.Context) *Todo {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TodoUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TodoUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TodoUpdateOne) check() error {
	if v, ok := tuo.mutation.Status(); ok {
		if err := todo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Todo.status": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Text(); ok {
		if err := todo.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Todo.text": %w`, err)}
		}
	}
	return nil
}

func (tuo *TodoUpdateOne) sqlSave(ctx context.Context) (_node *Todo, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Todo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todo.FieldID)
		for _, f := range fields {
			if !todo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != todo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(todo.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.Priority(); ok {
		_spec.SetField(todo.FieldPriority, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedPriority(); ok {
		_spec.AddField(todo.FieldPriority, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Text(); ok {
		_spec.SetField(todo.FieldText, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Blob(); ok {
		_spec.SetField(todo.FieldBlob, field.TypeBytes, value)
	}
	if tuo.mutation.BlobCleared() {
		_spec.ClearField(todo.FieldBlob, field.TypeBytes)
	}
	if value, ok := tuo.mutation.Init(); ok {
		_spec.SetField(todo.FieldInit, field.TypeJSON, value)
	}
	if tuo.mutation.InitCleared() {
		_spec.ClearField(todo.FieldInit, field.TypeJSON)
	}
	if value, ok := tuo.mutation.Custom(); ok {
		_spec.SetField(todo.FieldCustom, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.AppendedCustom(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, todo.FieldCustom, value)
		})
	}
	if tuo.mutation.CustomCleared() {
		_spec.ClearField(todo.FieldCustom, field.TypeJSON)
	}
	if value, ok := tuo.mutation.Customp(); ok {
		_spec.SetField(todo.FieldCustomp, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.AppendedCustomp(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, todo.FieldCustomp, value)
		})
	}
	if tuo.mutation.CustompCleared() {
		_spec.ClearField(todo.FieldCustomp, field.TypeJSON)
	}
	if value, ok := tuo.mutation.Value(); ok {
		_spec.SetField(todo.FieldValue, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedValue(); ok {
		_spec.AddField(todo.FieldValue, field.TypeInt, value)
	}
	if tuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.ParentTable,
			Columns: []string{todo.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.ParentTable,
			Columns: []string{todo.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !tuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.SecretCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.SecretTable,
			Columns: []string{todo.SecretColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(verysecret.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.SecretIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.SecretTable,
			Columns: []string{todo.SecretColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(verysecret.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Todo{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
