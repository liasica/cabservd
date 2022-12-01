// Code generated by ent, DO NOT EDIT.

package cabinet

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/auroraride/cabservd/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Brand applies equality check predicate on the "brand" field. It's identical to BrandEQ.
func Brand(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBrand), v))
	})
}

// Serial applies equality check predicate on the "serial" field. It's identical to SerialEQ.
func Serial(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSerial), v))
	})
}

// Enable applies equality check predicate on the "enable" field. It's identical to EnableEQ.
func Enable(v bool) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnable), v))
	})
}

// Lng applies equality check predicate on the "lng" field. It's identical to LngEQ.
func Lng(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLng), v))
	})
}

// Lat applies equality check predicate on the "lat" field. It's identical to LatEQ.
func Lat(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLat), v))
	})
}

// Gsm applies equality check predicate on the "gsm" field. It's identical to GsmEQ.
func Gsm(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGsm), v))
	})
}

// Voltage applies equality check predicate on the "voltage" field. It's identical to VoltageEQ.
func Voltage(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVoltage), v))
	})
}

// Current applies equality check predicate on the "current" field. It's identical to CurrentEQ.
func Current(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCurrent), v))
	})
}

// Temperature applies equality check predicate on the "temperature" field. It's identical to TemperatureEQ.
func Temperature(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTemperature), v))
	})
}

// Electricity applies equality check predicate on the "electricity" field. It's identical to ElectricityEQ.
func Electricity(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldElectricity), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// BrandEQ applies the EQ predicate on the "brand" field.
func BrandEQ(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBrand), v))
	})
}

// BrandNEQ applies the NEQ predicate on the "brand" field.
func BrandNEQ(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBrand), v))
	})
}

// BrandIn applies the In predicate on the "brand" field.
func BrandIn(vs ...string) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBrand), v...))
	})
}

// BrandNotIn applies the NotIn predicate on the "brand" field.
func BrandNotIn(vs ...string) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBrand), v...))
	})
}

// BrandGT applies the GT predicate on the "brand" field.
func BrandGT(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBrand), v))
	})
}

// BrandGTE applies the GTE predicate on the "brand" field.
func BrandGTE(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBrand), v))
	})
}

// BrandLT applies the LT predicate on the "brand" field.
func BrandLT(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBrand), v))
	})
}

// BrandLTE applies the LTE predicate on the "brand" field.
func BrandLTE(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBrand), v))
	})
}

// BrandContains applies the Contains predicate on the "brand" field.
func BrandContains(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBrand), v))
	})
}

// BrandHasPrefix applies the HasPrefix predicate on the "brand" field.
func BrandHasPrefix(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBrand), v))
	})
}

// BrandHasSuffix applies the HasSuffix predicate on the "brand" field.
func BrandHasSuffix(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBrand), v))
	})
}

// BrandEqualFold applies the EqualFold predicate on the "brand" field.
func BrandEqualFold(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBrand), v))
	})
}

// BrandContainsFold applies the ContainsFold predicate on the "brand" field.
func BrandContainsFold(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBrand), v))
	})
}

// SerialEQ applies the EQ predicate on the "serial" field.
func SerialEQ(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSerial), v))
	})
}

// SerialNEQ applies the NEQ predicate on the "serial" field.
func SerialNEQ(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSerial), v))
	})
}

// SerialIn applies the In predicate on the "serial" field.
func SerialIn(vs ...string) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSerial), v...))
	})
}

// SerialNotIn applies the NotIn predicate on the "serial" field.
func SerialNotIn(vs ...string) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSerial), v...))
	})
}

// SerialGT applies the GT predicate on the "serial" field.
func SerialGT(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSerial), v))
	})
}

// SerialGTE applies the GTE predicate on the "serial" field.
func SerialGTE(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSerial), v))
	})
}

// SerialLT applies the LT predicate on the "serial" field.
func SerialLT(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSerial), v))
	})
}

// SerialLTE applies the LTE predicate on the "serial" field.
func SerialLTE(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSerial), v))
	})
}

// SerialContains applies the Contains predicate on the "serial" field.
func SerialContains(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSerial), v))
	})
}

// SerialHasPrefix applies the HasPrefix predicate on the "serial" field.
func SerialHasPrefix(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSerial), v))
	})
}

// SerialHasSuffix applies the HasSuffix predicate on the "serial" field.
func SerialHasSuffix(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSerial), v))
	})
}

// SerialEqualFold applies the EqualFold predicate on the "serial" field.
func SerialEqualFold(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSerial), v))
	})
}

// SerialContainsFold applies the ContainsFold predicate on the "serial" field.
func SerialContainsFold(v string) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSerial), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// EnableEQ applies the EQ predicate on the "enable" field.
func EnableEQ(v bool) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnable), v))
	})
}

// EnableNEQ applies the NEQ predicate on the "enable" field.
func EnableNEQ(v bool) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnable), v))
	})
}

// LngEQ applies the EQ predicate on the "lng" field.
func LngEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLng), v))
	})
}

// LngNEQ applies the NEQ predicate on the "lng" field.
func LngNEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLng), v))
	})
}

// LngIn applies the In predicate on the "lng" field.
func LngIn(vs ...float64) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLng), v...))
	})
}

// LngNotIn applies the NotIn predicate on the "lng" field.
func LngNotIn(vs ...float64) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLng), v...))
	})
}

// LngGT applies the GT predicate on the "lng" field.
func LngGT(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLng), v))
	})
}

// LngGTE applies the GTE predicate on the "lng" field.
func LngGTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLng), v))
	})
}

// LngLT applies the LT predicate on the "lng" field.
func LngLT(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLng), v))
	})
}

// LngLTE applies the LTE predicate on the "lng" field.
func LngLTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLng), v))
	})
}

// LngIsNil applies the IsNil predicate on the "lng" field.
func LngIsNil() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLng)))
	})
}

// LngNotNil applies the NotNil predicate on the "lng" field.
func LngNotNil() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLng)))
	})
}

// LatEQ applies the EQ predicate on the "lat" field.
func LatEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLat), v))
	})
}

// LatNEQ applies the NEQ predicate on the "lat" field.
func LatNEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLat), v))
	})
}

// LatIn applies the In predicate on the "lat" field.
func LatIn(vs ...float64) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLat), v...))
	})
}

// LatNotIn applies the NotIn predicate on the "lat" field.
func LatNotIn(vs ...float64) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLat), v...))
	})
}

// LatGT applies the GT predicate on the "lat" field.
func LatGT(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLat), v))
	})
}

// LatGTE applies the GTE predicate on the "lat" field.
func LatGTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLat), v))
	})
}

// LatLT applies the LT predicate on the "lat" field.
func LatLT(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLat), v))
	})
}

// LatLTE applies the LTE predicate on the "lat" field.
func LatLTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLat), v))
	})
}

// LatIsNil applies the IsNil predicate on the "lat" field.
func LatIsNil() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLat)))
	})
}

// LatNotNil applies the NotNil predicate on the "lat" field.
func LatNotNil() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLat)))
	})
}

// GsmEQ applies the EQ predicate on the "gsm" field.
func GsmEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGsm), v))
	})
}

// GsmNEQ applies the NEQ predicate on the "gsm" field.
func GsmNEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGsm), v))
	})
}

// GsmIn applies the In predicate on the "gsm" field.
func GsmIn(vs ...float64) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGsm), v...))
	})
}

// GsmNotIn applies the NotIn predicate on the "gsm" field.
func GsmNotIn(vs ...float64) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGsm), v...))
	})
}

// GsmGT applies the GT predicate on the "gsm" field.
func GsmGT(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGsm), v))
	})
}

// GsmGTE applies the GTE predicate on the "gsm" field.
func GsmGTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGsm), v))
	})
}

// GsmLT applies the LT predicate on the "gsm" field.
func GsmLT(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGsm), v))
	})
}

// GsmLTE applies the LTE predicate on the "gsm" field.
func GsmLTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGsm), v))
	})
}

// GsmIsNil applies the IsNil predicate on the "gsm" field.
func GsmIsNil() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGsm)))
	})
}

// GsmNotNil applies the NotNil predicate on the "gsm" field.
func GsmNotNil() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGsm)))
	})
}

// VoltageEQ applies the EQ predicate on the "voltage" field.
func VoltageEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVoltage), v))
	})
}

// VoltageNEQ applies the NEQ predicate on the "voltage" field.
func VoltageNEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVoltage), v))
	})
}

// VoltageIn applies the In predicate on the "voltage" field.
func VoltageIn(vs ...float64) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldVoltage), v...))
	})
}

// VoltageNotIn applies the NotIn predicate on the "voltage" field.
func VoltageNotIn(vs ...float64) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldVoltage), v...))
	})
}

// VoltageGT applies the GT predicate on the "voltage" field.
func VoltageGT(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVoltage), v))
	})
}

// VoltageGTE applies the GTE predicate on the "voltage" field.
func VoltageGTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVoltage), v))
	})
}

// VoltageLT applies the LT predicate on the "voltage" field.
func VoltageLT(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVoltage), v))
	})
}

// VoltageLTE applies the LTE predicate on the "voltage" field.
func VoltageLTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVoltage), v))
	})
}

// VoltageIsNil applies the IsNil predicate on the "voltage" field.
func VoltageIsNil() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldVoltage)))
	})
}

// VoltageNotNil applies the NotNil predicate on the "voltage" field.
func VoltageNotNil() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldVoltage)))
	})
}

// CurrentEQ applies the EQ predicate on the "current" field.
func CurrentEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCurrent), v))
	})
}

// CurrentNEQ applies the NEQ predicate on the "current" field.
func CurrentNEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCurrent), v))
	})
}

// CurrentIn applies the In predicate on the "current" field.
func CurrentIn(vs ...float64) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCurrent), v...))
	})
}

// CurrentNotIn applies the NotIn predicate on the "current" field.
func CurrentNotIn(vs ...float64) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCurrent), v...))
	})
}

// CurrentGT applies the GT predicate on the "current" field.
func CurrentGT(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCurrent), v))
	})
}

// CurrentGTE applies the GTE predicate on the "current" field.
func CurrentGTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCurrent), v))
	})
}

// CurrentLT applies the LT predicate on the "current" field.
func CurrentLT(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCurrent), v))
	})
}

// CurrentLTE applies the LTE predicate on the "current" field.
func CurrentLTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCurrent), v))
	})
}

// CurrentIsNil applies the IsNil predicate on the "current" field.
func CurrentIsNil() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCurrent)))
	})
}

// CurrentNotNil applies the NotNil predicate on the "current" field.
func CurrentNotNil() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCurrent)))
	})
}

// TemperatureEQ applies the EQ predicate on the "temperature" field.
func TemperatureEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTemperature), v))
	})
}

// TemperatureNEQ applies the NEQ predicate on the "temperature" field.
func TemperatureNEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTemperature), v))
	})
}

// TemperatureIn applies the In predicate on the "temperature" field.
func TemperatureIn(vs ...float64) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTemperature), v...))
	})
}

// TemperatureNotIn applies the NotIn predicate on the "temperature" field.
func TemperatureNotIn(vs ...float64) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTemperature), v...))
	})
}

// TemperatureGT applies the GT predicate on the "temperature" field.
func TemperatureGT(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTemperature), v))
	})
}

// TemperatureGTE applies the GTE predicate on the "temperature" field.
func TemperatureGTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTemperature), v))
	})
}

// TemperatureLT applies the LT predicate on the "temperature" field.
func TemperatureLT(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTemperature), v))
	})
}

// TemperatureLTE applies the LTE predicate on the "temperature" field.
func TemperatureLTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTemperature), v))
	})
}

// TemperatureIsNil applies the IsNil predicate on the "temperature" field.
func TemperatureIsNil() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTemperature)))
	})
}

// TemperatureNotNil applies the NotNil predicate on the "temperature" field.
func TemperatureNotNil() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTemperature)))
	})
}

// ElectricityEQ applies the EQ predicate on the "electricity" field.
func ElectricityEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldElectricity), v))
	})
}

// ElectricityNEQ applies the NEQ predicate on the "electricity" field.
func ElectricityNEQ(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldElectricity), v))
	})
}

// ElectricityIn applies the In predicate on the "electricity" field.
func ElectricityIn(vs ...float64) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldElectricity), v...))
	})
}

// ElectricityNotIn applies the NotIn predicate on the "electricity" field.
func ElectricityNotIn(vs ...float64) predicate.Cabinet {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldElectricity), v...))
	})
}

// ElectricityGT applies the GT predicate on the "electricity" field.
func ElectricityGT(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldElectricity), v))
	})
}

// ElectricityGTE applies the GTE predicate on the "electricity" field.
func ElectricityGTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldElectricity), v))
	})
}

// ElectricityLT applies the LT predicate on the "electricity" field.
func ElectricityLT(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldElectricity), v))
	})
}

// ElectricityLTE applies the LTE predicate on the "electricity" field.
func ElectricityLTE(v float64) predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldElectricity), v))
	})
}

// ElectricityIsNil applies the IsNil predicate on the "electricity" field.
func ElectricityIsNil() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldElectricity)))
	})
}

// ElectricityNotNil applies the NotNil predicate on the "electricity" field.
func ElectricityNotNil() predicate.Cabinet {
	return predicate.Cabinet(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldElectricity)))
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