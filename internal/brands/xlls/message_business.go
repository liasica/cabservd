// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-16
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import "errors"

type ExchangeStatus float64

const (
	ExchangeStatusUnknown             ExchangeStatus = 0   // 未知状态
	ExchangeInitOrder                 ExchangeStatus = 1.0 // 初始化订单
	ExchangeInitCheckParamsFail       ExchangeStatus = 1.1 // 设备检查下发的参数失败 (中断流程)
	ExchangeInitCheckBatteryNotExists ExchangeStatus = 1.2 // 指定格挡不存在电池 (中断流程)
	ExchangeInitCheckCellNotEmpty     ExchangeStatus = 1.3 // 指定的格挡不是空仓 (中断流程)
	ExchangePlaceOpenSuccess          ExchangeStatus = 2.0 // 开门成功
	ExchangePlaceOpenFail             ExchangeStatus = 2.1 // 开门失败 (中断流程)
	ExchangePlaceBatteryCheckSuccess  ExchangeStatus = 3.0 // 放入电池检测成功
	ExchangePlaceBatteryCheckFail     ExchangeStatus = 3.1 // 检测电池直接失败 (中断流程) 例如: 放入的电池不是上一次的换电取走的电池 、放入的电池不属于第三方平台、因为开了strictMode模式, 网络不佳, 一直请求不到服务器, 导致检测电池失败
	ExchangePlaceBatteryCheckTimeout  ExchangeStatus = 3.2 // 检测电池超时 (中断流程) 默认两分钟
	ExchangeTakeOpenSuccess           ExchangeStatus = 4.0 // 取电池开门成功
	ExchangeTakeOpenFail              ExchangeStatus = 4.1 // 开门失败 (中断流程)
	ExchangeTakeBatterySuccess        ExchangeStatus = 5.0 // 电池取走成功 (流程结束)
	ExchangeTakeBatteryTimeout        ExchangeStatus = 5.1 // 电池取走超时 (中断流程并且流程结束)
)

func (s ExchangeStatus) String() string {
	switch s {
	default:
		return "未知状态"
	case ExchangeInitOrder:
		return "初始化订单"
	case ExchangeInitCheckParamsFail:
		return "设备检查下发的参数失败"
	case ExchangeInitCheckBatteryNotExists:
		return "指定格挡不存在电池"
	case ExchangeInitCheckCellNotEmpty:
		return "指定的格挡不是空仓"
	case ExchangePlaceOpenSuccess:
		return "开门成功"
	case ExchangePlaceOpenFail:
		return "开门失败"
	case ExchangePlaceBatteryCheckSuccess:
		return "放入电池检测成功"
	case ExchangePlaceBatteryCheckFail:
		return "检测电池直接失败"
	case ExchangePlaceBatteryCheckTimeout:
		return "检测电池超时"
	case ExchangeTakeOpenSuccess:
		return "取电池开门成功"
	case ExchangeTakeOpenFail:
		return "开门失败"
	case ExchangeTakeBatterySuccess:
		return "电池取走成功"
	case ExchangeTakeBatteryTimeout:
		return "电池取走超时"
	}
}

// 判定是否失败并返回失败消息
func (s ExchangeStatus) error() error {
	switch s {
	case ExchangeInitCheckParamsFail, ExchangeInitCheckBatteryNotExists, ExchangeInitCheckCellNotEmpty, ExchangePlaceOpenFail,
		ExchangePlaceBatteryCheckFail, ExchangePlaceBatteryCheckTimeout, ExchangeTakeOpenFail, ExchangeTakeBatteryTimeout:
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
