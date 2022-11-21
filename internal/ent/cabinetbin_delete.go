// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/auroraride/cabservd/internal/ent/cabinetbin"
	"github.com/auroraride/cabservd/internal/ent/predicate"
)

// CabinetBinDelete is the builder for deleting a CabinetBin entity.
type CabinetBinDelete struct {
	config
	hooks    []Hook
	mutation *CabinetBinMutation
}

// Where appends a list predicates to the CabinetBinDelete builder.
func (cbd *CabinetBinDelete) Where(ps ...predicate.CabinetBin) *CabinetBinDelete {
	cbd.mutation.Where(ps...)
	return cbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cbd *CabinetBinDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cbd.hooks) == 0 {
		affected, err = cbd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CabinetBinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cbd.mutation = mutation
			affected, err = cbd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cbd.hooks) - 1; i >= 0; i-- {
			if cbd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cbd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cbd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbd *CabinetBinDelete) ExecX(ctx context.Context) int {
	n, err := cbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cbd *CabinetBinDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: cabinetbin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: cabinetbin.FieldID,
			},
		},
	}
	if ps := cbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CabinetBinDeleteOne is the builder for deleting a single CabinetBin entity.
type CabinetBinDeleteOne struct {
	cbd *CabinetBinDelete
}

// Exec executes the deletion query.
func (cbdo *CabinetBinDeleteOne) Exec(ctx context.Context) error {
	n, err := cbdo.cbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cabinetbin.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cbdo *CabinetBinDeleteOne) ExecX(ctx context.Context) {
	cbdo.cbd.ExecX(ctx)
}
