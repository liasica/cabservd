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
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/ent/predicate"
)

// CabinetUpdate is the builder for updating Cabinet entities.
type CabinetUpdate struct {
	config
	hooks     []Hook
	mutation  *CabinetMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CabinetUpdate builder.
func (cu *CabinetUpdate) Where(ps ...predicate.Cabinet) *CabinetUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CabinetUpdate) SetUpdatedAt(t time.Time) *CabinetUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetOnline sets the "online" field.
func (cu *CabinetUpdate) SetOnline(b bool) *CabinetUpdate {
	cu.mutation.SetOnline(b)
	return cu
}

// SetNillableOnline sets the "online" field if the given value is not nil.
func (cu *CabinetUpdate) SetNillableOnline(b *bool) *CabinetUpdate {
	if b != nil {
		cu.SetOnline(*b)
	}
	return cu
}

// SetBrand sets the "brand" field.
func (cu *CabinetUpdate) SetBrand(s string) *CabinetUpdate {
	cu.mutation.SetBrand(s)
	return cu
}

// SetSerial sets the "serial" field.
func (cu *CabinetUpdate) SetSerial(s string) *CabinetUpdate {
	cu.mutation.SetSerial(s)
	return cu
}

// SetStatus sets the "status" field.
func (cu *CabinetUpdate) SetStatus(c cabinet.Status) *CabinetUpdate {
	cu.mutation.SetStatus(c)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *CabinetUpdate) SetNillableStatus(c *cabinet.Status) *CabinetUpdate {
	if c != nil {
		cu.SetStatus(*c)
	}
	return cu
}

// SetEnable sets the "enable" field.
func (cu *CabinetUpdate) SetEnable(b bool) *CabinetUpdate {
	cu.mutation.SetEnable(b)
	return cu
}

// SetNillableEnable sets the "enable" field if the given value is not nil.
func (cu *CabinetUpdate) SetNillableEnable(b *bool) *CabinetUpdate {
	if b != nil {
		cu.SetEnable(*b)
	}
	return cu
}

// SetLng sets the "lng" field.
func (cu *CabinetUpdate) SetLng(f float64) *CabinetUpdate {
	cu.mutation.ResetLng()
	cu.mutation.SetLng(f)
	return cu
}

// SetNillableLng sets the "lng" field if the given value is not nil.
func (cu *CabinetUpdate) SetNillableLng(f *float64) *CabinetUpdate {
	if f != nil {
		cu.SetLng(*f)
	}
	return cu
}

// AddLng adds f to the "lng" field.
func (cu *CabinetUpdate) AddLng(f float64) *CabinetUpdate {
	cu.mutation.AddLng(f)
	return cu
}

// ClearLng clears the value of the "lng" field.
func (cu *CabinetUpdate) ClearLng() *CabinetUpdate {
	cu.mutation.ClearLng()
	return cu
}

// SetLat sets the "lat" field.
func (cu *CabinetUpdate) SetLat(f float64) *CabinetUpdate {
	cu.mutation.ResetLat()
	cu.mutation.SetLat(f)
	return cu
}

// SetNillableLat sets the "lat" field if the given value is not nil.
func (cu *CabinetUpdate) SetNillableLat(f *float64) *CabinetUpdate {
	if f != nil {
		cu.SetLat(*f)
	}
	return cu
}

// AddLat adds f to the "lat" field.
func (cu *CabinetUpdate) AddLat(f float64) *CabinetUpdate {
	cu.mutation.AddLat(f)
	return cu
}

// ClearLat clears the value of the "lat" field.
func (cu *CabinetUpdate) ClearLat() *CabinetUpdate {
	cu.mutation.ClearLat()
	return cu
}

// SetGsm sets the "gsm" field.
func (cu *CabinetUpdate) SetGsm(f float64) *CabinetUpdate {
	cu.mutation.ResetGsm()
	cu.mutation.SetGsm(f)
	return cu
}

// SetNillableGsm sets the "gsm" field if the given value is not nil.
func (cu *CabinetUpdate) SetNillableGsm(f *float64) *CabinetUpdate {
	if f != nil {
		cu.SetGsm(*f)
	}
	return cu
}

// AddGsm adds f to the "gsm" field.
func (cu *CabinetUpdate) AddGsm(f float64) *CabinetUpdate {
	cu.mutation.AddGsm(f)
	return cu
}

// ClearGsm clears the value of the "gsm" field.
func (cu *CabinetUpdate) ClearGsm() *CabinetUpdate {
	cu.mutation.ClearGsm()
	return cu
}

// SetVoltage sets the "voltage" field.
func (cu *CabinetUpdate) SetVoltage(f float64) *CabinetUpdate {
	cu.mutation.ResetVoltage()
	cu.mutation.SetVoltage(f)
	return cu
}

// SetNillableVoltage sets the "voltage" field if the given value is not nil.
func (cu *CabinetUpdate) SetNillableVoltage(f *float64) *CabinetUpdate {
	if f != nil {
		cu.SetVoltage(*f)
	}
	return cu
}

// AddVoltage adds f to the "voltage" field.
func (cu *CabinetUpdate) AddVoltage(f float64) *CabinetUpdate {
	cu.mutation.AddVoltage(f)
	return cu
}

// ClearVoltage clears the value of the "voltage" field.
func (cu *CabinetUpdate) ClearVoltage() *CabinetUpdate {
	cu.mutation.ClearVoltage()
	return cu
}

// SetCurrent sets the "current" field.
func (cu *CabinetUpdate) SetCurrent(f float64) *CabinetUpdate {
	cu.mutation.ResetCurrent()
	cu.mutation.SetCurrent(f)
	return cu
}

// SetNillableCurrent sets the "current" field if the given value is not nil.
func (cu *CabinetUpdate) SetNillableCurrent(f *float64) *CabinetUpdate {
	if f != nil {
		cu.SetCurrent(*f)
	}
	return cu
}

// AddCurrent adds f to the "current" field.
func (cu *CabinetUpdate) AddCurrent(f float64) *CabinetUpdate {
	cu.mutation.AddCurrent(f)
	return cu
}

// ClearCurrent clears the value of the "current" field.
func (cu *CabinetUpdate) ClearCurrent() *CabinetUpdate {
	cu.mutation.ClearCurrent()
	return cu
}

// SetTemperature sets the "temperature" field.
func (cu *CabinetUpdate) SetTemperature(f float64) *CabinetUpdate {
	cu.mutation.ResetTemperature()
	cu.mutation.SetTemperature(f)
	return cu
}

// SetNillableTemperature sets the "temperature" field if the given value is not nil.
func (cu *CabinetUpdate) SetNillableTemperature(f *float64) *CabinetUpdate {
	if f != nil {
		cu.SetTemperature(*f)
	}
	return cu
}

// AddTemperature adds f to the "temperature" field.
func (cu *CabinetUpdate) AddTemperature(f float64) *CabinetUpdate {
	cu.mutation.AddTemperature(f)
	return cu
}

// ClearTemperature clears the value of the "temperature" field.
func (cu *CabinetUpdate) ClearTemperature() *CabinetUpdate {
	cu.mutation.ClearTemperature()
	return cu
}

// SetElectricity sets the "electricity" field.
func (cu *CabinetUpdate) SetElectricity(f float64) *CabinetUpdate {
	cu.mutation.ResetElectricity()
	cu.mutation.SetElectricity(f)
	return cu
}

// SetNillableElectricity sets the "electricity" field if the given value is not nil.
func (cu *CabinetUpdate) SetNillableElectricity(f *float64) *CabinetUpdate {
	if f != nil {
		cu.SetElectricity(*f)
	}
	return cu
}

// AddElectricity adds f to the "electricity" field.
func (cu *CabinetUpdate) AddElectricity(f float64) *CabinetUpdate {
	cu.mutation.AddElectricity(f)
	return cu
}

// ClearElectricity clears the value of the "electricity" field.
func (cu *CabinetUpdate) ClearElectricity() *CabinetUpdate {
	cu.mutation.ClearElectricity()
	return cu
}

// Mutation returns the CabinetMutation object of the builder.
func (cu *CabinetUpdate) Mutation() *CabinetMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CabinetUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CabinetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CabinetUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CabinetUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CabinetUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CabinetUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := cabinet.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CabinetUpdate) check() error {
	if v, ok := cu.mutation.Status(); ok {
		if err := cabinet.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Cabinet.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CabinetUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CabinetUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CabinetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cabinet.Table,
			Columns: cabinet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: cabinet.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(cabinet.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Online(); ok {
		_spec.SetField(cabinet.FieldOnline, field.TypeBool, value)
	}
	if value, ok := cu.mutation.Brand(); ok {
		_spec.SetField(cabinet.FieldBrand, field.TypeString, value)
	}
	if value, ok := cu.mutation.Serial(); ok {
		_spec.SetField(cabinet.FieldSerial, field.TypeString, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(cabinet.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.Enable(); ok {
		_spec.SetField(cabinet.FieldEnable, field.TypeBool, value)
	}
	if value, ok := cu.mutation.Lng(); ok {
		_spec.SetField(cabinet.FieldLng, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedLng(); ok {
		_spec.AddField(cabinet.FieldLng, field.TypeFloat64, value)
	}
	if cu.mutation.LngCleared() {
		_spec.ClearField(cabinet.FieldLng, field.TypeFloat64)
	}
	if value, ok := cu.mutation.Lat(); ok {
		_spec.SetField(cabinet.FieldLat, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedLat(); ok {
		_spec.AddField(cabinet.FieldLat, field.TypeFloat64, value)
	}
	if cu.mutation.LatCleared() {
		_spec.ClearField(cabinet.FieldLat, field.TypeFloat64)
	}
	if value, ok := cu.mutation.Gsm(); ok {
		_spec.SetField(cabinet.FieldGsm, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedGsm(); ok {
		_spec.AddField(cabinet.FieldGsm, field.TypeFloat64, value)
	}
	if cu.mutation.GsmCleared() {
		_spec.ClearField(cabinet.FieldGsm, field.TypeFloat64)
	}
	if value, ok := cu.mutation.Voltage(); ok {
		_spec.SetField(cabinet.FieldVoltage, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedVoltage(); ok {
		_spec.AddField(cabinet.FieldVoltage, field.TypeFloat64, value)
	}
	if cu.mutation.VoltageCleared() {
		_spec.ClearField(cabinet.FieldVoltage, field.TypeFloat64)
	}
	if value, ok := cu.mutation.Current(); ok {
		_spec.SetField(cabinet.FieldCurrent, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedCurrent(); ok {
		_spec.AddField(cabinet.FieldCurrent, field.TypeFloat64, value)
	}
	if cu.mutation.CurrentCleared() {
		_spec.ClearField(cabinet.FieldCurrent, field.TypeFloat64)
	}
	if value, ok := cu.mutation.Temperature(); ok {
		_spec.SetField(cabinet.FieldTemperature, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedTemperature(); ok {
		_spec.AddField(cabinet.FieldTemperature, field.TypeFloat64, value)
	}
	if cu.mutation.TemperatureCleared() {
		_spec.ClearField(cabinet.FieldTemperature, field.TypeFloat64)
	}
	if value, ok := cu.mutation.Electricity(); ok {
		_spec.SetField(cabinet.FieldElectricity, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedElectricity(); ok {
		_spec.AddField(cabinet.FieldElectricity, field.TypeFloat64, value)
	}
	if cu.mutation.ElectricityCleared() {
		_spec.ClearField(cabinet.FieldElectricity, field.TypeFloat64)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cabinet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CabinetUpdateOne is the builder for updating a single Cabinet entity.
type CabinetUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CabinetMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CabinetUpdateOne) SetUpdatedAt(t time.Time) *CabinetUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetOnline sets the "online" field.
func (cuo *CabinetUpdateOne) SetOnline(b bool) *CabinetUpdateOne {
	cuo.mutation.SetOnline(b)
	return cuo
}

// SetNillableOnline sets the "online" field if the given value is not nil.
func (cuo *CabinetUpdateOne) SetNillableOnline(b *bool) *CabinetUpdateOne {
	if b != nil {
		cuo.SetOnline(*b)
	}
	return cuo
}

// SetBrand sets the "brand" field.
func (cuo *CabinetUpdateOne) SetBrand(s string) *CabinetUpdateOne {
	cuo.mutation.SetBrand(s)
	return cuo
}

// SetSerial sets the "serial" field.
func (cuo *CabinetUpdateOne) SetSerial(s string) *CabinetUpdateOne {
	cuo.mutation.SetSerial(s)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CabinetUpdateOne) SetStatus(c cabinet.Status) *CabinetUpdateOne {
	cuo.mutation.SetStatus(c)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *CabinetUpdateOne) SetNillableStatus(c *cabinet.Status) *CabinetUpdateOne {
	if c != nil {
		cuo.SetStatus(*c)
	}
	return cuo
}

// SetEnable sets the "enable" field.
func (cuo *CabinetUpdateOne) SetEnable(b bool) *CabinetUpdateOne {
	cuo.mutation.SetEnable(b)
	return cuo
}

// SetNillableEnable sets the "enable" field if the given value is not nil.
func (cuo *CabinetUpdateOne) SetNillableEnable(b *bool) *CabinetUpdateOne {
	if b != nil {
		cuo.SetEnable(*b)
	}
	return cuo
}

// SetLng sets the "lng" field.
func (cuo *CabinetUpdateOne) SetLng(f float64) *CabinetUpdateOne {
	cuo.mutation.ResetLng()
	cuo.mutation.SetLng(f)
	return cuo
}

// SetNillableLng sets the "lng" field if the given value is not nil.
func (cuo *CabinetUpdateOne) SetNillableLng(f *float64) *CabinetUpdateOne {
	if f != nil {
		cuo.SetLng(*f)
	}
	return cuo
}

// AddLng adds f to the "lng" field.
func (cuo *CabinetUpdateOne) AddLng(f float64) *CabinetUpdateOne {
	cuo.mutation.AddLng(f)
	return cuo
}

// ClearLng clears the value of the "lng" field.
func (cuo *CabinetUpdateOne) ClearLng() *CabinetUpdateOne {
	cuo.mutation.ClearLng()
	return cuo
}

// SetLat sets the "lat" field.
func (cuo *CabinetUpdateOne) SetLat(f float64) *CabinetUpdateOne {
	cuo.mutation.ResetLat()
	cuo.mutation.SetLat(f)
	return cuo
}

// SetNillableLat sets the "lat" field if the given value is not nil.
func (cuo *CabinetUpdateOne) SetNillableLat(f *float64) *CabinetUpdateOne {
	if f != nil {
		cuo.SetLat(*f)
	}
	return cuo
}

// AddLat adds f to the "lat" field.
func (cuo *CabinetUpdateOne) AddLat(f float64) *CabinetUpdateOne {
	cuo.mutation.AddLat(f)
	return cuo
}

// ClearLat clears the value of the "lat" field.
func (cuo *CabinetUpdateOne) ClearLat() *CabinetUpdateOne {
	cuo.mutation.ClearLat()
	return cuo
}

// SetGsm sets the "gsm" field.
func (cuo *CabinetUpdateOne) SetGsm(f float64) *CabinetUpdateOne {
	cuo.mutation.ResetGsm()
	cuo.mutation.SetGsm(f)
	return cuo
}

// SetNillableGsm sets the "gsm" field if the given value is not nil.
func (cuo *CabinetUpdateOne) SetNillableGsm(f *float64) *CabinetUpdateOne {
	if f != nil {
		cuo.SetGsm(*f)
	}
	return cuo
}

// AddGsm adds f to the "gsm" field.
func (cuo *CabinetUpdateOne) AddGsm(f float64) *CabinetUpdateOne {
	cuo.mutation.AddGsm(f)
	return cuo
}

// ClearGsm clears the value of the "gsm" field.
func (cuo *CabinetUpdateOne) ClearGsm() *CabinetUpdateOne {
	cuo.mutation.ClearGsm()
	return cuo
}

// SetVoltage sets the "voltage" field.
func (cuo *CabinetUpdateOne) SetVoltage(f float64) *CabinetUpdateOne {
	cuo.mutation.ResetVoltage()
	cuo.mutation.SetVoltage(f)
	return cuo
}

// SetNillableVoltage sets the "voltage" field if the given value is not nil.
func (cuo *CabinetUpdateOne) SetNillableVoltage(f *float64) *CabinetUpdateOne {
	if f != nil {
		cuo.SetVoltage(*f)
	}
	return cuo
}

// AddVoltage adds f to the "voltage" field.
func (cuo *CabinetUpdateOne) AddVoltage(f float64) *CabinetUpdateOne {
	cuo.mutation.AddVoltage(f)
	return cuo
}

// ClearVoltage clears the value of the "voltage" field.
func (cuo *CabinetUpdateOne) ClearVoltage() *CabinetUpdateOne {
	cuo.mutation.ClearVoltage()
	return cuo
}

// SetCurrent sets the "current" field.
func (cuo *CabinetUpdateOne) SetCurrent(f float64) *CabinetUpdateOne {
	cuo.mutation.ResetCurrent()
	cuo.mutation.SetCurrent(f)
	return cuo
}

// SetNillableCurrent sets the "current" field if the given value is not nil.
func (cuo *CabinetUpdateOne) SetNillableCurrent(f *float64) *CabinetUpdateOne {
	if f != nil {
		cuo.SetCurrent(*f)
	}
	return cuo
}

// AddCurrent adds f to the "current" field.
func (cuo *CabinetUpdateOne) AddCurrent(f float64) *CabinetUpdateOne {
	cuo.mutation.AddCurrent(f)
	return cuo
}

// ClearCurrent clears the value of the "current" field.
func (cuo *CabinetUpdateOne) ClearCurrent() *CabinetUpdateOne {
	cuo.mutation.ClearCurrent()
	return cuo
}

// SetTemperature sets the "temperature" field.
func (cuo *CabinetUpdateOne) SetTemperature(f float64) *CabinetUpdateOne {
	cuo.mutation.ResetTemperature()
	cuo.mutation.SetTemperature(f)
	return cuo
}

// SetNillableTemperature sets the "temperature" field if the given value is not nil.
func (cuo *CabinetUpdateOne) SetNillableTemperature(f *float64) *CabinetUpdateOne {
	if f != nil {
		cuo.SetTemperature(*f)
	}
	return cuo
}

// AddTemperature adds f to the "temperature" field.
func (cuo *CabinetUpdateOne) AddTemperature(f float64) *CabinetUpdateOne {
	cuo.mutation.AddTemperature(f)
	return cuo
}

// ClearTemperature clears the value of the "temperature" field.
func (cuo *CabinetUpdateOne) ClearTemperature() *CabinetUpdateOne {
	cuo.mutation.ClearTemperature()
	return cuo
}

// SetElectricity sets the "electricity" field.
func (cuo *CabinetUpdateOne) SetElectricity(f float64) *CabinetUpdateOne {
	cuo.mutation.ResetElectricity()
	cuo.mutation.SetElectricity(f)
	return cuo
}

// SetNillableElectricity sets the "electricity" field if the given value is not nil.
func (cuo *CabinetUpdateOne) SetNillableElectricity(f *float64) *CabinetUpdateOne {
	if f != nil {
		cuo.SetElectricity(*f)
	}
	return cuo
}

// AddElectricity adds f to the "electricity" field.
func (cuo *CabinetUpdateOne) AddElectricity(f float64) *CabinetUpdateOne {
	cuo.mutation.AddElectricity(f)
	return cuo
}

// ClearElectricity clears the value of the "electricity" field.
func (cuo *CabinetUpdateOne) ClearElectricity() *CabinetUpdateOne {
	cuo.mutation.ClearElectricity()
	return cuo
}

// Mutation returns the CabinetMutation object of the builder.
func (cuo *CabinetUpdateOne) Mutation() *CabinetMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CabinetUpdateOne) Select(field string, fields ...string) *CabinetUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cabinet entity.
func (cuo *CabinetUpdateOne) Save(ctx context.Context) (*Cabinet, error) {
	var (
		err  error
		node *Cabinet
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CabinetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Cabinet)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CabinetMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CabinetUpdateOne) SaveX(ctx context.Context) *Cabinet {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CabinetUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CabinetUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CabinetUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := cabinet.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CabinetUpdateOne) check() error {
	if v, ok := cuo.mutation.Status(); ok {
		if err := cabinet.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Cabinet.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CabinetUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CabinetUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CabinetUpdateOne) sqlSave(ctx context.Context) (_node *Cabinet, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cabinet.Table,
			Columns: cabinet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: cabinet.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cabinet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cabinet.FieldID)
		for _, f := range fields {
			if !cabinet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cabinet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(cabinet.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Online(); ok {
		_spec.SetField(cabinet.FieldOnline, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.Brand(); ok {
		_spec.SetField(cabinet.FieldBrand, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Serial(); ok {
		_spec.SetField(cabinet.FieldSerial, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(cabinet.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.Enable(); ok {
		_spec.SetField(cabinet.FieldEnable, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.Lng(); ok {
		_spec.SetField(cabinet.FieldLng, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedLng(); ok {
		_spec.AddField(cabinet.FieldLng, field.TypeFloat64, value)
	}
	if cuo.mutation.LngCleared() {
		_spec.ClearField(cabinet.FieldLng, field.TypeFloat64)
	}
	if value, ok := cuo.mutation.Lat(); ok {
		_spec.SetField(cabinet.FieldLat, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedLat(); ok {
		_spec.AddField(cabinet.FieldLat, field.TypeFloat64, value)
	}
	if cuo.mutation.LatCleared() {
		_spec.ClearField(cabinet.FieldLat, field.TypeFloat64)
	}
	if value, ok := cuo.mutation.Gsm(); ok {
		_spec.SetField(cabinet.FieldGsm, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedGsm(); ok {
		_spec.AddField(cabinet.FieldGsm, field.TypeFloat64, value)
	}
	if cuo.mutation.GsmCleared() {
		_spec.ClearField(cabinet.FieldGsm, field.TypeFloat64)
	}
	if value, ok := cuo.mutation.Voltage(); ok {
		_spec.SetField(cabinet.FieldVoltage, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedVoltage(); ok {
		_spec.AddField(cabinet.FieldVoltage, field.TypeFloat64, value)
	}
	if cuo.mutation.VoltageCleared() {
		_spec.ClearField(cabinet.FieldVoltage, field.TypeFloat64)
	}
	if value, ok := cuo.mutation.Current(); ok {
		_spec.SetField(cabinet.FieldCurrent, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedCurrent(); ok {
		_spec.AddField(cabinet.FieldCurrent, field.TypeFloat64, value)
	}
	if cuo.mutation.CurrentCleared() {
		_spec.ClearField(cabinet.FieldCurrent, field.TypeFloat64)
	}
	if value, ok := cuo.mutation.Temperature(); ok {
		_spec.SetField(cabinet.FieldTemperature, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedTemperature(); ok {
		_spec.AddField(cabinet.FieldTemperature, field.TypeFloat64, value)
	}
	if cuo.mutation.TemperatureCleared() {
		_spec.ClearField(cabinet.FieldTemperature, field.TypeFloat64)
	}
	if value, ok := cuo.mutation.Electricity(); ok {
		_spec.SetField(cabinet.FieldElectricity, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedElectricity(); ok {
		_spec.AddField(cabinet.FieldElectricity, field.TypeFloat64, value)
	}
	if cuo.mutation.ElectricityCleared() {
		_spec.ClearField(cabinet.FieldElectricity, field.TypeFloat64)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Cabinet{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cabinet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
