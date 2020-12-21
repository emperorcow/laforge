// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/competition"
	"github.com/gen0cide/laforge/ent/dns"
	"github.com/gen0cide/laforge/ent/predicate"
)

// CompetitionUpdate is the builder for updating Competition entities.
type CompetitionUpdate struct {
	config
	hooks    []Hook
	mutation *CompetitionMutation
}

// Where adds a new predicate for the builder.
func (cu *CompetitionUpdate) Where(ps ...predicate.Competition) *CompetitionUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetRootPassword sets the root_password field.
func (cu *CompetitionUpdate) SetRootPassword(s string) *CompetitionUpdate {
	cu.mutation.SetRootPassword(s)
	return cu
}

// SetConfig sets the config field.
func (cu *CompetitionUpdate) SetConfig(m map[string]string) *CompetitionUpdate {
	cu.mutation.SetConfig(m)
	return cu
}

// AddDnIDs adds the dns edge to DNS by ids.
func (cu *CompetitionUpdate) AddDnIDs(ids ...int) *CompetitionUpdate {
	cu.mutation.AddDnIDs(ids...)
	return cu
}

// AddDNS adds the dns edges to DNS.
func (cu *CompetitionUpdate) AddDNS(d ...*DNS) *CompetitionUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cu.AddDnIDs(ids...)
}

// Mutation returns the CompetitionMutation object of the builder.
func (cu *CompetitionUpdate) Mutation() *CompetitionMutation {
	return cu.mutation
}

// ClearDNS clears all "dns" edges to type DNS.
func (cu *CompetitionUpdate) ClearDNS() *CompetitionUpdate {
	cu.mutation.ClearDNS()
	return cu
}

// RemoveDnIDs removes the dns edge to DNS by ids.
func (cu *CompetitionUpdate) RemoveDnIDs(ids ...int) *CompetitionUpdate {
	cu.mutation.RemoveDnIDs(ids...)
	return cu
}

// RemoveDNS removes dns edges to DNS.
func (cu *CompetitionUpdate) RemoveDNS(d ...*DNS) *CompetitionUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cu.RemoveDnIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CompetitionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompetitionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CompetitionUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CompetitionUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CompetitionUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CompetitionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   competition.Table,
			Columns: competition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: competition.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.RootPassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: competition.FieldRootPassword,
		})
	}
	if value, ok := cu.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: competition.FieldConfig,
		})
	}
	if cu.mutation.DNSCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.DNSTable,
			Columns: []string{competition.DNSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dns.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedDNSIDs(); len(nodes) > 0 && !cu.mutation.DNSCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.DNSTable,
			Columns: []string{competition.DNSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dns.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.DNSIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.DNSTable,
			Columns: []string{competition.DNSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dns.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{competition.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CompetitionUpdateOne is the builder for updating a single Competition entity.
type CompetitionUpdateOne struct {
	config
	hooks    []Hook
	mutation *CompetitionMutation
}

// SetRootPassword sets the root_password field.
func (cuo *CompetitionUpdateOne) SetRootPassword(s string) *CompetitionUpdateOne {
	cuo.mutation.SetRootPassword(s)
	return cuo
}

// SetConfig sets the config field.
func (cuo *CompetitionUpdateOne) SetConfig(m map[string]string) *CompetitionUpdateOne {
	cuo.mutation.SetConfig(m)
	return cuo
}

// AddDnIDs adds the dns edge to DNS by ids.
func (cuo *CompetitionUpdateOne) AddDnIDs(ids ...int) *CompetitionUpdateOne {
	cuo.mutation.AddDnIDs(ids...)
	return cuo
}

// AddDNS adds the dns edges to DNS.
func (cuo *CompetitionUpdateOne) AddDNS(d ...*DNS) *CompetitionUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cuo.AddDnIDs(ids...)
}

// Mutation returns the CompetitionMutation object of the builder.
func (cuo *CompetitionUpdateOne) Mutation() *CompetitionMutation {
	return cuo.mutation
}

// ClearDNS clears all "dns" edges to type DNS.
func (cuo *CompetitionUpdateOne) ClearDNS() *CompetitionUpdateOne {
	cuo.mutation.ClearDNS()
	return cuo
}

// RemoveDnIDs removes the dns edge to DNS by ids.
func (cuo *CompetitionUpdateOne) RemoveDnIDs(ids ...int) *CompetitionUpdateOne {
	cuo.mutation.RemoveDnIDs(ids...)
	return cuo
}

// RemoveDNS removes dns edges to DNS.
func (cuo *CompetitionUpdateOne) RemoveDNS(d ...*DNS) *CompetitionUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cuo.RemoveDnIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *CompetitionUpdateOne) Save(ctx context.Context) (*Competition, error) {
	var (
		err  error
		node *Competition
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompetitionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CompetitionUpdateOne) SaveX(ctx context.Context) *Competition {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CompetitionUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CompetitionUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CompetitionUpdateOne) sqlSave(ctx context.Context) (_node *Competition, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   competition.Table,
			Columns: competition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: competition.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Competition.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.RootPassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: competition.FieldRootPassword,
		})
	}
	if value, ok := cuo.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: competition.FieldConfig,
		})
	}
	if cuo.mutation.DNSCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.DNSTable,
			Columns: []string{competition.DNSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dns.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedDNSIDs(); len(nodes) > 0 && !cuo.mutation.DNSCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.DNSTable,
			Columns: []string{competition.DNSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dns.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.DNSIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.DNSTable,
			Columns: []string{competition.DNSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dns.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Competition{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{competition.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
