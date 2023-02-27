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

type BinMixin struct {
    mixin.Schema
    Optional     bool
    DisableIndex bool
}

func (m BinMixin) Fields() []ent.Field {
    relate := field.Uint64("bin_id")
    if m.Optional {
        relate.Optional().Nillable()
    }
    return []ent.Field{
        relate,
    }
}

func (m BinMixin) Edges() []ent.Edge {
    e := edge.To("bin", Bin.Type).Unique().Field("bin_id")
    if !m.Optional {
        e.Required()
    }
    return []ent.Edge{e}
}

func (m BinMixin) Indexes() (arr []ent.Index) {
    if !m.DisableIndex {
        arr = append(arr, index.Fields("bin_id"))
    }
    return
}

// Bin holds the schema definition for the Bin entity.
type Bin struct {
    ent.Schema
}

// Annotations of the Bin.
func (Bin) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "bin"},
        entsql.WithComments(true),
    }
}

// Fields of the Bin.
func (Bin) Fields() []ent.Field {
    return []ent.Field{
        field.Uint64("cabinet_id"),
        field.String("uuid").Unique().MaxLen(32).Comment("唯一标识"),

        // 电柜信息
        field.String("serial").Comment("电柜设备序列号"),

        // 仓位信息
        field.String("name").Comment("仓位名称(N号仓)"),
        field.Int("ordinal").Comment("仓位序号(从1开始)"),
        field.Bool("open").Default(false).Comment("仓门是否开启"),
        field.Bool("enable").Default(true).Comment("仓位是否启用"),
        field.Bool("health").Default(true).Comment("仓位是否健康"),

        // 电池信息
        field.Bool("battery_exists").Default(false).Comment("是否有电池"),
        field.String("battery_sn").Default("").Comment("电池序列号"),
        field.Float("voltage").Default(0).Comment("当前电压"),
        field.Float("current").Default(0).Comment("当前电流"),
        field.Float("soc").Default(0).Comment("电池电量"),
        field.Float("soh").Default(0).Comment("电池健康程度"),

        // 操作信息
        field.String("remark").Optional().Nillable().Comment("仓位备注"),
    }
}

// Edges of the Bin.
func (Bin) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("cabinet", Cabinet.Type).
            Ref("bins").
            Unique().
            Required().
            Field("cabinet_id"),
    }
}

func (Bin) Mixin() []ent.Mixin {
    return []ent.Mixin{
        internal.TimeMixin{},
    }
}

func (Bin) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("cabinet_id"),
        index.Fields("serial", "ordinal").Unique(),
        index.Fields("battery_exists"),
        index.Fields("ordinal"),
        index.Fields("battery_sn"),
        index.Fields("soc"),
    }
}
