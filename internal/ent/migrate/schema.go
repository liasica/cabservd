// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BinColumns holds the columns for the "bin" table.
	BinColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "uuid", Type: field.TypeString, Unique: true, Size: 32},
		{Name: "brand", Type: field.TypeString},
		{Name: "serial", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "ordinal", Type: field.TypeInt},
		{Name: "open", Type: field.TypeBool, Default: false},
		{Name: "enable", Type: field.TypeBool, Default: true},
		{Name: "health", Type: field.TypeBool, Default: true},
		{Name: "battery_exists", Type: field.TypeBool, Default: false},
		{Name: "battery_sn", Type: field.TypeString, Default: ""},
		{Name: "voltage", Type: field.TypeFloat64, Default: 0},
		{Name: "current", Type: field.TypeFloat64, Default: 0},
		{Name: "soc", Type: field.TypeFloat64, Default: 0},
		{Name: "soh", Type: field.TypeFloat64, Default: 0},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "cabinet_id", Type: field.TypeUint64},
	}
	// BinTable holds the schema information for the "bin" table.
	BinTable = &schema.Table{
		Name:       "bin",
		Columns:    BinColumns,
		PrimaryKey: []*schema.Column{BinColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bin_cabinet_bins",
				Columns:    []*schema.Column{BinColumns[18]},
				RefColumns: []*schema.Column{CabinetColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "bin_created_at",
				Unique:  false,
				Columns: []*schema.Column{BinColumns[1]},
			},
			{
				Name:    "bin_cabinet_id",
				Unique:  false,
				Columns: []*schema.Column{BinColumns[18]},
			},
			{
				Name:    "bin_serial_brand",
				Unique:  false,
				Columns: []*schema.Column{BinColumns[5], BinColumns[4]},
			},
			{
				Name:    "bin_battery_exists",
				Unique:  false,
				Columns: []*schema.Column{BinColumns[11]},
			},
			{
				Name:    "bin_ordinal",
				Unique:  false,
				Columns: []*schema.Column{BinColumns[7]},
			},
			{
				Name:    "bin_battery_sn",
				Unique:  false,
				Columns: []*schema.Column{BinColumns[12]},
			},
			{
				Name:    "bin_soc",
				Unique:  false,
				Columns: []*schema.Column{BinColumns[15]},
			},
		},
	}
	// CabinetColumns holds the columns for the "cabinet" table.
	CabinetColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "online", Type: field.TypeBool, Default: false},
		{Name: "brand", Type: field.TypeString},
		{Name: "serial", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"initializing", "idle", "busy", "exchange", "abnormal"}, Default: "initializing"},
		{Name: "enable", Type: field.TypeBool, Default: false},
		{Name: "lng", Type: field.TypeFloat64, Nullable: true},
		{Name: "lat", Type: field.TypeFloat64, Nullable: true},
		{Name: "gsm", Type: field.TypeFloat64, Nullable: true},
		{Name: "voltage", Type: field.TypeFloat64, Nullable: true},
		{Name: "current", Type: field.TypeFloat64, Nullable: true},
		{Name: "temperature", Type: field.TypeFloat64, Nullable: true},
		{Name: "electricity", Type: field.TypeFloat64, Nullable: true},
	}
	// CabinetTable holds the schema information for the "cabinet" table.
	CabinetTable = &schema.Table{
		Name:       "cabinet",
		Columns:    CabinetColumns,
		PrimaryKey: []*schema.Column{CabinetColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "cabinet_created_at",
				Unique:  false,
				Columns: []*schema.Column{CabinetColumns[1]},
			},
			{
				Name:    "cabinet_brand",
				Unique:  false,
				Columns: []*schema.Column{CabinetColumns[4]},
			},
			{
				Name:    "cabinet_status",
				Unique:  false,
				Columns: []*schema.Column{CabinetColumns[6]},
			},
			{
				Name:    "cabinet_enable",
				Unique:  false,
				Columns: []*schema.Column{CabinetColumns[7]},
			},
			{
				Name:    "cabinet_lng",
				Unique:  false,
				Columns: []*schema.Column{CabinetColumns[8]},
			},
			{
				Name:    "cabinet_lat",
				Unique:  false,
				Columns: []*schema.Column{CabinetColumns[9]},
			},
		},
	}
	// ConsoleColumns holds the columns for the "console" table.
	ConsoleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"exchange", "control", "cabinet"}},
		{Name: "user", Type: field.TypeJSON},
		{Name: "step", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"postgres": "smallint"}},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "running", "success", "failed"}},
		{Name: "before_bin", Type: field.TypeJSON, Nullable: true},
		{Name: "after_bin", Type: field.TypeJSON, Nullable: true},
		{Name: "message", Type: field.TypeString, Nullable: true},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "stop_at", Type: field.TypeTime, Nullable: true},
		{Name: "cabinet_id", Type: field.TypeUint64},
		{Name: "bin_id", Type: field.TypeUint64},
	}
	// ConsoleTable holds the schema information for the "console" table.
	ConsoleTable = &schema.Table{
		Name:       "console",
		Columns:    ConsoleColumns,
		PrimaryKey: []*schema.Column{ConsoleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "console_cabinet_cabinet",
				Columns:    []*schema.Column{ConsoleColumns[11]},
				RefColumns: []*schema.Column{CabinetColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "console_bin_bin",
				Columns:    []*schema.Column{ConsoleColumns[12]},
				RefColumns: []*schema.Column{BinColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "console_cabinet_id",
				Unique:  false,
				Columns: []*schema.Column{ConsoleColumns[11]},
			},
			{
				Name:    "console_bin_id",
				Unique:  false,
				Columns: []*schema.Column{ConsoleColumns[12]},
			},
			{
				Name:    "console_uuid",
				Unique:  false,
				Columns: []*schema.Column{ConsoleColumns[1]},
			},
			{
				Name:    "console_user",
				Unique:  false,
				Columns: []*schema.Column{ConsoleColumns[3]},
				Annotation: &entsql.IndexAnnotation{
					Types: map[string]string{
						"postgres": "GIN",
					},
				},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BinTable,
		CabinetTable,
		ConsoleTable,
	}
)

func init() {
	BinTable.ForeignKeys[0].RefTable = CabinetTable
	BinTable.Annotation = &entsql.Annotation{
		Table: "bin",
	}
	CabinetTable.Annotation = &entsql.Annotation{
		Table: "cabinet",
	}
	ConsoleTable.ForeignKeys[0].RefTable = CabinetTable
	ConsoleTable.ForeignKeys[1].RefTable = BinTable
	ConsoleTable.Annotation = &entsql.Annotation{
		Table: "console",
	}
}
