// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/GRTheory/my-model/ent/predicate"
	"github.com/GRTheory/my-model/ent/student"
	"github.com/GRTheory/my-model/ent/teacher"
)

// StudentQuery is the builder for querying Student entities.
type StudentQuery struct {
	config
	ctx          *QueryContext
	order        []OrderFunc
	inters       []Interceptor
	predicates   []predicate.Student
	withTeachers *TeacherQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StudentQuery builder.
func (sq *StudentQuery) Where(ps ...predicate.Student) *StudentQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *StudentQuery) Limit(limit int) *StudentQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *StudentQuery) Offset(offset int) *StudentQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *StudentQuery) Unique(unique bool) *StudentQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *StudentQuery) Order(o ...OrderFunc) *StudentQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryTeachers chains the current query on the "teachers" edge.
func (sq *StudentQuery) QueryTeachers() *TeacherQuery {
	query := (&TeacherClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(student.Table, student.FieldID, selector),
			sqlgraph.To(teacher.Table, teacher.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, student.TeachersTable, student.TeachersColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Student entity from the query.
// Returns a *NotFoundError when no Student was found.
func (sq *StudentQuery) First(ctx context.Context) (*Student, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{student.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *StudentQuery) FirstX(ctx context.Context) *Student {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Student ID from the query.
// Returns a *NotFoundError when no Student ID was found.
func (sq *StudentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{student.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *StudentQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Student entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Student entity is found.
// Returns a *NotFoundError when no Student entities are found.
func (sq *StudentQuery) Only(ctx context.Context) (*Student, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{student.Label}
	default:
		return nil, &NotSingularError{student.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *StudentQuery) OnlyX(ctx context.Context) *Student {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Student ID in the query.
// Returns a *NotSingularError when more than one Student ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *StudentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{student.Label}
	default:
		err = &NotSingularError{student.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *StudentQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Students.
func (sq *StudentQuery) All(ctx context.Context) ([]*Student, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Student, *StudentQuery]()
	return withInterceptors[[]*Student](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *StudentQuery) AllX(ctx context.Context) []*Student {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Student IDs.
func (sq *StudentQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err := sq.Select(student.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *StudentQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *StudentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*StudentQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *StudentQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *StudentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
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
func (sq *StudentQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StudentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *StudentQuery) Clone() *StudentQuery {
	if sq == nil {
		return nil
	}
	return &StudentQuery{
		config:       sq.config,
		ctx:          sq.ctx.Clone(),
		order:        append([]OrderFunc{}, sq.order...),
		inters:       append([]Interceptor{}, sq.inters...),
		predicates:   append([]predicate.Student{}, sq.predicates...),
		withTeachers: sq.withTeachers.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithTeachers tells the query-builder to eager-load the nodes that are connected to
// the "teachers" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StudentQuery) WithTeachers(opts ...func(*TeacherQuery)) *StudentQuery {
	query := (&TeacherClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withTeachers = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SID int `json:"s_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Student.Query().
//		GroupBy(student.FieldSID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *StudentQuery) GroupBy(field string, fields ...string) *StudentGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &StudentGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = student.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SID int `json:"s_id,omitempty"`
//	}
//
//	client.Student.Query().
//		Select(student.FieldSID).
//		Scan(ctx, &v)
func (sq *StudentQuery) Select(fields ...string) *StudentSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &StudentSelect{StudentQuery: sq}
	sbuild.label = student.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a StudentSelect configured with the given aggregations.
func (sq *StudentQuery) Aggregate(fns ...AggregateFunc) *StudentSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *StudentQuery) prepareQuery(ctx context.Context) error {
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
	for _, f := range sq.ctx.Fields {
		if !student.ValidColumn(f) {
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

func (sq *StudentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Student, error) {
	var (
		nodes       = []*Student{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withTeachers != nil,
		}
	)
	if sq.withTeachers != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, student.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Student).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Student{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
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
	if query := sq.withTeachers; query != nil {
		if err := sq.loadTeachers(ctx, query, nodes, nil,
			func(n *Student, e *Teacher) { n.Edges.Teachers = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *StudentQuery) loadTeachers(ctx context.Context, query *TeacherQuery, nodes []*Student, init func(*Student), assign func(*Student, *Teacher)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Student)
	for i := range nodes {
		if nodes[i].teacher_students == nil {
			continue
		}
		fk := *nodes[i].teacher_students
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(teacher.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "teacher_students" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *StudentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *StudentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   student.Table,
			Columns: student.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: student.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, student.FieldID)
		for i := range fields {
			if fields[i] != student.FieldID {
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
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
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

func (sq *StudentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(student.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = student.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StudentGroupBy is the group-by builder for Student entities.
type StudentGroupBy struct {
	selector
	build *StudentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *StudentGroupBy) Aggregate(fns ...AggregateFunc) *StudentGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *StudentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StudentQuery, *StudentGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *StudentGroupBy) sqlScan(ctx context.Context, root *StudentQuery, v any) error {
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

// StudentSelect is the builder for selecting fields of Student entities.
type StudentSelect struct {
	*StudentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *StudentSelect) Aggregate(fns ...AggregateFunc) *StudentSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *StudentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StudentQuery, *StudentSelect](ctx, ss.StudentQuery, ss, ss.inters, v)
}

func (ss *StudentSelect) sqlScan(ctx context.Context, root *StudentQuery, v any) error {
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
