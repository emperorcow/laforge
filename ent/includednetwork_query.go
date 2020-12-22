// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/includednetwork"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/tag"
)

// IncludedNetworkQuery is the builder for querying IncludedNetwork entities.
type IncludedNetworkQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.IncludedNetwork
	// eager-loading edges.
	withTag                          *TagQuery
	withIncludedNetworkToEnvironment *EnvironmentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (inq *IncludedNetworkQuery) Where(ps ...predicate.IncludedNetwork) *IncludedNetworkQuery {
	inq.predicates = append(inq.predicates, ps...)
	return inq
}

// Limit adds a limit step to the query.
func (inq *IncludedNetworkQuery) Limit(limit int) *IncludedNetworkQuery {
	inq.limit = &limit
	return inq
}

// Offset adds an offset step to the query.
func (inq *IncludedNetworkQuery) Offset(offset int) *IncludedNetworkQuery {
	inq.offset = &offset
	return inq
}

// Order adds an order step to the query.
func (inq *IncludedNetworkQuery) Order(o ...OrderFunc) *IncludedNetworkQuery {
	inq.order = append(inq.order, o...)
	return inq
}

// QueryTag chains the current query on the tag edge.
func (inq *IncludedNetworkQuery) QueryTag() *TagQuery {
	query := &TagQuery{config: inq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := inq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := inq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(includednetwork.Table, includednetwork.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, includednetwork.TagTable, includednetwork.TagColumn),
		)
		fromU = sqlgraph.SetNeighbors(inq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIncludedNetworkToEnvironment chains the current query on the IncludedNetworkToEnvironment edge.
func (inq *IncludedNetworkQuery) QueryIncludedNetworkToEnvironment() *EnvironmentQuery {
	query := &EnvironmentQuery{config: inq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := inq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := inq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(includednetwork.Table, includednetwork.FieldID, selector),
			sqlgraph.To(environment.Table, environment.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, includednetwork.IncludedNetworkToEnvironmentTable, includednetwork.IncludedNetworkToEnvironmentPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(inq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first IncludedNetwork entity in the query. Returns *NotFoundError when no includednetwork was found.
func (inq *IncludedNetworkQuery) First(ctx context.Context) (*IncludedNetwork, error) {
	nodes, err := inq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{includednetwork.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (inq *IncludedNetworkQuery) FirstX(ctx context.Context) *IncludedNetwork {
	node, err := inq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first IncludedNetwork id in the query. Returns *NotFoundError when no id was found.
func (inq *IncludedNetworkQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = inq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{includednetwork.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (inq *IncludedNetworkQuery) FirstIDX(ctx context.Context) int {
	id, err := inq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only IncludedNetwork entity in the query, returns an error if not exactly one entity was returned.
func (inq *IncludedNetworkQuery) Only(ctx context.Context) (*IncludedNetwork, error) {
	nodes, err := inq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{includednetwork.Label}
	default:
		return nil, &NotSingularError{includednetwork.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (inq *IncludedNetworkQuery) OnlyX(ctx context.Context) *IncludedNetwork {
	node, err := inq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only IncludedNetwork id in the query, returns an error if not exactly one id was returned.
func (inq *IncludedNetworkQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = inq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{includednetwork.Label}
	default:
		err = &NotSingularError{includednetwork.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (inq *IncludedNetworkQuery) OnlyIDX(ctx context.Context) int {
	id, err := inq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of IncludedNetworks.
func (inq *IncludedNetworkQuery) All(ctx context.Context) ([]*IncludedNetwork, error) {
	if err := inq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return inq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (inq *IncludedNetworkQuery) AllX(ctx context.Context) []*IncludedNetwork {
	nodes, err := inq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of IncludedNetwork ids.
func (inq *IncludedNetworkQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := inq.Select(includednetwork.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (inq *IncludedNetworkQuery) IDsX(ctx context.Context) []int {
	ids, err := inq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (inq *IncludedNetworkQuery) Count(ctx context.Context) (int, error) {
	if err := inq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return inq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (inq *IncludedNetworkQuery) CountX(ctx context.Context) int {
	count, err := inq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (inq *IncludedNetworkQuery) Exist(ctx context.Context) (bool, error) {
	if err := inq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return inq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (inq *IncludedNetworkQuery) ExistX(ctx context.Context) bool {
	exist, err := inq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (inq *IncludedNetworkQuery) Clone() *IncludedNetworkQuery {
	if inq == nil {
		return nil
	}
	return &IncludedNetworkQuery{
		config:                           inq.config,
		limit:                            inq.limit,
		offset:                           inq.offset,
		order:                            append([]OrderFunc{}, inq.order...),
		predicates:                       append([]predicate.IncludedNetwork{}, inq.predicates...),
		withTag:                          inq.withTag.Clone(),
		withIncludedNetworkToEnvironment: inq.withIncludedNetworkToEnvironment.Clone(),
		// clone intermediate query.
		sql:  inq.sql.Clone(),
		path: inq.path,
	}
}

//  WithTag tells the query-builder to eager-loads the nodes that are connected to
// the "tag" edge. The optional arguments used to configure the query builder of the edge.
func (inq *IncludedNetworkQuery) WithTag(opts ...func(*TagQuery)) *IncludedNetworkQuery {
	query := &TagQuery{config: inq.config}
	for _, opt := range opts {
		opt(query)
	}
	inq.withTag = query
	return inq
}

//  WithIncludedNetworkToEnvironment tells the query-builder to eager-loads the nodes that are connected to
// the "IncludedNetworkToEnvironment" edge. The optional arguments used to configure the query builder of the edge.
func (inq *IncludedNetworkQuery) WithIncludedNetworkToEnvironment(opts ...func(*EnvironmentQuery)) *IncludedNetworkQuery {
	query := &EnvironmentQuery{config: inq.config}
	for _, opt := range opts {
		opt(query)
	}
	inq.withIncludedNetworkToEnvironment = query
	return inq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.IncludedNetwork.Query().
//		GroupBy(includednetwork.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (inq *IncludedNetworkQuery) GroupBy(field string, fields ...string) *IncludedNetworkGroupBy {
	group := &IncludedNetworkGroupBy{config: inq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := inq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return inq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.IncludedNetwork.Query().
//		Select(includednetwork.FieldName).
//		Scan(ctx, &v)
//
func (inq *IncludedNetworkQuery) Select(field string, fields ...string) *IncludedNetworkSelect {
	selector := &IncludedNetworkSelect{config: inq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := inq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return inq.sqlQuery(), nil
	}
	return selector
}

func (inq *IncludedNetworkQuery) prepareQuery(ctx context.Context) error {
	if inq.path != nil {
		prev, err := inq.path(ctx)
		if err != nil {
			return err
		}
		inq.sql = prev
	}
	return nil
}

func (inq *IncludedNetworkQuery) sqlAll(ctx context.Context) ([]*IncludedNetwork, error) {
	var (
		nodes       = []*IncludedNetwork{}
		_spec       = inq.querySpec()
		loadedTypes = [2]bool{
			inq.withTag != nil,
			inq.withIncludedNetworkToEnvironment != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &IncludedNetwork{config: inq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, inq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := inq.withTag; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*IncludedNetwork)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Tag = []*Tag{}
		}
		query.withFKs = true
		query.Where(predicate.Tag(func(s *sql.Selector) {
			s.Where(sql.InValues(includednetwork.TagColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.included_network_tag
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "included_network_tag" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "included_network_tag" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Tag = append(node.Edges.Tag, n)
		}
	}

	if query := inq.withIncludedNetworkToEnvironment; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*IncludedNetwork, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.IncludedNetworkToEnvironment = []*Environment{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*IncludedNetwork)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   includednetwork.IncludedNetworkToEnvironmentTable,
				Columns: includednetwork.IncludedNetworkToEnvironmentPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(includednetwork.IncludedNetworkToEnvironmentPrimaryKey[0], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, inq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "IncludedNetworkToEnvironment": %v`, err)
		}
		query.Where(environment.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "IncludedNetworkToEnvironment" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.IncludedNetworkToEnvironment = append(nodes[i].Edges.IncludedNetworkToEnvironment, n)
			}
		}
	}

	return nodes, nil
}

func (inq *IncludedNetworkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := inq.querySpec()
	return sqlgraph.CountNodes(ctx, inq.driver, _spec)
}

func (inq *IncludedNetworkQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := inq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (inq *IncludedNetworkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   includednetwork.Table,
			Columns: includednetwork.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: includednetwork.FieldID,
			},
		},
		From:   inq.sql,
		Unique: true,
	}
	if ps := inq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := inq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := inq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := inq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, includednetwork.ValidColumn)
			}
		}
	}
	return _spec
}

func (inq *IncludedNetworkQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(inq.driver.Dialect())
	t1 := builder.Table(includednetwork.Table)
	selector := builder.Select(t1.Columns(includednetwork.Columns...)...).From(t1)
	if inq.sql != nil {
		selector = inq.sql
		selector.Select(selector.Columns(includednetwork.Columns...)...)
	}
	for _, p := range inq.predicates {
		p(selector)
	}
	for _, p := range inq.order {
		p(selector, includednetwork.ValidColumn)
	}
	if offset := inq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := inq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IncludedNetworkGroupBy is the builder for group-by IncludedNetwork entities.
type IncludedNetworkGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ingb *IncludedNetworkGroupBy) Aggregate(fns ...AggregateFunc) *IncludedNetworkGroupBy {
	ingb.fns = append(ingb.fns, fns...)
	return ingb
}

// Scan applies the group-by query and scan the result into the given value.
func (ingb *IncludedNetworkGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ingb.path(ctx)
	if err != nil {
		return err
	}
	ingb.sql = query
	return ingb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ingb *IncludedNetworkGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ingb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ingb *IncludedNetworkGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ingb.fields) > 1 {
		return nil, errors.New("ent: IncludedNetworkGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ingb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ingb *IncludedNetworkGroupBy) StringsX(ctx context.Context) []string {
	v, err := ingb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (ingb *IncludedNetworkGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ingb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{includednetwork.Label}
	default:
		err = fmt.Errorf("ent: IncludedNetworkGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ingb *IncludedNetworkGroupBy) StringX(ctx context.Context) string {
	v, err := ingb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ingb *IncludedNetworkGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ingb.fields) > 1 {
		return nil, errors.New("ent: IncludedNetworkGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ingb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ingb *IncludedNetworkGroupBy) IntsX(ctx context.Context) []int {
	v, err := ingb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (ingb *IncludedNetworkGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ingb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{includednetwork.Label}
	default:
		err = fmt.Errorf("ent: IncludedNetworkGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ingb *IncludedNetworkGroupBy) IntX(ctx context.Context) int {
	v, err := ingb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ingb *IncludedNetworkGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ingb.fields) > 1 {
		return nil, errors.New("ent: IncludedNetworkGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ingb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ingb *IncludedNetworkGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ingb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (ingb *IncludedNetworkGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ingb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{includednetwork.Label}
	default:
		err = fmt.Errorf("ent: IncludedNetworkGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ingb *IncludedNetworkGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ingb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ingb *IncludedNetworkGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ingb.fields) > 1 {
		return nil, errors.New("ent: IncludedNetworkGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ingb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ingb *IncludedNetworkGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ingb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (ingb *IncludedNetworkGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ingb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{includednetwork.Label}
	default:
		err = fmt.Errorf("ent: IncludedNetworkGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ingb *IncludedNetworkGroupBy) BoolX(ctx context.Context) bool {
	v, err := ingb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ingb *IncludedNetworkGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ingb.fields {
		if !includednetwork.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ingb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ingb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ingb *IncludedNetworkGroupBy) sqlQuery() *sql.Selector {
	selector := ingb.sql
	columns := make([]string, 0, len(ingb.fields)+len(ingb.fns))
	columns = append(columns, ingb.fields...)
	for _, fn := range ingb.fns {
		columns = append(columns, fn(selector, includednetwork.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(ingb.fields...)
}

// IncludedNetworkSelect is the builder for select fields of IncludedNetwork entities.
type IncludedNetworkSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ins *IncludedNetworkSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ins.path(ctx)
	if err != nil {
		return err
	}
	ins.sql = query
	return ins.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ins *IncludedNetworkSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ins.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ins *IncludedNetworkSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ins.fields) > 1 {
		return nil, errors.New("ent: IncludedNetworkSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ins.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ins *IncludedNetworkSelect) StringsX(ctx context.Context) []string {
	v, err := ins.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ins *IncludedNetworkSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ins.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{includednetwork.Label}
	default:
		err = fmt.Errorf("ent: IncludedNetworkSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ins *IncludedNetworkSelect) StringX(ctx context.Context) string {
	v, err := ins.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ins *IncludedNetworkSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ins.fields) > 1 {
		return nil, errors.New("ent: IncludedNetworkSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ins.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ins *IncludedNetworkSelect) IntsX(ctx context.Context) []int {
	v, err := ins.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ins *IncludedNetworkSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ins.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{includednetwork.Label}
	default:
		err = fmt.Errorf("ent: IncludedNetworkSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ins *IncludedNetworkSelect) IntX(ctx context.Context) int {
	v, err := ins.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ins *IncludedNetworkSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ins.fields) > 1 {
		return nil, errors.New("ent: IncludedNetworkSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ins.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ins *IncludedNetworkSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ins.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ins *IncludedNetworkSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ins.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{includednetwork.Label}
	default:
		err = fmt.Errorf("ent: IncludedNetworkSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ins *IncludedNetworkSelect) Float64X(ctx context.Context) float64 {
	v, err := ins.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ins *IncludedNetworkSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ins.fields) > 1 {
		return nil, errors.New("ent: IncludedNetworkSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ins.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ins *IncludedNetworkSelect) BoolsX(ctx context.Context) []bool {
	v, err := ins.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ins *IncludedNetworkSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ins.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{includednetwork.Label}
	default:
		err = fmt.Errorf("ent: IncludedNetworkSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ins *IncludedNetworkSelect) BoolX(ctx context.Context) bool {
	v, err := ins.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ins *IncludedNetworkSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ins.fields {
		if !includednetwork.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := ins.sqlQuery().Query()
	if err := ins.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ins *IncludedNetworkSelect) sqlQuery() sql.Querier {
	selector := ins.sql
	selector.Select(selector.Columns(ins.fields...)...)
	return selector
}
