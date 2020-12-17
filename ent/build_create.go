// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/tag"
	"github.com/gen0cide/laforge/ent/team"
	"github.com/gen0cide/laforge/ent/user"
)

// BuildCreate is the builder for creating a Build entity.
type BuildCreate struct {
	config
	mutation *BuildMutation
	hooks    []Hook
}

// SetRevision sets the revision field.
func (bc *BuildCreate) SetRevision(i int) *BuildCreate {
	bc.mutation.SetRevision(i)
	return bc
}

// SetConfig sets the config field.
func (bc *BuildCreate) SetConfig(s []string) *BuildCreate {
	bc.mutation.SetConfig(s)
	return bc
}

// AddUserIDs adds the user edge to User by ids.
func (bc *BuildCreate) AddUserIDs(ids ...int) *BuildCreate {
	bc.mutation.AddUserIDs(ids...)
	return bc
}

// AddUser adds the user edges to User.
func (bc *BuildCreate) AddUser(u ...*User) *BuildCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bc.AddUserIDs(ids...)
}

// AddTagIDs adds the tag edge to Tag by ids.
func (bc *BuildCreate) AddTagIDs(ids ...int) *BuildCreate {
	bc.mutation.AddTagIDs(ids...)
	return bc
}

// AddTag adds the tag edges to Tag.
func (bc *BuildCreate) AddTag(t ...*Tag) *BuildCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bc.AddTagIDs(ids...)
}

// AddTeamIDs adds the team edge to Team by ids.
func (bc *BuildCreate) AddTeamIDs(ids ...int) *BuildCreate {
	bc.mutation.AddTeamIDs(ids...)
	return bc
}

// AddTeam adds the team edges to Team.
func (bc *BuildCreate) AddTeam(t ...*Team) *BuildCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bc.AddTeamIDs(ids...)
}

// Mutation returns the BuildMutation object of the builder.
func (bc *BuildCreate) Mutation() *BuildMutation {
	return bc.mutation
}

// Save creates the Build in the database.
func (bc *BuildCreate) Save(ctx context.Context) (*Build, error) {
	var (
		err  error
		node *Build
	)
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BuildMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BuildCreate) SaveX(ctx context.Context) *Build {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (bc *BuildCreate) check() error {
	if _, ok := bc.mutation.Revision(); !ok {
		return &ValidationError{Name: "revision", err: errors.New("ent: missing required field \"revision\"")}
	}
	if v, ok := bc.mutation.Revision(); ok {
		if err := build.RevisionValidator(v); err != nil {
			return &ValidationError{Name: "revision", err: fmt.Errorf("ent: validator failed for field \"revision\": %w", err)}
		}
	}
	if _, ok := bc.mutation.Config(); !ok {
		return &ValidationError{Name: "config", err: errors.New("ent: missing required field \"config\"")}
	}
	return nil
}

func (bc *BuildCreate) sqlSave(ctx context.Context) (*Build, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BuildCreate) createSpec() (*Build, *sqlgraph.CreateSpec) {
	var (
		_node = &Build{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: build.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: build.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.Revision(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: build.FieldRevision,
		})
		_node.Revision = value
	}
	if value, ok := bc.mutation.Config(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: build.FieldConfig,
		})
		_node.Config = value
	}
	if nodes := bc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   build.UserTable,
			Columns: []string{build.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   build.TagTable,
			Columns: []string{build.TagColumn},
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
	if nodes := bc.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.TeamTable,
			Columns: build.TeamPrimaryKey,
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
	return _node, _spec
}

// BuildCreateBulk is the builder for creating a bulk of Build entities.
type BuildCreateBulk struct {
	config
	builders []*BuildCreate
}

// Save creates the Build entities in the database.
func (bcb *BuildCreateBulk) Save(ctx context.Context) ([]*Build, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Build, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BuildMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (bcb *BuildCreateBulk) SaveX(ctx context.Context) []*Build {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
