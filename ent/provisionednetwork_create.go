// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/network"
	"github.com/gen0cide/laforge/ent/provisionedhost"
	"github.com/gen0cide/laforge/ent/provisionednetwork"
	"github.com/gen0cide/laforge/ent/status"
	"github.com/gen0cide/laforge/ent/tag"
	"github.com/gen0cide/laforge/ent/team"
)

// ProvisionedNetworkCreate is the builder for creating a ProvisionedNetwork entity.
type ProvisionedNetworkCreate struct {
	config
	mutation *ProvisionedNetworkMutation
	hooks    []Hook
}

// SetName sets the name field.
func (pnc *ProvisionedNetworkCreate) SetName(s string) *ProvisionedNetworkCreate {
	pnc.mutation.SetName(s)
	return pnc
}

// SetCidr sets the cidr field.
func (pnc *ProvisionedNetworkCreate) SetCidr(s string) *ProvisionedNetworkCreate {
	pnc.mutation.SetCidr(s)
	return pnc
}

// SetVars sets the vars field.
func (pnc *ProvisionedNetworkCreate) SetVars(s []string) *ProvisionedNetworkCreate {
	pnc.mutation.SetVars(s)
	return pnc
}

// AddTagIDs adds the tag edge to Tag by ids.
func (pnc *ProvisionedNetworkCreate) AddTagIDs(ids ...int) *ProvisionedNetworkCreate {
	pnc.mutation.AddTagIDs(ids...)
	return pnc
}

// AddTag adds the tag edges to Tag.
func (pnc *ProvisionedNetworkCreate) AddTag(t ...*Tag) *ProvisionedNetworkCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pnc.AddTagIDs(ids...)
}

// AddStatuIDs adds the status edge to Status by ids.
func (pnc *ProvisionedNetworkCreate) AddStatuIDs(ids ...int) *ProvisionedNetworkCreate {
	pnc.mutation.AddStatuIDs(ids...)
	return pnc
}

// AddStatus adds the status edges to Status.
func (pnc *ProvisionedNetworkCreate) AddStatus(s ...*Status) *ProvisionedNetworkCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pnc.AddStatuIDs(ids...)
}

// AddNetworkIDs adds the network edge to Network by ids.
func (pnc *ProvisionedNetworkCreate) AddNetworkIDs(ids ...int) *ProvisionedNetworkCreate {
	pnc.mutation.AddNetworkIDs(ids...)
	return pnc
}

// AddNetwork adds the network edges to Network.
func (pnc *ProvisionedNetworkCreate) AddNetwork(n ...*Network) *ProvisionedNetworkCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return pnc.AddNetworkIDs(ids...)
}

// AddBuildIDs adds the build edge to Build by ids.
func (pnc *ProvisionedNetworkCreate) AddBuildIDs(ids ...int) *ProvisionedNetworkCreate {
	pnc.mutation.AddBuildIDs(ids...)
	return pnc
}

// AddBuild adds the build edges to Build.
func (pnc *ProvisionedNetworkCreate) AddBuild(b ...*Build) *ProvisionedNetworkCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pnc.AddBuildIDs(ids...)
}

// AddProvisionedNetworkToTeamIDs adds the ProvisionedNetworkToTeam edge to Team by ids.
func (pnc *ProvisionedNetworkCreate) AddProvisionedNetworkToTeamIDs(ids ...int) *ProvisionedNetworkCreate {
	pnc.mutation.AddProvisionedNetworkToTeamIDs(ids...)
	return pnc
}

// AddProvisionedNetworkToTeam adds the ProvisionedNetworkToTeam edges to Team.
func (pnc *ProvisionedNetworkCreate) AddProvisionedNetworkToTeam(t ...*Team) *ProvisionedNetworkCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pnc.AddProvisionedNetworkToTeamIDs(ids...)
}

// AddProvisionedHostIDs adds the provisioned_hosts edge to ProvisionedHost by ids.
func (pnc *ProvisionedNetworkCreate) AddProvisionedHostIDs(ids ...int) *ProvisionedNetworkCreate {
	pnc.mutation.AddProvisionedHostIDs(ids...)
	return pnc
}

// AddProvisionedHosts adds the provisioned_hosts edges to ProvisionedHost.
func (pnc *ProvisionedNetworkCreate) AddProvisionedHosts(p ...*ProvisionedHost) *ProvisionedNetworkCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pnc.AddProvisionedHostIDs(ids...)
}

// Mutation returns the ProvisionedNetworkMutation object of the builder.
func (pnc *ProvisionedNetworkCreate) Mutation() *ProvisionedNetworkMutation {
	return pnc.mutation
}

// Save creates the ProvisionedNetwork in the database.
func (pnc *ProvisionedNetworkCreate) Save(ctx context.Context) (*ProvisionedNetwork, error) {
	var (
		err  error
		node *ProvisionedNetwork
	)
	if len(pnc.hooks) == 0 {
		if err = pnc.check(); err != nil {
			return nil, err
		}
		node, err = pnc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProvisionedNetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pnc.check(); err != nil {
				return nil, err
			}
			pnc.mutation = mutation
			node, err = pnc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pnc.hooks) - 1; i >= 0; i-- {
			mut = pnc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pnc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pnc *ProvisionedNetworkCreate) SaveX(ctx context.Context) *ProvisionedNetwork {
	v, err := pnc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (pnc *ProvisionedNetworkCreate) check() error {
	if _, ok := pnc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := pnc.mutation.Cidr(); !ok {
		return &ValidationError{Name: "cidr", err: errors.New("ent: missing required field \"cidr\"")}
	}
	if _, ok := pnc.mutation.Vars(); !ok {
		return &ValidationError{Name: "vars", err: errors.New("ent: missing required field \"vars\"")}
	}
	return nil
}

func (pnc *ProvisionedNetworkCreate) sqlSave(ctx context.Context) (*ProvisionedNetwork, error) {
	_node, _spec := pnc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pnc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pnc *ProvisionedNetworkCreate) createSpec() (*ProvisionedNetwork, *sqlgraph.CreateSpec) {
	var (
		_node = &ProvisionedNetwork{config: pnc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: provisionednetwork.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: provisionednetwork.FieldID,
			},
		}
	)
	if value, ok := pnc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provisionednetwork.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pnc.mutation.Cidr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provisionednetwork.FieldCidr,
		})
		_node.Cidr = value
	}
	if value, ok := pnc.mutation.Vars(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: provisionednetwork.FieldVars,
		})
		_node.Vars = value
	}
	if nodes := pnc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provisionednetwork.TagTable,
			Columns: []string{provisionednetwork.TagColumn},
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
	if nodes := pnc.mutation.StatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provisionednetwork.StatusTable,
			Columns: []string{provisionednetwork.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: status.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pnc.mutation.NetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provisionednetwork.NetworkTable,
			Columns: []string{provisionednetwork.NetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: network.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pnc.mutation.BuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provisionednetwork.BuildTable,
			Columns: []string{provisionednetwork.BuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: build.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pnc.mutation.ProvisionedNetworkToTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   provisionednetwork.ProvisionedNetworkToTeamTable,
			Columns: provisionednetwork.ProvisionedNetworkToTeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pnc.mutation.ProvisionedHostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   provisionednetwork.ProvisionedHostsTable,
			Columns: provisionednetwork.ProvisionedHostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: provisionedhost.FieldID,
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

// ProvisionedNetworkCreateBulk is the builder for creating a bulk of ProvisionedNetwork entities.
type ProvisionedNetworkCreateBulk struct {
	config
	builders []*ProvisionedNetworkCreate
}

// Save creates the ProvisionedNetwork entities in the database.
func (pncb *ProvisionedNetworkCreateBulk) Save(ctx context.Context) ([]*ProvisionedNetwork, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pncb.builders))
	nodes := make([]*ProvisionedNetwork, len(pncb.builders))
	mutators := make([]Mutator, len(pncb.builders))
	for i := range pncb.builders {
		func(i int, root context.Context) {
			builder := pncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProvisionedNetworkMutation)
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
					_, err = mutators[i+1].Mutate(root, pncb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pncb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (pncb *ProvisionedNetworkCreateBulk) SaveX(ctx context.Context) []*ProvisionedNetwork {
	v, err := pncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
