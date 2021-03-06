// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/finding"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/script"
	"github.com/gen0cide/laforge/ent/tag"
	"github.com/gen0cide/laforge/ent/user"
)

// ScriptUpdate is the builder for updating Script entities.
type ScriptUpdate struct {
	config
	hooks    []Hook
	mutation *ScriptMutation
}

// Where adds a new predicate for the builder.
func (su *ScriptUpdate) Where(ps ...predicate.Script) *ScriptUpdate {
	su.mutation.predicates = append(su.mutation.predicates, ps...)
	return su
}

// SetName sets the name field.
func (su *ScriptUpdate) SetName(s string) *ScriptUpdate {
	su.mutation.SetName(s)
	return su
}

// SetLanguage sets the language field.
func (su *ScriptUpdate) SetLanguage(s string) *ScriptUpdate {
	su.mutation.SetLanguage(s)
	return su
}

// SetDescription sets the description field.
func (su *ScriptUpdate) SetDescription(s string) *ScriptUpdate {
	su.mutation.SetDescription(s)
	return su
}

// SetSource sets the source field.
func (su *ScriptUpdate) SetSource(s string) *ScriptUpdate {
	su.mutation.SetSource(s)
	return su
}

// SetSourceType sets the source_type field.
func (su *ScriptUpdate) SetSourceType(s string) *ScriptUpdate {
	su.mutation.SetSourceType(s)
	return su
}

// SetCooldown sets the cooldown field.
func (su *ScriptUpdate) SetCooldown(i int) *ScriptUpdate {
	su.mutation.ResetCooldown()
	su.mutation.SetCooldown(i)
	return su
}

// AddCooldown adds i to cooldown.
func (su *ScriptUpdate) AddCooldown(i int) *ScriptUpdate {
	su.mutation.AddCooldown(i)
	return su
}

// SetTimeout sets the timeout field.
func (su *ScriptUpdate) SetTimeout(i int) *ScriptUpdate {
	su.mutation.ResetTimeout()
	su.mutation.SetTimeout(i)
	return su
}

// AddTimeout adds i to timeout.
func (su *ScriptUpdate) AddTimeout(i int) *ScriptUpdate {
	su.mutation.AddTimeout(i)
	return su
}

// SetIgnoreErrors sets the ignore_errors field.
func (su *ScriptUpdate) SetIgnoreErrors(b bool) *ScriptUpdate {
	su.mutation.SetIgnoreErrors(b)
	return su
}

// SetArgs sets the args field.
func (su *ScriptUpdate) SetArgs(s []string) *ScriptUpdate {
	su.mutation.SetArgs(s)
	return su
}

// SetDisabled sets the disabled field.
func (su *ScriptUpdate) SetDisabled(b bool) *ScriptUpdate {
	su.mutation.SetDisabled(b)
	return su
}

// SetVars sets the vars field.
func (su *ScriptUpdate) SetVars(m map[string]string) *ScriptUpdate {
	su.mutation.SetVars(m)
	return su
}

// SetAbsPath sets the abs_path field.
func (su *ScriptUpdate) SetAbsPath(s string) *ScriptUpdate {
	su.mutation.SetAbsPath(s)
	return su
}

// AddTagIDs adds the tag edge to Tag by ids.
func (su *ScriptUpdate) AddTagIDs(ids ...int) *ScriptUpdate {
	su.mutation.AddTagIDs(ids...)
	return su
}

// AddTag adds the tag edges to Tag.
func (su *ScriptUpdate) AddTag(t ...*Tag) *ScriptUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.AddTagIDs(ids...)
}

// AddMaintainerIDs adds the maintainer edge to User by ids.
func (su *ScriptUpdate) AddMaintainerIDs(ids ...int) *ScriptUpdate {
	su.mutation.AddMaintainerIDs(ids...)
	return su
}

// AddMaintainer adds the maintainer edges to User.
func (su *ScriptUpdate) AddMaintainer(u ...*User) *ScriptUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.AddMaintainerIDs(ids...)
}

// AddFindingIDs adds the finding edge to Finding by ids.
func (su *ScriptUpdate) AddFindingIDs(ids ...int) *ScriptUpdate {
	su.mutation.AddFindingIDs(ids...)
	return su
}

// AddFinding adds the finding edges to Finding.
func (su *ScriptUpdate) AddFinding(f ...*Finding) *ScriptUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return su.AddFindingIDs(ids...)
}

// Mutation returns the ScriptMutation object of the builder.
func (su *ScriptUpdate) Mutation() *ScriptMutation {
	return su.mutation
}

// ClearTag clears all "tag" edges to type Tag.
func (su *ScriptUpdate) ClearTag() *ScriptUpdate {
	su.mutation.ClearTag()
	return su
}

// RemoveTagIDs removes the tag edge to Tag by ids.
func (su *ScriptUpdate) RemoveTagIDs(ids ...int) *ScriptUpdate {
	su.mutation.RemoveTagIDs(ids...)
	return su
}

// RemoveTag removes tag edges to Tag.
func (su *ScriptUpdate) RemoveTag(t ...*Tag) *ScriptUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.RemoveTagIDs(ids...)
}

// ClearMaintainer clears all "maintainer" edges to type User.
func (su *ScriptUpdate) ClearMaintainer() *ScriptUpdate {
	su.mutation.ClearMaintainer()
	return su
}

// RemoveMaintainerIDs removes the maintainer edge to User by ids.
func (su *ScriptUpdate) RemoveMaintainerIDs(ids ...int) *ScriptUpdate {
	su.mutation.RemoveMaintainerIDs(ids...)
	return su
}

// RemoveMaintainer removes maintainer edges to User.
func (su *ScriptUpdate) RemoveMaintainer(u ...*User) *ScriptUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.RemoveMaintainerIDs(ids...)
}

// ClearFinding clears all "finding" edges to type Finding.
func (su *ScriptUpdate) ClearFinding() *ScriptUpdate {
	su.mutation.ClearFinding()
	return su
}

// RemoveFindingIDs removes the finding edge to Finding by ids.
func (su *ScriptUpdate) RemoveFindingIDs(ids ...int) *ScriptUpdate {
	su.mutation.RemoveFindingIDs(ids...)
	return su
}

// RemoveFinding removes finding edges to Finding.
func (su *ScriptUpdate) RemoveFinding(f ...*Finding) *ScriptUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return su.RemoveFindingIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ScriptUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScriptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ScriptUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ScriptUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ScriptUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *ScriptUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   script.Table,
			Columns: script.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: script.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldName,
		})
	}
	if value, ok := su.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldLanguage,
		})
	}
	if value, ok := su.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldDescription,
		})
	}
	if value, ok := su.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldSource,
		})
	}
	if value, ok := su.mutation.SourceType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldSourceType,
		})
	}
	if value, ok := su.mutation.Cooldown(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldCooldown,
		})
	}
	if value, ok := su.mutation.AddedCooldown(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldCooldown,
		})
	}
	if value, ok := su.mutation.Timeout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldTimeout,
		})
	}
	if value, ok := su.mutation.AddedTimeout(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldTimeout,
		})
	}
	if value, ok := su.mutation.IgnoreErrors(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: script.FieldIgnoreErrors,
		})
	}
	if value, ok := su.mutation.Args(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: script.FieldArgs,
		})
	}
	if value, ok := su.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: script.FieldDisabled,
		})
	}
	if value, ok := su.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: script.FieldVars,
		})
	}
	if value, ok := su.mutation.AbsPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldAbsPath,
		})
	}
	if su.mutation.TagCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedTagIDs(); len(nodes) > 0 && !su.mutation.TagCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.TagIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.MaintainerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedMaintainerIDs(); len(nodes) > 0 && !su.mutation.MaintainerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.MaintainerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.FindingCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedFindingIDs(); len(nodes) > 0 && !su.mutation.FindingCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.FindingIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{script.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ScriptUpdateOne is the builder for updating a single Script entity.
type ScriptUpdateOne struct {
	config
	hooks    []Hook
	mutation *ScriptMutation
}

// SetName sets the name field.
func (suo *ScriptUpdateOne) SetName(s string) *ScriptUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetLanguage sets the language field.
func (suo *ScriptUpdateOne) SetLanguage(s string) *ScriptUpdateOne {
	suo.mutation.SetLanguage(s)
	return suo
}

// SetDescription sets the description field.
func (suo *ScriptUpdateOne) SetDescription(s string) *ScriptUpdateOne {
	suo.mutation.SetDescription(s)
	return suo
}

// SetSource sets the source field.
func (suo *ScriptUpdateOne) SetSource(s string) *ScriptUpdateOne {
	suo.mutation.SetSource(s)
	return suo
}

// SetSourceType sets the source_type field.
func (suo *ScriptUpdateOne) SetSourceType(s string) *ScriptUpdateOne {
	suo.mutation.SetSourceType(s)
	return suo
}

// SetCooldown sets the cooldown field.
func (suo *ScriptUpdateOne) SetCooldown(i int) *ScriptUpdateOne {
	suo.mutation.ResetCooldown()
	suo.mutation.SetCooldown(i)
	return suo
}

// AddCooldown adds i to cooldown.
func (suo *ScriptUpdateOne) AddCooldown(i int) *ScriptUpdateOne {
	suo.mutation.AddCooldown(i)
	return suo
}

// SetTimeout sets the timeout field.
func (suo *ScriptUpdateOne) SetTimeout(i int) *ScriptUpdateOne {
	suo.mutation.ResetTimeout()
	suo.mutation.SetTimeout(i)
	return suo
}

// AddTimeout adds i to timeout.
func (suo *ScriptUpdateOne) AddTimeout(i int) *ScriptUpdateOne {
	suo.mutation.AddTimeout(i)
	return suo
}

// SetIgnoreErrors sets the ignore_errors field.
func (suo *ScriptUpdateOne) SetIgnoreErrors(b bool) *ScriptUpdateOne {
	suo.mutation.SetIgnoreErrors(b)
	return suo
}

// SetArgs sets the args field.
func (suo *ScriptUpdateOne) SetArgs(s []string) *ScriptUpdateOne {
	suo.mutation.SetArgs(s)
	return suo
}

// SetDisabled sets the disabled field.
func (suo *ScriptUpdateOne) SetDisabled(b bool) *ScriptUpdateOne {
	suo.mutation.SetDisabled(b)
	return suo
}

// SetVars sets the vars field.
func (suo *ScriptUpdateOne) SetVars(m map[string]string) *ScriptUpdateOne {
	suo.mutation.SetVars(m)
	return suo
}

// SetAbsPath sets the abs_path field.
func (suo *ScriptUpdateOne) SetAbsPath(s string) *ScriptUpdateOne {
	suo.mutation.SetAbsPath(s)
	return suo
}

// AddTagIDs adds the tag edge to Tag by ids.
func (suo *ScriptUpdateOne) AddTagIDs(ids ...int) *ScriptUpdateOne {
	suo.mutation.AddTagIDs(ids...)
	return suo
}

// AddTag adds the tag edges to Tag.
func (suo *ScriptUpdateOne) AddTag(t ...*Tag) *ScriptUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.AddTagIDs(ids...)
}

// AddMaintainerIDs adds the maintainer edge to User by ids.
func (suo *ScriptUpdateOne) AddMaintainerIDs(ids ...int) *ScriptUpdateOne {
	suo.mutation.AddMaintainerIDs(ids...)
	return suo
}

// AddMaintainer adds the maintainer edges to User.
func (suo *ScriptUpdateOne) AddMaintainer(u ...*User) *ScriptUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.AddMaintainerIDs(ids...)
}

// AddFindingIDs adds the finding edge to Finding by ids.
func (suo *ScriptUpdateOne) AddFindingIDs(ids ...int) *ScriptUpdateOne {
	suo.mutation.AddFindingIDs(ids...)
	return suo
}

// AddFinding adds the finding edges to Finding.
func (suo *ScriptUpdateOne) AddFinding(f ...*Finding) *ScriptUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return suo.AddFindingIDs(ids...)
}

// Mutation returns the ScriptMutation object of the builder.
func (suo *ScriptUpdateOne) Mutation() *ScriptMutation {
	return suo.mutation
}

// ClearTag clears all "tag" edges to type Tag.
func (suo *ScriptUpdateOne) ClearTag() *ScriptUpdateOne {
	suo.mutation.ClearTag()
	return suo
}

// RemoveTagIDs removes the tag edge to Tag by ids.
func (suo *ScriptUpdateOne) RemoveTagIDs(ids ...int) *ScriptUpdateOne {
	suo.mutation.RemoveTagIDs(ids...)
	return suo
}

// RemoveTag removes tag edges to Tag.
func (suo *ScriptUpdateOne) RemoveTag(t ...*Tag) *ScriptUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.RemoveTagIDs(ids...)
}

// ClearMaintainer clears all "maintainer" edges to type User.
func (suo *ScriptUpdateOne) ClearMaintainer() *ScriptUpdateOne {
	suo.mutation.ClearMaintainer()
	return suo
}

// RemoveMaintainerIDs removes the maintainer edge to User by ids.
func (suo *ScriptUpdateOne) RemoveMaintainerIDs(ids ...int) *ScriptUpdateOne {
	suo.mutation.RemoveMaintainerIDs(ids...)
	return suo
}

// RemoveMaintainer removes maintainer edges to User.
func (suo *ScriptUpdateOne) RemoveMaintainer(u ...*User) *ScriptUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.RemoveMaintainerIDs(ids...)
}

// ClearFinding clears all "finding" edges to type Finding.
func (suo *ScriptUpdateOne) ClearFinding() *ScriptUpdateOne {
	suo.mutation.ClearFinding()
	return suo
}

// RemoveFindingIDs removes the finding edge to Finding by ids.
func (suo *ScriptUpdateOne) RemoveFindingIDs(ids ...int) *ScriptUpdateOne {
	suo.mutation.RemoveFindingIDs(ids...)
	return suo
}

// RemoveFinding removes finding edges to Finding.
func (suo *ScriptUpdateOne) RemoveFinding(f ...*Finding) *ScriptUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return suo.RemoveFindingIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *ScriptUpdateOne) Save(ctx context.Context) (*Script, error) {
	var (
		err  error
		node *Script
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScriptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ScriptUpdateOne) SaveX(ctx context.Context) *Script {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ScriptUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ScriptUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *ScriptUpdateOne) sqlSave(ctx context.Context) (_node *Script, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   script.Table,
			Columns: script.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: script.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Script.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldName,
		})
	}
	if value, ok := suo.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldLanguage,
		})
	}
	if value, ok := suo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldDescription,
		})
	}
	if value, ok := suo.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldSource,
		})
	}
	if value, ok := suo.mutation.SourceType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldSourceType,
		})
	}
	if value, ok := suo.mutation.Cooldown(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldCooldown,
		})
	}
	if value, ok := suo.mutation.AddedCooldown(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldCooldown,
		})
	}
	if value, ok := suo.mutation.Timeout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldTimeout,
		})
	}
	if value, ok := suo.mutation.AddedTimeout(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldTimeout,
		})
	}
	if value, ok := suo.mutation.IgnoreErrors(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: script.FieldIgnoreErrors,
		})
	}
	if value, ok := suo.mutation.Args(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: script.FieldArgs,
		})
	}
	if value, ok := suo.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: script.FieldDisabled,
		})
	}
	if value, ok := suo.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: script.FieldVars,
		})
	}
	if value, ok := suo.mutation.AbsPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldAbsPath,
		})
	}
	if suo.mutation.TagCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedTagIDs(); len(nodes) > 0 && !suo.mutation.TagCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.TagIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.MaintainerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedMaintainerIDs(); len(nodes) > 0 && !suo.mutation.MaintainerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.MaintainerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.FindingCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedFindingIDs(); len(nodes) > 0 && !suo.mutation.FindingCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.FindingIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Script{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{script.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
