// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-16
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"errors"

	"github.com/auroraride/cabservd/internal/types"
)

type bizStatus interface {
	last() bool
	step() int
	binStep() *types.BinStep
	error() error
}

type exchangeStatus float64

const (
	exchangeStatusUnknown             exchangeStatus = 0   // 未知状态 (中断流程)
	exchangeInitOrder                 exchangeStatus = 1.0 // 初始化订单
	exchangeInitCheckParamsFail       exchangeStatus = 1.1 // 设备检查下发的参数失败 (中断流程)
	exchangeInitCheckBatteryNotExists exchangeStatus = 1.2 // 指定格挡不存在电池 (中断流程)
	exchangeInitCheckCellNotEmpty     exchangeStatus = 1.3 // 指定的格挡不是空仓 (中断流程)
	exchangePlaceOpenSuccess          exchangeStatus = 2.0 // 开门成功
	exchangePlaceOpenFail             exchangeStatus = 2.1 // 开门失败 (中断流程)
	exchangePlaceBatteryCheckSuccess  exchangeStatus = 3.0 // 放入电池检测成功
	exchangePlaceBatteryCheckFail     exchangeStatus = 3.1 // 检测电池直接失败 (中断流程) 例如: 放入的电池不是上一次的换电取走的电池 、放入的电池不属于第三方平台、因为开了strictMode模式, 网络不佳, 一直请求不到服务器, 导致检测电池失败
	exchangePlaceBatteryCheckTimeout  exchangeStatus = 3.2 // 检测电池超时 (中断流程) 默认两分钟
	exchangeTakeOpenSuccess           exchangeStatus = 4.0 // 取电池开门成功
	exchangeTakeOpenFail              exchangeStatus = 4.1 // 开门失败 (中断流程)
	exchangeTakeBatterySuccess        exchangeStatus = 5.0 // 电池取走成功 (换电成功, 流程结束)
	exchangeTakeBatteryTimeout        exchangeStatus = 5.1 // 电池取走超时 (中断流程并且流程结束)
)

func (s exchangeStatus) last() bool {
	return s.step() == 4
}

func (s exchangeStatus) step() int {
	return int(s - 1)
}

func (s exchangeStatus) binStep() *types.BinStep {
	return types.ExchangeConfigure[1-(4-s.step())/2][1-(4-s.step())%2]
}

// 判定是否失败并返回失败消息
func (s exchangeStatus) error() error {
	switch s {
	case exchangeInitOrder, exchangePlaceOpenSuccess, exchangePlaceBatteryCheckSuccess, exchangeTakeOpenSuccess, exchangeTakeBatterySuccess:
		return nil
	}
	return errors.New(s.String())
}

func (s exchangeStatus) String() string {
	switch s {
	default:
		return "未知状态"
	case exchangeInitOrder:
		return "初始化订单"
	case exchangeInitCheckParamsFail:
		return "设备检查下发的参数失败"
	case exchangeInitCheckBatteryNotExists:
		return "指定格挡不存在电池"
	case exchangeInitCheckCellNotEmpty:
		return "指定的格挡不是空仓"
	case exchangePlaceOpenSuccess:
		return "开门成功"
	case exchangePlaceOpenFail:
		return "开门失败"
	case exchangePlaceBatteryCheckSuccess:
		return "放入电池检测成功"
	case exchangePlaceBatteryCheckFail:
		return "检测电池直接失败"
	case exchangePlaceBatteryCheckTimeout:
		return "检测电池超时"
	case exchangeTakeOpenSuccess:
		return "取电池开门成功"
	case exchangeTakeOpenFail:
		return "开门失败"
	case exchangeTakeBatterySuccess:
		return "电池取走成功"
	case exchangeTakeBatteryTimeout:
		return "电池取走超时"
	}
}

type bizExchangeRequest struct {
	Sn            string `json:"sn"`
	OrderNo       string `json:"orderNo"`             // 第三方订单号, 第三方必须保证唯一性, 如果有多个重复订单号, 那么查询订单时, 默认返回第一个。
	EmptyCellNo   int    `json:"emptyCellNo"`         // 空仓格挡号
	BatteryCellNo int    `json:"batteryCellNo"`       // 满电电池格挡号
	ModelType     string `json:"modelType,omitempty"` // 多型号的名称
	// 智能换仓模式 0--不开启 1--开启
	// 开启此模式, 如果柜机发现下发的仓门不满足条件, 就会换成另一个满足条件的仓门
	// 比如: 下发的满电电池1号门, 结果柜机检测到1号门没有电池, 则会查找另一个满电电池仓门, 并开门。
	SmartMode int `json:"smartMode,omitempty"`
	// 严格模式, 0--不开启, 1--开启
	// 与bindingBatterySn必须二选一如果都传, 默认使用bindingBatterySn
	// 作用是用来控制当网络不佳时, 柜机访问开放平台检测电池是否属于第三方平台电池的http请求超时失败, 是否还继续订单流程。如果开启, 则直接报错订单结束。
	StrictMode int `json:"strictMode,omitempty"`
	// 用户绑定的电池编号, 如果此参数不为空, 那么柜机就会检测用户放入的电池sn是否跟此参数的值一致
	BindingBatterySn string `json:"bindingBatterySn,omitempty"`
}

// 业务请求返回
type bizResponse struct {
	OrderNo string  `json:"orderNo"` // 第三方订单号
	Status  float64 `json:"status"`  // 订单状态
}

type bizQueryRequest struct {
	OrderNo     string `json:"orderNo"`               // 第三方订单号
	NeedRecords int    `json:"needRecords,omitempty"` // 是否返回订单的所有操作记录
}

type bizOperateRecord struct {
	Status     string  `json:"status"`
	Seq        float64 `json:"seq"`
	Msg        string  `json:"msg"`
	CreateTime int64   `json:"createTime"`
}

type bizQueryResult struct {
	OrderNo         string             `json:"orderNo,omitempty"`
	Status          string             `json:"status,omitempty"`
	EndTime         int64              `json:"endTime,omitempty"`
	ExchangeEndTime int64              `json:"exchangeEndTime,omitempty"`
	EmptyCellNo     int                `json:"emptyCellNo,omitempty"`
	BatteryCellNo   int                `json:"batteryCellNo,omitempty"`
	PlaceBattery    string             `json:"placeBattery,omitempty"`
	TakeBattery     string             `json:"takeBattery,omitempty"`
	ConfirmTime     int64              `json:"confirmTime,omitempty"`
	OperRecords     []bizOperateRecord `json:"operRecords,omitempty"`
}
