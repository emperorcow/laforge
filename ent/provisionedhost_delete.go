// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/provisionedhost"
)

// ProvisionedHostDelete is the builder for deleting a ProvisionedHost entity.
type ProvisionedHostDelete struct {
	config
	hooks    []Hook
	mutation *ProvisionedHostMutation
}

// Where adds a new predicate to the delete builder.
func (phd *ProvisionedHostDelete) Where(ps ...predicate.ProvisionedHost) *ProvisionedHostDelete {
	phd.mutation.predicates = append(phd.mutation.predicates, ps...)
	return phd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (phd *ProvisionedHostDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(phd.hooks) == 0 {
		affected, err = phd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProvisionedHostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			phd.mutation = mutation
			affected, err = phd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(phd.hooks) - 1; i >= 0; i-- {
			mut = phd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, phd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (phd *ProvisionedHostDelete) ExecX(ctx context.Context) int {
	n, err := phd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (phd *ProvisionedHostDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: provisionedhost.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: provisionedhost.FieldID,
			},
		},
	}
	if ps := phd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, phd.driver, _spec)
}

// ProvisionedHostDeleteOne is the builder for deleting a single ProvisionedHost entity.
type ProvisionedHostDeleteOne struct {
	phd *ProvisionedHostDelete
}

// Exec executes the deletion query.
func (phdo *ProvisionedHostDeleteOne) Exec(ctx context.Context) error {
	n, err := phdo.phd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{provisionedhost.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (phdo *ProvisionedHostDeleteOne) ExecX(ctx context.Context) {
	phdo.phd.ExecX(ctx)
}
