// Code generated by ent, DO NOT EDIT.

package cabinet

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/auroraride/cabservd/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldUpdatedAt, v))
}

// Online applies equality check predicate on the "online" field. It's identical to OnlineEQ.
func Online(v bool) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldOnline, v))
}

// Power applies equality check predicate on the "power" field. It's identical to PowerEQ.
func Power(v bool) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldPower, v))
}

// Serial applies equality check predicate on the "serial" field. It's identical to SerialEQ.
func Serial(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldSerial, v))
}

// Enable applies equality check predicate on the "enable" field. It's identical to EnableEQ.
func Enable(v bool) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldEnable, v))
}

// Lng applies equality check predicate on the "lng" field. It's identical to LngEQ.
func Lng(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldLng, v))
}

// Lat applies equality check predicate on the "lat" field. It's identical to LatEQ.
func Lat(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldLat, v))
}

// Gsm applies equality check predicate on the "gsm" field. It's identical to GsmEQ.
func Gsm(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldGsm, v))
}

// Voltage applies equality check predicate on the "voltage" field. It's identical to VoltageEQ.
func Voltage(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldVoltage, v))
}

// Current applies equality check predicate on the "current" field. It's identical to CurrentEQ.
func Current(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldCurrent, v))
}

// Temperature applies equality check predicate on the "temperature" field. It's identical to TemperatureEQ.
func Temperature(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldTemperature, v))
}

// Electricity applies equality check predicate on the "electricity" field. It's identical to ElectricityEQ.
func Electricity(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldElectricity, v))
}

// Sim applies equality check predicate on the "sim" field. It's identical to SimEQ.
func Sim(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldSim, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLTE(FieldUpdatedAt, v))
}

// OnlineEQ applies the EQ predicate on the "online" field.
func OnlineEQ(v bool) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldOnline, v))
}

// OnlineNEQ applies the NEQ predicate on the "online" field.
func OnlineNEQ(v bool) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldOnline, v))
}

// PowerEQ applies the EQ predicate on the "power" field.
func PowerEQ(v bool) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldPower, v))
}

// PowerNEQ applies the NEQ predicate on the "power" field.
func PowerNEQ(v bool) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldPower, v))
}

// SerialEQ applies the EQ predicate on the "serial" field.
func SerialEQ(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldSerial, v))
}

// SerialNEQ applies the NEQ predicate on the "serial" field.
func SerialNEQ(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldSerial, v))
}

// SerialIn applies the In predicate on the "serial" field.
func SerialIn(vs ...string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIn(FieldSerial, vs...))
}

// SerialNotIn applies the NotIn predicate on the "serial" field.
func SerialNotIn(vs ...string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotIn(FieldSerial, vs...))
}

// SerialGT applies the GT predicate on the "serial" field.
func SerialGT(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGT(FieldSerial, v))
}

// SerialGTE applies the GTE predicate on the "serial" field.
func SerialGTE(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGTE(FieldSerial, v))
}

// SerialLT applies the LT predicate on the "serial" field.
func SerialLT(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLT(FieldSerial, v))
}

// SerialLTE applies the LTE predicate on the "serial" field.
func SerialLTE(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLTE(FieldSerial, v))
}

// SerialContains applies the Contains predicate on the "serial" field.
func SerialContains(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldContains(FieldSerial, v))
}

// SerialHasPrefix applies the HasPrefix predicate on the "serial" field.
func SerialHasPrefix(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldHasPrefix(FieldSerial, v))
}

// SerialHasSuffix applies the HasSuffix predicate on the "serial" field.
func SerialHasSuffix(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldHasSuffix(FieldSerial, v))
}

// SerialEqualFold applies the EqualFold predicate on the "serial" field.
func SerialEqualFold(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEqualFold(FieldSerial, v))
}

// SerialContainsFold applies the ContainsFold predicate on the "serial" field.
func SerialContainsFold(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldContainsFold(FieldSerial, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotIn(FieldStatus, vs...))
}

// EnableEQ applies the EQ predicate on the "enable" field.
func EnableEQ(v bool) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldEnable, v))
}

// EnableNEQ applies the NEQ predicate on the "enable" field.
func EnableNEQ(v bool) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldEnable, v))
}

// LngEQ applies the EQ predicate on the "lng" field.
func LngEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldLng, v))
}

// LngNEQ applies the NEQ predicate on the "lng" field.
func LngNEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldLng, v))
}

// LngIn applies the In predicate on the "lng" field.
func LngIn(vs ...float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIn(FieldLng, vs...))
}

// LngNotIn applies the NotIn predicate on the "lng" field.
func LngNotIn(vs ...float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotIn(FieldLng, vs...))
}

// LngGT applies the GT predicate on the "lng" field.
func LngGT(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGT(FieldLng, v))
}

// LngGTE applies the GTE predicate on the "lng" field.
func LngGTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGTE(FieldLng, v))
}

// LngLT applies the LT predicate on the "lng" field.
func LngLT(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLT(FieldLng, v))
}

// LngLTE applies the LTE predicate on the "lng" field.
func LngLTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLTE(FieldLng, v))
}

// LngIsNil applies the IsNil predicate on the "lng" field.
func LngIsNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIsNull(FieldLng))
}

// LngNotNil applies the NotNil predicate on the "lng" field.
func LngNotNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotNull(FieldLng))
}

// LatEQ applies the EQ predicate on the "lat" field.
func LatEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldLat, v))
}

// LatNEQ applies the NEQ predicate on the "lat" field.
func LatNEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldLat, v))
}

// LatIn applies the In predicate on the "lat" field.
func LatIn(vs ...float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIn(FieldLat, vs...))
}

// LatNotIn applies the NotIn predicate on the "lat" field.
func LatNotIn(vs ...float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotIn(FieldLat, vs...))
}

// LatGT applies the GT predicate on the "lat" field.
func LatGT(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGT(FieldLat, v))
}

// LatGTE applies the GTE predicate on the "lat" field.
func LatGTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGTE(FieldLat, v))
}

// LatLT applies the LT predicate on the "lat" field.
func LatLT(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLT(FieldLat, v))
}

// LatLTE applies the LTE predicate on the "lat" field.
func LatLTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLTE(FieldLat, v))
}

// LatIsNil applies the IsNil predicate on the "lat" field.
func LatIsNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIsNull(FieldLat))
}

// LatNotNil applies the NotNil predicate on the "lat" field.
func LatNotNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotNull(FieldLat))
}

// GsmEQ applies the EQ predicate on the "gsm" field.
func GsmEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldGsm, v))
}

// GsmNEQ applies the NEQ predicate on the "gsm" field.
func GsmNEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldGsm, v))
}

// GsmIn applies the In predicate on the "gsm" field.
func GsmIn(vs ...float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIn(FieldGsm, vs...))
}

// GsmNotIn applies the NotIn predicate on the "gsm" field.
func GsmNotIn(vs ...float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotIn(FieldGsm, vs...))
}

// GsmGT applies the GT predicate on the "gsm" field.
func GsmGT(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGT(FieldGsm, v))
}

// GsmGTE applies the GTE predicate on the "gsm" field.
func GsmGTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGTE(FieldGsm, v))
}

// GsmLT applies the LT predicate on the "gsm" field.
func GsmLT(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLT(FieldGsm, v))
}

// GsmLTE applies the LTE predicate on the "gsm" field.
func GsmLTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLTE(FieldGsm, v))
}

// GsmIsNil applies the IsNil predicate on the "gsm" field.
func GsmIsNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIsNull(FieldGsm))
}

// GsmNotNil applies the NotNil predicate on the "gsm" field.
func GsmNotNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotNull(FieldGsm))
}

// VoltageEQ applies the EQ predicate on the "voltage" field.
func VoltageEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldVoltage, v))
}

// VoltageNEQ applies the NEQ predicate on the "voltage" field.
func VoltageNEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldVoltage, v))
}

// VoltageIn applies the In predicate on the "voltage" field.
func VoltageIn(vs ...float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIn(FieldVoltage, vs...))
}

// VoltageNotIn applies the NotIn predicate on the "voltage" field.
func VoltageNotIn(vs ...float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotIn(FieldVoltage, vs...))
}

// VoltageGT applies the GT predicate on the "voltage" field.
func VoltageGT(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGT(FieldVoltage, v))
}

// VoltageGTE applies the GTE predicate on the "voltage" field.
func VoltageGTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGTE(FieldVoltage, v))
}

// VoltageLT applies the LT predicate on the "voltage" field.
func VoltageLT(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLT(FieldVoltage, v))
}

// VoltageLTE applies the LTE predicate on the "voltage" field.
func VoltageLTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLTE(FieldVoltage, v))
}

// VoltageIsNil applies the IsNil predicate on the "voltage" field.
func VoltageIsNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIsNull(FieldVoltage))
}

// VoltageNotNil applies the NotNil predicate on the "voltage" field.
func VoltageNotNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotNull(FieldVoltage))
}

// CurrentEQ applies the EQ predicate on the "current" field.
func CurrentEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldCurrent, v))
}

// CurrentNEQ applies the NEQ predicate on the "current" field.
func CurrentNEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldCurrent, v))
}

// CurrentIn applies the In predicate on the "current" field.
func CurrentIn(vs ...float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIn(FieldCurrent, vs...))
}

// CurrentNotIn applies the NotIn predicate on the "current" field.
func CurrentNotIn(vs ...float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotIn(FieldCurrent, vs...))
}

// CurrentGT applies the GT predicate on the "current" field.
func CurrentGT(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGT(FieldCurrent, v))
}

// CurrentGTE applies the GTE predicate on the "current" field.
func CurrentGTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGTE(FieldCurrent, v))
}

// CurrentLT applies the LT predicate on the "current" field.
func CurrentLT(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLT(FieldCurrent, v))
}

// CurrentLTE applies the LTE predicate on the "current" field.
func CurrentLTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLTE(FieldCurrent, v))
}

// CurrentIsNil applies the IsNil predicate on the "current" field.
func CurrentIsNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIsNull(FieldCurrent))
}

// CurrentNotNil applies the NotNil predicate on the "current" field.
func CurrentNotNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotNull(FieldCurrent))
}

// TemperatureEQ applies the EQ predicate on the "temperature" field.
func TemperatureEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldTemperature, v))
}

// TemperatureNEQ applies the NEQ predicate on the "temperature" field.
func TemperatureNEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldTemperature, v))
}

// TemperatureIn applies the In predicate on the "temperature" field.
func TemperatureIn(vs ...float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIn(FieldTemperature, vs...))
}

// TemperatureNotIn applies the NotIn predicate on the "temperature" field.
func TemperatureNotIn(vs ...float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotIn(FieldTemperature, vs...))
}

// TemperatureGT applies the GT predicate on the "temperature" field.
func TemperatureGT(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGT(FieldTemperature, v))
}

// TemperatureGTE applies the GTE predicate on the "temperature" field.
func TemperatureGTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGTE(FieldTemperature, v))
}

// TemperatureLT applies the LT predicate on the "temperature" field.
func TemperatureLT(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLT(FieldTemperature, v))
}

// TemperatureLTE applies the LTE predicate on the "temperature" field.
func TemperatureLTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLTE(FieldTemperature, v))
}

// TemperatureIsNil applies the IsNil predicate on the "temperature" field.
func TemperatureIsNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIsNull(FieldTemperature))
}

// TemperatureNotNil applies the NotNil predicate on the "temperature" field.
func TemperatureNotNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotNull(FieldTemperature))
}

// ElectricityEQ applies the EQ predicate on the "electricity" field.
func ElectricityEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldElectricity, v))
}

// ElectricityNEQ applies the NEQ predicate on the "electricity" field.
func ElectricityNEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldElectricity, v))
}

// ElectricityIn applies the In predicate on the "electricity" field.
func ElectricityIn(vs ...float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIn(FieldElectricity, vs...))
}

// ElectricityNotIn applies the NotIn predicate on the "electricity" field.
func ElectricityNotIn(vs ...float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotIn(FieldElectricity, vs...))
}

// ElectricityGT applies the GT predicate on the "electricity" field.
func ElectricityGT(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGT(FieldElectricity, v))
}

// ElectricityGTE applies the GTE predicate on the "electricity" field.
func ElectricityGTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGTE(FieldElectricity, v))
}

// ElectricityLT applies the LT predicate on the "electricity" field.
func ElectricityLT(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLT(FieldElectricity, v))
}

// ElectricityLTE applies the LTE predicate on the "electricity" field.
func ElectricityLTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLTE(FieldElectricity, v))
}

// ElectricityIsNil applies the IsNil predicate on the "electricity" field.
func ElectricityIsNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIsNull(FieldElectricity))
}

// ElectricityNotNil applies the NotNil predicate on the "electricity" field.
func ElectricityNotNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotNull(FieldElectricity))
}

// SimEQ applies the EQ predicate on the "sim" field.
func SimEQ(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEQ(FieldSim, v))
}

// SimNEQ applies the NEQ predicate on the "sim" field.
func SimNEQ(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNEQ(FieldSim, v))
}

// SimIn applies the In predicate on the "sim" field.
func SimIn(vs ...string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIn(FieldSim, vs...))
}

// SimNotIn applies the NotIn predicate on the "sim" field.
func SimNotIn(vs ...string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotIn(FieldSim, vs...))
}

// SimGT applies the GT predicate on the "sim" field.
func SimGT(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGT(FieldSim, v))
}

// SimGTE applies the GTE predicate on the "sim" field.
func SimGTE(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldGTE(FieldSim, v))
}

// SimLT applies the LT predicate on the "sim" field.
func SimLT(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLT(FieldSim, v))
}

// SimLTE applies the LTE predicate on the "sim" field.
func SimLTE(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldLTE(FieldSim, v))
}

// SimContains applies the Contains predicate on the "sim" field.
func SimContains(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldContains(FieldSim, v))
}

// SimHasPrefix applies the HasPrefix predicate on the "sim" field.
func SimHasPrefix(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldHasPrefix(FieldSim, v))
}

// SimHasSuffix applies the HasSuffix predicate on the "sim" field.
func SimHasSuffix(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldHasSuffix(FieldSim, v))
}

// SimIsNil applies the IsNil predicate on the "sim" field.
func SimIsNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldIsNull(FieldSim))
}

// SimNotNil applies the NotNil predicate on the "sim" field.
func SimNotNil() predicate.Cabinet {
	return predicate.Cabinet(sql.FieldNotNull(FieldSim))
}

// SimEqualFold applies the EqualFold predicate on the "sim" field.
func SimEqualFold(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldEqualFold(FieldSim, v))
}

// SimContainsFold applies the ContainsFold predicate on the "sim" field.
func SimContainsFold(v string) predicate.Cabinet {
	return predicate.Cabinet(sql.FieldContainsFold(FieldSim, v))
}

// HasBins applies the HasEdge predicate on the "bins" edge.
func HasBins() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BinsTable, BinsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBinsWith applies the HasEdge predicate on the "bins" edge with a given conditions (other predicates).
func HasBinsWith(preds ...predicate.Bin) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		step := newBinsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Cabinet) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Cabinet) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Cabinet) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		p(s.Not())
	})
}
