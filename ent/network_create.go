// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/network"
	"github.com/gen0cide/laforge/ent/tag"
)

// NetworkCreate is the builder for creating a Network entity.
type NetworkCreate struct {
	config
	mutation *NetworkMutation
	hooks    []Hook
}

// SetName sets the name field.
func (nc *NetworkCreate) SetName(s string) *NetworkCreate {
	nc.mutation.SetName(s)
	return nc
}

// SetCidr sets the cidr field.
func (nc *NetworkCreate) SetCidr(s string) *NetworkCreate {
	nc.mutation.SetCidr(s)
	return nc
}

// SetVdiVisible sets the vdi_visible field.
func (nc *NetworkCreate) SetVdiVisible(b bool) *NetworkCreate {
	nc.mutation.SetVdiVisible(b)
	return nc
}

// SetVars sets the vars field.
func (nc *NetworkCreate) SetVars(s []string) *NetworkCreate {
	nc.mutation.SetVars(s)
	return nc
}

// AddTagIDs adds the tag edge to Tag by ids.
func (nc *NetworkCreate) AddTagIDs(ids ...int) *NetworkCreate {
	nc.mutation.AddTagIDs(ids...)
	return nc
}

// AddTag adds the tag edges to Tag.
func (nc *NetworkCreate) AddTag(t ...*Tag) *NetworkCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return nc.AddTagIDs(ids...)
}

// AddNetworkToEnvironmentIDs adds the NetworkToEnvironment edge to Environment by ids.
func (nc *NetworkCreate) AddNetworkToEnvironmentIDs(ids ...int) *NetworkCreate {
	nc.mutation.AddNetworkToEnvironmentIDs(ids...)
	return nc
}

// AddNetworkToEnvironment adds the NetworkToEnvironment edges to Environment.
func (nc *NetworkCreate) AddNetworkToEnvironment(e ...*Environment) *NetworkCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return nc.AddNetworkToEnvironmentIDs(ids...)
}

// Mutation returns the NetworkMutation object of the builder.
func (nc *NetworkCreate) Mutation() *NetworkMutation {
	return nc.mutation
}

// Save creates the Network in the database.
func (nc *NetworkCreate) Save(ctx context.Context) (*Network, error) {
	var (
		err  error
		node *Network
	)
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			node, err = nc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			mut = nc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NetworkCreate) SaveX(ctx context.Context) *Network {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (nc *NetworkCreate) check() error {
	if _, ok := nc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := nc.mutation.Cidr(); !ok {
		return &ValidationError{Name: "cidr", err: errors.New("ent: missing required field \"cidr\"")}
	}
	if _, ok := nc.mutation.VdiVisible(); !ok {
		return &ValidationError{Name: "vdi_visible", err: errors.New("ent: missing required field \"vdi_visible\"")}
	}
	if _, ok := nc.mutation.Vars(); !ok {
		return &ValidationError{Name: "vars", err: errors.New("ent: missing required field \"vars\"")}
	}
	return nil
}

func (nc *NetworkCreate) sqlSave(ctx context.Context) (*Network, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (nc *NetworkCreate) createSpec() (*Network, *sqlgraph.CreateSpec) {
	var (
		_node = &Network{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: network.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: network.FieldID,
			},
		}
	)
	if value, ok := nc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: network.FieldName,
		})
		_node.Name = value
	}
	if value, ok := nc.mutation.Cidr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: network.FieldCidr,
		})
		_node.Cidr = value
	}
	if value, ok := nc.mutation.VdiVisible(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: network.FieldVdiVisible,
		})
		_node.VdiVisible = value
	}
	if value, ok := nc.mutation.Vars(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: network.FieldVars,
		})
		_node.Vars = value
	}
	if nodes := nc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   network.TagTable,
			Columns: []string{network.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.NetworkToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   network.NetworkToEnvironmentTable,
			Columns: network.NetworkToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NetworkCreateBulk is the builder for creating a bulk of Network entities.
type NetworkCreateBulk struct {
	config
	builders []*NetworkCreate
}

// Save creates the Network entities in the database.
func (ncb *NetworkCreateBulk) Save(ctx context.Context) ([]*Network, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Network, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NetworkMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ncb *NetworkCreateBulk) SaveX(ctx context.Context) []*Network {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
