// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/disk"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/tag"
)

// DiskUpdate is the builder for updating Disk entities.
type DiskUpdate struct {
	config
	hooks    []Hook
	mutation *DiskMutation
}

// Where adds a new predicate for the builder.
func (du *DiskUpdate) Where(ps ...predicate.Disk) *DiskUpdate {
	du.mutation.predicates = append(du.mutation.predicates, ps...)
	return du
}

// SetSize sets the size field.
func (du *DiskUpdate) SetSize(i int) *DiskUpdate {
	du.mutation.ResetSize()
	du.mutation.SetSize(i)
	return du
}

// AddSize adds i to size.
func (du *DiskUpdate) AddSize(i int) *DiskUpdate {
	du.mutation.AddSize(i)
	return du
}

// AddTagIDs adds the tag edge to Tag by ids.
func (du *DiskUpdate) AddTagIDs(ids ...int) *DiskUpdate {
	du.mutation.AddTagIDs(ids...)
	return du
}

// AddTag adds the tag edges to Tag.
func (du *DiskUpdate) AddTag(t ...*Tag) *DiskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return du.AddTagIDs(ids...)
}

// Mutation returns the DiskMutation object of the builder.
func (du *DiskUpdate) Mutation() *DiskMutation {
	return du.mutation
}

// ClearTag clears all "tag" edges to type Tag.
func (du *DiskUpdate) ClearTag() *DiskUpdate {
	du.mutation.ClearTag()
	return du
}

// RemoveTagIDs removes the tag edge to Tag by ids.
func (du *DiskUpdate) RemoveTagIDs(ids ...int) *DiskUpdate {
	du.mutation.RemoveTagIDs(ids...)
	return du
}

// RemoveTag removes tag edges to Tag.
func (du *DiskUpdate) RemoveTag(t ...*Tag) *DiskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return du.RemoveTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DiskUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DiskUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DiskUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DiskUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DiskUpdate) check() error {
	if v, ok := du.mutation.Size(); ok {
		if err := disk.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf("ent: validator failed for field \"size\": %w", err)}
		}
	}
	return nil
}

func (du *DiskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   disk.Table,
			Columns: disk.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: disk.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: disk.FieldSize,
		})
	}
	if value, ok := du.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: disk.FieldSize,
		})
	}
	if du.mutation.TagCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedTagIDs(); len(nodes) > 0 && !du.mutation.TagCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.TagIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{disk.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DiskUpdateOne is the builder for updating a single Disk entity.
type DiskUpdateOne struct {
	config
	hooks    []Hook
	mutation *DiskMutation
}

// SetSize sets the size field.
func (duo *DiskUpdateOne) SetSize(i int) *DiskUpdateOne {
	duo.mutation.ResetSize()
	duo.mutation.SetSize(i)
	return duo
}

// AddSize adds i to size.
func (duo *DiskUpdateOne) AddSize(i int) *DiskUpdateOne {
	duo.mutation.AddSize(i)
	return duo
}

// AddTagIDs adds the tag edge to Tag by ids.
func (duo *DiskUpdateOne) AddTagIDs(ids ...int) *DiskUpdateOne {
	duo.mutation.AddTagIDs(ids...)
	return duo
}

// AddTag adds the tag edges to Tag.
func (duo *DiskUpdateOne) AddTag(t ...*Tag) *DiskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return duo.AddTagIDs(ids...)
}

// Mutation returns the DiskMutation object of the builder.
func (duo *DiskUpdateOne) Mutation() *DiskMutation {
	return duo.mutation
}

// ClearTag clears all "tag" edges to type Tag.
func (duo *DiskUpdateOne) ClearTag() *DiskUpdateOne {
	duo.mutation.ClearTag()
	return duo
}

// RemoveTagIDs removes the tag edge to Tag by ids.
func (duo *DiskUpdateOne) RemoveTagIDs(ids ...int) *DiskUpdateOne {
	duo.mutation.RemoveTagIDs(ids...)
	return duo
}

// RemoveTag removes tag edges to Tag.
func (duo *DiskUpdateOne) RemoveTag(t ...*Tag) *DiskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return duo.RemoveTagIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DiskUpdateOne) Save(ctx context.Context) (*Disk, error) {
	var (
		err  error
		node *Disk
	)
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DiskUpdateOne) SaveX(ctx context.Context) *Disk {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DiskUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DiskUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DiskUpdateOne) check() error {
	if v, ok := duo.mutation.Size(); ok {
		if err := disk.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf("ent: validator failed for field \"size\": %w", err)}
		}
	}
	return nil
}

func (duo *DiskUpdateOne) sqlSave(ctx context.Context) (_node *Disk, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   disk.Table,
			Columns: disk.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: disk.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Disk.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: disk.FieldSize,
		})
	}
	if value, ok := duo.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: disk.FieldSize,
		})
	}
	if duo.mutation.TagCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedTagIDs(); len(nodes) > 0 && !duo.mutation.TagCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.TagIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Disk{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{disk.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
