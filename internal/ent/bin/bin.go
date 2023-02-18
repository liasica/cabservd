// Code generated by ent, DO NOT EDIT.

package bin

import (
	"time"
)

const (
	// Label holds the string label denoting the bin type in the database.
	Label = "bin"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldCabinetID holds the string denoting the cabinet_id field in the database.
	FieldCabinetID = "cabinet_id"
	// FieldSerial holds the string denoting the serial field in the database.
	FieldSerial = "serial"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldOrdinal holds the string denoting the ordinal field in the database.
	FieldOrdinal = "ordinal"
	// FieldOpen holds the string denoting the open field in the database.
	FieldOpen = "open"
	// FieldEnable holds the string denoting the enable field in the database.
	FieldEnable = "enable"
	// FieldHealth holds the string denoting the health field in the database.
	FieldHealth = "health"
	// FieldBatteryExists holds the string denoting the battery_exists field in the database.
	FieldBatteryExists = "battery_exists"
	// FieldBatterySn holds the string denoting the battery_sn field in the database.
	FieldBatterySn = "battery_sn"
	// FieldVoltage holds the string denoting the voltage field in the database.
	FieldVoltage = "voltage"
	// FieldCurrent holds the string denoting the current field in the database.
	FieldCurrent = "current"
	// FieldSoc holds the string denoting the soc field in the database.
	FieldSoc = "soc"
	// FieldSoh holds the string denoting the soh field in the database.
	FieldSoh = "soh"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// EdgeCabinet holds the string denoting the cabinet edge name in mutations.
	EdgeCabinet = "cabinet"
	// Table holds the table name of the bin in the database.
	Table = "bin"
	// CabinetTable is the table that holds the cabinet relation/edge.
	CabinetTable = "bin"
	// CabinetInverseTable is the table name for the Cabinet entity.
	// It exists in this package in order to avoid circular dependency with the "cabinet" package.
	CabinetInverseTable = "cabinet"
	// CabinetColumn is the table column denoting the cabinet relation/edge.
	CabinetColumn = "cabinet_id"
)

// Columns holds all SQL columns for bin fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUUID,
	FieldCabinetID,
	FieldSerial,
	FieldName,
	FieldOrdinal,
	FieldOpen,
	FieldEnable,
	FieldHealth,
	FieldBatteryExists,
	FieldBatterySn,
	FieldVoltage,
	FieldCurrent,
	FieldSoc,
	FieldSoh,
	FieldRemark,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	UUIDValidator func(string) error
	// DefaultOpen holds the default value on creation for the "open" field.
	DefaultOpen bool
	// DefaultEnable holds the default value on creation for the "enable" field.
	DefaultEnable bool
	// DefaultHealth holds the default value on creation for the "health" field.
	DefaultHealth bool
	// DefaultBatteryExists holds the default value on creation for the "battery_exists" field.
	DefaultBatteryExists bool
	// DefaultBatterySn holds the default value on creation for the "battery_sn" field.
	DefaultBatterySn string
	// DefaultVoltage holds the default value on creation for the "voltage" field.
	DefaultVoltage float64
	// DefaultCurrent holds the default value on creation for the "current" field.
	DefaultCurrent float64
	// DefaultSoc holds the default value on creation for the "soc" field.
	DefaultSoc float64
	// DefaultSoh holds the default value on creation for the "soh" field.
	DefaultSoh float64
)
