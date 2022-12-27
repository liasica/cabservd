// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/auroraride/cabservd/internal/ent/bin"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/ent/predicate"
)

// CabinetQuery is the builder for querying Cabinet entities.
type CabinetQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.Cabinet
	withBins   *BinQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CabinetQuery builder.
func (cq *CabinetQuery) Where(ps ...predicate.Cabinet) *CabinetQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CabinetQuery) Limit(limit int) *CabinetQuery {
	cq.limit = &limit
	return cq
}

// Offset to start from.
func (cq *CabinetQuery) Offset(offset int) *CabinetQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CabinetQuery) Unique(unique bool) *CabinetQuery {
	cq.unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CabinetQuery) Order(o ...OrderFunc) *CabinetQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryBins chains the current query on the "bins" edge.
func (cq *CabinetQuery) QueryBins() *BinQuery {
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
			sqlgraph.From(cabinet.Table, cabinet.FieldID, selector),
			sqlgraph.To(bin.Table, bin.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, cabinet.BinsTable, cabinet.BinsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Cabinet entity from the query.
// Returns a *NotFoundError when no Cabinet was found.
func (cq *CabinetQuery) First(ctx context.Context) (*Cabinet, error) {
	nodes, err := cq.Limit(1).All(newQueryContext(ctx, TypeCabinet, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cabinet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CabinetQuery) FirstX(ctx context.Context) *Cabinet {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Cabinet ID from the query.
// Returns a *NotFoundError when no Cabinet ID was found.
func (cq *CabinetQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = cq.Limit(1).IDs(newQueryContext(ctx, TypeCabinet, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cabinet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CabinetQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Cabinet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Cabinet entity is found.
// Returns a *NotFoundError when no Cabinet entities are found.
func (cq *CabinetQuery) Only(ctx context.Context) (*Cabinet, error) {
	nodes, err := cq.Limit(2).All(newQueryContext(ctx, TypeCabinet, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cabinet.Label}
	default:
		return nil, &NotSingularError{cabinet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CabinetQuery) OnlyX(ctx context.Context) *Cabinet {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Cabinet ID in the query.
// Returns a *NotSingularError when more than one Cabinet ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CabinetQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = cq.Limit(2).IDs(newQueryContext(ctx, TypeCabinet, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cabinet.Label}
	default:
		err = &NotSingularError{cabinet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CabinetQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Cabinets.
func (cq *CabinetQuery) All(ctx context.Context) ([]*Cabinet, error) {
	ctx = newQueryContext(ctx, TypeCabinet, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Cabinet, *CabinetQuery]()
	return withInterceptors[[]*Cabinet](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CabinetQuery) AllX(ctx context.Context) []*Cabinet {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Cabinet IDs.
func (cq *CabinetQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	ctx = newQueryContext(ctx, TypeCabinet, "IDs")
	if err := cq.Select(cabinet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CabinetQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CabinetQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeCabinet, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CabinetQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CabinetQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CabinetQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeCabinet, "Exist")
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
func (cq *CabinetQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CabinetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CabinetQuery) Clone() *CabinetQuery {
	if cq == nil {
		return nil
	}
	return &CabinetQuery{
		config:     cq.config,
		limit:      cq.limit,
		offset:     cq.offset,
		order:      append([]OrderFunc{}, cq.order...),
		predicates: append([]predicate.Cabinet{}, cq.predicates...),
		withBins:   cq.withBins.Clone(),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
	}
}

// WithBins tells the query-builder to eager-load the nodes that are connected to
// the "bins" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CabinetQuery) WithBins(opts ...func(*BinQuery)) *CabinetQuery {
	query := (&BinClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withBins = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Cabinet.Query().
//		GroupBy(cabinet.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CabinetQuery) GroupBy(field string, fields ...string) *CabinetGroupBy {
	cq.fields = append([]string{field}, fields...)
	grbuild := &CabinetGroupBy{build: cq}
	grbuild.flds = &cq.fields
	grbuild.label = cabinet.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Cabinet.Query().
//		Select(cabinet.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *CabinetQuery) Select(fields ...string) *CabinetSelect {
	cq.fields = append(cq.fields, fields...)
	sbuild := &CabinetSelect{CabinetQuery: cq}
	sbuild.label = cabinet.Label
	sbuild.flds, sbuild.scan = &cq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CabinetSelect configured with the given aggregations.
func (cq *CabinetQuery) Aggregate(fns ...AggregateFunc) *CabinetSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CabinetQuery) prepareQuery(ctx context.Context) error {
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
	for _, f := range cq.fields {
		if !cabinet.ValidColumn(f) {
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

func (cq *CabinetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Cabinet, error) {
	var (
		nodes       = []*Cabinet{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withBins != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Cabinet).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Cabinet{config: cq.config}
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
	if query := cq.withBins; query != nil {
		if err := cq.loadBins(ctx, query, nodes,
			func(n *Cabinet) { n.Edges.Bins = []*Bin{} },
			func(n *Cabinet, e *Bin) { n.Edges.Bins = append(n.Edges.Bins, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CabinetQuery) loadBins(ctx context.Context, query *BinQuery, nodes []*Cabinet, init func(*Cabinet), assign func(*Cabinet, *Bin)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Cabinet)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Bin(func(s *sql.Selector) {
		s.Where(sql.InValues(cabinet.BinsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.CabinetID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "cabinet_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CabinetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.fields
	if len(cq.fields) > 0 {
		_spec.Unique = cq.unique != nil && *cq.unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CabinetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cabinet.Table,
			Columns: cabinet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: cabinet.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cabinet.FieldID)
		for i := range fields {
			if fields[i] != cabinet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
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

func (cq *CabinetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(cabinet.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = cabinet.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.unique != nil && *cq.unique {
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
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cq *CabinetQuery) Modify(modifiers ...func(s *sql.Selector)) *CabinetSelect {
	cq.modifiers = append(cq.modifiers, modifiers...)
	return cq.Select()
}

// CabinetGroupBy is the group-by builder for Cabinet entities.
type CabinetGroupBy struct {
	selector
	build *CabinetQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CabinetGroupBy) Aggregate(fns ...AggregateFunc) *CabinetGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CabinetGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeCabinet, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CabinetQuery, *CabinetGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CabinetGroupBy) sqlScan(ctx context.Context, root *CabinetQuery, v any) error {
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

// CabinetSelect is the builder for selecting fields of Cabinet entities.
type CabinetSelect struct {
	*CabinetQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CabinetSelect) Aggregate(fns ...AggregateFunc) *CabinetSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CabinetSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeCabinet, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CabinetQuery, *CabinetSelect](ctx, cs.CabinetQuery, cs, cs.inters, v)
}

func (cs *CabinetSelect) sqlScan(ctx context.Context, root *CabinetQuery, v any) error {
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
func (cs *CabinetSelect) Modify(modifiers ...func(s *sql.Selector)) *CabinetSelect {
	cs.modifiers = append(cs.modifiers, modifiers...)
	return cs
}
