// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/disk"
	"github.com/gen0cide/laforge/ent/tag"
)

// DiskCreate is the builder for creating a Disk entity.
type DiskCreate struct {
	config
	mutation *DiskMutation
	hooks    []Hook
}

// SetSize sets the size field.
func (dc *DiskCreate) SetSize(i int) *DiskCreate {
	dc.mutation.SetSize(i)
	return dc
}

// AddTagIDs adds the tag edge to Tag by ids.
func (dc *DiskCreate) AddTagIDs(ids ...int) *DiskCreate {
	dc.mutation.AddTagIDs(ids...)
	return dc
}

// AddTag adds the tag edges to Tag.
func (dc *DiskCreate) AddTag(t ...*Tag) *DiskCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return dc.AddTagIDs(ids...)
}

// Mutation returns the DiskMutation object of the builder.
func (dc *DiskCreate) Mutation() *DiskMutation {
	return dc.mutation
}

// Save creates the Disk in the database.
func (dc *DiskCreate) Save(ctx context.Context) (*Disk, error) {
	var (
		err  error
		node *Disk
	)
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DiskCreate) SaveX(ctx context.Context) *Disk {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (dc *DiskCreate) check() error {
	if _, ok := dc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New("ent: missing required field \"size\"")}
	}
	if v, ok := dc.mutation.Size(); ok {
		if err := disk.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf("ent: validator failed for field \"size\": %w", err)}
		}
	}
	return nil
}

func (dc *DiskCreate) sqlSave(ctx context.Context) (*Disk, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dc *DiskCreate) createSpec() (*Disk, *sqlgraph.CreateSpec) {
	var (
		_node = &Disk{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: disk.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: disk.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Size(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: disk.FieldSize,
		})
		_node.Size = value
	}
	if nodes := dc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   disk.TagTable,
			Columns: []string{disk.TagColumn},
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
	return _node, _spec
}

// DiskCreateBulk is the builder for creating a bulk of Disk entities.
type DiskCreateBulk struct {
	config
	builders []*DiskCreate
}

// Save creates the Disk entities in the database.
func (dcb *DiskCreateBulk) Save(ctx context.Context) ([]*Disk, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Disk, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DiskMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (dcb *DiskCreateBulk) SaveX(ctx context.Context) []*Disk {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
