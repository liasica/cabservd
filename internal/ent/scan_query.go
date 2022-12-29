// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/ent/predicate"
	"github.com/auroraride/cabservd/internal/ent/scan"
	"github.com/google/uuid"
)

// ScanQuery is the builder for querying Scan entities.
type ScanQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	inters      []Interceptor
	predicates  []predicate.Scan
	withCabinet *CabinetQuery
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ScanQuery builder.
func (sq *ScanQuery) Where(ps ...predicate.Scan) *ScanQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *ScanQuery) Limit(limit int) *ScanQuery {
	sq.limit = &limit
	return sq
}

// Offset to start from.
func (sq *ScanQuery) Offset(offset int) *ScanQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *ScanQuery) Unique(unique bool) *ScanQuery {
	sq.unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *ScanQuery) Order(o ...OrderFunc) *ScanQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryCabinet chains the current query on the "cabinet" edge.
func (sq *ScanQuery) QueryCabinet() *CabinetQuery {
	query := (&CabinetClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scan.Table, scan.FieldID, selector),
			sqlgraph.To(cabinet.Table, cabinet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, scan.CabinetTable, scan.CabinetColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Scan entity from the query.
// Returns a *NotFoundError when no Scan was found.
func (sq *ScanQuery) First(ctx context.Context) (*Scan, error) {
	nodes, err := sq.Limit(1).All(newQueryContext(ctx, TypeScan, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{scan.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *ScanQuery) FirstX(ctx context.Context) *Scan {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Scan ID from the query.
// Returns a *NotFoundError when no Scan ID was found.
func (sq *ScanQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(1).IDs(newQueryContext(ctx, TypeScan, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{scan.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *ScanQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Scan entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Scan entity is found.
// Returns a *NotFoundError when no Scan entities are found.
func (sq *ScanQuery) Only(ctx context.Context) (*Scan, error) {
	nodes, err := sq.Limit(2).All(newQueryContext(ctx, TypeScan, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{scan.Label}
	default:
		return nil, &NotSingularError{scan.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *ScanQuery) OnlyX(ctx context.Context) *Scan {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Scan ID in the query.
// Returns a *NotSingularError when more than one Scan ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *ScanQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(2).IDs(newQueryContext(ctx, TypeScan, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{scan.Label}
	default:
		err = &NotSingularError{scan.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *ScanQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Scans.
func (sq *ScanQuery) All(ctx context.Context) ([]*Scan, error) {
	ctx = newQueryContext(ctx, TypeScan, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Scan, *ScanQuery]()
	return withInterceptors[[]*Scan](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *ScanQuery) AllX(ctx context.Context) []*Scan {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Scan IDs.
func (sq *ScanQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeScan, "IDs")
	if err := sq.Select(scan.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *ScanQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *ScanQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeScan, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*ScanQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *ScanQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *ScanQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeScan, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *ScanQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ScanQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *ScanQuery) Clone() *ScanQuery {
	if sq == nil {
		return nil
	}
	return &ScanQuery{
		config:      sq.config,
		limit:       sq.limit,
		offset:      sq.offset,
		order:       append([]OrderFunc{}, sq.order...),
		predicates:  append([]predicate.Scan{}, sq.predicates...),
		withCabinet: sq.withCabinet.Clone(),
		// clone intermediate query.
		sql:    sq.sql.Clone(),
		path:   sq.path,
		unique: sq.unique,
	}
}

// WithCabinet tells the query-builder to eager-load the nodes that are connected to
// the "cabinet" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ScanQuery) WithCabinet(opts ...func(*CabinetQuery)) *ScanQuery {
	query := (&CabinetClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withCabinet = query
	return sq
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
//	client.Scan.Query().
//		GroupBy(scan.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *ScanQuery) GroupBy(field string, fields ...string) *ScanGroupBy {
	sq.fields = append([]string{field}, fields...)
	grbuild := &ScanGroupBy{build: sq}
	grbuild.flds = &sq.fields
	grbuild.label = scan.Label
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
//	client.Scan.Query().
//		Select(scan.FieldCreatedAt).
//		Scan(ctx, &v)
func (sq *ScanQuery) Select(fields ...string) *ScanSelect {
	sq.fields = append(sq.fields, fields...)
	sbuild := &ScanSelect{ScanQuery: sq}
	sbuild.label = scan.Label
	sbuild.flds, sbuild.scan = &sq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ScanSelect configured with the given aggregations.
func (sq *ScanQuery) Aggregate(fns ...AggregateFunc) *ScanSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *ScanQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.fields {
		if !scan.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *ScanQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Scan, error) {
	var (
		nodes       = []*Scan{}
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withCabinet != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Scan).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Scan{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withCabinet; query != nil {
		if err := sq.loadCabinet(ctx, query, nodes, nil,
			func(n *Scan, e *Cabinet) { n.Edges.Cabinet = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *ScanQuery) loadCabinet(ctx context.Context, query *CabinetQuery, nodes []*Scan, init func(*Scan), assign func(*Scan, *Cabinet)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*Scan)
	for i := range nodes {
		fk := nodes[i].CabinetID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
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

func (sq *ScanQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.fields
	if len(sq.fields) > 0 {
		_spec.Unique = sq.unique != nil && *sq.unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *ScanQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   scan.Table,
			Columns: scan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: scan.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, scan.FieldID)
		for i := range fields {
			if fields[i] != scan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *ScanQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(scan.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = scan.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.unique != nil && *sq.unique {
		selector.Distinct()
	}
	for _, m := range sq.modifiers {
		m(selector)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sq *ScanQuery) Modify(modifiers ...func(s *sql.Selector)) *ScanSelect {
	sq.modifiers = append(sq.modifiers, modifiers...)
	return sq.Select()
}

// ScanGroupBy is the group-by builder for Scan entities.
type ScanGroupBy struct {
	selector
	build *ScanQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *ScanGroupBy) Aggregate(fns ...AggregateFunc) *ScanGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *ScanGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeScan, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScanQuery, *ScanGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *ScanGroupBy) sqlScan(ctx context.Context, root *ScanQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ScanSelect is the builder for selecting fields of Scan entities.
type ScanSelect struct {
	*ScanQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *ScanSelect) Aggregate(fns ...AggregateFunc) *ScanSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *ScanSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeScan, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScanQuery, *ScanSelect](ctx, ss.ScanQuery, ss, ss.inters, v)
}

func (ss *ScanSelect) sqlScan(ctx context.Context, root *ScanQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ss *ScanSelect) Modify(modifiers ...func(s *sql.Selector)) *ScanSelect {
	ss.modifiers = append(ss.modifiers, modifiers...)
	return ss
}