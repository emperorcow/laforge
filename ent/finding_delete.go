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
)

// FindingDelete is the builder for deleting a Finding entity.
type FindingDelete struct {
	config
	hooks    []Hook
	mutation *FindingMutation
}

// Where adds a new predicate to the delete builder.
func (fd *FindingDelete) Where(ps ...predicate.Finding) *FindingDelete {
	fd.mutation.predicates = append(fd.mutation.predicates, ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FindingDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fd.hooks) == 0 {
		affected, err = fd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FindingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fd.mutation = mutation
			affected, err = fd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fd.hooks) - 1; i >= 0; i-- {
			mut = fd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FindingDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FindingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: finding.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: finding.FieldID,
			},
		},
	}
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
}

// FindingDeleteOne is the builder for deleting a single Finding entity.
type FindingDeleteOne struct {
	fd *FindingDelete
}

// Exec executes the deletion query.
func (fdo *FindingDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{finding.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FindingDeleteOne) ExecX(ctx context.Context) {
	fdo.fd.ExecX(ctx)
}
