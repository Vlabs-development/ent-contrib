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
	"fmt"
	"math"

	"entgo.io/contrib/entgql/internal/todogotype/ent/predicate"
	"entgo.io/contrib/entgql/internal/todogotype/ent/verysecret"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VerySecretQuery is the builder for querying VerySecret entities.
type VerySecretQuery struct {
	config
	ctx        *QueryContext
	order      []verysecret.OrderOption
	inters     []Interceptor
	predicates []predicate.VerySecret
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*VerySecret) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VerySecretQuery builder.
func (vsq *VerySecretQuery) Where(ps ...predicate.VerySecret) *VerySecretQuery {
	vsq.predicates = append(vsq.predicates, ps...)
	return vsq
}

// Limit the number of records to be returned by this query.
func (vsq *VerySecretQuery) Limit(limit int) *VerySecretQuery {
	vsq.ctx.Limit = &limit
	return vsq
}

// Offset to start from.
func (vsq *VerySecretQuery) Offset(offset int) *VerySecretQuery {
	vsq.ctx.Offset = &offset
	return vsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vsq *VerySecretQuery) Unique(unique bool) *VerySecretQuery {
	vsq.ctx.Unique = &unique
	return vsq
}

// Order specifies how the records should be ordered.
func (vsq *VerySecretQuery) Order(o ...verysecret.OrderOption) *VerySecretQuery {
	vsq.order = append(vsq.order, o...)
	return vsq
}

// First returns the first VerySecret entity from the query.
// Returns a *NotFoundError when no VerySecret was found.
func (vsq *VerySecretQuery) First(ctx context.Context) (*VerySecret, error) {
	nodes, err := vsq.Limit(1).All(setContextOp(ctx, vsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{verysecret.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vsq *VerySecretQuery) FirstX(ctx context.Context) *VerySecret {
	node, err := vsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VerySecret ID from the query.
// Returns a *NotFoundError when no VerySecret ID was found.
func (vsq *VerySecretQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = vsq.Limit(1).IDs(setContextOp(ctx, vsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{verysecret.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vsq *VerySecretQuery) FirstIDX(ctx context.Context) string {
	id, err := vsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VerySecret entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VerySecret entity is found.
// Returns a *NotFoundError when no VerySecret entities are found.
func (vsq *VerySecretQuery) Only(ctx context.Context) (*VerySecret, error) {
	nodes, err := vsq.Limit(2).All(setContextOp(ctx, vsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{verysecret.Label}
	default:
		return nil, &NotSingularError{verysecret.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vsq *VerySecretQuery) OnlyX(ctx context.Context) *VerySecret {
	node, err := vsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VerySecret ID in the query.
// Returns a *NotSingularError when more than one VerySecret ID is found.
// Returns a *NotFoundError when no entities are found.
func (vsq *VerySecretQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = vsq.Limit(2).IDs(setContextOp(ctx, vsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{verysecret.Label}
	default:
		err = &NotSingularError{verysecret.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vsq *VerySecretQuery) OnlyIDX(ctx context.Context) string {
	id, err := vsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VerySecrets.
func (vsq *VerySecretQuery) All(ctx context.Context) ([]*VerySecret, error) {
	ctx = setContextOp(ctx, vsq.ctx, "All")
	if err := vsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VerySecret, *VerySecretQuery]()
	return withInterceptors[[]*VerySecret](ctx, vsq, qr, vsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vsq *VerySecretQuery) AllX(ctx context.Context) []*VerySecret {
	nodes, err := vsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VerySecret IDs.
func (vsq *VerySecretQuery) IDs(ctx context.Context) (ids []string, err error) {
	if vsq.ctx.Unique == nil && vsq.path != nil {
		vsq.Unique(true)
	}
	ctx = setContextOp(ctx, vsq.ctx, "IDs")
	if err = vsq.Select(verysecret.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vsq *VerySecretQuery) IDsX(ctx context.Context) []string {
	ids, err := vsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vsq *VerySecretQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vsq.ctx, "Count")
	if err := vsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vsq, querierCount[*VerySecretQuery](), vsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vsq *VerySecretQuery) CountX(ctx context.Context) int {
	count, err := vsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vsq *VerySecretQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vsq.ctx, "Exist")
	switch _, err := vsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vsq *VerySecretQuery) ExistX(ctx context.Context) bool {
	exist, err := vsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VerySecretQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vsq *VerySecretQuery) Clone() *VerySecretQuery {
	if vsq == nil {
		return nil
	}
	return &VerySecretQuery{
		config:     vsq.config,
		ctx:        vsq.ctx.Clone(),
		order:      append([]verysecret.OrderOption{}, vsq.order...),
		inters:     append([]Interceptor{}, vsq.inters...),
		predicates: append([]predicate.VerySecret{}, vsq.predicates...),
		// clone intermediate query.
		sql:  vsq.sql.Clone(),
		path: vsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Password string `json:"password,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VerySecret.Query().
//		GroupBy(verysecret.FieldPassword).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vsq *VerySecretQuery) GroupBy(field string, fields ...string) *VerySecretGroupBy {
	vsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VerySecretGroupBy{build: vsq}
	grbuild.flds = &vsq.ctx.Fields
	grbuild.label = verysecret.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Password string `json:"password,omitempty"`
//	}
//
//	client.VerySecret.Query().
//		Select(verysecret.FieldPassword).
//		Scan(ctx, &v)
func (vsq *VerySecretQuery) Select(fields ...string) *VerySecretSelect {
	vsq.ctx.Fields = append(vsq.ctx.Fields, fields...)
	sbuild := &VerySecretSelect{VerySecretQuery: vsq}
	sbuild.label = verysecret.Label
	sbuild.flds, sbuild.scan = &vsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VerySecretSelect configured with the given aggregations.
func (vsq *VerySecretQuery) Aggregate(fns ...AggregateFunc) *VerySecretSelect {
	return vsq.Select().Aggregate(fns...)
}

func (vsq *VerySecretQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vsq); err != nil {
				return err
			}
		}
	}
	for _, f := range vsq.ctx.Fields {
		if !verysecret.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vsq.path != nil {
		prev, err := vsq.path(ctx)
		if err != nil {
			return err
		}
		vsq.sql = prev
	}
	return nil
}

func (vsq *VerySecretQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VerySecret, error) {
	var (
		nodes = []*VerySecret{}
		_spec = vsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VerySecret).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VerySecret{config: vsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(vsq.modifiers) > 0 {
		_spec.Modifiers = vsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range vsq.loadTotal {
		if err := vsq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vsq *VerySecretQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vsq.querySpec()
	if len(vsq.modifiers) > 0 {
		_spec.Modifiers = vsq.modifiers
	}
	_spec.Node.Columns = vsq.ctx.Fields
	if len(vsq.ctx.Fields) > 0 {
		_spec.Unique = vsq.ctx.Unique != nil && *vsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vsq.driver, _spec)
}

func (vsq *VerySecretQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(verysecret.Table, verysecret.Columns, sqlgraph.NewFieldSpec(verysecret.FieldID, field.TypeString))
	_spec.From = vsq.sql
	if unique := vsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vsq.path != nil {
		_spec.Unique = true
	}
	if fields := vsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, verysecret.FieldID)
		for i := range fields {
			if fields[i] != verysecret.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vsq *VerySecretQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vsq.driver.Dialect())
	t1 := builder.Table(verysecret.Table)
	columns := vsq.ctx.Fields
	if len(columns) == 0 {
		columns = verysecret.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vsq.sql != nil {
		selector = vsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vsq.ctx.Unique != nil && *vsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vsq.predicates {
		p(selector)
	}
	for _, p := range vsq.order {
		p(selector)
	}
	if offset := vsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VerySecretGroupBy is the group-by builder for VerySecret entities.
type VerySecretGroupBy struct {
	selector
	build *VerySecretQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vsgb *VerySecretGroupBy) Aggregate(fns ...AggregateFunc) *VerySecretGroupBy {
	vsgb.fns = append(vsgb.fns, fns...)
	return vsgb
}

// Scan applies the selector query and scans the result into the given value.
func (vsgb *VerySecretGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vsgb.build.ctx, "GroupBy")
	if err := vsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VerySecretQuery, *VerySecretGroupBy](ctx, vsgb.build, vsgb, vsgb.build.inters, v)
}

func (vsgb *VerySecretGroupBy) sqlScan(ctx context.Context, root *VerySecretQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vsgb.fns))
	for _, fn := range vsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vsgb.flds)+len(vsgb.fns))
		for _, f := range *vsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VerySecretSelect is the builder for selecting fields of VerySecret entities.
type VerySecretSelect struct {
	*VerySecretQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vss *VerySecretSelect) Aggregate(fns ...AggregateFunc) *VerySecretSelect {
	vss.fns = append(vss.fns, fns...)
	return vss
}

// Scan applies the selector query and scans the result into the given value.
func (vss *VerySecretSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vss.ctx, "Select")
	if err := vss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VerySecretQuery, *VerySecretSelect](ctx, vss.VerySecretQuery, vss, vss.inters, v)
}

func (vss *VerySecretSelect) sqlScan(ctx context.Context, root *VerySecretQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vss.fns))
	for _, fn := range vss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
