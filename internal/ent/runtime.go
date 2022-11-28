// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/auroraride/cabservd/internal/ent/bin"
	"github.com/auroraride/cabservd/internal/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	binMixin := schema.Bin{}.Mixin()
	binMixinFields0 := binMixin[0].Fields()
	_ = binMixinFields0
	binFields := schema.Bin{}.Fields()
	_ = binFields
	// binDescCreatedAt is the schema descriptor for created_at field.
	binDescCreatedAt := binMixinFields0[0].Descriptor()
	// bin.DefaultCreatedAt holds the default value on creation for the created_at field.
	bin.DefaultCreatedAt = binDescCreatedAt.Default.(func() time.Time)
	// binDescUpdatedAt is the schema descriptor for updated_at field.
	binDescUpdatedAt := binMixinFields0[1].Descriptor()
	// bin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	bin.DefaultUpdatedAt = binDescUpdatedAt.Default.(func() time.Time)
	// bin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	bin.UpdateDefaultUpdatedAt = binDescUpdatedAt.UpdateDefault.(func() time.Time)
	// binDescUUID is the schema descriptor for uuid field.
	binDescUUID := binFields[0].Descriptor()
	// bin.UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	bin.UUIDValidator = binDescUUID.Validators[0].(func(string) error)
	// binDescOpen is the schema descriptor for open field.
	binDescOpen := binFields[5].Descriptor()
	// bin.DefaultOpen holds the default value on creation for the open field.
	bin.DefaultOpen = binDescOpen.Default.(bool)
	// binDescEnable is the schema descriptor for enable field.
	binDescEnable := binFields[6].Descriptor()
	// bin.DefaultEnable holds the default value on creation for the enable field.
	bin.DefaultEnable = binDescEnable.Default.(bool)
	// binDescBatterySn is the schema descriptor for battery_sn field.
	binDescBatterySn := binFields[7].Descriptor()
	// bin.DefaultBatterySn holds the default value on creation for the battery_sn field.
	bin.DefaultBatterySn = binDescBatterySn.Default.(string)
	// binDescVoltage is the schema descriptor for voltage field.
	binDescVoltage := binFields[8].Descriptor()
	// bin.DefaultVoltage holds the default value on creation for the voltage field.
	bin.DefaultVoltage = binDescVoltage.Default.(float64)
	// binDescCurrent is the schema descriptor for current field.
	binDescCurrent := binFields[9].Descriptor()
	// bin.DefaultCurrent holds the default value on creation for the current field.
	bin.DefaultCurrent = binDescCurrent.Default.(float64)
	// binDescSoc is the schema descriptor for soc field.
	binDescSoc := binFields[10].Descriptor()
	// bin.DefaultSoc holds the default value on creation for the soc field.
	bin.DefaultSoc = binDescSoc.Default.(float64)
	// binDescSoh is the schema descriptor for soh field.
	binDescSoh := binFields[11].Descriptor()
	// bin.DefaultSoh holds the default value on creation for the soh field.
	bin.DefaultSoh = binDescSoh.Default.(float64)
}
