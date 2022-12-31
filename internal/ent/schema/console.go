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
    "github.com/google/uuid"
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
        entsql.WithComments(true),
    }
}

// Fields of the Console.
func (Console) Fields() []ent.Field {
    return []ent.Field{
        field.Other("operate", adapter.OperateUnknown).SchemaType(map[string]string{dialect.Postgres: postgres.TypeVarChar}).Comment("操作"),

        field.String("serial").Comment("电柜设备序列号"),
        field.UUID("uuid", uuid.UUID{}).Immutable().Comment("标识符"),
        field.Enum("type").Values("exchange", "operate", "cabinet").Comment("日志类别 exchange:换电控制 operate:手动操作 cabinet:电柜日志"),

        field.String("user_id").Comment("用户ID"),
        field.Other("user_type", adapter.UserTypeUnknown).SchemaType(map[string]string{dialect.Postgres: postgres.TypeVarChar}).Comment("用户类别"),

        field.Other("step", adapter.ExchangeStepFirst).SchemaType(map[string]string{dialect.Postgres: postgres.TypeSmallInt}).Optional().Nillable().Comment("换电步骤"),

        field.Enum("status").Values("invalid", "pending", "running", "success", "failed").Comment("状态 invalid:无效 pending:未开始 running:执行中 success:成功 failed:失败"),
        field.JSON("before_bin", &adapter.BinInfo{}).Optional().Comment("变化前仓位信息"),
        field.JSON("after_bin", &adapter.BinInfo{}).Optional().Comment("变化后仓位信息"),

        field.String("message").Optional().Nillable().Comment("消息"),
        field.Time("startAt").Optional().Nillable().Comment("开始时间"),
        field.Time("stopAt").Optional().Nillable().Comment("结束时间"),
        field.Float("duration").Optional().Nillable().Comment("耗时"),
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
    return []ent.Index{
        index.Fields("serial"),
        index.Fields("uuid"),
        index.Fields("user_id"),
        index.Fields("user_type"),
        index.Fields("startAt"),
        index.Fields("stopAt"),
    }
}
