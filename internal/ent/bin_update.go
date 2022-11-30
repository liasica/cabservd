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
	"github.com/auroraride/cabservd/internal/ent/bin"
	"github.com/auroraride/cabservd/internal/ent/predicate"
)

// BinUpdate is the builder for updating Bin entities.
type BinUpdate struct {
	config
	hooks     []Hook
	mutation  *BinMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the BinUpdate builder.
func (bu *BinUpdate) Where(ps ...predicate.Bin) *BinUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BinUpdate) SetUpdatedAt(t time.Time) *BinUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetUUID sets the "uuid" field.
func (bu *BinUpdate) SetUUID(s string) *BinUpdate {
	bu.mutation.SetUUID(s)
	return bu
}

// SetBrand sets the "brand" field.
func (bu *BinUpdate) SetBrand(s string) *BinUpdate {
	bu.mutation.SetBrand(s)
	return bu
}

// SetSerial sets the "serial" field.
func (bu *BinUpdate) SetSerial(s string) *BinUpdate {
	bu.mutation.SetSerial(s)
	return bu
}

// SetLock sets the "lock" field.
func (bu *BinUpdate) SetLock(b bool) *BinUpdate {
	bu.mutation.SetLock(b)
	return bu
}

// SetNillableLock sets the "lock" field if the given value is not nil.
func (bu *BinUpdate) SetNillableLock(b *bool) *BinUpdate {
	if b != nil {
		bu.SetLock(*b)
	}
	return bu
}

// SetName sets the "name" field.
func (bu *BinUpdate) SetName(s string) *BinUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetIndex sets the "index" field.
func (bu *BinUpdate) SetIndex(i int) *BinUpdate {
	bu.mutation.ResetIndex()
	bu.mutation.SetIndex(i)
	return bu
}

// AddIndex adds i to the "index" field.
func (bu *BinUpdate) AddIndex(i int) *BinUpdate {
	bu.mutation.AddIndex(i)
	return bu
}

// SetOpen sets the "open" field.
func (bu *BinUpdate) SetOpen(b bool) *BinUpdate {
	bu.mutation.SetOpen(b)
	return bu
}

// SetNillableOpen sets the "open" field if the given value is not nil.
func (bu *BinUpdate) SetNillableOpen(b *bool) *BinUpdate {
	if b != nil {
		bu.SetOpen(*b)
	}
	return bu
}

// SetEnable sets the "enable" field.
func (bu *BinUpdate) SetEnable(b bool) *BinUpdate {
	bu.mutation.SetEnable(b)
	return bu
}

// SetNillableEnable sets the "enable" field if the given value is not nil.
func (bu *BinUpdate) SetNillableEnable(b *bool) *BinUpdate {
	if b != nil {
		bu.SetEnable(*b)
	}
	return bu
}

// SetHealth sets the "health" field.
func (bu *BinUpdate) SetHealth(b bool) *BinUpdate {
	bu.mutation.SetHealth(b)
	return bu
}

// SetNillableHealth sets the "health" field if the given value is not nil.
func (bu *BinUpdate) SetNillableHealth(b *bool) *BinUpdate {
	if b != nil {
		bu.SetHealth(*b)
	}
	return bu
}

// SetBatterySn sets the "battery_sn" field.
func (bu *BinUpdate) SetBatterySn(s string) *BinUpdate {
	bu.mutation.SetBatterySn(s)
	return bu
}

// SetNillableBatterySn sets the "battery_sn" field if the given value is not nil.
func (bu *BinUpdate) SetNillableBatterySn(s *string) *BinUpdate {
	if s != nil {
		bu.SetBatterySn(*s)
	}
	return bu
}

// SetVoltage sets the "voltage" field.
func (bu *BinUpdate) SetVoltage(f float64) *BinUpdate {
	bu.mutation.ResetVoltage()
	bu.mutation.SetVoltage(f)
	return bu
}

// SetNillableVoltage sets the "voltage" field if the given value is not nil.
func (bu *BinUpdate) SetNillableVoltage(f *float64) *BinUpdate {
	if f != nil {
		bu.SetVoltage(*f)
	}
	return bu
}

// AddVoltage adds f to the "voltage" field.
func (bu *BinUpdate) AddVoltage(f float64) *BinUpdate {
	bu.mutation.AddVoltage(f)
	return bu
}

// SetCurrent sets the "current" field.
func (bu *BinUpdate) SetCurrent(f float64) *BinUpdate {
	bu.mutation.ResetCurrent()
	bu.mutation.SetCurrent(f)
	return bu
}

// SetNillableCurrent sets the "current" field if the given value is not nil.
func (bu *BinUpdate) SetNillableCurrent(f *float64) *BinUpdate {
	if f != nil {
		bu.SetCurrent(*f)
	}
	return bu
}

// AddCurrent adds f to the "current" field.
func (bu *BinUpdate) AddCurrent(f float64) *BinUpdate {
	bu.mutation.AddCurrent(f)
	return bu
}

// SetSoc sets the "soc" field.
func (bu *BinUpdate) SetSoc(f float64) *BinUpdate {
	bu.mutation.ResetSoc()
	bu.mutation.SetSoc(f)
	return bu
}

// SetNillableSoc sets the "soc" field if the given value is not nil.
func (bu *BinUpdate) SetNillableSoc(f *float64) *BinUpdate {
	if f != nil {
		bu.SetSoc(*f)
	}
	return bu
}

// AddSoc adds f to the "soc" field.
func (bu *BinUpdate) AddSoc(f float64) *BinUpdate {
	bu.mutation.AddSoc(f)
	return bu
}

// SetSoh sets the "soh" field.
func (bu *BinUpdate) SetSoh(f float64) *BinUpdate {
	bu.mutation.ResetSoh()
	bu.mutation.SetSoh(f)
	return bu
}

// SetNillableSoh sets the "soh" field if the given value is not nil.
func (bu *BinUpdate) SetNillableSoh(f *float64) *BinUpdate {
	if f != nil {
		bu.SetSoh(*f)
	}
	return bu
}

// AddSoh adds f to the "soh" field.
func (bu *BinUpdate) AddSoh(f float64) *BinUpdate {
	bu.mutation.AddSoh(f)
	return bu
}

// Mutation returns the BinMutation object of the builder.
func (bu *BinUpdate) Mutation() *BinMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BinUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	bu.defaults()
	if len(bu.hooks) == 0 {
		if err = bu.check(); err != nil {
			return 0, err
		}
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bu.check(); err != nil {
				return 0, err
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BinUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BinUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BinUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BinUpdate) defaults() {
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		v := bin.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BinUpdate) check() error {
	if v, ok := bu.mutation.UUID(); ok {
		if err := bin.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf(`ent: validator failed for field "Bin.uuid": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (bu *BinUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *BinUpdate {
	bu.modifiers = append(bu.modifiers, modifiers...)
	return bu
}

func (bu *BinUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bin.Table,
			Columns: bin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: bin.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(bin.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.UUID(); ok {
		_spec.SetField(bin.FieldUUID, field.TypeString, value)
	}
	if value, ok := bu.mutation.Brand(); ok {
		_spec.SetField(bin.FieldBrand, field.TypeString, value)
	}
	if value, ok := bu.mutation.Serial(); ok {
		_spec.SetField(bin.FieldSerial, field.TypeString, value)
	}
	if value, ok := bu.mutation.Lock(); ok {
		_spec.SetField(bin.FieldLock, field.TypeBool, value)
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.SetField(bin.FieldName, field.TypeString, value)
	}
	if value, ok := bu.mutation.Index(); ok {
		_spec.SetField(bin.FieldIndex, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedIndex(); ok {
		_spec.AddField(bin.FieldIndex, field.TypeInt, value)
	}
	if value, ok := bu.mutation.Open(); ok {
		_spec.SetField(bin.FieldOpen, field.TypeBool, value)
	}
	if value, ok := bu.mutation.Enable(); ok {
		_spec.SetField(bin.FieldEnable, field.TypeBool, value)
	}
	if value, ok := bu.mutation.Health(); ok {
		_spec.SetField(bin.FieldHealth, field.TypeBool, value)
	}
	if value, ok := bu.mutation.BatterySn(); ok {
		_spec.SetField(bin.FieldBatterySn, field.TypeString, value)
	}
	if value, ok := bu.mutation.Voltage(); ok {
		_spec.SetField(bin.FieldVoltage, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.AddedVoltage(); ok {
		_spec.AddField(bin.FieldVoltage, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.Current(); ok {
		_spec.SetField(bin.FieldCurrent, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.AddedCurrent(); ok {
		_spec.AddField(bin.FieldCurrent, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.Soc(); ok {
		_spec.SetField(bin.FieldSoc, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.AddedSoc(); ok {
		_spec.AddField(bin.FieldSoc, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.Soh(); ok {
		_spec.SetField(bin.FieldSoh, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.AddedSoh(); ok {
		_spec.AddField(bin.FieldSoh, field.TypeFloat64, value)
	}
	_spec.AddModifiers(bu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// BinUpdateOne is the builder for updating a single Bin entity.
type BinUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *BinMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BinUpdateOne) SetUpdatedAt(t time.Time) *BinUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetUUID sets the "uuid" field.
func (buo *BinUpdateOne) SetUUID(s string) *BinUpdateOne {
	buo.mutation.SetUUID(s)
	return buo
}

// SetBrand sets the "brand" field.
func (buo *BinUpdateOne) SetBrand(s string) *BinUpdateOne {
	buo.mutation.SetBrand(s)
	return buo
}

// SetSerial sets the "serial" field.
func (buo *BinUpdateOne) SetSerial(s string) *BinUpdateOne {
	buo.mutation.SetSerial(s)
	return buo
}

// SetLock sets the "lock" field.
func (buo *BinUpdateOne) SetLock(b bool) *BinUpdateOne {
	buo.mutation.SetLock(b)
	return buo
}

// SetNillableLock sets the "lock" field if the given value is not nil.
func (buo *BinUpdateOne) SetNillableLock(b *bool) *BinUpdateOne {
	if b != nil {
		buo.SetLock(*b)
	}
	return buo
}

// SetName sets the "name" field.
func (buo *BinUpdateOne) SetName(s string) *BinUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetIndex sets the "index" field.
func (buo *BinUpdateOne) SetIndex(i int) *BinUpdateOne {
	buo.mutation.ResetIndex()
	buo.mutation.SetIndex(i)
	return buo
}

// AddIndex adds i to the "index" field.
func (buo *BinUpdateOne) AddIndex(i int) *BinUpdateOne {
	buo.mutation.AddIndex(i)
	return buo
}

// SetOpen sets the "open" field.
func (buo *BinUpdateOne) SetOpen(b bool) *BinUpdateOne {
	buo.mutation.SetOpen(b)
	return buo
}

// SetNillableOpen sets the "open" field if the given value is not nil.
func (buo *BinUpdateOne) SetNillableOpen(b *bool) *BinUpdateOne {
	if b != nil {
		buo.SetOpen(*b)
	}
	return buo
}

// SetEnable sets the "enable" field.
func (buo *BinUpdateOne) SetEnable(b bool) *BinUpdateOne {
	buo.mutation.SetEnable(b)
	return buo
}

// SetNillableEnable sets the "enable" field if the given value is not nil.
func (buo *BinUpdateOne) SetNillableEnable(b *bool) *BinUpdateOne {
	if b != nil {
		buo.SetEnable(*b)
	}
	return buo
}

// SetHealth sets the "health" field.
func (buo *BinUpdateOne) SetHealth(b bool) *BinUpdateOne {
	buo.mutation.SetHealth(b)
	return buo
}

// SetNillableHealth sets the "health" field if the given value is not nil.
func (buo *BinUpdateOne) SetNillableHealth(b *bool) *BinUpdateOne {
	if b != nil {
		buo.SetHealth(*b)
	}
	return buo
}

// SetBatterySn sets the "battery_sn" field.
func (buo *BinUpdateOne) SetBatterySn(s string) *BinUpdateOne {
	buo.mutation.SetBatterySn(s)
	return buo
}

// SetNillableBatterySn sets the "battery_sn" field if the given value is not nil.
func (buo *BinUpdateOne) SetNillableBatterySn(s *string) *BinUpdateOne {
	if s != nil {
		buo.SetBatterySn(*s)
	}
	return buo
}

// SetVoltage sets the "voltage" field.
func (buo *BinUpdateOne) SetVoltage(f float64) *BinUpdateOne {
	buo.mutation.ResetVoltage()
	buo.mutation.SetVoltage(f)
	return buo
}

// SetNillableVoltage sets the "voltage" field if the given value is not nil.
func (buo *BinUpdateOne) SetNillableVoltage(f *float64) *BinUpdateOne {
	if f != nil {
		buo.SetVoltage(*f)
	}
	return buo
}

// AddVoltage adds f to the "voltage" field.
func (buo *BinUpdateOne) AddVoltage(f float64) *BinUpdateOne {
	buo.mutation.AddVoltage(f)
	return buo
}

// SetCurrent sets the "current" field.
func (buo *BinUpdateOne) SetCurrent(f float64) *BinUpdateOne {
	buo.mutation.ResetCurrent()
	buo.mutation.SetCurrent(f)
	return buo
}

// SetNillableCurrent sets the "current" field if the given value is not nil.
func (buo *BinUpdateOne) SetNillableCurrent(f *float64) *BinUpdateOne {
	if f != nil {
		buo.SetCurrent(*f)
	}
	return buo
}

// AddCurrent adds f to the "current" field.
func (buo *BinUpdateOne) AddCurrent(f float64) *BinUpdateOne {
	buo.mutation.AddCurrent(f)
	return buo
}

// SetSoc sets the "soc" field.
func (buo *BinUpdateOne) SetSoc(f float64) *BinUpdateOne {
	buo.mutation.ResetSoc()
	buo.mutation.SetSoc(f)
	return buo
}

// SetNillableSoc sets the "soc" field if the given value is not nil.
func (buo *BinUpdateOne) SetNillableSoc(f *float64) *BinUpdateOne {
	if f != nil {
		buo.SetSoc(*f)
	}
	return buo
}

// AddSoc adds f to the "soc" field.
func (buo *BinUpdateOne) AddSoc(f float64) *BinUpdateOne {
	buo.mutation.AddSoc(f)
	return buo
}

// SetSoh sets the "soh" field.
func (buo *BinUpdateOne) SetSoh(f float64) *BinUpdateOne {
	buo.mutation.ResetSoh()
	buo.mutation.SetSoh(f)
	return buo
}

// SetNillableSoh sets the "soh" field if the given value is not nil.
func (buo *BinUpdateOne) SetNillableSoh(f *float64) *BinUpdateOne {
	if f != nil {
		buo.SetSoh(*f)
	}
	return buo
}

// AddSoh adds f to the "soh" field.
func (buo *BinUpdateOne) AddSoh(f float64) *BinUpdateOne {
	buo.mutation.AddSoh(f)
	return buo
}

// Mutation returns the BinMutation object of the builder.
func (buo *BinUpdateOne) Mutation() *BinMutation {
	return buo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BinUpdateOne) Select(field string, fields ...string) *BinUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Bin entity.
func (buo *BinUpdateOne) Save(ctx context.Context) (*Bin, error) {
	var (
		err  error
		node *Bin
	)
	buo.defaults()
	if len(buo.hooks) == 0 {
		if err = buo.check(); err != nil {
			return nil, err
		}
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buo.check(); err != nil {
				return nil, err
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, buo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Bin)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BinMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BinUpdateOne) SaveX(ctx context.Context) *Bin {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BinUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BinUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BinUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		v := bin.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BinUpdateOne) check() error {
	if v, ok := buo.mutation.UUID(); ok {
		if err := bin.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf(`ent: validator failed for field "Bin.uuid": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (buo *BinUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *BinUpdateOne {
	buo.modifiers = append(buo.modifiers, modifiers...)
	return buo
}

func (buo *BinUpdateOne) sqlSave(ctx context.Context) (_node *Bin, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bin.Table,
			Columns: bin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: bin.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Bin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bin.FieldID)
		for _, f := range fields {
			if !bin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(bin.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.UUID(); ok {
		_spec.SetField(bin.FieldUUID, field.TypeString, value)
	}
	if value, ok := buo.mutation.Brand(); ok {
		_spec.SetField(bin.FieldBrand, field.TypeString, value)
	}
	if value, ok := buo.mutation.Serial(); ok {
		_spec.SetField(bin.FieldSerial, field.TypeString, value)
	}
	if value, ok := buo.mutation.Lock(); ok {
		_spec.SetField(bin.FieldLock, field.TypeBool, value)
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.SetField(bin.FieldName, field.TypeString, value)
	}
	if value, ok := buo.mutation.Index(); ok {
		_spec.SetField(bin.FieldIndex, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedIndex(); ok {
		_spec.AddField(bin.FieldIndex, field.TypeInt, value)
	}
	if value, ok := buo.mutation.Open(); ok {
		_spec.SetField(bin.FieldOpen, field.TypeBool, value)
	}
	if value, ok := buo.mutation.Enable(); ok {
		_spec.SetField(bin.FieldEnable, field.TypeBool, value)
	}
	if value, ok := buo.mutation.Health(); ok {
		_spec.SetField(bin.FieldHealth, field.TypeBool, value)
	}
	if value, ok := buo.mutation.BatterySn(); ok {
		_spec.SetField(bin.FieldBatterySn, field.TypeString, value)
	}
	if value, ok := buo.mutation.Voltage(); ok {
		_spec.SetField(bin.FieldVoltage, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.AddedVoltage(); ok {
		_spec.AddField(bin.FieldVoltage, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.Current(); ok {
		_spec.SetField(bin.FieldCurrent, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.AddedCurrent(); ok {
		_spec.AddField(bin.FieldCurrent, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.Soc(); ok {
		_spec.SetField(bin.FieldSoc, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.AddedSoc(); ok {
		_spec.AddField(bin.FieldSoc, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.Soh(); ok {
		_spec.SetField(bin.FieldSoh, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.AddedSoh(); ok {
		_spec.AddField(bin.FieldSoh, field.TypeFloat64, value)
	}
	_spec.AddModifiers(buo.modifiers...)
	_node = &Bin{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
