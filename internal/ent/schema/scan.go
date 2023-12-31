package schema

import (
	"ariga.io/atlas/sql/postgres"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/defs/cabdef"
	"github.com/google/uuid"

	"github.com/auroraride/cabservd/internal/ent/internal"
)

type ScanMixin struct {
	mixin.Schema
	Optional     bool
	DisableIndex bool
}

func (m ScanMixin) Fields() []ent.Field {
	relate := field.Uint64("scan_id")
	if m.Optional {
		relate.Optional().Nillable()
	}
	return []ent.Field{
		relate,
	}
}

func (m ScanMixin) Edges() []ent.Edge {
	e := edge.To("scan", Scan.Type).Unique().Field("scan_id")
	if !m.Optional {
		e.Required()
	}
	return []ent.Edge{e}
}

func (m ScanMixin) Indexes() (arr []ent.Index) {
	if !m.DisableIndex {
		arr = append(arr, index.Fields("scan_id"))
	}
	return
}

// Scan holds the schema definition for the Scan entity.
type Scan struct {
	ent.Schema
}

// Annotations of the Scan.
func (Scan) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "scan"},
		entsql.WithComments(true),
	}
}

// Fields of the Scan.
func (Scan) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Unique().Default(uuid.New),
		field.String("order_no").Optional().Nillable().Immutable().Comment("订单编号"),
		field.Enum("business").GoType(adapter.Business("")).Comment("业务 operate:运维操作 exchange:换电 active:激活 pause:寄存 continue:结束寄存 unsubscribe:退订"),
		field.Bool("efficient").Default(true).Comment("是否有效"),
		field.String("user_id").Comment("用户ID"),
		field.Other("user_type", adapter.UserTypeUnknown).SchemaType(map[string]string{dialect.Postgres: postgres.TypeVarChar}).Comment("用户类别"),
		field.String("serial").Comment("电柜编号"),
		field.JSON("data", &cabdef.CabinetBinUsableResponse{}).Optional().Comment("换电信息"),
	}
}

// Edges of the Scan.
func (Scan) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Scan) Mixin() []ent.Mixin {
	return []ent.Mixin{
		internal.TimeMixin{},
		CabinetMixin{},
	}
}

func (Scan) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("user_type"),
		index.Fields("serial"),
	}
}
