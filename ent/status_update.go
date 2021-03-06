// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/status"
	"github.com/gen0cide/laforge/ent/tag"
)

// StatusUpdate is the builder for updating Status entities.
type StatusUpdate struct {
	config
	hooks    []Hook
	mutation *StatusMutation
}

// Where adds a new predicate for the builder.
func (su *StatusUpdate) Where(ps ...predicate.Status) *StatusUpdate {
	su.mutation.predicates = append(su.mutation.predicates, ps...)
	return su
}

// SetState sets the state field.
func (su *StatusUpdate) SetState(s status.State) *StatusUpdate {
	su.mutation.SetState(s)
	return su
}

// SetStartedAt sets the started_at field.
func (su *StatusUpdate) SetStartedAt(t time.Time) *StatusUpdate {
	su.mutation.SetStartedAt(t)
	return su
}

// SetEndedAt sets the ended_at field.
func (su *StatusUpdate) SetEndedAt(t time.Time) *StatusUpdate {
	su.mutation.SetEndedAt(t)
	return su
}

// SetFailed sets the failed field.
func (su *StatusUpdate) SetFailed(b bool) *StatusUpdate {
	su.mutation.SetFailed(b)
	return su
}

// SetCompleted sets the completed field.
func (su *StatusUpdate) SetCompleted(b bool) *StatusUpdate {
	su.mutation.SetCompleted(b)
	return su
}

// SetError sets the error field.
func (su *StatusUpdate) SetError(s string) *StatusUpdate {
	su.mutation.SetError(s)
	return su
}

// AddTagIDs adds the tag edge to Tag by ids.
func (su *StatusUpdate) AddTagIDs(ids ...int) *StatusUpdate {
	su.mutation.AddTagIDs(ids...)
	return su
}

// AddTag adds the tag edges to Tag.
func (su *StatusUpdate) AddTag(t ...*Tag) *StatusUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.AddTagIDs(ids...)
}

// Mutation returns the StatusMutation object of the builder.
func (su *StatusUpdate) Mutation() *StatusMutation {
	return su.mutation
}

// ClearTag clears all "tag" edges to type Tag.
func (su *StatusUpdate) ClearTag() *StatusUpdate {
	su.mutation.ClearTag()
	return su
}

// RemoveTagIDs removes the tag edge to Tag by ids.
func (su *StatusUpdate) RemoveTagIDs(ids ...int) *StatusUpdate {
	su.mutation.RemoveTagIDs(ids...)
	return su
}

// RemoveTag removes tag edges to Tag.
func (su *StatusUpdate) RemoveTag(t ...*Tag) *StatusUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.RemoveTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StatusUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
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
func (su *StatusUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StatusUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StatusUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StatusUpdate) check() error {
	if v, ok := su.mutation.State(); ok {
		if err := status.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	return nil
}

func (su *StatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   status.Table,
			Columns: status.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: status.FieldID,
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
	if value, ok := su.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: status.FieldState,
		})
	}
	if value, ok := su.mutation.StartedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: status.FieldStartedAt,
		})
	}
	if value, ok := su.mutation.EndedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: status.FieldEndedAt,
		})
	}
	if value, ok := su.mutation.Failed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: status.FieldFailed,
		})
	}
	if value, ok := su.mutation.Completed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: status.FieldCompleted,
		})
	}
	if value, ok := su.mutation.Error(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: status.FieldError,
		})
	}
	if su.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   status.TagTable,
			Columns: []string{status.TagColumn},
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
			Table:   status.TagTable,
			Columns: []string{status.TagColumn},
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
			Table:   status.TagTable,
			Columns: []string{status.TagColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{status.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// StatusUpdateOne is the builder for updating a single Status entity.
type StatusUpdateOne struct {
	config
	hooks    []Hook
	mutation *StatusMutation
}

// SetState sets the state field.
func (suo *StatusUpdateOne) SetState(s status.State) *StatusUpdateOne {
	suo.mutation.SetState(s)
	return suo
}

// SetStartedAt sets the started_at field.
func (suo *StatusUpdateOne) SetStartedAt(t time.Time) *StatusUpdateOne {
	suo.mutation.SetStartedAt(t)
	return suo
}

// SetEndedAt sets the ended_at field.
func (suo *StatusUpdateOne) SetEndedAt(t time.Time) *StatusUpdateOne {
	suo.mutation.SetEndedAt(t)
	return suo
}

// SetFailed sets the failed field.
func (suo *StatusUpdateOne) SetFailed(b bool) *StatusUpdateOne {
	suo.mutation.SetFailed(b)
	return suo
}

// SetCompleted sets the completed field.
func (suo *StatusUpdateOne) SetCompleted(b bool) *StatusUpdateOne {
	suo.mutation.SetCompleted(b)
	return suo
}

// SetError sets the error field.
func (suo *StatusUpdateOne) SetError(s string) *StatusUpdateOne {
	suo.mutation.SetError(s)
	return suo
}

// AddTagIDs adds the tag edge to Tag by ids.
func (suo *StatusUpdateOne) AddTagIDs(ids ...int) *StatusUpdateOne {
	suo.mutation.AddTagIDs(ids...)
	return suo
}

// AddTag adds the tag edges to Tag.
func (suo *StatusUpdateOne) AddTag(t ...*Tag) *StatusUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.AddTagIDs(ids...)
}

// Mutation returns the StatusMutation object of the builder.
func (suo *StatusUpdateOne) Mutation() *StatusMutation {
	return suo.mutation
}

// ClearTag clears all "tag" edges to type Tag.
func (suo *StatusUpdateOne) ClearTag() *StatusUpdateOne {
	suo.mutation.ClearTag()
	return suo
}

// RemoveTagIDs removes the tag edge to Tag by ids.
func (suo *StatusUpdateOne) RemoveTagIDs(ids ...int) *StatusUpdateOne {
	suo.mutation.RemoveTagIDs(ids...)
	return suo
}

// RemoveTag removes tag edges to Tag.
func (suo *StatusUpdateOne) RemoveTag(t ...*Tag) *StatusUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.RemoveTagIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *StatusUpdateOne) Save(ctx context.Context) (*Status, error) {
	var (
		err  error
		node *Status
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
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
func (suo *StatusUpdateOne) SaveX(ctx context.Context) *Status {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StatusUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StatusUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StatusUpdateOne) check() error {
	if v, ok := suo.mutation.State(); ok {
		if err := status.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	return nil
}

func (suo *StatusUpdateOne) sqlSave(ctx context.Context) (_node *Status, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   status.Table,
			Columns: status.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: status.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Status.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: status.FieldState,
		})
	}
	if value, ok := suo.mutation.StartedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: status.FieldStartedAt,
		})
	}
	if value, ok := suo.mutation.EndedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: status.FieldEndedAt,
		})
	}
	if value, ok := suo.mutation.Failed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: status.FieldFailed,
		})
	}
	if value, ok := suo.mutation.Completed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: status.FieldCompleted,
		})
	}
	if value, ok := suo.mutation.Error(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: status.FieldError,
		})
	}
	if suo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   status.TagTable,
			Columns: []string{status.TagColumn},
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
			Table:   status.TagTable,
			Columns: []string{status.TagColumn},
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
			Table:   status.TagTable,
			Columns: []string{status.TagColumn},
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
	_node = &Status{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{status.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
