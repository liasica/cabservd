// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/auroraride/cabservd/internal/ent/predicate"
	"github.com/auroraride/cabservd/internal/ent/scan"
)

// ScanDelete is the builder for deleting a Scan entity.
type ScanDelete struct {
	config
	hooks    []Hook
	mutation *ScanMutation
}

// Where appends a list predicates to the ScanDelete builder.
func (sd *ScanDelete) Where(ps ...predicate.Scan) *ScanDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *ScanDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ScanMutation](ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *ScanDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *ScanDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(scan.Table, sqlgraph.NewFieldSpec(scan.FieldID, field.TypeUint64))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// ScanDeleteOne is the builder for deleting a single Scan entity.
type ScanDeleteOne struct {
	sd *ScanDelete
}

// Where appends a list predicates to the ScanDelete builder.
func (sdo *ScanDeleteOne) Where(ps ...predicate.Scan) *ScanDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *ScanDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{scan.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *ScanDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
