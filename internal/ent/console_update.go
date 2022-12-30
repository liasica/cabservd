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
	"github.com/auroraride/adapter/model"
	"github.com/auroraride/cabservd/internal/ent/bin"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/ent/console"
	"github.com/auroraride/cabservd/internal/ent/predicate"
)

// ConsoleUpdate is the builder for updating Console entities.
type ConsoleUpdate struct {
	config
	hooks     []Hook
	mutation  *ConsoleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ConsoleUpdate builder.
func (cu *ConsoleUpdate) Where(ps ...predicate.Console) *ConsoleUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCabinetID sets the "cabinet_id" field.
func (cu *ConsoleUpdate) SetCabinetID(u uint64) *ConsoleUpdate {
	cu.mutation.SetCabinetID(u)
	return cu
}

// SetBinID sets the "bin_id" field.
func (cu *ConsoleUpdate) SetBinID(u uint64) *ConsoleUpdate {
	cu.mutation.SetBinID(u)
	return cu
}

// SetOperate sets the "operate" field.
func (cu *ConsoleUpdate) SetOperate(mt model.Operator) *ConsoleUpdate {
	cu.mutation.SetOperate(mt)
	return cu
}

// SetNillableOperate sets the "operate" field if the given value is not nil.
func (cu *ConsoleUpdate) SetNillableOperate(mt *model.Operator) *ConsoleUpdate {
	if mt != nil {
		cu.SetOperate(*mt)
	}
	return cu
}

// ClearOperate clears the value of the "operate" field.
func (cu *ConsoleUpdate) ClearOperate() *ConsoleUpdate {
	cu.mutation.ClearOperate()
	return cu
}

// SetSerial sets the "serial" field.
func (cu *ConsoleUpdate) SetSerial(s string) *ConsoleUpdate {
	cu.mutation.SetSerial(s)
	return cu
}

// SetType sets the "type" field.
func (cu *ConsoleUpdate) SetType(c console.Type) *ConsoleUpdate {
	cu.mutation.SetType(c)
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *ConsoleUpdate) SetUserID(s string) *ConsoleUpdate {
	cu.mutation.SetUserID(s)
	return cu
}

// SetUserType sets the "user_type" field.
func (cu *ConsoleUpdate) SetUserType(mt model.UserType) *ConsoleUpdate {
	cu.mutation.SetUserType(mt)
	return cu
}

// SetStep sets the "step" field.
func (cu *ConsoleUpdate) SetStep(ms model.ExchangeStep) *ConsoleUpdate {
	cu.mutation.SetStep(ms)
	return cu
}

// SetNillableStep sets the "step" field if the given value is not nil.
func (cu *ConsoleUpdate) SetNillableStep(ms *model.ExchangeStep) *ConsoleUpdate {
	if ms != nil {
		cu.SetStep(*ms)
	}
	return cu
}

// ClearStep clears the value of the "step" field.
func (cu *ConsoleUpdate) ClearStep() *ConsoleUpdate {
	cu.mutation.ClearStep()
	return cu
}

// SetStatus sets the "status" field.
func (cu *ConsoleUpdate) SetStatus(c console.Status) *ConsoleUpdate {
	cu.mutation.SetStatus(c)
	return cu
}

// SetBeforeBin sets the "before_bin" field.
func (cu *ConsoleUpdate) SetBeforeBin(mi *model.BinInfo) *ConsoleUpdate {
	cu.mutation.SetBeforeBin(mi)
	return cu
}

// ClearBeforeBin clears the value of the "before_bin" field.
func (cu *ConsoleUpdate) ClearBeforeBin() *ConsoleUpdate {
	cu.mutation.ClearBeforeBin()
	return cu
}

// SetAfterBin sets the "after_bin" field.
func (cu *ConsoleUpdate) SetAfterBin(mi *model.BinInfo) *ConsoleUpdate {
	cu.mutation.SetAfterBin(mi)
	return cu
}

// ClearAfterBin clears the value of the "after_bin" field.
func (cu *ConsoleUpdate) ClearAfterBin() *ConsoleUpdate {
	cu.mutation.ClearAfterBin()
	return cu
}

// SetMessage sets the "message" field.
func (cu *ConsoleUpdate) SetMessage(s string) *ConsoleUpdate {
	cu.mutation.SetMessage(s)
	return cu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cu *ConsoleUpdate) SetNillableMessage(s *string) *ConsoleUpdate {
	if s != nil {
		cu.SetMessage(*s)
	}
	return cu
}

// ClearMessage clears the value of the "message" field.
func (cu *ConsoleUpdate) ClearMessage() *ConsoleUpdate {
	cu.mutation.ClearMessage()
	return cu
}

// SetStartAt sets the "startAt" field.
func (cu *ConsoleUpdate) SetStartAt(t time.Time) *ConsoleUpdate {
	cu.mutation.SetStartAt(t)
	return cu
}

// SetNillableStartAt sets the "startAt" field if the given value is not nil.
func (cu *ConsoleUpdate) SetNillableStartAt(t *time.Time) *ConsoleUpdate {
	if t != nil {
		cu.SetStartAt(*t)
	}
	return cu
}

// ClearStartAt clears the value of the "startAt" field.
func (cu *ConsoleUpdate) ClearStartAt() *ConsoleUpdate {
	cu.mutation.ClearStartAt()
	return cu
}

// SetStopAt sets the "stopAt" field.
func (cu *ConsoleUpdate) SetStopAt(t time.Time) *ConsoleUpdate {
	cu.mutation.SetStopAt(t)
	return cu
}

// SetNillableStopAt sets the "stopAt" field if the given value is not nil.
func (cu *ConsoleUpdate) SetNillableStopAt(t *time.Time) *ConsoleUpdate {
	if t != nil {
		cu.SetStopAt(*t)
	}
	return cu
}

// ClearStopAt clears the value of the "stopAt" field.
func (cu *ConsoleUpdate) ClearStopAt() *ConsoleUpdate {
	cu.mutation.ClearStopAt()
	return cu
}

// SetDuration sets the "duration" field.
func (cu *ConsoleUpdate) SetDuration(f float64) *ConsoleUpdate {
	cu.mutation.ResetDuration()
	cu.mutation.SetDuration(f)
	return cu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (cu *ConsoleUpdate) SetNillableDuration(f *float64) *ConsoleUpdate {
	if f != nil {
		cu.SetDuration(*f)
	}
	return cu
}

// AddDuration adds f to the "duration" field.
func (cu *ConsoleUpdate) AddDuration(f float64) *ConsoleUpdate {
	cu.mutation.AddDuration(f)
	return cu
}

// ClearDuration clears the value of the "duration" field.
func (cu *ConsoleUpdate) ClearDuration() *ConsoleUpdate {
	cu.mutation.ClearDuration()
	return cu
}

// SetCabinet sets the "cabinet" edge to the Cabinet entity.
func (cu *ConsoleUpdate) SetCabinet(c *Cabinet) *ConsoleUpdate {
	return cu.SetCabinetID(c.ID)
}

// SetBin sets the "bin" edge to the Bin entity.
func (cu *ConsoleUpdate) SetBin(b *Bin) *ConsoleUpdate {
	return cu.SetBinID(b.ID)
}

// Mutation returns the ConsoleMutation object of the builder.
func (cu *ConsoleUpdate) Mutation() *ConsoleMutation {
	return cu.mutation
}

// ClearCabinet clears the "cabinet" edge to the Cabinet entity.
func (cu *ConsoleUpdate) ClearCabinet() *ConsoleUpdate {
	cu.mutation.ClearCabinet()
	return cu
}

// ClearBin clears the "bin" edge to the Bin entity.
func (cu *ConsoleUpdate) ClearBin() *ConsoleUpdate {
	cu.mutation.ClearBin()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConsoleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ConsoleMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConsoleUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConsoleUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConsoleUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ConsoleUpdate) check() error {
	if v, ok := cu.mutation.GetType(); ok {
		if err := console.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Console.type": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Status(); ok {
		if err := console.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Console.status": %w`, err)}
		}
	}
	if _, ok := cu.mutation.CabinetID(); cu.mutation.CabinetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Console.cabinet"`)
	}
	if _, ok := cu.mutation.BinID(); cu.mutation.BinCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Console.bin"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *ConsoleUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ConsoleUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *ConsoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   console.Table,
			Columns: console.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: console.FieldID,
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
	if value, ok := cu.mutation.Operate(); ok {
		_spec.SetField(console.FieldOperate, field.TypeOther, value)
	}
	if cu.mutation.OperateCleared() {
		_spec.ClearField(console.FieldOperate, field.TypeOther)
	}
	if value, ok := cu.mutation.Serial(); ok {
		_spec.SetField(console.FieldSerial, field.TypeString, value)
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.SetField(console.FieldType, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.UserID(); ok {
		_spec.SetField(console.FieldUserID, field.TypeString, value)
	}
	if value, ok := cu.mutation.UserType(); ok {
		_spec.SetField(console.FieldUserType, field.TypeOther, value)
	}
	if value, ok := cu.mutation.Step(); ok {
		_spec.SetField(console.FieldStep, field.TypeOther, value)
	}
	if cu.mutation.StepCleared() {
		_spec.ClearField(console.FieldStep, field.TypeOther)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(console.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.BeforeBin(); ok {
		_spec.SetField(console.FieldBeforeBin, field.TypeJSON, value)
	}
	if cu.mutation.BeforeBinCleared() {
		_spec.ClearField(console.FieldBeforeBin, field.TypeJSON)
	}
	if value, ok := cu.mutation.AfterBin(); ok {
		_spec.SetField(console.FieldAfterBin, field.TypeJSON, value)
	}
	if cu.mutation.AfterBinCleared() {
		_spec.ClearField(console.FieldAfterBin, field.TypeJSON)
	}
	if value, ok := cu.mutation.Message(); ok {
		_spec.SetField(console.FieldMessage, field.TypeString, value)
	}
	if cu.mutation.MessageCleared() {
		_spec.ClearField(console.FieldMessage, field.TypeString)
	}
	if value, ok := cu.mutation.StartAt(); ok {
		_spec.SetField(console.FieldStartAt, field.TypeTime, value)
	}
	if cu.mutation.StartAtCleared() {
		_spec.ClearField(console.FieldStartAt, field.TypeTime)
	}
	if value, ok := cu.mutation.StopAt(); ok {
		_spec.SetField(console.FieldStopAt, field.TypeTime, value)
	}
	if cu.mutation.StopAtCleared() {
		_spec.ClearField(console.FieldStopAt, field.TypeTime)
	}
	if value, ok := cu.mutation.Duration(); ok {
		_spec.SetField(console.FieldDuration, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedDuration(); ok {
		_spec.AddField(console.FieldDuration, field.TypeFloat64, value)
	}
	if cu.mutation.DurationCleared() {
		_spec.ClearField(console.FieldDuration, field.TypeFloat64)
	}
	if cu.mutation.CabinetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   console.CabinetTable,
			Columns: []string{console.CabinetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: cabinet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CabinetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   console.CabinetTable,
			Columns: []string{console.CabinetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: cabinet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.BinCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   console.BinTable,
			Columns: []string{console.BinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: bin.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.BinIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   console.BinTable,
			Columns: []string{console.BinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: bin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{console.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ConsoleUpdateOne is the builder for updating a single Console entity.
type ConsoleUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ConsoleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCabinetID sets the "cabinet_id" field.
func (cuo *ConsoleUpdateOne) SetCabinetID(u uint64) *ConsoleUpdateOne {
	cuo.mutation.SetCabinetID(u)
	return cuo
}

// SetBinID sets the "bin_id" field.
func (cuo *ConsoleUpdateOne) SetBinID(u uint64) *ConsoleUpdateOne {
	cuo.mutation.SetBinID(u)
	return cuo
}

// SetOperate sets the "operate" field.
func (cuo *ConsoleUpdateOne) SetOperate(mt model.Operator) *ConsoleUpdateOne {
	cuo.mutation.SetOperate(mt)
	return cuo
}

// SetNillableOperate sets the "operate" field if the given value is not nil.
func (cuo *ConsoleUpdateOne) SetNillableOperate(mt *model.Operator) *ConsoleUpdateOne {
	if mt != nil {
		cuo.SetOperate(*mt)
	}
	return cuo
}

// ClearOperate clears the value of the "operate" field.
func (cuo *ConsoleUpdateOne) ClearOperate() *ConsoleUpdateOne {
	cuo.mutation.ClearOperate()
	return cuo
}

// SetSerial sets the "serial" field.
func (cuo *ConsoleUpdateOne) SetSerial(s string) *ConsoleUpdateOne {
	cuo.mutation.SetSerial(s)
	return cuo
}

// SetType sets the "type" field.
func (cuo *ConsoleUpdateOne) SetType(c console.Type) *ConsoleUpdateOne {
	cuo.mutation.SetType(c)
	return cuo
}

// SetUserID sets the "user_id" field.
func (cuo *ConsoleUpdateOne) SetUserID(s string) *ConsoleUpdateOne {
	cuo.mutation.SetUserID(s)
	return cuo
}

// SetUserType sets the "user_type" field.
func (cuo *ConsoleUpdateOne) SetUserType(mt model.UserType) *ConsoleUpdateOne {
	cuo.mutation.SetUserType(mt)
	return cuo
}

// SetStep sets the "step" field.
func (cuo *ConsoleUpdateOne) SetStep(ms model.ExchangeStep) *ConsoleUpdateOne {
	cuo.mutation.SetStep(ms)
	return cuo
}

// SetNillableStep sets the "step" field if the given value is not nil.
func (cuo *ConsoleUpdateOne) SetNillableStep(ms *model.ExchangeStep) *ConsoleUpdateOne {
	if ms != nil {
		cuo.SetStep(*ms)
	}
	return cuo
}

// ClearStep clears the value of the "step" field.
func (cuo *ConsoleUpdateOne) ClearStep() *ConsoleUpdateOne {
	cuo.mutation.ClearStep()
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *ConsoleUpdateOne) SetStatus(c console.Status) *ConsoleUpdateOne {
	cuo.mutation.SetStatus(c)
	return cuo
}

// SetBeforeBin sets the "before_bin" field.
func (cuo *ConsoleUpdateOne) SetBeforeBin(mi *model.BinInfo) *ConsoleUpdateOne {
	cuo.mutation.SetBeforeBin(mi)
	return cuo
}

// ClearBeforeBin clears the value of the "before_bin" field.
func (cuo *ConsoleUpdateOne) ClearBeforeBin() *ConsoleUpdateOne {
	cuo.mutation.ClearBeforeBin()
	return cuo
}

// SetAfterBin sets the "after_bin" field.
func (cuo *ConsoleUpdateOne) SetAfterBin(mi *model.BinInfo) *ConsoleUpdateOne {
	cuo.mutation.SetAfterBin(mi)
	return cuo
}

// ClearAfterBin clears the value of the "after_bin" field.
func (cuo *ConsoleUpdateOne) ClearAfterBin() *ConsoleUpdateOne {
	cuo.mutation.ClearAfterBin()
	return cuo
}

// SetMessage sets the "message" field.
func (cuo *ConsoleUpdateOne) SetMessage(s string) *ConsoleUpdateOne {
	cuo.mutation.SetMessage(s)
	return cuo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cuo *ConsoleUpdateOne) SetNillableMessage(s *string) *ConsoleUpdateOne {
	if s != nil {
		cuo.SetMessage(*s)
	}
	return cuo
}

// ClearMessage clears the value of the "message" field.
func (cuo *ConsoleUpdateOne) ClearMessage() *ConsoleUpdateOne {
	cuo.mutation.ClearMessage()
	return cuo
}

// SetStartAt sets the "startAt" field.
func (cuo *ConsoleUpdateOne) SetStartAt(t time.Time) *ConsoleUpdateOne {
	cuo.mutation.SetStartAt(t)
	return cuo
}

// SetNillableStartAt sets the "startAt" field if the given value is not nil.
func (cuo *ConsoleUpdateOne) SetNillableStartAt(t *time.Time) *ConsoleUpdateOne {
	if t != nil {
		cuo.SetStartAt(*t)
	}
	return cuo
}

// ClearStartAt clears the value of the "startAt" field.
func (cuo *ConsoleUpdateOne) ClearStartAt() *ConsoleUpdateOne {
	cuo.mutation.ClearStartAt()
	return cuo
}

// SetStopAt sets the "stopAt" field.
func (cuo *ConsoleUpdateOne) SetStopAt(t time.Time) *ConsoleUpdateOne {
	cuo.mutation.SetStopAt(t)
	return cuo
}

// SetNillableStopAt sets the "stopAt" field if the given value is not nil.
func (cuo *ConsoleUpdateOne) SetNillableStopAt(t *time.Time) *ConsoleUpdateOne {
	if t != nil {
		cuo.SetStopAt(*t)
	}
	return cuo
}

// ClearStopAt clears the value of the "stopAt" field.
func (cuo *ConsoleUpdateOne) ClearStopAt() *ConsoleUpdateOne {
	cuo.mutation.ClearStopAt()
	return cuo
}

// SetDuration sets the "duration" field.
func (cuo *ConsoleUpdateOne) SetDuration(f float64) *ConsoleUpdateOne {
	cuo.mutation.ResetDuration()
	cuo.mutation.SetDuration(f)
	return cuo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (cuo *ConsoleUpdateOne) SetNillableDuration(f *float64) *ConsoleUpdateOne {
	if f != nil {
		cuo.SetDuration(*f)
	}
	return cuo
}

// AddDuration adds f to the "duration" field.
func (cuo *ConsoleUpdateOne) AddDuration(f float64) *ConsoleUpdateOne {
	cuo.mutation.AddDuration(f)
	return cuo
}

// ClearDuration clears the value of the "duration" field.
func (cuo *ConsoleUpdateOne) ClearDuration() *ConsoleUpdateOne {
	cuo.mutation.ClearDuration()
	return cuo
}

// SetCabinet sets the "cabinet" edge to the Cabinet entity.
func (cuo *ConsoleUpdateOne) SetCabinet(c *Cabinet) *ConsoleUpdateOne {
	return cuo.SetCabinetID(c.ID)
}

// SetBin sets the "bin" edge to the Bin entity.
func (cuo *ConsoleUpdateOne) SetBin(b *Bin) *ConsoleUpdateOne {
	return cuo.SetBinID(b.ID)
}

// Mutation returns the ConsoleMutation object of the builder.
func (cuo *ConsoleUpdateOne) Mutation() *ConsoleMutation {
	return cuo.mutation
}

// ClearCabinet clears the "cabinet" edge to the Cabinet entity.
func (cuo *ConsoleUpdateOne) ClearCabinet() *ConsoleUpdateOne {
	cuo.mutation.ClearCabinet()
	return cuo
}

// ClearBin clears the "bin" edge to the Bin entity.
func (cuo *ConsoleUpdateOne) ClearBin() *ConsoleUpdateOne {
	cuo.mutation.ClearBin()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConsoleUpdateOne) Select(field string, fields ...string) *ConsoleUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Console entity.
func (cuo *ConsoleUpdateOne) Save(ctx context.Context) (*Console, error) {
	return withHooks[*Console, ConsoleMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConsoleUpdateOne) SaveX(ctx context.Context) *Console {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConsoleUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConsoleUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ConsoleUpdateOne) check() error {
	if v, ok := cuo.mutation.GetType(); ok {
		if err := console.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Console.type": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Status(); ok {
		if err := console.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Console.status": %w`, err)}
		}
	}
	if _, ok := cuo.mutation.CabinetID(); cuo.mutation.CabinetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Console.cabinet"`)
	}
	if _, ok := cuo.mutation.BinID(); cuo.mutation.BinCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Console.bin"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *ConsoleUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ConsoleUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *ConsoleUpdateOne) sqlSave(ctx context.Context) (_node *Console, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   console.Table,
			Columns: console.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: console.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Console.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, console.FieldID)
		for _, f := range fields {
			if !console.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != console.FieldID {
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
	if value, ok := cuo.mutation.Operate(); ok {
		_spec.SetField(console.FieldOperate, field.TypeOther, value)
	}
	if cuo.mutation.OperateCleared() {
		_spec.ClearField(console.FieldOperate, field.TypeOther)
	}
	if value, ok := cuo.mutation.Serial(); ok {
		_spec.SetField(console.FieldSerial, field.TypeString, value)
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.SetField(console.FieldType, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.UserID(); ok {
		_spec.SetField(console.FieldUserID, field.TypeString, value)
	}
	if value, ok := cuo.mutation.UserType(); ok {
		_spec.SetField(console.FieldUserType, field.TypeOther, value)
	}
	if value, ok := cuo.mutation.Step(); ok {
		_spec.SetField(console.FieldStep, field.TypeOther, value)
	}
	if cuo.mutation.StepCleared() {
		_spec.ClearField(console.FieldStep, field.TypeOther)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(console.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.BeforeBin(); ok {
		_spec.SetField(console.FieldBeforeBin, field.TypeJSON, value)
	}
	if cuo.mutation.BeforeBinCleared() {
		_spec.ClearField(console.FieldBeforeBin, field.TypeJSON)
	}
	if value, ok := cuo.mutation.AfterBin(); ok {
		_spec.SetField(console.FieldAfterBin, field.TypeJSON, value)
	}
	if cuo.mutation.AfterBinCleared() {
		_spec.ClearField(console.FieldAfterBin, field.TypeJSON)
	}
	if value, ok := cuo.mutation.Message(); ok {
		_spec.SetField(console.FieldMessage, field.TypeString, value)
	}
	if cuo.mutation.MessageCleared() {
		_spec.ClearField(console.FieldMessage, field.TypeString)
	}
	if value, ok := cuo.mutation.StartAt(); ok {
		_spec.SetField(console.FieldStartAt, field.TypeTime, value)
	}
	if cuo.mutation.StartAtCleared() {
		_spec.ClearField(console.FieldStartAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.StopAt(); ok {
		_spec.SetField(console.FieldStopAt, field.TypeTime, value)
	}
	if cuo.mutation.StopAtCleared() {
		_spec.ClearField(console.FieldStopAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.Duration(); ok {
		_spec.SetField(console.FieldDuration, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedDuration(); ok {
		_spec.AddField(console.FieldDuration, field.TypeFloat64, value)
	}
	if cuo.mutation.DurationCleared() {
		_spec.ClearField(console.FieldDuration, field.TypeFloat64)
	}
	if cuo.mutation.CabinetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   console.CabinetTable,
			Columns: []string{console.CabinetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: cabinet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CabinetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   console.CabinetTable,
			Columns: []string{console.CabinetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: cabinet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.BinCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   console.BinTable,
			Columns: []string{console.BinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: bin.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.BinIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   console.BinTable,
			Columns: []string{console.BinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: bin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Console{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{console.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
