// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/remotefile"
	"github.com/gen0cide/laforge/ent/tag"
)

// RemoteFileUpdate is the builder for updating RemoteFile entities.
type RemoteFileUpdate struct {
	config
	hooks    []Hook
	mutation *RemoteFileMutation
}

// Where adds a new predicate for the builder.
func (rfu *RemoteFileUpdate) Where(ps ...predicate.RemoteFile) *RemoteFileUpdate {
	rfu.mutation.predicates = append(rfu.mutation.predicates, ps...)
	return rfu
}

// SetSourceType sets the source_type field.
func (rfu *RemoteFileUpdate) SetSourceType(s string) *RemoteFileUpdate {
	rfu.mutation.SetSourceType(s)
	return rfu
}

// SetSource sets the source field.
func (rfu *RemoteFileUpdate) SetSource(s string) *RemoteFileUpdate {
	rfu.mutation.SetSource(s)
	return rfu
}

// SetDestination sets the destination field.
func (rfu *RemoteFileUpdate) SetDestination(s string) *RemoteFileUpdate {
	rfu.mutation.SetDestination(s)
	return rfu
}

// SetVars sets the vars field.
func (rfu *RemoteFileUpdate) SetVars(s []string) *RemoteFileUpdate {
	rfu.mutation.SetVars(s)
	return rfu
}

// SetTemplate sets the template field.
func (rfu *RemoteFileUpdate) SetTemplate(b bool) *RemoteFileUpdate {
	rfu.mutation.SetTemplate(b)
	return rfu
}

// SetPerms sets the perms field.
func (rfu *RemoteFileUpdate) SetPerms(s string) *RemoteFileUpdate {
	rfu.mutation.SetPerms(s)
	return rfu
}

// SetDisabled sets the disabled field.
func (rfu *RemoteFileUpdate) SetDisabled(b bool) *RemoteFileUpdate {
	rfu.mutation.SetDisabled(b)
	return rfu
}

// SetMd5 sets the md5 field.
func (rfu *RemoteFileUpdate) SetMd5(s string) *RemoteFileUpdate {
	rfu.mutation.SetMd5(s)
	return rfu
}

// SetAbsPath sets the abs_path field.
func (rfu *RemoteFileUpdate) SetAbsPath(s string) *RemoteFileUpdate {
	rfu.mutation.SetAbsPath(s)
	return rfu
}

// SetExt sets the ext field.
func (rfu *RemoteFileUpdate) SetExt(s string) *RemoteFileUpdate {
	rfu.mutation.SetExt(s)
	return rfu
}

// AddTagIDs adds the tag edge to Tag by ids.
func (rfu *RemoteFileUpdate) AddTagIDs(ids ...int) *RemoteFileUpdate {
	rfu.mutation.AddTagIDs(ids...)
	return rfu
}

// AddTag adds the tag edges to Tag.
func (rfu *RemoteFileUpdate) AddTag(t ...*Tag) *RemoteFileUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rfu.AddTagIDs(ids...)
}

// Mutation returns the RemoteFileMutation object of the builder.
func (rfu *RemoteFileUpdate) Mutation() *RemoteFileMutation {
	return rfu.mutation
}

// ClearTag clears all "tag" edges to type Tag.
func (rfu *RemoteFileUpdate) ClearTag() *RemoteFileUpdate {
	rfu.mutation.ClearTag()
	return rfu
}

// RemoveTagIDs removes the tag edge to Tag by ids.
func (rfu *RemoteFileUpdate) RemoveTagIDs(ids ...int) *RemoteFileUpdate {
	rfu.mutation.RemoveTagIDs(ids...)
	return rfu
}

// RemoveTag removes tag edges to Tag.
func (rfu *RemoteFileUpdate) RemoveTag(t ...*Tag) *RemoteFileUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rfu.RemoveTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rfu *RemoteFileUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rfu.hooks) == 0 {
		affected, err = rfu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RemoteFileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rfu.mutation = mutation
			affected, err = rfu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rfu.hooks) - 1; i >= 0; i-- {
			mut = rfu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rfu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rfu *RemoteFileUpdate) SaveX(ctx context.Context) int {
	affected, err := rfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rfu *RemoteFileUpdate) Exec(ctx context.Context) error {
	_, err := rfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rfu *RemoteFileUpdate) ExecX(ctx context.Context) {
	if err := rfu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rfu *RemoteFileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   remotefile.Table,
			Columns: remotefile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: remotefile.FieldID,
			},
		},
	}
	if ps := rfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rfu.mutation.SourceType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: remotefile.FieldSourceType,
		})
	}
	if value, ok := rfu.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: remotefile.FieldSource,
		})
	}
	if value, ok := rfu.mutation.Destination(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: remotefile.FieldDestination,
		})
	}
	if value, ok := rfu.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: remotefile.FieldVars,
		})
	}
	if value, ok := rfu.mutation.Template(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: remotefile.FieldTemplate,
		})
	}
	if value, ok := rfu.mutation.Perms(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: remotefile.FieldPerms,
		})
	}
	if value, ok := rfu.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: remotefile.FieldDisabled,
		})
	}
	if value, ok := rfu.mutation.Md5(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: remotefile.FieldMd5,
		})
	}
	if value, ok := rfu.mutation.AbsPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: remotefile.FieldAbsPath,
		})
	}
	if value, ok := rfu.mutation.Ext(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: remotefile.FieldExt,
		})
	}
	if rfu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   remotefile.TagTable,
			Columns: []string{remotefile.TagColumn},
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
	if nodes := rfu.mutation.RemovedTagIDs(); len(nodes) > 0 && !rfu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   remotefile.TagTable,
			Columns: []string{remotefile.TagColumn},
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
	if nodes := rfu.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   remotefile.TagTable,
			Columns: []string{remotefile.TagColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, rfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{remotefile.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RemoteFileUpdateOne is the builder for updating a single RemoteFile entity.
type RemoteFileUpdateOne struct {
	config
	hooks    []Hook
	mutation *RemoteFileMutation
}

// SetSourceType sets the source_type field.
func (rfuo *RemoteFileUpdateOne) SetSourceType(s string) *RemoteFileUpdateOne {
	rfuo.mutation.SetSourceType(s)
	return rfuo
}

// SetSource sets the source field.
func (rfuo *RemoteFileUpdateOne) SetSource(s string) *RemoteFileUpdateOne {
	rfuo.mutation.SetSource(s)
	return rfuo
}

// SetDestination sets the destination field.
func (rfuo *RemoteFileUpdateOne) SetDestination(s string) *RemoteFileUpdateOne {
	rfuo.mutation.SetDestination(s)
	return rfuo
}

// SetVars sets the vars field.
func (rfuo *RemoteFileUpdateOne) SetVars(s []string) *RemoteFileUpdateOne {
	rfuo.mutation.SetVars(s)
	return rfuo
}

// SetTemplate sets the template field.
func (rfuo *RemoteFileUpdateOne) SetTemplate(b bool) *RemoteFileUpdateOne {
	rfuo.mutation.SetTemplate(b)
	return rfuo
}

// SetPerms sets the perms field.
func (rfuo *RemoteFileUpdateOne) SetPerms(s string) *RemoteFileUpdateOne {
	rfuo.mutation.SetPerms(s)
	return rfuo
}

// SetDisabled sets the disabled field.
func (rfuo *RemoteFileUpdateOne) SetDisabled(b bool) *RemoteFileUpdateOne {
	rfuo.mutation.SetDisabled(b)
	return rfuo
}

// SetMd5 sets the md5 field.
func (rfuo *RemoteFileUpdateOne) SetMd5(s string) *RemoteFileUpdateOne {
	rfuo.mutation.SetMd5(s)
	return rfuo
}

// SetAbsPath sets the abs_path field.
func (rfuo *RemoteFileUpdateOne) SetAbsPath(s string) *RemoteFileUpdateOne {
	rfuo.mutation.SetAbsPath(s)
	return rfuo
}

// SetExt sets the ext field.
func (rfuo *RemoteFileUpdateOne) SetExt(s string) *RemoteFileUpdateOne {
	rfuo.mutation.SetExt(s)
	return rfuo
}

// AddTagIDs adds the tag edge to Tag by ids.
func (rfuo *RemoteFileUpdateOne) AddTagIDs(ids ...int) *RemoteFileUpdateOne {
	rfuo.mutation.AddTagIDs(ids...)
	return rfuo
}

// AddTag adds the tag edges to Tag.
func (rfuo *RemoteFileUpdateOne) AddTag(t ...*Tag) *RemoteFileUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rfuo.AddTagIDs(ids...)
}

// Mutation returns the RemoteFileMutation object of the builder.
func (rfuo *RemoteFileUpdateOne) Mutation() *RemoteFileMutation {
	return rfuo.mutation
}

// ClearTag clears all "tag" edges to type Tag.
func (rfuo *RemoteFileUpdateOne) ClearTag() *RemoteFileUpdateOne {
	rfuo.mutation.ClearTag()
	return rfuo
}

// RemoveTagIDs removes the tag edge to Tag by ids.
func (rfuo *RemoteFileUpdateOne) RemoveTagIDs(ids ...int) *RemoteFileUpdateOne {
	rfuo.mutation.RemoveTagIDs(ids...)
	return rfuo
}

// RemoveTag removes tag edges to Tag.
func (rfuo *RemoteFileUpdateOne) RemoveTag(t ...*Tag) *RemoteFileUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rfuo.RemoveTagIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (rfuo *RemoteFileUpdateOne) Save(ctx context.Context) (*RemoteFile, error) {
	var (
		err  error
		node *RemoteFile
	)
	if len(rfuo.hooks) == 0 {
		node, err = rfuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RemoteFileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rfuo.mutation = mutation
			node, err = rfuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rfuo.hooks) - 1; i >= 0; i-- {
			mut = rfuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rfuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rfuo *RemoteFileUpdateOne) SaveX(ctx context.Context) *RemoteFile {
	node, err := rfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rfuo *RemoteFileUpdateOne) Exec(ctx context.Context) error {
	_, err := rfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rfuo *RemoteFileUpdateOne) ExecX(ctx context.Context) {
	if err := rfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rfuo *RemoteFileUpdateOne) sqlSave(ctx context.Context) (_node *RemoteFile, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   remotefile.Table,
			Columns: remotefile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: remotefile.FieldID,
			},
		},
	}
	id, ok := rfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing RemoteFile.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := rfuo.mutation.SourceType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: remotefile.FieldSourceType,
		})
	}
	if value, ok := rfuo.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: remotefile.FieldSource,
		})
	}
	if value, ok := rfuo.mutation.Destination(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: remotefile.FieldDestination,
		})
	}
	if value, ok := rfuo.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: remotefile.FieldVars,
		})
	}
	if value, ok := rfuo.mutation.Template(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: remotefile.FieldTemplate,
		})
	}
	if value, ok := rfuo.mutation.Perms(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: remotefile.FieldPerms,
		})
	}
	if value, ok := rfuo.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: remotefile.FieldDisabled,
		})
	}
	if value, ok := rfuo.mutation.Md5(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: remotefile.FieldMd5,
		})
	}
	if value, ok := rfuo.mutation.AbsPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: remotefile.FieldAbsPath,
		})
	}
	if value, ok := rfuo.mutation.Ext(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: remotefile.FieldExt,
		})
	}
	if rfuo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   remotefile.TagTable,
			Columns: []string{remotefile.TagColumn},
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
	if nodes := rfuo.mutation.RemovedTagIDs(); len(nodes) > 0 && !rfuo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   remotefile.TagTable,
			Columns: []string{remotefile.TagColumn},
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
	if nodes := rfuo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   remotefile.TagTable,
			Columns: []string{remotefile.TagColumn},
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
	_node = &RemoteFile{config: rfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, rfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{remotefile.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
