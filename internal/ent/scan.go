// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/auroraride/adapter/model"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/ent/scan"
	"github.com/google/uuid"
)

// Scan is the model entity for the Scan schema.
type Scan struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CabinetID holds the value of the "cabinet_id" field.
	CabinetID uint64 `json:"cabinet_id,omitempty"`
	// 是否有效
	Efficient bool `json:"efficient,omitempty"`
	// 用户ID
	UserID string `json:"user_id,omitempty"`
	// 用户类别
	UserType model.UserType `json:"user_type,omitempty"`
	// 电柜编号
	Serial string `json:"serial,omitempty"`
	// 换电信息
	Data *model.ExchangeUsableResponse `json:"data,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ScanQuery when eager-loading is set.
	Edges ScanEdges `json:"edges"`
}

// ScanEdges holds the relations/edges for other nodes in the graph.
type ScanEdges struct {
	// Cabinet holds the value of the cabinet edge.
	Cabinet *Cabinet `json:"cabinet,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CabinetOrErr returns the Cabinet value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ScanEdges) CabinetOrErr() (*Cabinet, error) {
	if e.loadedTypes[0] {
		if e.Cabinet == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: cabinet.Label}
		}
		return e.Cabinet, nil
	}
	return nil, &NotLoadedError{edge: "cabinet"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Scan) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case scan.FieldData:
			values[i] = new([]byte)
		case scan.FieldUserType:
			values[i] = new(model.UserType)
		case scan.FieldEfficient:
			values[i] = new(sql.NullBool)
		case scan.FieldCabinetID:
			values[i] = new(sql.NullInt64)
		case scan.FieldUserID, scan.FieldSerial:
			values[i] = new(sql.NullString)
		case scan.FieldCreatedAt, scan.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case scan.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Scan", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Scan fields.
func (s *Scan) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case scan.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case scan.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case scan.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case scan.FieldCabinetID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cabinet_id", values[i])
			} else if value.Valid {
				s.CabinetID = uint64(value.Int64)
			}
		case scan.FieldEfficient:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field efficient", values[i])
			} else if value.Valid {
				s.Efficient = value.Bool
			}
		case scan.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				s.UserID = value.String
			}
		case scan.FieldUserType:
			if value, ok := values[i].(*model.UserType); !ok {
				return fmt.Errorf("unexpected type %T for field user_type", values[i])
			} else if value != nil {
				s.UserType = *value
			}
		case scan.FieldSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial", values[i])
			} else if value.Valid {
				s.Serial = value.String
			}
		case scan.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Data); err != nil {
					return fmt.Errorf("unmarshal field data: %w", err)
				}
			}
		}
	}
	return nil
}

// QueryCabinet queries the "cabinet" edge of the Scan entity.
func (s *Scan) QueryCabinet() *CabinetQuery {
	return (&ScanClient{config: s.config}).QueryCabinet(s)
}

// Update returns a builder for updating this Scan.
// Note that you need to call Scan.Unwrap() before calling this method if this Scan
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Scan) Update() *ScanUpdateOne {
	return (&ScanClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Scan entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Scan) Unwrap() *Scan {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Scan is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Scan) String() string {
	var builder strings.Builder
	builder.WriteString("Scan(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("cabinet_id=")
	builder.WriteString(fmt.Sprintf("%v", s.CabinetID))
	builder.WriteString(", ")
	builder.WriteString("efficient=")
	builder.WriteString(fmt.Sprintf("%v", s.Efficient))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(s.UserID)
	builder.WriteString(", ")
	builder.WriteString("user_type=")
	builder.WriteString(fmt.Sprintf("%v", s.UserType))
	builder.WriteString(", ")
	builder.WriteString("serial=")
	builder.WriteString(s.Serial)
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(fmt.Sprintf("%v", s.Data))
	builder.WriteByte(')')
	return builder.String()
}

// Scans is a parsable slice of Scan.
type Scans []*Scan

func (s Scans) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
