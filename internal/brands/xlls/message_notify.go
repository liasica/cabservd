// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-17
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

type Notifyer interface {
	GetRequestID() string
}

type Notify struct {
	RequestID string `json:"requestId"`
}

func (n Notify) GetRequestID() string {
	return n.RequestID
}

// HardwareOperation 硬件操作结果通知, 西六楼 -> 平台
type HardwareOperation[T any] struct {
	Notify
	Command string `json:"command"`        // 指令
	Result  int    `json:"result"`         // 0:成功 1:失败
	Code    string `json:"code,omitempty"` // 错误编号
	Data    T      `json:"data,omitempty"` // 结果
}

// BusinessNotify 业务类结果通知
type BusinessNotify struct {
	Notify
	Sn             string  `json:"sn"`
	OrderNo        string  `json:"orderNo"`                  //  第三方生成订单号
	OrderStatus    string  `json:"orderStatus"`              //  订单状态
	Confirm        *int    `json:"confirm,omitempty"`        //  1:需要确认。当订单结束时（无论异常结束还是正常结束），此参数值不为空，需要第三方确认订单，否则开放平台会一直发送，直到第三方确认为止。
	OperateStep    float64 `json:"operateStep"`              //  操作步骤，默认从1.0开始，参见附录七
	Status         int     `json:"status"`                   //  操作结果 0:成功 1:失败
	EmptyCellNo    *int    `json:"emptyCellNo,omitempty"`    //  空格挡号
	BatteryCellNo  *int    `json:"batteryCellNo,omitempty"`  //  电池格挡号
	PlaceBatterySn *string `json:"placeBatterySn,omitempty"` //  放入的电池编号
	TakeBatterySn  *string `json:"takeBatterySn,omitempty"`  //  取走的电池编号
	Type           int     `json:"type"`                     //  业务类型 1:租电 2:退电 3:换电
	Msg            *string `json:"msg,omitempty"`            //  一些辅助订单提示信息
	OperateTime    int64   `json:"operateTime"`              //  柜机的操作时间
}

// BinNotify 仓位状态变化通知
type BinNotify struct {
	Notify
	Sn         string `json:"sn"`
	ReportTime int64  `json:"reportTime"`
	BinAttr
}

type CabNotify struct {
	Notify
	CabAttr
	ReportTime int64  `json:"reportTime"`
	OccurTime  *int64 `json:"occurTime,omitempty"` // 离线或在线时间戳
}
