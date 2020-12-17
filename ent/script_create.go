// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/finding"
	"github.com/gen0cide/laforge/ent/script"
	"github.com/gen0cide/laforge/ent/tag"
	"github.com/gen0cide/laforge/ent/user"
)

// ScriptCreate is the builder for creating a Script entity.
type ScriptCreate struct {
	config
	mutation *ScriptMutation
	hooks    []Hook
}

// SetName sets the name field.
func (sc *ScriptCreate) SetName(s string) *ScriptCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetLanguage sets the language field.
func (sc *ScriptCreate) SetLanguage(s string) *ScriptCreate {
	sc.mutation.SetLanguage(s)
	return sc
}

// SetDescription sets the description field.
func (sc *ScriptCreate) SetDescription(s string) *ScriptCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetSource sets the source field.
func (sc *ScriptCreate) SetSource(s string) *ScriptCreate {
	sc.mutation.SetSource(s)
	return sc
}

// SetSourceType sets the source_type field.
func (sc *ScriptCreate) SetSourceType(s string) *ScriptCreate {
	sc.mutation.SetSourceType(s)
	return sc
}

// SetCooldown sets the cooldown field.
func (sc *ScriptCreate) SetCooldown(i int) *ScriptCreate {
	sc.mutation.SetCooldown(i)
	return sc
}

// SetTimeout sets the timeout field.
func (sc *ScriptCreate) SetTimeout(i int) *ScriptCreate {
	sc.mutation.SetTimeout(i)
	return sc
}

// SetIgnoreErrors sets the ignore_errors field.
func (sc *ScriptCreate) SetIgnoreErrors(b bool) *ScriptCreate {
	sc.mutation.SetIgnoreErrors(b)
	return sc
}

// SetArgs sets the args field.
func (sc *ScriptCreate) SetArgs(s []string) *ScriptCreate {
	sc.mutation.SetArgs(s)
	return sc
}

// SetDisabled sets the disabled field.
func (sc *ScriptCreate) SetDisabled(b bool) *ScriptCreate {
	sc.mutation.SetDisabled(b)
	return sc
}

// SetVars sets the vars field.
func (sc *ScriptCreate) SetVars(s []string) *ScriptCreate {
	sc.mutation.SetVars(s)
	return sc
}

// SetAbsPath sets the abs_path field.
func (sc *ScriptCreate) SetAbsPath(s string) *ScriptCreate {
	sc.mutation.SetAbsPath(s)
	return sc
}

// AddTagIDs adds the tag edge to Tag by ids.
func (sc *ScriptCreate) AddTagIDs(ids ...int) *ScriptCreate {
	sc.mutation.AddTagIDs(ids...)
	return sc
}

// AddTag adds the tag edges to Tag.
func (sc *ScriptCreate) AddTag(t ...*Tag) *ScriptCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddTagIDs(ids...)
}

// AddMaintainerIDs adds the maintainer edge to User by ids.
func (sc *ScriptCreate) AddMaintainerIDs(ids ...int) *ScriptCreate {
	sc.mutation.AddMaintainerIDs(ids...)
	return sc
}

// AddMaintainer adds the maintainer edges to User.
func (sc *ScriptCreate) AddMaintainer(u ...*User) *ScriptCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return sc.AddMaintainerIDs(ids...)
}

// AddFindingIDs adds the finding edge to Finding by ids.
func (sc *ScriptCreate) AddFindingIDs(ids ...int) *ScriptCreate {
	sc.mutation.AddFindingIDs(ids...)
	return sc
}

// AddFinding adds the finding edges to Finding.
func (sc *ScriptCreate) AddFinding(f ...*Finding) *ScriptCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return sc.AddFindingIDs(ids...)
}

// Mutation returns the ScriptMutation object of the builder.
func (sc *ScriptCreate) Mutation() *ScriptMutation {
	return sc.mutation
}

// Save creates the Script in the database.
func (sc *ScriptCreate) Save(ctx context.Context) (*Script, error) {
	var (
		err  error
		node *Script
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScriptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScriptCreate) SaveX(ctx context.Context) *Script {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScriptCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := sc.mutation.Language(); !ok {
		return &ValidationError{Name: "language", err: errors.New("ent: missing required field \"language\"")}
	}
	if _, ok := sc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	if _, ok := sc.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New("ent: missing required field \"source\"")}
	}
	if _, ok := sc.mutation.SourceType(); !ok {
		return &ValidationError{Name: "source_type", err: errors.New("ent: missing required field \"source_type\"")}
	}
	if _, ok := sc.mutation.Cooldown(); !ok {
		return &ValidationError{Name: "cooldown", err: errors.New("ent: missing required field \"cooldown\"")}
	}
	if _, ok := sc.mutation.Timeout(); !ok {
		return &ValidationError{Name: "timeout", err: errors.New("ent: missing required field \"timeout\"")}
	}
	if _, ok := sc.mutation.IgnoreErrors(); !ok {
		return &ValidationError{Name: "ignore_errors", err: errors.New("ent: missing required field \"ignore_errors\"")}
	}
	if _, ok := sc.mutation.Args(); !ok {
		return &ValidationError{Name: "args", err: errors.New("ent: missing required field \"args\"")}
	}
	if _, ok := sc.mutation.Disabled(); !ok {
		return &ValidationError{Name: "disabled", err: errors.New("ent: missing required field \"disabled\"")}
	}
	if _, ok := sc.mutation.Vars(); !ok {
		return &ValidationError{Name: "vars", err: errors.New("ent: missing required field \"vars\"")}
	}
	if _, ok := sc.mutation.AbsPath(); !ok {
		return &ValidationError{Name: "abs_path", err: errors.New("ent: missing required field \"abs_path\"")}
	}
	return nil
}

func (sc *ScriptCreate) sqlSave(ctx context.Context) (*Script, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *ScriptCreate) createSpec() (*Script, *sqlgraph.CreateSpec) {
	var (
		_node = &Script{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: script.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: script.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldName,
		})
		_node.Name = value
	}
	if value, ok := sc.mutation.Language(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldLanguage,
		})
		_node.Language = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := sc.mutation.Source(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldSource,
		})
		_node.Source = value
	}
	if value, ok := sc.mutation.SourceType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldSourceType,
		})
		_node.SourceType = value
	}
	if value, ok := sc.mutation.Cooldown(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldCooldown,
		})
		_node.Cooldown = value
	}
	if value, ok := sc.mutation.Timeout(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldTimeout,
		})
		_node.Timeout = value
	}
	if value, ok := sc.mutation.IgnoreErrors(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: script.FieldIgnoreErrors,
		})
		_node.IgnoreErrors = value
	}
	if value, ok := sc.mutation.Args(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: script.FieldArgs,
		})
		_node.Args = value
	}
	if value, ok := sc.mutation.Disabled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: script.FieldDisabled,
		})
		_node.Disabled = value
	}
	if value, ok := sc.mutation.Vars(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: script.FieldVars,
		})
		_node.Vars = value
	}
	if value, ok := sc.mutation.AbsPath(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldAbsPath,
		})
		_node.AbsPath = value
	}
	if nodes := sc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   script.TagTable,
			Columns: []string{script.TagColumn},
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
	if nodes := sc.mutation.MaintainerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   script.MaintainerTable,
			Columns: []string{script.MaintainerColumn},
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
	if nodes := sc.mutation.FindingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   script.FindingTable,
			Columns: script.FindingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: finding.FieldID,
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

// ScriptCreateBulk is the builder for creating a bulk of Script entities.
type ScriptCreateBulk struct {
	config
	builders []*ScriptCreate
}

// Save creates the Script entities in the database.
func (scb *ScriptCreateBulk) Save(ctx context.Context) ([]*Script, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Script, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScriptMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (scb *ScriptCreateBulk) SaveX(ctx context.Context) []*Script {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
