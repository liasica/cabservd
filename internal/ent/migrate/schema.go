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
		{Name: "serial", Type: field.TypeString, Comment: "电柜设备序列号"},
		{Name: "name", Type: field.TypeString, Comment: "仓位名称(N号仓)"},
		{Name: "ordinal", Type: field.TypeInt, Comment: "仓位序号(从1开始)"},
		{Name: "open", Type: field.TypeBool, Comment: "仓门是否开启", Default: false},
		{Name: "enable", Type: field.TypeBool, Comment: "仓位是否启用", Default: true},
		{Name: "health", Type: field.TypeBool, Comment: "仓位是否健康", Default: true},
		{Name: "battery_exists", Type: field.TypeBool, Comment: "是否有电池", Default: false},
		{Name: "battery_sn", Type: field.TypeString, Comment: "电池序列号", Default: ""},
		{Name: "voltage", Type: field.TypeFloat64, Comment: "当前电压", Default: 0},
		{Name: "current", Type: field.TypeFloat64, Comment: "当前电流", Default: 0},
		{Name: "soc", Type: field.TypeFloat64, Comment: "电池电量", Default: 0},
		{Name: "soh", Type: field.TypeFloat64, Comment: "电池健康程度", Default: 0},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "仓位备注"},
		{Name: "deactivate", Type: field.TypeBool, Comment: "是否停用", Default: false},
		{Name: "deactivate_reason", Type: field.TypeString, Nullable: true, Comment: "停用信息"},
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
				Name:    "bin_serial_ordinal",
				Unique:  true,
				Columns: []*schema.Column{BinColumns[3], BinColumns[5]},
			},
			{
				Name:    "bin_battery_exists",
				Unique:  false,
				Columns: []*schema.Column{BinColumns[9]},
			},
			{
				Name:    "bin_battery_sn",
				Unique:  false,
				Columns: []*schema.Column{BinColumns[10]},
			},
			{
				Name:    "bin_soc",
				Unique:  false,
				Columns: []*schema.Column{BinColumns[13]},
			},
		},
	}
	// CabinetColumns holds the columns for the "cabinet" table.
	CabinetColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "online", Type: field.TypeBool, Comment: "是否在线", Default: false},
		{Name: "power", Type: field.TypeBool, Comment: "市电是否正常", Default: true},
		{Name: "serial", Type: field.TypeString, Unique: true, Comment: "电柜编号"},
		{Name: "status", Type: field.TypeEnum, Comment: "状态", Enums: []string{"initializing", "normal", "abnormal"}, Default: "initializing"},
		{Name: "enable", Type: field.TypeBool, Comment: "电柜是否启用", Default: false},
		{Name: "lng", Type: field.TypeFloat64, Nullable: true, Comment: "经度"},
		{Name: "lat", Type: field.TypeFloat64, Nullable: true, Comment: "纬度"},
		{Name: "gsm", Type: field.TypeFloat64, Nullable: true, Comment: "GSM信号强度"},
		{Name: "voltage", Type: field.TypeFloat64, Nullable: true, Comment: "换电柜总电压 (V)"},
		{Name: "current", Type: field.TypeFloat64, Nullable: true, Comment: "换电柜总电流 (A)"},
		{Name: "temperature", Type: field.TypeFloat64, Nullable: true, Comment: "柜体温度值 (换电柜温度)"},
		{Name: "electricity", Type: field.TypeFloat64, Nullable: true, Comment: "总用电量"},
		{Name: "sim", Type: field.TypeString, Nullable: true, Comment: "SIM卡号"},
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
		{Name: "operate", Type: field.TypeOther, Comment: "操作", SchemaType: map[string]string{"postgres": "varchar"}},
		{Name: "serial", Type: field.TypeString, Comment: "电柜设备序列号"},
		{Name: "uuid", Type: field.TypeUUID, Comment: "标识符"},
		{Name: "business", Type: field.TypeEnum, Comment: "业务 operate:运维操作 exchange:换电 active:激活 pause:寄存 continue:结束寄存 unsubscribe:退订", Enums: []string{"operate", "exchange", "active", "pause", "continue", "unsubscribe"}},
		{Name: "user_id", Type: field.TypeString, Comment: "用户ID"},
		{Name: "user_type", Type: field.TypeOther, Comment: "用户类别", SchemaType: map[string]string{"postgres": "varchar"}},
		{Name: "step", Type: field.TypeInt, Comment: "步骤", Default: 1},
		{Name: "status", Type: field.TypeEnum, Comment: "状态 invalid:无效 pending:未开始 running:执行中 success:成功 failed:失败", Enums: []string{"invalid", "pending", "running", "success", "failed"}},
		{Name: "before_bin", Type: field.TypeJSON, Nullable: true, Comment: "变化前仓位信息"},
		{Name: "after_bin", Type: field.TypeJSON, Nullable: true, Comment: "变化后仓位信息"},
		{Name: "message", Type: field.TypeString, Nullable: true, Comment: "消息"},
		{Name: "start_at", Type: field.TypeTime, Nullable: true, Comment: "开始时间"},
		{Name: "stop_at", Type: field.TypeTime, Nullable: true, Comment: "结束时间"},
		{Name: "duration", Type: field.TypeFloat64, Nullable: true, Comment: "耗时"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "备注信息"},
		{Name: "command_retry_times", Type: field.TypeInt, Comment: "指令重试次数", Default: 1},
		{Name: "order_sn", Type: field.TypeString, Nullable: true, Comment: "第三方平台订单号"},
		{Name: "cabinet_id", Type: field.TypeUint64},
		{Name: "bin_id", Type: field.TypeUint64, Nullable: true},
	}
	// ConsoleTable holds the schema information for the "console" table.
	ConsoleTable = &schema.Table{
		Name:       "console",
		Columns:    ConsoleColumns,
		PrimaryKey: []*schema.Column{ConsoleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "console_cabinet_cabinet",
				Columns:    []*schema.Column{ConsoleColumns[18]},
				RefColumns: []*schema.Column{CabinetColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "console_bin_bin",
				Columns:    []*schema.Column{ConsoleColumns[19]},
				RefColumns: []*schema.Column{BinColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "console_cabinet_id",
				Unique:  false,
				Columns: []*schema.Column{ConsoleColumns[18]},
			},
			{
				Name:    "console_bin_id",
				Unique:  false,
				Columns: []*schema.Column{ConsoleColumns[19]},
			},
			{
				Name:    "console_serial",
				Unique:  false,
				Columns: []*schema.Column{ConsoleColumns[2]},
			},
			{
				Name:    "console_uuid",
				Unique:  false,
				Columns: []*schema.Column{ConsoleColumns[3]},
			},
			{
				Name:    "console_user_id",
				Unique:  false,
				Columns: []*schema.Column{ConsoleColumns[5]},
			},
			{
				Name:    "console_user_type",
				Unique:  false,
				Columns: []*schema.Column{ConsoleColumns[6]},
			},
			{
				Name:    "console_start_at",
				Unique:  false,
				Columns: []*schema.Column{ConsoleColumns[12]},
			},
			{
				Name:    "console_stop_at",
				Unique:  false,
				Columns: []*schema.Column{ConsoleColumns[13]},
			},
			{
				Name:    "console_order_sn",
				Unique:  false,
				Columns: []*schema.Column{ConsoleColumns[17]},
			},
		},
	}
	// ScanColumns holds the columns for the "scan" table.
	ScanColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "order_no", Type: field.TypeString, Nullable: true, Comment: "订单编号"},
		{Name: "business", Type: field.TypeEnum, Comment: "业务 operate:运维操作 exchange:换电 active:激活 pause:寄存 continue:结束寄存 unsubscribe:退订", Enums: []string{"operate", "exchange", "active", "pause", "continue", "unsubscribe"}},
		{Name: "efficient", Type: field.TypeBool, Comment: "是否有效", Default: true},
		{Name: "user_id", Type: field.TypeString, Comment: "用户ID"},
		{Name: "user_type", Type: field.TypeOther, Comment: "用户类别", SchemaType: map[string]string{"postgres": "varchar"}},
		{Name: "serial", Type: field.TypeString, Comment: "电柜编号"},
		{Name: "data", Type: field.TypeJSON, Nullable: true, Comment: "换电信息"},
		{Name: "cabinet_id", Type: field.TypeUint64},
	}
	// ScanTable holds the schema information for the "scan" table.
	ScanTable = &schema.Table{
		Name:       "scan",
		Columns:    ScanColumns,
		PrimaryKey: []*schema.Column{ScanColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "scan_cabinet_cabinet",
				Columns:    []*schema.Column{ScanColumns[11]},
				RefColumns: []*schema.Column{CabinetColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "scan_created_at",
				Unique:  false,
				Columns: []*schema.Column{ScanColumns[1]},
			},
			{
				Name:    "scan_cabinet_id",
				Unique:  false,
				Columns: []*schema.Column{ScanColumns[11]},
			},
			{
				Name:    "scan_user_id",
				Unique:  false,
				Columns: []*schema.Column{ScanColumns[7]},
			},
			{
				Name:    "scan_user_type",
				Unique:  false,
				Columns: []*schema.Column{ScanColumns[8]},
			},
			{
				Name:    "scan_serial",
				Unique:  false,
				Columns: []*schema.Column{ScanColumns[9]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BinTable,
		CabinetTable,
		ConsoleTable,
		ScanTable,
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
	ScanTable.ForeignKeys[0].RefTable = CabinetTable
	ScanTable.Annotation = &entsql.Annotation{
		Table: "scan",
	}
}
