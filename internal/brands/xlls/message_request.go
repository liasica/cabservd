// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-16
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap/zapcore"
)

type BusinessAttrRequest struct {
	NeedCellAttr int      `json:"needCellAttr,omitempty"` // 是否同时获取仓位信息 1:获取
	SnList       []string `json:"snList"`
}

// SnListRequest 批量sn请求
type SnListRequest struct {
	SnList []string `json:"snList"`
}

func (a *SnListRequest) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	b, _ := jsoniter.Marshal(a)
	enc.AddByteString("snList", b)
	return nil
}

// CabinetModelCreateRequest 批量创建柜机型号
type CabinetModelCreateRequest struct {
	ModelList []*CabinetModel `json:"modelList"` // 型号的创建所需属性，最大5个
}

type CabinetCreateRequest struct {
	CabinetList []*CabinetMeta `json:"cabinetList"` // 柜机列表
}

// ModelTypeRequest 型号请求
type ModelTypeRequest struct {
	ModelType string `json:"modelType"`
}

// CabinetCommandRequest 「ASYNC」柜机指令请求
// 此接口属于通过柜机控制硬件的接口，所以只能通过轮询或者回调查看响应的操作接口，其他控制硬件接口同理。
// 如果修改了电压和电流相关的属性，app会自动重启（硬件不会重启），从而让设置参数生效。
type CabinetCommandRequest[T any] struct {
	Sn      string         `json:"sn"`
	Command CabinetCommand `json:"command"` // SettingApp(设置app参数: AppAttr), SettingPassword(设置柜机密码: CabinetPassword), RestartCabinet(重启屏幕: CabinetPassword)
	Data    T              `json:"data"`
}

// CellAttrRequest 获取格挡相关属性请求
type CellAttrRequest struct {
	Sn      string `json:"sn"`
	CellNos []int  `json:"cellNos,omitempty"` // 柜机的格挡号，可以为多个，如果传入不存在的，那么最终只会返回存在的格挡，不会报错
	IsAll   int    `json:"isAll,omitempty"`   // 和cellNos必须二选一，如果两个参数都传了，系统默认取isAll,如果这个值不为空(任意的int数值)，那么就会返回所有的格挡相关属性
}

// CellCommandRequest 「ASYNC」发送柜机格挡硬件控制指令
type CellCommandRequest struct {
	Sn      string      `json:"sn"`
	CellNos []int       `json:"cellNos,omitempty"` // 柜机的格挡号，可以为多个，如果传入不存在的，那么最终只会返回存在的格挡，不会报错
	IsAll   int         `json:"isAll,omitempty"`   // 和cellNos必须二选一，如果两个参数都传了，系统默认取isAll,如果这个值不为空(任意的int数值)，那么就会返回所有的格挡相关属性
	Command CellCommand `json:"command"`           // 格挡控制指令
	Data    any         `json:"data,omitempty"`    // 如果是开指示灯，那么必须传入指示灯的颜色
}

// BatteryCreateRequest 批量创建电池请求
type BatteryCreateRequest struct {
	BatteryList             []BatterySnData `json:"batteryList"`                       // 电池的创建所需属性，最大50个
	NeedSyncBatteryPlatform int             `json:"needSyncBatteryPlatform,omitempty"` // 是否同步到电池BMS平台，1:同步，不传或者其他值为不同步
}

// BatteryDeleteRequest 批量删除电池请求
type BatteryDeleteRequest struct {
	BatterySnList           []string `json:"batterySnList"`                     // 电池SN列表
	NeedSyncBatteryPlatform int      `json:"needSyncBatteryPlatform,omitempty"` // 是否同步到电池BMS平台，1:同步，不传或者其他值为不同步
}

// BatteryModifySnRequest 修改电池编号
type BatteryModifySnRequest struct {
	OldBatterySn            string `json:"oldBatterySn"`                      // 旧编号
	NewBatterySn            string `json:"newBatterySn"`                      // 新编号
	NeedSyncBatteryPlatform int    `json:"needSyncBatteryPlatform,omitempty"` // 是否同步到电池BMS平台，1:同步，不传或者其他值为不同步
}

// BatteryAttrRequest 获取电池属性
type BatteryAttrRequest struct {
	BatterySn string `json:"batterySn"`
}
