// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-04
// Based on cabservd by liasica, magicrolan@qq.com.

package tower

import (
    "context"
    "github.com/auroraride/adapter"
    "github.com/auroraride/adapter/log"
    "github.com/auroraride/cabservd/internal/core"
    jsoniter "github.com/json-iterator/go"
    "go.uber.org/zap"
)

var (
    messageType    *MessageTypeList
    cabinetSignals map[Signal]CabinetSignalFunc
    binSignals     map[Signal]BinSignalFunc
)

type Handler struct {
    core.Bean
}

func New(mt *MessageTypeList, params ...any) *Handler {
    cabinetSignals = make(map[Signal]CabinetSignalFunc)
    binSignals = make(map[Signal]BinSignalFunc)

    for _, param := range params {
        switch m := param.(type) {
        case map[Signal]CabinetSignalFunc:
            for k, v := range m {
                CabinetSignalMap[k] = struct{}{}
                cabinetSignals[k] = v
            }
        case map[Signal]BinSignalFunc:
            for k, v := range m {
                binSignals[k] = v
            }
        }
    }

    messageType = mt

    return &Handler{}
}

// GetEmptyDeviation TODO 后续做在数据库中
func (h *Handler) GetEmptyDeviation() (voltage, current float64) {
    voltage = 40
    current = 1
    return
}

// OnMessage 解析消息
func (h *Handler) OnMessage(_ *core.Client, b []byte) (serial string, res core.ResponseMessenger, fields []zap.Field, err error) {
    fields = []zap.Field{
        zap.ByteString("decoded", b),
    }

    req := new(Request)
    err = jsoniter.Unmarshal(b, req)
    if err != nil {
        return
    }

    serial = req.DevID
    fields = append(fields, log.Payload(req))

    switch req.MsgType {
    case messageType.LoginRequest:
        err = h.LoginHandle(req)
    case messageType.ReportRequest:
        err = h.ReportHandle(req)
    case messageType.NoticeRequest:
        err = h.NoticeHandle(req)
    case messageType.ControlResponse:
        // TODO 控制成功逻辑
        // 收到成功逻辑处理完成后, 不发送反馈消息
        return
    }

    // 发送失败响应
    if err != nil {
        res = req.Fail()
        return
    }

    res = req.Success()

    return
}

// LoginHandle 登录请求
func (h *Handler) LoginHandle(req *Request) (err error) {
    if req.DevID == "" {
        return adapter.ErrorCabinetSerialRequired
    }

    // // 清除仓位电池信息
    // // TODO 清除的时候会不会后来的消息先到
    // err = core.ResetBins(req.DevID)
    // if err != nil {
    //     return
    // }

    // 查找或创建电柜
    go core.LoadOrStoreCabinet(context.Background(), req.DevID)

    // TODO: 保存其他信息
    return
}

// ReportHandle 状态上报请求
func (h *Handler) ReportHandle(req *Request) (err error) {
    if req.DevID == "" {
        return adapter.ErrorCabinetSerialRequired
    }
    core.UpdateCabinet(req)
    return
}

// NoticeHandle 告警上报请求
func (h *Handler) NoticeHandle(req *Request) (err error) {
    // TODO 解读并保存所有告警信息
    return
}
