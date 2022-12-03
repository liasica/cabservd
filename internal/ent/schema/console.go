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
    "github.com/auroraride/cabservd/internal/types"
)

type ConsoleMixin struct {
    mixin.Schema
    Optional     bool
    DisableIndex bool
}

func (m ConsoleMixin) Fields() []ent.Field {
    relate := field.Uint64("console_id")
    if m.Optional {
        relate.Optional().Nillable()
    }
    return []ent.Field{
        relate,
    }
}

func (m ConsoleMixin) Edges() []ent.Edge {
    e := edge.To("console", Console.Type).Unique().Field("console_id")
    if !m.Optional {
        e.Required()
    }
    return []ent.Edge{e}
}

func (m ConsoleMixin) Indexes() (arr []ent.Index) {
    if !m.DisableIndex {
        arr = append(arr, index.Fields("console_id"))
    }
    return
}

// Console holds the schema definition for the Console entity.
type Console struct {
    ent.Schema
}

// Annotations of the Console.
func (Console) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "console"},
    }
}

// Fields of the Console.
func (Console) Fields() []ent.Field {
    return []ent.Field{
        field.Enum("type").Values("exchange", "control").Comment("类别"),

        field.Uint64("user_id").Comment("用户ID"),
        field.Enum("user_type").Values("manager", "rider").Comment("用户类别"),
        field.String("phone").Optional().Nillable().Comment("用户电话"),

        field.Other("step", types.ExchangeStepFirst).SchemaType(map[string]string{dialect.Postgres: postgres.TypeSmallInt}).Optional().Nillable().Comment("换电步骤"),

        field.Enum("status").Values("pending", "running", "success", "failed").Comment("状态"), // pending:未开始 running:执行中 success:成功 failed:失败
        field.JSON("before_bin", &types.BinInfo{}).Optional().Comment("操作前仓位信息"),
        field.JSON("after_bin", &types.BinInfo{}).Optional().Comment("操作后仓位信息"),

        field.String("message").Optional().Nillable().Comment("消息"),
        field.Time("startAt").Comment("开始时间"),
        field.Time("stopAt").Comment("结束时间"),
    }
}

// Edges of the Console.
func (Console) Edges() []ent.Edge {
    return []ent.Edge{}
}

func (Console) Mixin() []ent.Mixin {
    return []ent.Mixin{
        CabinetMixin{},
        BinMixin{},
    }
}

func (Console) Indexes() []ent.Index {
    return []ent.Index{}
}
