// Code generated by ent, DO NOT EDIT.

package console

import (
	"fmt"
)

const (
	// Label holds the string label denoting the console type in the database.
	Label = "console"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCabinetID holds the string denoting the cabinet_id field in the database.
	FieldCabinetID = "cabinet_id"
	// FieldBinID holds the string denoting the bin_id field in the database.
	FieldBinID = "bin_id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldUserType holds the string denoting the user_type field in the database.
	FieldUserType = "user_type"
	// FieldStep holds the string denoting the step field in the database.
	FieldStep = "step"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldBeforeBin holds the string denoting the before_bin field in the database.
	FieldBeforeBin = "before_bin"
	// FieldAfterBin holds the string denoting the after_bin field in the database.
	FieldAfterBin = "after_bin"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldStartAt holds the string denoting the startat field in the database.
	FieldStartAt = "start_at"
	// FieldStopAt holds the string denoting the stopat field in the database.
	FieldStopAt = "stop_at"
	// EdgeCabinet holds the string denoting the cabinet edge name in mutations.
	EdgeCabinet = "cabinet"
	// EdgeBin holds the string denoting the bin edge name in mutations.
	EdgeBin = "bin"
	// Table holds the table name of the console in the database.
	Table = "console"
	// CabinetTable is the table that holds the cabinet relation/edge.
	CabinetTable = "console"
	// CabinetInverseTable is the table name for the Cabinet entity.
	// It exists in this package in order to avoid circular dependency with the "cabinet" package.
	CabinetInverseTable = "cabinet"
	// CabinetColumn is the table column denoting the cabinet relation/edge.
	CabinetColumn = "cabinet_id"
	// BinTable is the table that holds the bin relation/edge.
	BinTable = "console"
	// BinInverseTable is the table name for the Bin entity.
	// It exists in this package in order to avoid circular dependency with the "bin" package.
	BinInverseTable = "bin"
	// BinColumn is the table column denoting the bin relation/edge.
	BinColumn = "bin_id"
)

// Columns holds all SQL columns for console fields.
var Columns = []string{
	FieldID,
	FieldCabinetID,
	FieldBinID,
	FieldUUID,
	FieldType,
	FieldUserID,
	FieldUserType,
	FieldStep,
	FieldStatus,
	FieldBeforeBin,
	FieldAfterBin,
	FieldMessage,
	FieldStartAt,
	FieldStopAt,
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

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeExchange Type = "exchange"
	TypeControl  Type = "control"
	TypeCabinet  Type = "cabinet"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeExchange, TypeControl, TypeCabinet:
		return nil
	default:
		return fmt.Errorf("console: invalid enum value for type field: %q", _type)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusInvalid Status = "invalid"
	StatusPending Status = "pending"
	StatusRunning Status = "running"
	StatusSuccess Status = "success"
	StatusFailed  Status = "failed"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusInvalid, StatusPending, StatusRunning, StatusSuccess, StatusFailed:
		return nil
	default:
		return fmt.Errorf("console: invalid enum value for status field: %q", s)
	}
}
