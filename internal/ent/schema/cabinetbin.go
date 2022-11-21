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
        field.String("uuid").Unique().MaxLen(32).Comment("唯一标识"),

        // 电柜信息
        field.String("brand").Comment("品牌"),
        field.String("sn").Comment("电柜设备序列号"),

        // 仓位信息
        field.String("name").Comment("仓位名称(N号仓)"),
        field.Int("index").Comment("仓位序号(从0开始)"),
        field.Bool("open").Default(false).Comment("仓门是否开启"),
        field.Bool("enable").Default(true).Comment("仓位是否启用"),

        // 电池信息
        field.String("battery_sn").Default("").Comment("电池序列号"),
        field.Float("voltage").Default(0).Comment("当前电压"),
        field.Float("current").Default(0).Comment("当前电流"),
        field.Float("soc").Default(0).Comment("电池电量"),
        field.Float("soh").Default(0).Comment("电池健康程度"),
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
        index.Fields("soc"),
    }
}
