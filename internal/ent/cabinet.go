// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
)

// Cabinet is the model entity for the Cabinet schema.
type Cabinet struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 是否在线
	Online bool `json:"online,omitempty"`
	// 品牌
	Brand string `json:"brand,omitempty"`
	// 电柜编号
	Serial string `json:"serial,omitempty"`
	// 状态
	Status cabinet.Status `json:"status,omitempty"`
	// 电柜是否启用
	Enable bool `json:"enable,omitempty"`
	// 经度
	Lng *float64 `json:"lng,omitempty"`
	// 纬度
	Lat *float64 `json:"lat,omitempty"`
	// GSM信号强度
	Gsm *float64 `json:"gsm,omitempty"`
	// 换电柜总电压 (V)
	Voltage *float64 `json:"voltage,omitempty"`
	// 换电柜总电流 (A)
	Current *float64 `json:"current,omitempty"`
	// 柜体温度值 (换电柜温度)
	Temperature *float64 `json:"temperature,omitempty"`
	// 总用电量
	Electricity *float64 `json:"electricity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CabinetQuery when eager-loading is set.
	Edges CabinetEdges `json:"edges"`
}

// CabinetEdges holds the relations/edges for other nodes in the graph.
type CabinetEdges struct {
	// Bins holds the value of the bins edge.
	Bins []*Bin `json:"bins,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BinsOrErr returns the Bins value or an error if the edge
// was not loaded in eager-loading.
func (e CabinetEdges) BinsOrErr() ([]*Bin, error) {
	if e.loadedTypes[0] {
		return e.Bins, nil
	}
	return nil, &NotLoadedError{edge: "bins"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cabinet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cabinet.FieldOnline, cabinet.FieldEnable:
			values[i] = new(sql.NullBool)
		case cabinet.FieldLng, cabinet.FieldLat, cabinet.FieldGsm, cabinet.FieldVoltage, cabinet.FieldCurrent, cabinet.FieldTemperature, cabinet.FieldElectricity:
			values[i] = new(sql.NullFloat64)
		case cabinet.FieldID:
			values[i] = new(sql.NullInt64)
		case cabinet.FieldBrand, cabinet.FieldSerial, cabinet.FieldStatus:
			values[i] = new(sql.NullString)
		case cabinet.FieldCreatedAt, cabinet.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Cabinet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cabinet fields.
func (c *Cabinet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cabinet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint64(value.Int64)
		case cabinet.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case cabinet.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case cabinet.FieldOnline:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field online", values[i])
			} else if value.Valid {
				c.Online = value.Bool
			}
		case cabinet.FieldBrand:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field brand", values[i])
			} else if value.Valid {
				c.Brand = value.String
			}
		case cabinet.FieldSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial", values[i])
			} else if value.Valid {
				c.Serial = value.String
			}
		case cabinet.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = cabinet.Status(value.String)
			}
		case cabinet.FieldEnable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enable", values[i])
			} else if value.Valid {
				c.Enable = value.Bool
			}
		case cabinet.FieldLng:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field lng", values[i])
			} else if value.Valid {
				c.Lng = new(float64)
				*c.Lng = value.Float64
			}
		case cabinet.FieldLat:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field lat", values[i])
			} else if value.Valid {
				c.Lat = new(float64)
				*c.Lat = value.Float64
			}
		case cabinet.FieldGsm:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field gsm", values[i])
			} else if value.Valid {
				c.Gsm = new(float64)
				*c.Gsm = value.Float64
			}
		case cabinet.FieldVoltage:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field voltage", values[i])
			} else if value.Valid {
				c.Voltage = new(float64)
				*c.Voltage = value.Float64
			}
		case cabinet.FieldCurrent:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field current", values[i])
			} else if value.Valid {
				c.Current = new(float64)
				*c.Current = value.Float64
			}
		case cabinet.FieldTemperature:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field temperature", values[i])
			} else if value.Valid {
				c.Temperature = new(float64)
				*c.Temperature = value.Float64
			}
		case cabinet.FieldElectricity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field electricity", values[i])
			} else if value.Valid {
				c.Electricity = new(float64)
				*c.Electricity = value.Float64
			}
		}
	}
	return nil
}

// QueryBins queries the "bins" edge of the Cabinet entity.
func (c *Cabinet) QueryBins() *BinQuery {
	return (&CabinetClient{config: c.config}).QueryBins(c)
}

// Update returns a builder for updating this Cabinet.
// Note that you need to call Cabinet.Unwrap() before calling this method if this Cabinet
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cabinet) Update() *CabinetUpdateOne {
	return (&CabinetClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Cabinet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cabinet) Unwrap() *Cabinet {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cabinet is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cabinet) String() string {
	var builder strings.Builder
	builder.WriteString("Cabinet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("online=")
	builder.WriteString(fmt.Sprintf("%v", c.Online))
	builder.WriteString(", ")
	builder.WriteString("brand=")
	builder.WriteString(c.Brand)
	builder.WriteString(", ")
	builder.WriteString("serial=")
	builder.WriteString(c.Serial)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", ")
	builder.WriteString("enable=")
	builder.WriteString(fmt.Sprintf("%v", c.Enable))
	builder.WriteString(", ")
	if v := c.Lng; v != nil {
		builder.WriteString("lng=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.Lat; v != nil {
		builder.WriteString("lat=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.Gsm; v != nil {
		builder.WriteString("gsm=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.Voltage; v != nil {
		builder.WriteString("voltage=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.Current; v != nil {
		builder.WriteString("current=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.Temperature; v != nil {
		builder.WriteString("temperature=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.Electricity; v != nil {
		builder.WriteString("electricity=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Cabinets is a parsable slice of Cabinet.
type Cabinets []*Cabinet

func (c Cabinets) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
