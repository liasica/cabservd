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
    }
}

// Fields of the Cabinet.
func (Cabinet) Fields() []ent.Field {
    return []ent.Field{
        field.String("sn").Unique().Comment("电柜编号"),
        field.Uint("status").Default(0).Comment("状态"),
    }
}

// Edges of the Cabinet.
func (Cabinet) Edges() []ent.Edge {
    return []ent.Edge{}
}

func (Cabinet) Mixin() []ent.Mixin {
    return []ent.Mixin{
        internal.TimeMixin{},
    }
}

func (Cabinet) Indexes() []ent.Index {
    return []ent.Index{}
}
