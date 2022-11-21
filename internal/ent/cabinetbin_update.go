// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/auroraride/cabservd/internal/ent/cabinetbin"
	"github.com/auroraride/cabservd/internal/ent/predicate"
)

// CabinetBinUpdate is the builder for updating CabinetBin entities.
type CabinetBinUpdate struct {
	config
	hooks     []Hook
	mutation  *CabinetBinMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CabinetBinUpdate builder.
func (cbu *CabinetBinUpdate) Where(ps ...predicate.CabinetBin) *CabinetBinUpdate {
	cbu.mutation.Where(ps...)
	return cbu
}

// SetUpdatedAt sets the "updated_at" field.
func (cbu *CabinetBinUpdate) SetUpdatedAt(t time.Time) *CabinetBinUpdate {
	cbu.mutation.SetUpdatedAt(t)
	return cbu
}

// SetUUID sets the "uuid" field.
func (cbu *CabinetBinUpdate) SetUUID(s string) *CabinetBinUpdate {
	cbu.mutation.SetUUID(s)
	return cbu
}

// SetBrand sets the "brand" field.
func (cbu *CabinetBinUpdate) SetBrand(s string) *CabinetBinUpdate {
	cbu.mutation.SetBrand(s)
	return cbu
}

// SetSn sets the "sn" field.
func (cbu *CabinetBinUpdate) SetSn(s string) *CabinetBinUpdate {
	cbu.mutation.SetSn(s)
	return cbu
}

// SetName sets the "name" field.
func (cbu *CabinetBinUpdate) SetName(s string) *CabinetBinUpdate {
	cbu.mutation.SetName(s)
	return cbu
}

// SetIndex sets the "index" field.
func (cbu *CabinetBinUpdate) SetIndex(i int) *CabinetBinUpdate {
	cbu.mutation.ResetIndex()
	cbu.mutation.SetIndex(i)
	return cbu
}

// AddIndex adds i to the "index" field.
func (cbu *CabinetBinUpdate) AddIndex(i int) *CabinetBinUpdate {
	cbu.mutation.AddIndex(i)
	return cbu
}

// SetOpen sets the "open" field.
func (cbu *CabinetBinUpdate) SetOpen(b bool) *CabinetBinUpdate {
	cbu.mutation.SetOpen(b)
	return cbu
}

// SetNillableOpen sets the "open" field if the given value is not nil.
func (cbu *CabinetBinUpdate) SetNillableOpen(b *bool) *CabinetBinUpdate {
	if b != nil {
		cbu.SetOpen(*b)
	}
	return cbu
}

// SetEnable sets the "enable" field.
func (cbu *CabinetBinUpdate) SetEnable(b bool) *CabinetBinUpdate {
	cbu.mutation.SetEnable(b)
	return cbu
}

// SetNillableEnable sets the "enable" field if the given value is not nil.
func (cbu *CabinetBinUpdate) SetNillableEnable(b *bool) *CabinetBinUpdate {
	if b != nil {
		cbu.SetEnable(*b)
	}
	return cbu
}

// SetBatterySn sets the "battery_sn" field.
func (cbu *CabinetBinUpdate) SetBatterySn(s string) *CabinetBinUpdate {
	cbu.mutation.SetBatterySn(s)
	return cbu
}

// SetNillableBatterySn sets the "battery_sn" field if the given value is not nil.
func (cbu *CabinetBinUpdate) SetNillableBatterySn(s *string) *CabinetBinUpdate {
	if s != nil {
		cbu.SetBatterySn(*s)
	}
	return cbu
}

// ClearBatterySn clears the value of the "battery_sn" field.
func (cbu *CabinetBinUpdate) ClearBatterySn() *CabinetBinUpdate {
	cbu.mutation.ClearBatterySn()
	return cbu
}

// SetVoltage sets the "voltage" field.
func (cbu *CabinetBinUpdate) SetVoltage(f float64) *CabinetBinUpdate {
	cbu.mutation.ResetVoltage()
	cbu.mutation.SetVoltage(f)
	return cbu
}

// SetNillableVoltage sets the "voltage" field if the given value is not nil.
func (cbu *CabinetBinUpdate) SetNillableVoltage(f *float64) *CabinetBinUpdate {
	if f != nil {
		cbu.SetVoltage(*f)
	}
	return cbu
}

// AddVoltage adds f to the "voltage" field.
func (cbu *CabinetBinUpdate) AddVoltage(f float64) *CabinetBinUpdate {
	cbu.mutation.AddVoltage(f)
	return cbu
}

// SetCurrent sets the "current" field.
func (cbu *CabinetBinUpdate) SetCurrent(f float64) *CabinetBinUpdate {
	cbu.mutation.ResetCurrent()
	cbu.mutation.SetCurrent(f)
	return cbu
}

// SetNillableCurrent sets the "current" field if the given value is not nil.
func (cbu *CabinetBinUpdate) SetNillableCurrent(f *float64) *CabinetBinUpdate {
	if f != nil {
		cbu.SetCurrent(*f)
	}
	return cbu
}

// AddCurrent adds f to the "current" field.
func (cbu *CabinetBinUpdate) AddCurrent(f float64) *CabinetBinUpdate {
	cbu.mutation.AddCurrent(f)
	return cbu
}

// Mutation returns the CabinetBinMutation object of the builder.
func (cbu *CabinetBinUpdate) Mutation() *CabinetBinMutation {
	return cbu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cbu *CabinetBinUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cbu.defaults()
	if len(cbu.hooks) == 0 {
		if err = cbu.check(); err != nil {
			return 0, err
		}
		affected, err = cbu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CabinetBinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cbu.check(); err != nil {
				return 0, err
			}
			cbu.mutation = mutation
			affected, err = cbu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cbu.hooks) - 1; i >= 0; i-- {
			if cbu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cbu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cbu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cbu *CabinetBinUpdate) SaveX(ctx context.Context) int {
	affected, err := cbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cbu *CabinetBinUpdate) Exec(ctx context.Context) error {
	_, err := cbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbu *CabinetBinUpdate) ExecX(ctx context.Context) {
	if err := cbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cbu *CabinetBinUpdate) defaults() {
	if _, ok := cbu.mutation.UpdatedAt(); !ok {
		v := cabinetbin.UpdateDefaultUpdatedAt()
		cbu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cbu *CabinetBinUpdate) check() error {
	if v, ok := cbu.mutation.UUID(); ok {
		if err := cabinetbin.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf(`ent: validator failed for field "CabinetBin.uuid": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cbu *CabinetBinUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CabinetBinUpdate {
	cbu.modifiers = append(cbu.modifiers, modifiers...)
	return cbu
}

func (cbu *CabinetBinUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cabinetbin.Table,
			Columns: cabinetbin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: cabinetbin.FieldID,
			},
		},
	}
	if ps := cbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cbu.mutation.UpdatedAt(); ok {
		_spec.SetField(cabinetbin.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cbu.mutation.UUID(); ok {
		_spec.SetField(cabinetbin.FieldUUID, field.TypeString, value)
	}
	if value, ok := cbu.mutation.Brand(); ok {
		_spec.SetField(cabinetbin.FieldBrand, field.TypeString, value)
	}
	if value, ok := cbu.mutation.Sn(); ok {
		_spec.SetField(cabinetbin.FieldSn, field.TypeString, value)
	}
	if value, ok := cbu.mutation.Name(); ok {
		_spec.SetField(cabinetbin.FieldName, field.TypeString, value)
	}
	if value, ok := cbu.mutation.Index(); ok {
		_spec.SetField(cabinetbin.FieldIndex, field.TypeInt, value)
	}
	if value, ok := cbu.mutation.AddedIndex(); ok {
		_spec.AddField(cabinetbin.FieldIndex, field.TypeInt, value)
	}
	if value, ok := cbu.mutation.Open(); ok {
		_spec.SetField(cabinetbin.FieldOpen, field.TypeBool, value)
	}
	if value, ok := cbu.mutation.Enable(); ok {
		_spec.SetField(cabinetbin.FieldEnable, field.TypeBool, value)
	}
	if value, ok := cbu.mutation.BatterySn(); ok {
		_spec.SetField(cabinetbin.FieldBatterySn, field.TypeString, value)
	}
	if cbu.mutation.BatterySnCleared() {
		_spec.ClearField(cabinetbin.FieldBatterySn, field.TypeString)
	}
	if value, ok := cbu.mutation.Voltage(); ok {
		_spec.SetField(cabinetbin.FieldVoltage, field.TypeFloat64, value)
	}
	if value, ok := cbu.mutation.AddedVoltage(); ok {
		_spec.AddField(cabinetbin.FieldVoltage, field.TypeFloat64, value)
	}
	if value, ok := cbu.mutation.Current(); ok {
		_spec.SetField(cabinetbin.FieldCurrent, field.TypeFloat64, value)
	}
	if value, ok := cbu.mutation.AddedCurrent(); ok {
		_spec.AddField(cabinetbin.FieldCurrent, field.TypeFloat64, value)
	}
	_spec.AddModifiers(cbu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cabinetbin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CabinetBinUpdateOne is the builder for updating a single CabinetBin entity.
type CabinetBinUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CabinetBinMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (cbuo *CabinetBinUpdateOne) SetUpdatedAt(t time.Time) *CabinetBinUpdateOne {
	cbuo.mutation.SetUpdatedAt(t)
	return cbuo
}

// SetUUID sets the "uuid" field.
func (cbuo *CabinetBinUpdateOne) SetUUID(s string) *CabinetBinUpdateOne {
	cbuo.mutation.SetUUID(s)
	return cbuo
}

// SetBrand sets the "brand" field.
func (cbuo *CabinetBinUpdateOne) SetBrand(s string) *CabinetBinUpdateOne {
	cbuo.mutation.SetBrand(s)
	return cbuo
}

// SetSn sets the "sn" field.
func (cbuo *CabinetBinUpdateOne) SetSn(s string) *CabinetBinUpdateOne {
	cbuo.mutation.SetSn(s)
	return cbuo
}

// SetName sets the "name" field.
func (cbuo *CabinetBinUpdateOne) SetName(s string) *CabinetBinUpdateOne {
	cbuo.mutation.SetName(s)
	return cbuo
}

// SetIndex sets the "index" field.
func (cbuo *CabinetBinUpdateOne) SetIndex(i int) *CabinetBinUpdateOne {
	cbuo.mutation.ResetIndex()
	cbuo.mutation.SetIndex(i)
	return cbuo
}

// AddIndex adds i to the "index" field.
func (cbuo *CabinetBinUpdateOne) AddIndex(i int) *CabinetBinUpdateOne {
	cbuo.mutation.AddIndex(i)
	return cbuo
}

// SetOpen sets the "open" field.
func (cbuo *CabinetBinUpdateOne) SetOpen(b bool) *CabinetBinUpdateOne {
	cbuo.mutation.SetOpen(b)
	return cbuo
}

// SetNillableOpen sets the "open" field if the given value is not nil.
func (cbuo *CabinetBinUpdateOne) SetNillableOpen(b *bool) *CabinetBinUpdateOne {
	if b != nil {
		cbuo.SetOpen(*b)
	}
	return cbuo
}

// SetEnable sets the "enable" field.
func (cbuo *CabinetBinUpdateOne) SetEnable(b bool) *CabinetBinUpdateOne {
	cbuo.mutation.SetEnable(b)
	return cbuo
}

// SetNillableEnable sets the "enable" field if the given value is not nil.
func (cbuo *CabinetBinUpdateOne) SetNillableEnable(b *bool) *CabinetBinUpdateOne {
	if b != nil {
		cbuo.SetEnable(*b)
	}
	return cbuo
}

// SetBatterySn sets the "battery_sn" field.
func (cbuo *CabinetBinUpdateOne) SetBatterySn(s string) *CabinetBinUpdateOne {
	cbuo.mutation.SetBatterySn(s)
	return cbuo
}

// SetNillableBatterySn sets the "battery_sn" field if the given value is not nil.
func (cbuo *CabinetBinUpdateOne) SetNillableBatterySn(s *string) *CabinetBinUpdateOne {
	if s != nil {
		cbuo.SetBatterySn(*s)
	}
	return cbuo
}

// ClearBatterySn clears the value of the "battery_sn" field.
func (cbuo *CabinetBinUpdateOne) ClearBatterySn() *CabinetBinUpdateOne {
	cbuo.mutation.ClearBatterySn()
	return cbuo
}

// SetVoltage sets the "voltage" field.
func (cbuo *CabinetBinUpdateOne) SetVoltage(f float64) *CabinetBinUpdateOne {
	cbuo.mutation.ResetVoltage()
	cbuo.mutation.SetVoltage(f)
	return cbuo
}

// SetNillableVoltage sets the "voltage" field if the given value is not nil.
func (cbuo *CabinetBinUpdateOne) SetNillableVoltage(f *float64) *CabinetBinUpdateOne {
	if f != nil {
		cbuo.SetVoltage(*f)
	}
	return cbuo
}

// AddVoltage adds f to the "voltage" field.
func (cbuo *CabinetBinUpdateOne) AddVoltage(f float64) *CabinetBinUpdateOne {
	cbuo.mutation.AddVoltage(f)
	return cbuo
}

// SetCurrent sets the "current" field.
func (cbuo *CabinetBinUpdateOne) SetCurrent(f float64) *CabinetBinUpdateOne {
	cbuo.mutation.ResetCurrent()
	cbuo.mutation.SetCurrent(f)
	return cbuo
}

// SetNillableCurrent sets the "current" field if the given value is not nil.
func (cbuo *CabinetBinUpdateOne) SetNillableCurrent(f *float64) *CabinetBinUpdateOne {
	if f != nil {
		cbuo.SetCurrent(*f)
	}
	return cbuo
}

// AddCurrent adds f to the "current" field.
func (cbuo *CabinetBinUpdateOne) AddCurrent(f float64) *CabinetBinUpdateOne {
	cbuo.mutation.AddCurrent(f)
	return cbuo
}

// Mutation returns the CabinetBinMutation object of the builder.
func (cbuo *CabinetBinUpdateOne) Mutation() *CabinetBinMutation {
	return cbuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cbuo *CabinetBinUpdateOne) Select(field string, fields ...string) *CabinetBinUpdateOne {
	cbuo.fields = append([]string{field}, fields...)
	return cbuo
}

// Save executes the query and returns the updated CabinetBin entity.
func (cbuo *CabinetBinUpdateOne) Save(ctx context.Context) (*CabinetBin, error) {
	var (
		err  error
		node *CabinetBin
	)
	cbuo.defaults()
	if len(cbuo.hooks) == 0 {
		if err = cbuo.check(); err != nil {
			return nil, err
		}
		node, err = cbuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CabinetBinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cbuo.check(); err != nil {
				return nil, err
			}
			cbuo.mutation = mutation
			node, err = cbuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cbuo.hooks) - 1; i >= 0; i-- {
			if cbuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cbuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cbuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CabinetBin)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CabinetBinMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cbuo *CabinetBinUpdateOne) SaveX(ctx context.Context) *CabinetBin {
	node, err := cbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cbuo *CabinetBinUpdateOne) Exec(ctx context.Context) error {
	_, err := cbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbuo *CabinetBinUpdateOne) ExecX(ctx context.Context) {
	if err := cbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cbuo *CabinetBinUpdateOne) defaults() {
	if _, ok := cbuo.mutation.UpdatedAt(); !ok {
		v := cabinetbin.UpdateDefaultUpdatedAt()
		cbuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cbuo *CabinetBinUpdateOne) check() error {
	if v, ok := cbuo.mutation.UUID(); ok {
		if err := cabinetbin.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf(`ent: validator failed for field "CabinetBin.uuid": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cbuo *CabinetBinUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CabinetBinUpdateOne {
	cbuo.modifiers = append(cbuo.modifiers, modifiers...)
	return cbuo
}

func (cbuo *CabinetBinUpdateOne) sqlSave(ctx context.Context) (_node *CabinetBin, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cabinetbin.Table,
			Columns: cabinetbin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: cabinetbin.FieldID,
			},
		},
	}
	id, ok := cbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CabinetBin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cabinetbin.FieldID)
		for _, f := range fields {
			if !cabinetbin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cabinetbin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cbuo.mutation.UpdatedAt(); ok {
		_spec.SetField(cabinetbin.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cbuo.mutation.UUID(); ok {
		_spec.SetField(cabinetbin.FieldUUID, field.TypeString, value)
	}
	if value, ok := cbuo.mutation.Brand(); ok {
		_spec.SetField(cabinetbin.FieldBrand, field.TypeString, value)
	}
	if value, ok := cbuo.mutation.Sn(); ok {
		_spec.SetField(cabinetbin.FieldSn, field.TypeString, value)
	}
	if value, ok := cbuo.mutation.Name(); ok {
		_spec.SetField(cabinetbin.FieldName, field.TypeString, value)
	}
	if value, ok := cbuo.mutation.Index(); ok {
		_spec.SetField(cabinetbin.FieldIndex, field.TypeInt, value)
	}
	if value, ok := cbuo.mutation.AddedIndex(); ok {
		_spec.AddField(cabinetbin.FieldIndex, field.TypeInt, value)
	}
	if value, ok := cbuo.mutation.Open(); ok {
		_spec.SetField(cabinetbin.FieldOpen, field.TypeBool, value)
	}
	if value, ok := cbuo.mutation.Enable(); ok {
		_spec.SetField(cabinetbin.FieldEnable, field.TypeBool, value)
	}
	if value, ok := cbuo.mutation.BatterySn(); ok {
		_spec.SetField(cabinetbin.FieldBatterySn, field.TypeString, value)
	}
	if cbuo.mutation.BatterySnCleared() {
		_spec.ClearField(cabinetbin.FieldBatterySn, field.TypeString)
	}
	if value, ok := cbuo.mutation.Voltage(); ok {
		_spec.SetField(cabinetbin.FieldVoltage, field.TypeFloat64, value)
	}
	if value, ok := cbuo.mutation.AddedVoltage(); ok {
		_spec.AddField(cabinetbin.FieldVoltage, field.TypeFloat64, value)
	}
	if value, ok := cbuo.mutation.Current(); ok {
		_spec.SetField(cabinetbin.FieldCurrent, field.TypeFloat64, value)
	}
	if value, ok := cbuo.mutation.AddedCurrent(); ok {
		_spec.AddField(cabinetbin.FieldCurrent, field.TypeFloat64, value)
	}
	_spec.AddModifiers(cbuo.modifiers...)
	_node = &CabinetBin{config: cbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cabinetbin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
