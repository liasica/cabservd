// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/defs/cabdef"
	"github.com/auroraride/cabservd/internal/ent/bin"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/ent/console"
	"github.com/google/uuid"
)

// Console is the model entity for the Console schema.
type Console struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CabinetID holds the value of the "cabinet_id" field.
	CabinetID uint64 `json:"cabinet_id,omitempty"`
	// BinID holds the value of the "bin_id" field.
	BinID *uint64 `json:"bin_id,omitempty"`
	// 操作
	Operate cabdef.Operate `json:"operate,omitempty"`
	// 电柜设备序列号
	Serial string `json:"serial,omitempty"`
	// 标识符
	UUID uuid.UUID `json:"uuid,omitempty"`
	// 业务 operate:运维操作 exchange:换电 active:激活 pause:寄存 continue:结束寄存 unsubscribe:退订
	Business adapter.Business `json:"business,omitempty"`
	// 用户ID
	UserID string `json:"user_id,omitempty"`
	// 用户类别
	UserType adapter.UserType `json:"user_type,omitempty"`
	// 步骤
	Step int `json:"step,omitempty"`
	// 状态 invalid:无效 pending:未开始 running:执行中 success:成功 failed:失败
	Status console.Status `json:"status,omitempty"`
	// 变化前仓位信息
	BeforeBin *cabdef.BinInfo `json:"before_bin,omitempty"`
	// 变化后仓位信息
	AfterBin *cabdef.BinInfo `json:"after_bin,omitempty"`
	// 消息
	Message *string `json:"message,omitempty"`
	// 开始时间
	StartAt *time.Time `json:"startAt,omitempty"`
	// 结束时间
	StopAt *time.Time `json:"stopAt,omitempty"`
	// 耗时
	Duration *float64 `json:"duration,omitempty"`
	// 备注信息
	Remark *string `json:"remark,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ConsoleQuery when eager-loading is set.
	Edges ConsoleEdges `json:"edges"`
}

// ConsoleEdges holds the relations/edges for other nodes in the graph.
type ConsoleEdges struct {
	// Cabinet holds the value of the cabinet edge.
	Cabinet *Cabinet `json:"cabinet,omitempty"`
	// Bin holds the value of the bin edge.
	Bin *Bin `json:"bin,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CabinetOrErr returns the Cabinet value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ConsoleEdges) CabinetOrErr() (*Cabinet, error) {
	if e.loadedTypes[0] {
		if e.Cabinet == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: cabinet.Label}
		}
		return e.Cabinet, nil
	}
	return nil, &NotLoadedError{edge: "cabinet"}
}

// BinOrErr returns the Bin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ConsoleEdges) BinOrErr() (*Bin, error) {
	if e.loadedTypes[1] {
		if e.Bin == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: bin.Label}
		}
		return e.Bin, nil
	}
	return nil, &NotLoadedError{edge: "bin"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Console) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case console.FieldBeforeBin, console.FieldAfterBin:
			values[i] = new([]byte)
		case console.FieldBusiness:
			values[i] = new(adapter.Business)
		case console.FieldUserType:
			values[i] = new(adapter.UserType)
		case console.FieldOperate:
			values[i] = new(cabdef.Operate)
		case console.FieldDuration:
			values[i] = new(sql.NullFloat64)
		case console.FieldID, console.FieldCabinetID, console.FieldBinID, console.FieldStep:
			values[i] = new(sql.NullInt64)
		case console.FieldSerial, console.FieldUserID, console.FieldStatus, console.FieldMessage, console.FieldRemark:
			values[i] = new(sql.NullString)
		case console.FieldStartAt, console.FieldStopAt:
			values[i] = new(sql.NullTime)
		case console.FieldUUID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Console", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Console fields.
func (c *Console) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case console.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint64(value.Int64)
		case console.FieldCabinetID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cabinet_id", values[i])
			} else if value.Valid {
				c.CabinetID = uint64(value.Int64)
			}
		case console.FieldBinID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field bin_id", values[i])
			} else if value.Valid {
				c.BinID = new(uint64)
				*c.BinID = uint64(value.Int64)
			}
		case console.FieldOperate:
			if value, ok := values[i].(*cabdef.Operate); !ok {
				return fmt.Errorf("unexpected type %T for field operate", values[i])
			} else if value != nil {
				c.Operate = *value
			}
		case console.FieldSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial", values[i])
			} else if value.Valid {
				c.Serial = value.String
			}
		case console.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				c.UUID = *value
			}
		case console.FieldBusiness:
			if value, ok := values[i].(*adapter.Business); !ok {
				return fmt.Errorf("unexpected type %T for field business", values[i])
			} else if value != nil {
				c.Business = *value
			}
		case console.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				c.UserID = value.String
			}
		case console.FieldUserType:
			if value, ok := values[i].(*adapter.UserType); !ok {
				return fmt.Errorf("unexpected type %T for field user_type", values[i])
			} else if value != nil {
				c.UserType = *value
			}
		case console.FieldStep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field step", values[i])
			} else if value.Valid {
				c.Step = int(value.Int64)
			}
		case console.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = console.Status(value.String)
			}
		case console.FieldBeforeBin:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field before_bin", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.BeforeBin); err != nil {
					return fmt.Errorf("unmarshal field before_bin: %w", err)
				}
			}
		case console.FieldAfterBin:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field after_bin", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.AfterBin); err != nil {
					return fmt.Errorf("unmarshal field after_bin: %w", err)
				}
			}
		case console.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				c.Message = new(string)
				*c.Message = value.String
			}
		case console.FieldStartAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field startAt", values[i])
			} else if value.Valid {
				c.StartAt = new(time.Time)
				*c.StartAt = value.Time
			}
		case console.FieldStopAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field stopAt", values[i])
			} else if value.Valid {
				c.StopAt = new(time.Time)
				*c.StopAt = value.Time
			}
		case console.FieldDuration:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				c.Duration = new(float64)
				*c.Duration = value.Float64
			}
		case console.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				c.Remark = new(string)
				*c.Remark = value.String
			}
		}
	}
	return nil
}

// QueryCabinet queries the "cabinet" edge of the Console entity.
func (c *Console) QueryCabinet() *CabinetQuery {
	return (&ConsoleClient{config: c.config}).QueryCabinet(c)
}

// QueryBin queries the "bin" edge of the Console entity.
func (c *Console) QueryBin() *BinQuery {
	return (&ConsoleClient{config: c.config}).QueryBin(c)
}

// Update returns a builder for updating this Console.
// Note that you need to call Console.Unwrap() before calling this method if this Console
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Console) Update() *ConsoleUpdateOne {
	return (&ConsoleClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Console entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Console) Unwrap() *Console {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Console is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Console) String() string {
	var builder strings.Builder
	builder.WriteString("Console(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("cabinet_id=")
	builder.WriteString(fmt.Sprintf("%v", c.CabinetID))
	builder.WriteString(", ")
	if v := c.BinID; v != nil {
		builder.WriteString("bin_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("operate=")
	builder.WriteString(fmt.Sprintf("%v", c.Operate))
	builder.WriteString(", ")
	builder.WriteString("serial=")
	builder.WriteString(c.Serial)
	builder.WriteString(", ")
	builder.WriteString("uuid=")
	builder.WriteString(fmt.Sprintf("%v", c.UUID))
	builder.WriteString(", ")
	builder.WriteString("business=")
	builder.WriteString(fmt.Sprintf("%v", c.Business))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(c.UserID)
	builder.WriteString(", ")
	builder.WriteString("user_type=")
	builder.WriteString(fmt.Sprintf("%v", c.UserType))
	builder.WriteString(", ")
	builder.WriteString("step=")
	builder.WriteString(fmt.Sprintf("%v", c.Step))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", ")
	builder.WriteString("before_bin=")
	builder.WriteString(fmt.Sprintf("%v", c.BeforeBin))
	builder.WriteString(", ")
	builder.WriteString("after_bin=")
	builder.WriteString(fmt.Sprintf("%v", c.AfterBin))
	builder.WriteString(", ")
	if v := c.Message; v != nil {
		builder.WriteString("message=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := c.StartAt; v != nil {
		builder.WriteString("startAt=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := c.StopAt; v != nil {
		builder.WriteString("stopAt=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := c.Duration; v != nil {
		builder.WriteString("duration=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := c.Remark; v != nil {
		builder.WriteString("remark=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Consoles is a parsable slice of Console.
type Consoles []*Console

func (c Consoles) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
