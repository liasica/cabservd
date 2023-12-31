// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/auroraride/cabservd/internal/ent/bin"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/ent/console"
	"github.com/auroraride/cabservd/internal/ent/predicate"
)

// ConsoleQuery is the builder for querying Console entities.
type ConsoleQuery struct {
	config
	ctx         *QueryContext
	order       []console.OrderOption
	inters      []Interceptor
	predicates  []predicate.Console
	withCabinet *CabinetQuery
	withBin     *BinQuery
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ConsoleQuery builder.
func (cq *ConsoleQuery) Where(ps ...predicate.Console) *ConsoleQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ConsoleQuery) Limit(limit int) *ConsoleQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ConsoleQuery) Offset(offset int) *ConsoleQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ConsoleQuery) Unique(unique bool) *ConsoleQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ConsoleQuery) Order(o ...console.OrderOption) *ConsoleQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryCabinet chains the current query on the "cabinet" edge.
func (cq *ConsoleQuery) QueryCabinet() *CabinetQuery {
	query := (&CabinetClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(console.Table, console.FieldID, selector),
			sqlgraph.To(cabinet.Table, cabinet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, console.CabinetTable, console.CabinetColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBin chains the current query on the "bin" edge.
func (cq *ConsoleQuery) QueryBin() *BinQuery {
	query := (&BinClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(console.Table, console.FieldID, selector),
			sqlgraph.To(bin.Table, bin.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, console.BinTable, console.BinColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Console entity from the query.
// Returns a *NotFoundError when no Console was found.
func (cq *ConsoleQuery) First(ctx context.Context) (*Console, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{console.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ConsoleQuery) FirstX(ctx context.Context) *Console {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Console ID from the query.
// Returns a *NotFoundError when no Console ID was found.
func (cq *ConsoleQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{console.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ConsoleQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Console entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Console entity is found.
// Returns a *NotFoundError when no Console entities are found.
func (cq *ConsoleQuery) Only(ctx context.Context) (*Console, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{console.Label}
	default:
		return nil, &NotSingularError{console.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ConsoleQuery) OnlyX(ctx context.Context) *Console {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Console ID in the query.
// Returns a *NotSingularError when more than one Console ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ConsoleQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{console.Label}
	default:
		err = &NotSingularError{console.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ConsoleQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Consoles.
func (cq *ConsoleQuery) All(ctx context.Context) ([]*Console, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Console, *ConsoleQuery]()
	return withInterceptors[[]*Console](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ConsoleQuery) AllX(ctx context.Context) []*Console {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Console IDs.
func (cq *ConsoleQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(console.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ConsoleQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ConsoleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ConsoleQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ConsoleQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ConsoleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ConsoleQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ConsoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ConsoleQuery) Clone() *ConsoleQuery {
	if cq == nil {
		return nil
	}
	return &ConsoleQuery{
		config:      cq.config,
		ctx:         cq.ctx.Clone(),
		order:       append([]console.OrderOption{}, cq.order...),
		inters:      append([]Interceptor{}, cq.inters...),
		predicates:  append([]predicate.Console{}, cq.predicates...),
		withCabinet: cq.withCabinet.Clone(),
		withBin:     cq.withBin.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithCabinet tells the query-builder to eager-load the nodes that are connected to
// the "cabinet" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConsoleQuery) WithCabinet(opts ...func(*CabinetQuery)) *ConsoleQuery {
	query := (&CabinetClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCabinet = query
	return cq
}

// WithBin tells the query-builder to eager-load the nodes that are connected to
// the "bin" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConsoleQuery) WithBin(opts ...func(*BinQuery)) *ConsoleQuery {
	query := (&BinClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withBin = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CabinetID uint64 `json:"cabinet_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Console.Query().
//		GroupBy(console.FieldCabinetID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ConsoleQuery) GroupBy(field string, fields ...string) *ConsoleGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ConsoleGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = console.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CabinetID uint64 `json:"cabinet_id,omitempty"`
//	}
//
//	client.Console.Query().
//		Select(console.FieldCabinetID).
//		Scan(ctx, &v)
func (cq *ConsoleQuery) Select(fields ...string) *ConsoleSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ConsoleSelect{ConsoleQuery: cq}
	sbuild.label = console.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ConsoleSelect configured with the given aggregations.
func (cq *ConsoleQuery) Aggregate(fns ...AggregateFunc) *ConsoleSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ConsoleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !console.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ConsoleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Console, error) {
	var (
		nodes       = []*Console{}
		_spec       = cq.querySpec()
		loadedTypes = [2]bool{
			cq.withCabinet != nil,
			cq.withBin != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Console).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Console{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withCabinet; query != nil {
		if err := cq.loadCabinet(ctx, query, nodes, nil,
			func(n *Console, e *Cabinet) { n.Edges.Cabinet = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withBin; query != nil {
		if err := cq.loadBin(ctx, query, nodes, nil,
			func(n *Console, e *Bin) { n.Edges.Bin = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ConsoleQuery) loadCabinet(ctx context.Context, query *CabinetQuery, nodes []*Console, init func(*Console), assign func(*Console, *Cabinet)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*Console)
	for i := range nodes {
		fk := nodes[i].CabinetID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(cabinet.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "cabinet_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ConsoleQuery) loadBin(ctx context.Context, query *BinQuery, nodes []*Console, init func(*Console), assign func(*Console, *Bin)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*Console)
	for i := range nodes {
		if nodes[i].BinID == nil {
			continue
		}
		fk := *nodes[i].BinID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(bin.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "bin_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cq *ConsoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ConsoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(console.Table, console.Columns, sqlgraph.NewFieldSpec(console.FieldID, field.TypeUint64))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, console.FieldID)
		for i := range fields {
			if fields[i] != console.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if cq.withCabinet != nil {
			_spec.Node.AddColumnOnce(console.FieldCabinetID)
		}
		if cq.withBin != nil {
			_spec.Node.AddColumnOnce(console.FieldBinID)
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ConsoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(console.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = console.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cq.modifiers {
		m(selector)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cq *ConsoleQuery) Modify(modifiers ...func(s *sql.Selector)) *ConsoleSelect {
	cq.modifiers = append(cq.modifiers, modifiers...)
	return cq.Select()
}

type ConsoleQueryWith string

var (
	ConsoleQueryWithCabinet ConsoleQueryWith = "Cabinet"
	ConsoleQueryWithBin     ConsoleQueryWith = "Bin"
)

func (cq *ConsoleQuery) With(withEdges ...ConsoleQueryWith) *ConsoleQuery {
	for _, v := range withEdges {
		switch v {
		case ConsoleQueryWithCabinet:
			cq.WithCabinet()
		case ConsoleQueryWithBin:
			cq.WithBin()
		}
	}
	return cq
}

// ConsoleGroupBy is the group-by builder for Console entities.
type ConsoleGroupBy struct {
	selector
	build *ConsoleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ConsoleGroupBy) Aggregate(fns ...AggregateFunc) *ConsoleGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ConsoleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConsoleQuery, *ConsoleGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ConsoleGroupBy) sqlScan(ctx context.Context, root *ConsoleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ConsoleSelect is the builder for selecting fields of Console entities.
type ConsoleSelect struct {
	*ConsoleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ConsoleSelect) Aggregate(fns ...AggregateFunc) *ConsoleSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ConsoleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConsoleQuery, *ConsoleSelect](ctx, cs.ConsoleQuery, cs, cs.inters, v)
}

func (cs *ConsoleSelect) sqlScan(ctx context.Context, root *ConsoleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cs *ConsoleSelect) Modify(modifiers ...func(s *sql.Selector)) *ConsoleSelect {
	cs.modifiers = append(cs.modifiers, modifiers...)
	return cs
}
