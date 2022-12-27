// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/ent/predicate"
)

// CabinetDelete is the builder for deleting a Cabinet entity.
type CabinetDelete struct {
	config
	hooks    []Hook
	mutation *CabinetMutation
}

// Where appends a list predicates to the CabinetDelete builder.
func (cd *CabinetDelete) Where(ps ...predicate.Cabinet) *CabinetDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CabinetDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, CabinetMutation](ctx, cd.sqlExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CabinetDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CabinetDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: cabinet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: cabinet.FieldID,
			},
		},
	}
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cd.mutation.done = true
	return affected, err
}

// CabinetDeleteOne is the builder for deleting a single Cabinet entity.
type CabinetDeleteOne struct {
	cd *CabinetDelete
}

// Exec executes the deletion query.
func (cdo *CabinetDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cabinet.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CabinetDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
