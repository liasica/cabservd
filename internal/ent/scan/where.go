// Code generated by ent, DO NOT EDIT.

package scan

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/auroraride/adapter/model"
	"github.com/auroraride/cabservd/internal/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Scan {
	return predicate.Scan(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Scan {
	return predicate.Scan(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Scan {
	return predicate.Scan(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Scan {
	return predicate.Scan(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Scan {
	return predicate.Scan(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Scan {
	return predicate.Scan(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Scan {
	return predicate.Scan(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldUpdatedAt, v))
}

// CabinetID applies equality check predicate on the "cabinet_id" field. It's identical to CabinetIDEQ.
func CabinetID(v uint64) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldCabinetID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldUserID, v))
}

// UserType applies equality check predicate on the "user_type" field. It's identical to UserTypeEQ.
func UserType(v model.UserType) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldUserType, v))
}

// Serial applies equality check predicate on the "serial" field. It's identical to SerialEQ.
func Serial(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldSerial, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Scan {
	return predicate.Scan(sql.FieldLTE(FieldUpdatedAt, v))
}

// CabinetIDEQ applies the EQ predicate on the "cabinet_id" field.
func CabinetIDEQ(v uint64) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldCabinetID, v))
}

// CabinetIDNEQ applies the NEQ predicate on the "cabinet_id" field.
func CabinetIDNEQ(v uint64) predicate.Scan {
	return predicate.Scan(sql.FieldNEQ(FieldCabinetID, v))
}

// CabinetIDIn applies the In predicate on the "cabinet_id" field.
func CabinetIDIn(vs ...uint64) predicate.Scan {
	return predicate.Scan(sql.FieldIn(FieldCabinetID, vs...))
}

// CabinetIDNotIn applies the NotIn predicate on the "cabinet_id" field.
func CabinetIDNotIn(vs ...uint64) predicate.Scan {
	return predicate.Scan(sql.FieldNotIn(FieldCabinetID, vs...))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Scan {
	return predicate.Scan(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Scan {
	return predicate.Scan(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Scan {
	return predicate.Scan(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Scan {
	return predicate.Scan(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Scan {
	return predicate.Scan(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Scan {
	return predicate.Scan(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Scan {
	return predicate.Scan(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Scan {
	return predicate.Scan(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Scan {
	return predicate.Scan(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Scan {
	return predicate.Scan(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Scan {
	return predicate.Scan(sql.FieldContainsFold(FieldUserID, v))
}

// UserTypeEQ applies the EQ predicate on the "user_type" field.
func UserTypeEQ(v model.UserType) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldUserType, v))
}

// UserTypeNEQ applies the NEQ predicate on the "user_type" field.
func UserTypeNEQ(v model.UserType) predicate.Scan {
	return predicate.Scan(sql.FieldNEQ(FieldUserType, v))
}

// UserTypeIn applies the In predicate on the "user_type" field.
func UserTypeIn(vs ...model.UserType) predicate.Scan {
	return predicate.Scan(sql.FieldIn(FieldUserType, vs...))
}

// UserTypeNotIn applies the NotIn predicate on the "user_type" field.
func UserTypeNotIn(vs ...model.UserType) predicate.Scan {
	return predicate.Scan(sql.FieldNotIn(FieldUserType, vs...))
}

// UserTypeGT applies the GT predicate on the "user_type" field.
func UserTypeGT(v model.UserType) predicate.Scan {
	return predicate.Scan(sql.FieldGT(FieldUserType, v))
}

// UserTypeGTE applies the GTE predicate on the "user_type" field.
func UserTypeGTE(v model.UserType) predicate.Scan {
	return predicate.Scan(sql.FieldGTE(FieldUserType, v))
}

// UserTypeLT applies the LT predicate on the "user_type" field.
func UserTypeLT(v model.UserType) predicate.Scan {
	return predicate.Scan(sql.FieldLT(FieldUserType, v))
}

// UserTypeLTE applies the LTE predicate on the "user_type" field.
func UserTypeLTE(v model.UserType) predicate.Scan {
	return predicate.Scan(sql.FieldLTE(FieldUserType, v))
}

// SerialEQ applies the EQ predicate on the "serial" field.
func SerialEQ(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEQ(FieldSerial, v))
}

// SerialNEQ applies the NEQ predicate on the "serial" field.
func SerialNEQ(v string) predicate.Scan {
	return predicate.Scan(sql.FieldNEQ(FieldSerial, v))
}

// SerialIn applies the In predicate on the "serial" field.
func SerialIn(vs ...string) predicate.Scan {
	return predicate.Scan(sql.FieldIn(FieldSerial, vs...))
}

// SerialNotIn applies the NotIn predicate on the "serial" field.
func SerialNotIn(vs ...string) predicate.Scan {
	return predicate.Scan(sql.FieldNotIn(FieldSerial, vs...))
}

// SerialGT applies the GT predicate on the "serial" field.
func SerialGT(v string) predicate.Scan {
	return predicate.Scan(sql.FieldGT(FieldSerial, v))
}

// SerialGTE applies the GTE predicate on the "serial" field.
func SerialGTE(v string) predicate.Scan {
	return predicate.Scan(sql.FieldGTE(FieldSerial, v))
}

// SerialLT applies the LT predicate on the "serial" field.
func SerialLT(v string) predicate.Scan {
	return predicate.Scan(sql.FieldLT(FieldSerial, v))
}

// SerialLTE applies the LTE predicate on the "serial" field.
func SerialLTE(v string) predicate.Scan {
	return predicate.Scan(sql.FieldLTE(FieldSerial, v))
}

// SerialContains applies the Contains predicate on the "serial" field.
func SerialContains(v string) predicate.Scan {
	return predicate.Scan(sql.FieldContains(FieldSerial, v))
}

// SerialHasPrefix applies the HasPrefix predicate on the "serial" field.
func SerialHasPrefix(v string) predicate.Scan {
	return predicate.Scan(sql.FieldHasPrefix(FieldSerial, v))
}

// SerialHasSuffix applies the HasSuffix predicate on the "serial" field.
func SerialHasSuffix(v string) predicate.Scan {
	return predicate.Scan(sql.FieldHasSuffix(FieldSerial, v))
}

// SerialEqualFold applies the EqualFold predicate on the "serial" field.
func SerialEqualFold(v string) predicate.Scan {
	return predicate.Scan(sql.FieldEqualFold(FieldSerial, v))
}

// SerialContainsFold applies the ContainsFold predicate on the "serial" field.
func SerialContainsFold(v string) predicate.Scan {
	return predicate.Scan(sql.FieldContainsFold(FieldSerial, v))
}

// DataIsNil applies the IsNil predicate on the "data" field.
func DataIsNil() predicate.Scan {
	return predicate.Scan(sql.FieldIsNull(FieldData))
}

// DataNotNil applies the NotNil predicate on the "data" field.
func DataNotNil() predicate.Scan {
	return predicate.Scan(sql.FieldNotNull(FieldData))
}

// HasCabinet applies the HasEdge predicate on the "cabinet" edge.
func HasCabinet() predicate.Scan {
	return predicate.Scan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CabinetTable, CabinetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCabinetWith applies the HasEdge predicate on the "cabinet" edge with a given conditions (other predicates).
func HasCabinetWith(preds ...predicate.Cabinet) predicate.Scan {
	return predicate.Scan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CabinetInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CabinetTable, CabinetColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Scan) predicate.Scan {
	return predicate.Scan(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Scan) predicate.Scan {
	return predicate.Scan(func(s *sql.Selector) {
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
func Not(p predicate.Scan) predicate.Scan {
	return predicate.Scan(func(s *sql.Selector) {
		p(s.Not())
	})
}