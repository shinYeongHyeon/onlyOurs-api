// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shinYeongHyeon/onlyOurs-api/ent/predicate"
	"github.com/shinYeongHyeon/onlyOurs-api/ent/users"
)

// UsersQuery is the builder for querying Users entities.
type UsersQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Users
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UsersQuery builder.
func (uq *UsersQuery) Where(ps ...predicate.Users) *UsersQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit adds a limit step to the query.
func (uq *UsersQuery) Limit(limit int) *UsersQuery {
	uq.limit = &limit
	return uq
}

// Offset adds an offset step to the query.
func (uq *UsersQuery) Offset(offset int) *UsersQuery {
	uq.offset = &offset
	return uq
}

// Order adds an order step to the query.
func (uq *UsersQuery) Order(o ...OrderFunc) *UsersQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// First returns the first Users entity from the query.
// Returns a *NotFoundError when no Users was found.
func (uq *UsersQuery) First(ctx context.Context) (*Users, error) {
	nodes, err := uq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{users.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UsersQuery) FirstX(ctx context.Context) *Users {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Users ID from the query.
// Returns a *NotFoundError when no Users ID was found.
func (uq *UsersQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{users.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UsersQuery) FirstIDX(ctx context.Context) int {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Users entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Users entity is not found.
// Returns a *NotFoundError when no Users entities are found.
func (uq *UsersQuery) Only(ctx context.Context) (*Users, error) {
	nodes, err := uq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{users.Label}
	default:
		return nil, &NotSingularError{users.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UsersQuery) OnlyX(ctx context.Context) *Users {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Users ID in the query.
// Returns a *NotSingularError when exactly one Users ID is not found.
// Returns a *NotFoundError when no entities are found.
func (uq *UsersQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{users.Label}
	default:
		err = &NotSingularError{users.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UsersQuery) OnlyIDX(ctx context.Context) int {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UsersSlice.
func (uq *UsersQuery) All(ctx context.Context) ([]*Users, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uq *UsersQuery) AllX(ctx context.Context) []*Users {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Users IDs.
func (uq *UsersQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uq.Select(users.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UsersQuery) IDsX(ctx context.Context) []int {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UsersQuery) Count(ctx context.Context) (int, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UsersQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UsersQuery) Exist(ctx context.Context) (bool, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UsersQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UsersQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UsersQuery) Clone() *UsersQuery {
	if uq == nil {
		return nil
	}
	return &UsersQuery{
		config:     uq.config,
		limit:      uq.limit,
		offset:     uq.offset,
		order:      append([]OrderFunc{}, uq.order...),
		predicates: append([]predicate.Users{}, uq.predicates...),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Users.Query().
//		GroupBy(users.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uq *UsersQuery) GroupBy(field string, fields ...string) *UsersGroupBy {
	group := &UsersGroupBy{config: uq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Users.Query().
//		Select(users.FieldName).
//		Scan(ctx, &v)
//
func (uq *UsersQuery) Select(field string, fields ...string) *UsersSelect {
	uq.fields = append([]string{field}, fields...)
	return &UsersSelect{UsersQuery: uq}
}

func (uq *UsersQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uq.fields {
		if !users.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UsersQuery) sqlAll(ctx context.Context) ([]*Users, error) {
	var (
		nodes = []*Users{}
		_spec = uq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Users{config: uq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (uq *UsersQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UsersQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uq *UsersQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   users.Table,
			Columns: users.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: users.FieldID,
			},
		},
		From:   uq.sql,
		Unique: true,
	}
	if fields := uq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, users.FieldID)
		for i := range fields {
			if fields[i] != users.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, users.ValidColumn)
			}
		}
	}
	return _spec
}

func (uq *UsersQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(users.Table)
	selector := builder.Select(t1.Columns(users.Columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(users.Columns...)...)
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector, users.ValidColumn)
	}
	if offset := uq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UsersGroupBy is the group-by builder for Users entities.
type UsersGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UsersGroupBy) Aggregate(fns ...AggregateFunc) *UsersGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the group-by query and scans the result into the given value.
func (ugb *UsersGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ugb.path(ctx)
	if err != nil {
		return err
	}
	ugb.sql = query
	return ugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ugb *UsersGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ugb *UsersGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UsersGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ugb *UsersGroupBy) StringsX(ctx context.Context) []string {
	v, err := ugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ugb *UsersGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ugb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{users.Label}
	default:
		err = fmt.Errorf("ent: UsersGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ugb *UsersGroupBy) StringX(ctx context.Context) string {
	v, err := ugb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ugb *UsersGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UsersGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ugb *UsersGroupBy) IntsX(ctx context.Context) []int {
	v, err := ugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ugb *UsersGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ugb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{users.Label}
	default:
		err = fmt.Errorf("ent: UsersGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ugb *UsersGroupBy) IntX(ctx context.Context) int {
	v, err := ugb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ugb *UsersGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UsersGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ugb *UsersGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ugb *UsersGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ugb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{users.Label}
	default:
		err = fmt.Errorf("ent: UsersGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ugb *UsersGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ugb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ugb *UsersGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UsersGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ugb *UsersGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ugb *UsersGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ugb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{users.Label}
	default:
		err = fmt.Errorf("ent: UsersGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ugb *UsersGroupBy) BoolX(ctx context.Context) bool {
	v, err := ugb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ugb *UsersGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ugb.fields {
		if !users.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ugb *UsersGroupBy) sqlQuery() *sql.Selector {
	selector := ugb.sql
	columns := make([]string, 0, len(ugb.fields)+len(ugb.fns))
	columns = append(columns, ugb.fields...)
	for _, fn := range ugb.fns {
		columns = append(columns, fn(selector, users.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(ugb.fields...)
}

// UsersSelect is the builder for selecting fields of Users entities.
type UsersSelect struct {
	*UsersQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (us *UsersSelect) Scan(ctx context.Context, v interface{}) error {
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	us.sql = us.UsersQuery.sqlQuery(ctx)
	return us.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (us *UsersSelect) ScanX(ctx context.Context, v interface{}) {
	if err := us.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (us *UsersSelect) Strings(ctx context.Context) ([]string, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UsersSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (us *UsersSelect) StringsX(ctx context.Context) []string {
	v, err := us.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (us *UsersSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = us.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{users.Label}
	default:
		err = fmt.Errorf("ent: UsersSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (us *UsersSelect) StringX(ctx context.Context) string {
	v, err := us.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (us *UsersSelect) Ints(ctx context.Context) ([]int, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UsersSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (us *UsersSelect) IntsX(ctx context.Context) []int {
	v, err := us.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (us *UsersSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = us.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{users.Label}
	default:
		err = fmt.Errorf("ent: UsersSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (us *UsersSelect) IntX(ctx context.Context) int {
	v, err := us.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (us *UsersSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UsersSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (us *UsersSelect) Float64sX(ctx context.Context) []float64 {
	v, err := us.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (us *UsersSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = us.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{users.Label}
	default:
		err = fmt.Errorf("ent: UsersSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (us *UsersSelect) Float64X(ctx context.Context) float64 {
	v, err := us.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (us *UsersSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UsersSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (us *UsersSelect) BoolsX(ctx context.Context) []bool {
	v, err := us.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (us *UsersSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = us.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{users.Label}
	default:
		err = fmt.Errorf("ent: UsersSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (us *UsersSelect) BoolX(ctx context.Context) bool {
	v, err := us.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (us *UsersSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := us.sqlQuery().Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (us *UsersSelect) sqlQuery() sql.Querier {
	selector := us.sql
	selector.Select(selector.Columns(us.fields...)...)
	return selector
}
