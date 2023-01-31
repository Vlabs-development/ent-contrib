// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/contrib/entproto/internal/entprototest/ent/invalidfieldmessage"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InvalidFieldMessageQuery is the builder for querying InvalidFieldMessage entities.
type InvalidFieldMessageQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.InvalidFieldMessage
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InvalidFieldMessageQuery builder.
func (ifmq *InvalidFieldMessageQuery) Where(ps ...predicate.InvalidFieldMessage) *InvalidFieldMessageQuery {
	ifmq.predicates = append(ifmq.predicates, ps...)
	return ifmq
}

// Limit the number of records to be returned by this query.
func (ifmq *InvalidFieldMessageQuery) Limit(limit int) *InvalidFieldMessageQuery {
	ifmq.ctx.Limit = &limit
	return ifmq
}

// Offset to start from.
func (ifmq *InvalidFieldMessageQuery) Offset(offset int) *InvalidFieldMessageQuery {
	ifmq.ctx.Offset = &offset
	return ifmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ifmq *InvalidFieldMessageQuery) Unique(unique bool) *InvalidFieldMessageQuery {
	ifmq.ctx.Unique = &unique
	return ifmq
}

// Order specifies how the records should be ordered.
func (ifmq *InvalidFieldMessageQuery) Order(o ...OrderFunc) *InvalidFieldMessageQuery {
	ifmq.order = append(ifmq.order, o...)
	return ifmq
}

// First returns the first InvalidFieldMessage entity from the query.
// Returns a *NotFoundError when no InvalidFieldMessage was found.
func (ifmq *InvalidFieldMessageQuery) First(ctx context.Context) (*InvalidFieldMessage, error) {
	nodes, err := ifmq.Limit(1).All(setContextOp(ctx, ifmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{invalidfieldmessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ifmq *InvalidFieldMessageQuery) FirstX(ctx context.Context) *InvalidFieldMessage {
	node, err := ifmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first InvalidFieldMessage ID from the query.
// Returns a *NotFoundError when no InvalidFieldMessage ID was found.
func (ifmq *InvalidFieldMessageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ifmq.Limit(1).IDs(setContextOp(ctx, ifmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{invalidfieldmessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ifmq *InvalidFieldMessageQuery) FirstIDX(ctx context.Context) int {
	id, err := ifmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single InvalidFieldMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one InvalidFieldMessage entity is found.
// Returns a *NotFoundError when no InvalidFieldMessage entities are found.
func (ifmq *InvalidFieldMessageQuery) Only(ctx context.Context) (*InvalidFieldMessage, error) {
	nodes, err := ifmq.Limit(2).All(setContextOp(ctx, ifmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{invalidfieldmessage.Label}
	default:
		return nil, &NotSingularError{invalidfieldmessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ifmq *InvalidFieldMessageQuery) OnlyX(ctx context.Context) *InvalidFieldMessage {
	node, err := ifmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only InvalidFieldMessage ID in the query.
// Returns a *NotSingularError when more than one InvalidFieldMessage ID is found.
// Returns a *NotFoundError when no entities are found.
func (ifmq *InvalidFieldMessageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ifmq.Limit(2).IDs(setContextOp(ctx, ifmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{invalidfieldmessage.Label}
	default:
		err = &NotSingularError{invalidfieldmessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ifmq *InvalidFieldMessageQuery) OnlyIDX(ctx context.Context) int {
	id, err := ifmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of InvalidFieldMessages.
func (ifmq *InvalidFieldMessageQuery) All(ctx context.Context) ([]*InvalidFieldMessage, error) {
	ctx = setContextOp(ctx, ifmq.ctx, "All")
	if err := ifmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*InvalidFieldMessage, *InvalidFieldMessageQuery]()
	return withInterceptors[[]*InvalidFieldMessage](ctx, ifmq, qr, ifmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ifmq *InvalidFieldMessageQuery) AllX(ctx context.Context) []*InvalidFieldMessage {
	nodes, err := ifmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of InvalidFieldMessage IDs.
func (ifmq *InvalidFieldMessageQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, ifmq.ctx, "IDs")
	if err := ifmq.Select(invalidfieldmessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ifmq *InvalidFieldMessageQuery) IDsX(ctx context.Context) []int {
	ids, err := ifmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ifmq *InvalidFieldMessageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ifmq.ctx, "Count")
	if err := ifmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ifmq, querierCount[*InvalidFieldMessageQuery](), ifmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ifmq *InvalidFieldMessageQuery) CountX(ctx context.Context) int {
	count, err := ifmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ifmq *InvalidFieldMessageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ifmq.ctx, "Exist")
	switch _, err := ifmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ifmq *InvalidFieldMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := ifmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InvalidFieldMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ifmq *InvalidFieldMessageQuery) Clone() *InvalidFieldMessageQuery {
	if ifmq == nil {
		return nil
	}
	return &InvalidFieldMessageQuery{
		config:     ifmq.config,
		ctx:        ifmq.ctx.Clone(),
		order:      append([]OrderFunc{}, ifmq.order...),
		inters:     append([]Interceptor{}, ifmq.inters...),
		predicates: append([]predicate.InvalidFieldMessage{}, ifmq.predicates...),
		// clone intermediate query.
		sql:  ifmq.sql.Clone(),
		path: ifmq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		JSON *schema.SomeJSON `json:"json,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.InvalidFieldMessage.Query().
//		GroupBy(invalidfieldmessage.FieldJSON).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ifmq *InvalidFieldMessageQuery) GroupBy(field string, fields ...string) *InvalidFieldMessageGroupBy {
	ifmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InvalidFieldMessageGroupBy{build: ifmq}
	grbuild.flds = &ifmq.ctx.Fields
	grbuild.label = invalidfieldmessage.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		JSON *schema.SomeJSON `json:"json,omitempty"`
//	}
//
//	client.InvalidFieldMessage.Query().
//		Select(invalidfieldmessage.FieldJSON).
//		Scan(ctx, &v)
func (ifmq *InvalidFieldMessageQuery) Select(fields ...string) *InvalidFieldMessageSelect {
	ifmq.ctx.Fields = append(ifmq.ctx.Fields, fields...)
	sbuild := &InvalidFieldMessageSelect{InvalidFieldMessageQuery: ifmq}
	sbuild.label = invalidfieldmessage.Label
	sbuild.flds, sbuild.scan = &ifmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InvalidFieldMessageSelect configured with the given aggregations.
func (ifmq *InvalidFieldMessageQuery) Aggregate(fns ...AggregateFunc) *InvalidFieldMessageSelect {
	return ifmq.Select().Aggregate(fns...)
}

func (ifmq *InvalidFieldMessageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ifmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ifmq); err != nil {
				return err
			}
		}
	}
	for _, f := range ifmq.ctx.Fields {
		if !invalidfieldmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ifmq.path != nil {
		prev, err := ifmq.path(ctx)
		if err != nil {
			return err
		}
		ifmq.sql = prev
	}
	return nil
}

func (ifmq *InvalidFieldMessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*InvalidFieldMessage, error) {
	var (
		nodes = []*InvalidFieldMessage{}
		_spec = ifmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*InvalidFieldMessage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &InvalidFieldMessage{config: ifmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ifmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ifmq *InvalidFieldMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ifmq.querySpec()
	_spec.Node.Columns = ifmq.ctx.Fields
	if len(ifmq.ctx.Fields) > 0 {
		_spec.Unique = ifmq.ctx.Unique != nil && *ifmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ifmq.driver, _spec)
}

func (ifmq *InvalidFieldMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   invalidfieldmessage.Table,
			Columns: invalidfieldmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: invalidfieldmessage.FieldID,
			},
		},
		From:   ifmq.sql,
		Unique: true,
	}
	if unique := ifmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ifmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invalidfieldmessage.FieldID)
		for i := range fields {
			if fields[i] != invalidfieldmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ifmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ifmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ifmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ifmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ifmq *InvalidFieldMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ifmq.driver.Dialect())
	t1 := builder.Table(invalidfieldmessage.Table)
	columns := ifmq.ctx.Fields
	if len(columns) == 0 {
		columns = invalidfieldmessage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ifmq.sql != nil {
		selector = ifmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ifmq.ctx.Unique != nil && *ifmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ifmq.predicates {
		p(selector)
	}
	for _, p := range ifmq.order {
		p(selector)
	}
	if offset := ifmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ifmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InvalidFieldMessageGroupBy is the group-by builder for InvalidFieldMessage entities.
type InvalidFieldMessageGroupBy struct {
	selector
	build *InvalidFieldMessageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ifmgb *InvalidFieldMessageGroupBy) Aggregate(fns ...AggregateFunc) *InvalidFieldMessageGroupBy {
	ifmgb.fns = append(ifmgb.fns, fns...)
	return ifmgb
}

// Scan applies the selector query and scans the result into the given value.
func (ifmgb *InvalidFieldMessageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ifmgb.build.ctx, "GroupBy")
	if err := ifmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvalidFieldMessageQuery, *InvalidFieldMessageGroupBy](ctx, ifmgb.build, ifmgb, ifmgb.build.inters, v)
}

func (ifmgb *InvalidFieldMessageGroupBy) sqlScan(ctx context.Context, root *InvalidFieldMessageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ifmgb.fns))
	for _, fn := range ifmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ifmgb.flds)+len(ifmgb.fns))
		for _, f := range *ifmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ifmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ifmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InvalidFieldMessageSelect is the builder for selecting fields of InvalidFieldMessage entities.
type InvalidFieldMessageSelect struct {
	*InvalidFieldMessageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ifms *InvalidFieldMessageSelect) Aggregate(fns ...AggregateFunc) *InvalidFieldMessageSelect {
	ifms.fns = append(ifms.fns, fns...)
	return ifms
}

// Scan applies the selector query and scans the result into the given value.
func (ifms *InvalidFieldMessageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ifms.ctx, "Select")
	if err := ifms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvalidFieldMessageQuery, *InvalidFieldMessageSelect](ctx, ifms.InvalidFieldMessageQuery, ifms, ifms.inters, v)
}

func (ifms *InvalidFieldMessageSelect) sqlScan(ctx context.Context, root *InvalidFieldMessageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ifms.fns))
	for _, fn := range ifms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ifms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ifms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
