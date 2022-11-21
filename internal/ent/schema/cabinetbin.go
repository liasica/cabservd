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

type CabinetBinMixin struct {
    mixin.Schema
    Optional     bool
    DisableIndex bool
}

func (m CabinetBinMixin) Fields() []ent.Field {
    relate := field.Uint64("bin_id")
    if m.Optional {
        relate.Optional().Nillable()
    }
    return []ent.Field{
        relate,
    }
}

func (m CabinetBinMixin) Edges() []ent.Edge {
    e := edge.To("bin", CabinetBin.Type).Unique().Field("bin_id")
    if !m.Optional {
        e.Required()
    }
    return []ent.Edge{e}
}

func (m CabinetBinMixin) Indexes() (arr []ent.Index) {
    if !m.DisableIndex {
        arr = append(arr, index.Fields("bin_id"))
    }
    return
}

// CabinetBin holds the schema definition for the CabinetBin entity.
type CabinetBin struct {
    ent.Schema
}

// Annotations of the CabinetBin.
func (CabinetBin) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "cabinet_bin"},
    }
}

// Fields of the CabinetBin.
func (CabinetBin) Fields() []ent.Field {
    return []ent.Field{
        field.String("brand").Comment("品牌"),
        field.String("sn").Comment("电柜设备序列号"),
        field.String("name").Comment("仓位名称(N号仓)"),
        field.Int("index").Comment("仓位序号(从0开始)"),
        field.Bool("open").Comment("仓门是否开启"),

        field.String("battery_sn").Optional().Nillable().Comment("电池序列号"),
        field.Float("voltage").Optional().Nillable().Comment("当前电压"),
        field.Float("current").Optional().Nillable().Comment("当前电流"),
    }
}

// Edges of the CabinetBin.
func (CabinetBin) Edges() []ent.Edge {
    return []ent.Edge{}
}

func (CabinetBin) Mixin() []ent.Mixin {
    return []ent.Mixin{
        internal.TimeMixin{},
    }
}

func (CabinetBin) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("sn", "brand"),
        index.Fields("index"),
        index.Fields("battery_sn"),
    }
}
