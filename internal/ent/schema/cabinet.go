package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
    "entgo.io/ent/schema/mixin"
    "github.com/auroraride/cabservd/internal/ent/internal"
)

type CabinetMixin struct {
    mixin.Schema
    Optional     bool
    DisableIndex bool
}

func (m CabinetMixin) Fields() []ent.Field {
    relate := field.Uint64("cabinet_id")
    if m.Optional {
        relate.Optional().Nillable()
    }
    return []ent.Field{
        relate,
    }
}

func (m CabinetMixin) Edges() []ent.Edge {
    e := edge.To("cabinet", Cabinet.Type).Unique().Field("cabinet_id")
    if !m.Optional {
        e.Required()
    }
    return []ent.Edge{e}
}

func (m CabinetMixin) Indexes() (arr []ent.Index) {
    if !m.DisableIndex {
        arr = append(arr, index.Fields("cabinet_id"))
    }
    return
}

// Cabinet holds the schema definition for the Cabinet entity.
type Cabinet struct {
    ent.Schema
}

// Annotations of the Cabinet.
func (Cabinet) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "cabinet"},
        entsql.WithComments(true),
    }
}

// Fields of the Cabinet.
func (Cabinet) Fields() []ent.Field {
    return []ent.Field{
        field.Bool("online").Default(false).Comment("是否在线"),
        field.Bool("power").Default(true).Comment("市电是否正常"),
        field.String("serial").Unique().Comment("电柜编号"),
        field.Enum("status").Default("initializing").Values("initializing", "normal", "abnormal").Comment("状态"), // initializing:初始化中 normal:正常 abnormal:异常
        field.Bool("enable").Default(false).Comment("电柜是否启用"),
        field.Float("lng").Optional().Nillable().Comment("经度"),
        field.Float("lat").Optional().Nillable().Comment("纬度"),
        field.Float("gsm").Optional().Nillable().Comment("GSM信号强度"),
        field.Float("voltage").Optional().Nillable().Comment("换电柜总电压 (V)"),
        field.Float("current").Optional().Nillable().Comment("换电柜总电流 (A)"),
        field.Float("temperature").Optional().Nillable().Comment("柜体温度值 (换电柜温度)"),
        field.Float("electricity").Optional().Nillable().Comment("总用电量"),
    }
}

// Edges of the Cabinet.
func (Cabinet) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("bins", Bin.Type),
    }
}

func (Cabinet) Mixin() []ent.Mixin {
    return []ent.Mixin{
        internal.TimeMixin{},
    }
}

func (Cabinet) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("status"),
        index.Fields("enable"),
        index.Fields("lng"),
        index.Fields("lat"),
    }
}
