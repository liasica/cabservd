// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-16
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

type Business struct {
	RequestID      string  `json:"requestID"`
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
