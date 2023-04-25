// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/contrib/entproto/internal/entprototest/ent/implicitskippedmessage"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImplicitSkippedMessageQuery is the builder for querying ImplicitSkippedMessage entities.
type ImplicitSkippedMessageQuery struct {
	config
	ctx        *QueryContext
	order      []implicitskippedmessage.OrderOption
	inters     []Interceptor
	predicates []predicate.ImplicitSkippedMessage
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ImplicitSkippedMessageQuery builder.
func (ismq *ImplicitSkippedMessageQuery) Where(ps ...predicate.ImplicitSkippedMessage) *ImplicitSkippedMessageQuery {
	ismq.predicates = append(ismq.predicates, ps...)
	return ismq
}

// Limit the number of records to be returned by this query.
func (ismq *ImplicitSkippedMessageQuery) Limit(limit int) *ImplicitSkippedMessageQuery {
	ismq.ctx.Limit = &limit
	return ismq
}

// Offset to start from.
func (ismq *ImplicitSkippedMessageQuery) Offset(offset int) *ImplicitSkippedMessageQuery {
	ismq.ctx.Offset = &offset
	return ismq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ismq *ImplicitSkippedMessageQuery) Unique(unique bool) *ImplicitSkippedMessageQuery {
	ismq.ctx.Unique = &unique
	return ismq
}

// Order specifies how the records should be ordered.
func (ismq *ImplicitSkippedMessageQuery) Order(o ...implicitskippedmessage.OrderOption) *ImplicitSkippedMessageQuery {
	ismq.order = append(ismq.order, o...)
	return ismq
}

// First returns the first ImplicitSkippedMessage entity from the query.
// Returns a *NotFoundError when no ImplicitSkippedMessage was found.
func (ismq *ImplicitSkippedMessageQuery) First(ctx context.Context) (*ImplicitSkippedMessage, error) {
	nodes, err := ismq.Limit(1).All(setContextOp(ctx, ismq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{implicitskippedmessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) FirstX(ctx context.Context) *ImplicitSkippedMessage {
	node, err := ismq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ImplicitSkippedMessage ID from the query.
// Returns a *NotFoundError when no ImplicitSkippedMessage ID was found.
func (ismq *ImplicitSkippedMessageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ismq.Limit(1).IDs(setContextOp(ctx, ismq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{implicitskippedmessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) FirstIDX(ctx context.Context) int {
	id, err := ismq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ImplicitSkippedMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ImplicitSkippedMessage entity is found.
// Returns a *NotFoundError when no ImplicitSkippedMessage entities are found.
func (ismq *ImplicitSkippedMessageQuery) Only(ctx context.Context) (*ImplicitSkippedMessage, error) {
	nodes, err := ismq.Limit(2).All(setContextOp(ctx, ismq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{implicitskippedmessage.Label}
	default:
		return nil, &NotSingularError{implicitskippedmessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) OnlyX(ctx context.Context) *ImplicitSkippedMessage {
	node, err := ismq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ImplicitSkippedMessage ID in the query.
// Returns a *NotSingularError when more than one ImplicitSkippedMessage ID is found.
// Returns a *NotFoundError when no entities are found.
func (ismq *ImplicitSkippedMessageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ismq.Limit(2).IDs(setContextOp(ctx, ismq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{implicitskippedmessage.Label}
	default:
		err = &NotSingularError{implicitskippedmessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) OnlyIDX(ctx context.Context) int {
	id, err := ismq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ImplicitSkippedMessages.
func (ismq *ImplicitSkippedMessageQuery) All(ctx context.Context) ([]*ImplicitSkippedMessage, error) {
	ctx = setContextOp(ctx, ismq.ctx, "All")
	if err := ismq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ImplicitSkippedMessage, *ImplicitSkippedMessageQuery]()
	return withInterceptors[[]*ImplicitSkippedMessage](ctx, ismq, qr, ismq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) AllX(ctx context.Context) []*ImplicitSkippedMessage {
	nodes, err := ismq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ImplicitSkippedMessage IDs.
func (ismq *ImplicitSkippedMessageQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ismq.ctx.Unique == nil && ismq.path != nil {
		ismq.Unique(true)
	}
	ctx = setContextOp(ctx, ismq.ctx, "IDs")
	if err = ismq.Select(implicitskippedmessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) IDsX(ctx context.Context) []int {
	ids, err := ismq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ismq *ImplicitSkippedMessageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ismq.ctx, "Count")
	if err := ismq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ismq, querierCount[*ImplicitSkippedMessageQuery](), ismq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) CountX(ctx context.Context) int {
	count, err := ismq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ismq *ImplicitSkippedMessageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ismq.ctx, "Exist")
	switch _, err := ismq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := ismq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ImplicitSkippedMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ismq *ImplicitSkippedMessageQuery) Clone() *ImplicitSkippedMessageQuery {
	if ismq == nil {
		return nil
	}
	return &ImplicitSkippedMessageQuery{
		config:     ismq.config,
		ctx:        ismq.ctx.Clone(),
		order:      append([]implicitskippedmessage.OrderOption{}, ismq.order...),
		inters:     append([]Interceptor{}, ismq.inters...),
		predicates: append([]predicate.ImplicitSkippedMessage{}, ismq.predicates...),
		// clone intermediate query.
		sql:  ismq.sql.Clone(),
		path: ismq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (ismq *ImplicitSkippedMessageQuery) GroupBy(field string, fields ...string) *ImplicitSkippedMessageGroupBy {
	ismq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ImplicitSkippedMessageGroupBy{build: ismq}
	grbuild.flds = &ismq.ctx.Fields
	grbuild.label = implicitskippedmessage.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (ismq *ImplicitSkippedMessageQuery) Select(fields ...string) *ImplicitSkippedMessageSelect {
	ismq.ctx.Fields = append(ismq.ctx.Fields, fields...)
	sbuild := &ImplicitSkippedMessageSelect{ImplicitSkippedMessageQuery: ismq}
	sbuild.label = implicitskippedmessage.Label
	sbuild.flds, sbuild.scan = &ismq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ImplicitSkippedMessageSelect configured with the given aggregations.
func (ismq *ImplicitSkippedMessageQuery) Aggregate(fns ...AggregateFunc) *ImplicitSkippedMessageSelect {
	return ismq.Select().Aggregate(fns...)
}

func (ismq *ImplicitSkippedMessageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ismq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ismq); err != nil {
				return err
			}
		}
	}
	for _, f := range ismq.ctx.Fields {
		if !implicitskippedmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ismq.path != nil {
		prev, err := ismq.path(ctx)
		if err != nil {
			return err
		}
		ismq.sql = prev
	}
	return nil
}

func (ismq *ImplicitSkippedMessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ImplicitSkippedMessage, error) {
	var (
		nodes   = []*ImplicitSkippedMessage{}
		withFKs = ismq.withFKs
		_spec   = ismq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, implicitskippedmessage.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ImplicitSkippedMessage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ImplicitSkippedMessage{config: ismq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ismq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ismq *ImplicitSkippedMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ismq.querySpec()
	_spec.Node.Columns = ismq.ctx.Fields
	if len(ismq.ctx.Fields) > 0 {
		_spec.Unique = ismq.ctx.Unique != nil && *ismq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ismq.driver, _spec)
}

func (ismq *ImplicitSkippedMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(implicitskippedmessage.Table, implicitskippedmessage.Columns, sqlgraph.NewFieldSpec(implicitskippedmessage.FieldID, field.TypeInt))
	_spec.From = ismq.sql
	if unique := ismq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ismq.path != nil {
		_spec.Unique = true
	}
	if fields := ismq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, implicitskippedmessage.FieldID)
		for i := range fields {
			if fields[i] != implicitskippedmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ismq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ismq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ismq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ismq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ismq *ImplicitSkippedMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ismq.driver.Dialect())
	t1 := builder.Table(implicitskippedmessage.Table)
	columns := ismq.ctx.Fields
	if len(columns) == 0 {
		columns = implicitskippedmessage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ismq.sql != nil {
		selector = ismq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ismq.ctx.Unique != nil && *ismq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ismq.predicates {
		p(selector)
	}
	for _, p := range ismq.order {
		p(selector)
	}
	if offset := ismq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ismq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ImplicitSkippedMessageGroupBy is the group-by builder for ImplicitSkippedMessage entities.
type ImplicitSkippedMessageGroupBy struct {
	selector
	build *ImplicitSkippedMessageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ismgb *ImplicitSkippedMessageGroupBy) Aggregate(fns ...AggregateFunc) *ImplicitSkippedMessageGroupBy {
	ismgb.fns = append(ismgb.fns, fns...)
	return ismgb
}

// Scan applies the selector query and scans the result into the given value.
func (ismgb *ImplicitSkippedMessageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ismgb.build.ctx, "GroupBy")
	if err := ismgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ImplicitSkippedMessageQuery, *ImplicitSkippedMessageGroupBy](ctx, ismgb.build, ismgb, ismgb.build.inters, v)
}

func (ismgb *ImplicitSkippedMessageGroupBy) sqlScan(ctx context.Context, root *ImplicitSkippedMessageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ismgb.fns))
	for _, fn := range ismgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ismgb.flds)+len(ismgb.fns))
		for _, f := range *ismgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ismgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ismgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ImplicitSkippedMessageSelect is the builder for selecting fields of ImplicitSkippedMessage entities.
type ImplicitSkippedMessageSelect struct {
	*ImplicitSkippedMessageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (isms *ImplicitSkippedMessageSelect) Aggregate(fns ...AggregateFunc) *ImplicitSkippedMessageSelect {
	isms.fns = append(isms.fns, fns...)
	return isms
}

// Scan applies the selector query and scans the result into the given value.
func (isms *ImplicitSkippedMessageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, isms.ctx, "Select")
	if err := isms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ImplicitSkippedMessageQuery, *ImplicitSkippedMessageSelect](ctx, isms.ImplicitSkippedMessageQuery, isms, isms.inters, v)
}

func (isms *ImplicitSkippedMessageSelect) sqlScan(ctx context.Context, root *ImplicitSkippedMessageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(isms.fns))
	for _, fn := range isms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*isms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := isms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
