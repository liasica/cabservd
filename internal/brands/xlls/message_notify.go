// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-17
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

type Notifyer interface {
	GetRequestID() string
}

type Notify struct {
	RequestID string `json:"requestID"`
}

func (n Notify) GetRequestID() string {
	return n.RequestID
}

// CommandNotify 硬件操作结果通知, 西六楼 -> 平台
type CommandNotify[T any] struct {
	Notify
	Command string `json:"command"`        // 指令
	Result  int    `json:"result"`         // 0--成功 1--失败
	Code    string `json:"code,omitempty"` // 错误编号
	Data    T      `json:"data,omitempty"` // 结果
}

// BusinessNotify 业务类结果通知
type BusinessNotify struct {
	Notify
	Sn             string  `json:"sn"`
	OrderNo        string  `json:"orderNo"`                  //  第三方生成订单号
	OrderStatus    string  `json:"orderStatus"`              //  订单状态
	Confirm        *int    `json:"confirm,omitempty"`        //  1--需要确认。当订单结束时（无论异常结束还是正常结束），此参数值不为空，需要第三方确认订单，否则开放平台会一直发送，直到第三方确认为止。
	OperateStep    float64 `json:"operateStep"`              //  操作步骤，默认从1.0开始，参见附录七
	Status         int     `json:"status"`                   //  操作结果 0--成功 1--失败
	EmptyCellNo    *int    `json:"emptyCellNo,omitempty"`    //  空格挡号
	BatteryCellNo  *int    `json:"batteryCellNo,omitempty"`  //  电池格挡号
	PlaceBatterySn *string `json:"placeBatterySn,omitempty"` //  放入的电池编号
	TakeBatterySn  *string `json:"takeBatterySn,omitempty"`  //  取走的电池编号
	Type           int     `json:"type"`                     //  业务类型 1--租电 2--退电 3--换电
	Msg            *string `json:"msg,omitempty"`            //  一些辅助订单提示信息
	OperateTime    int64   `json:"operateTime"`              //  柜机的操作时间
}

// CellNotify 格挡状态变化通知
type CellNotify struct {
	Notify
	Sn         string `json:"sn"`
	ReportTime int64  `json:"reportTime"`
	CellAttr
}

type BatteryChargeNotify struct {
	Notify
	Sn            string   `json:"sn"`
	ReportTime    int64    `json:"reportTime"`
	ChargeStatus  *int     // 充电器的状态：0--关闭 1--开机中 2--充电中 3--充满电 4--限制充电 -128--过压充电   64--过流充电   32--短路  16--温度过高.   10--超压 11--电池反接 12--NTC故障停机 13--输出短路停机
	ChargeV       *float64 // 充电器的电压
	ChargeA       *float64 // 充电器的电流
	Soc           *int     // 电池的电量，0-100
	CellNo        *int     // 格挡号
	ExistsBattery *int     // 0--不存在 ,1--存在 是否真实存在电池,（保留）
	BatterySn     *string  // 电池名称，如果名称为空，代表舱内没有电池
	BatteryA      *float64 // 电池的电流
	BatteryV      *float64 // 电池的电压
	CoreNum       *int     // 电芯数量
	CoreV         *string  // 每个电芯的电压，根据电芯数量的不同，列表的数量也不同，默认从电芯1开始。 TODO: json的List<Double> 字段啥意思??
	BatteryHealth *string  // 电池的健康状态信息，详见附录六| TODO: json的List<String> 字段啥意思??
	Capacity      *float64 // 电池容量
	EnvTemp       *float64 // 电池环境温度
	CoreTemp      *float64 // 电池电芯温度
	BoardTemp     *float64 // 电池板卡温度
	ModelType     *string  // 电池型号名称，如果是多型号，就会返回
}

type CabinetNotify struct {
	Notify
	PhysicsAttr
	Online    *int   `json:"online,omitempty"`    // 柜机是否在线 0--在线 1--不在线
	OccurTime *int64 `json:"occurTime,omitempty"` // 离线或在线时间戳
}
