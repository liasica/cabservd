// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-16
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import "errors"

type ExchangeStatus float64

const (
	ExchangeStatusUnknown            ExchangeStatus = 0   // 未知状态
	ExchangeStatusInit               ExchangeStatus = 1.0 // SELF_INIT_ORDER                     初始化订单
	ExchangeStatusInitFail           ExchangeStatus = 1.1 // SELF_INIT_CHECK_PARAMS_FAIL         设备检查下发的参数失败
	ExchangeStatusNoBattery          ExchangeStatus = 1.2 // SELF_INIT_CHECK_BATTERY_NOT_EXISTS  指定格挡不存在电池
	ExchangeStatusBatteryTypeMismach ExchangeStatus = 1.3 // SELF_BATTERY_TYPE_MISMATCH          电池名称不匹配
	ExchangeStatusDoorOpenSuccess    ExchangeStatus = 2.0 // SELF_OPEN_DOOR_SUCCESS              开门成功
	ExchangeStatusDoorOpenFail       ExchangeStatus = 2.1 // SELF_OPEN_DOOR_ERROR                开门失败 (最终步骤, 中断流程)
	ExchangeStatusBatteryTimeout     ExchangeStatus = 3.0 // SELF_TAKE_BATTERY_TIMEOUT           电池超时未取走
	ExchangeStatusBatteryTook        ExchangeStatus = 3.1 // SELF_TAKE_BATTERY_SUCCESS           电池取走成功
	ExchangeStatusCloseTimeout       ExchangeStatus = 4.0 // SELF_CLOSE_DOOR_TIMEOUT             关门超时 (最终步骤)
	ExchangeStatusCloseSuccess       ExchangeStatus = 4.1 // SELF_CLOSE_DOOR_SUCCESS             关门成功 (最终步骤)
)

func (s ExchangeStatus) String() string {
	switch s {
	default:
		return "未知状态"
	case ExchangeStatusInit:
		return "初始化订单"
	case ExchangeStatusInitFail:
		return "设备检查下发的参数失败"
	case ExchangeStatusNoBattery:
		return "指定格挡不存在电池"
	case ExchangeStatusBatteryTypeMismach:
		return "电池名称不匹配"
	case ExchangeStatusDoorOpenSuccess:
		return "开门成功"
	case ExchangeStatusDoorOpenFail:
		return "开门失败"
	case ExchangeStatusBatteryTimeout:
		return "电池超时未取走"
	case ExchangeStatusBatteryTook:
		return "电池取走成功"
	case ExchangeStatusCloseTimeout:
		return "关门超时"
	case ExchangeStatusCloseSuccess:
		return "关门成功"
	}
}

// 判定是否失败并返回失败消息
func (s ExchangeStatus) error() error {
	switch s {
	case ExchangeStatusInit, ExchangeStatusDoorOpenSuccess, ExchangeStatusBatteryTook, ExchangeStatusCloseSuccess:
		return nil
	}
	return errors.New(s.String())
}

type BusinessExchangeRequest struct {
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

// BusinessResponse 业务请求返回
type BusinessResponse struct {
	OrderNo string  `json:"orderNo"` // 第三方订单号
	Status  float64 `json:"status"`  // 订单状态
}
