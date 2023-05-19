// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-16
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

type CabinetCommand string

const (
	CoreOpenFan     CabinetCommand = "CORE_OPEN_FAN"    // 开启核心风扇
	CoreCloseFan    CabinetCommand = "CORE_CLOSE_FAN"   // 关闭核心风扇
	CoreOpenLight   CabinetCommand = "CORE_OPEN_LIGHT"  // 开启核心灯
	CoreCloseLight  CabinetCommand = "CORE_CLOSE_LIGHT" // 关闭核心灯
	QueryPower      CabinetCommand = "QUERY_POWER"      // 查询耗电量
	QueryIccid      CabinetCommand = "QUERY_ICCID"      // 查询sim卡号
	RestartCabinet  CabinetCommand = "RESTART_CABINET"  // 重启柜机（硬件不重启，只是工控机） data: CabinetPassword
	UpdateApp       CabinetCommand = "UPDATE_APP"       // 更新app
	SettingApp      CabinetCommand = "SETTING_APP"      // 设置app的参数 data: AppAttr
	SettingRead     CabinetCommand = "SETTING_READ"     // 读取app的参数
	SettingPassword CabinetCommand = "SETTING_PASSWORD" // 修改柜机密码 data: CabinetPassword
)

type CellCommand string

const (
	CellOpenDoor            CellCommand = "CELL_OPEN_DOOR"             // 开门
	CellOpenIndicatorLight  CellCommand = "CELL_OPEN_INDICATOR_LIGHT"  // 开指示灯,颜色有：1:红色 2:绿色 3:黄色 data: CellOpenIndicatorLightData
	CellCloseIndicatorLight CellCommand = "CELL_CLOSE_INDICATOR_LIGHT" // 关指示灯
	CellOpenCellLight       CellCommand = "CELL_OPEN_CELL_LIGHT"       // 开仓内的指示灯
	CellCloseCellLight      CellCommand = "CELL_CLOSE_CELL_LIGHT"      // 关仓内的指示灯
	CellOpenHeat            CellCommand = "CELL_OPEN_HEAT"             // 开加热
	CellCloseHeat           CellCommand = "CELL_CLOSE_HEAT"            // 关加热
	CellOpenCharge          CellCommand = "CELL_OPEN_CHARGE"           // 打开充电器 data: CellOpenChargeData
	CellCloseCharge         CellCommand = "CELL_CLOSE_CHARGE"          // 关闭充电器
	CellForbid              CellCommand = "CELL_FORBID"                // 禁用格挡
	CellUnForbid            CellCommand = "CELL_UN_FORBID"             // 开启格挡
	CellOpenFan             CellCommand = "CELL_OPEN_FAN"              // 打开格挡风扇
	CellCloseFan            CellCommand = "CELL_CLOSE_FAN"             // 关闭格挡风扇
)
