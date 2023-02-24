// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-16
// Based on yundong by liasica, magicrolan@qq.com.

package yundong

import (
    "fmt"
    "github.com/auroraride/adapter"
    jsoniter "github.com/json-iterator/go"
    "go.uber.org/zap"
    "math"
)

const (
    defaultPassthroughSN byte = 0
)

const (
    PassthroughCommandChangeServer    PassthroughCommand = 11   // 修改服务器地址
    PassthroughCommandExchange        PassthroughCommand = 68   // 换电
    PassthroughCommandOpenDoor        PassthroughCommand = 69   // 开仓
    PassthroughCommandCharge          PassthroughCommand = 70   // 充电开关
    PassthroughCommandUPdateAPK       PassthroughCommand = 72   // 升级APK
    PassthroughCommandExchangeCancel  PassthroughCommand = 73   // 取消换电
    PassthroughCommandChargeAvailable PassthroughCommand = 74   // 充电是否启用
    PassthroughCommandBinUsable       PassthroughCommand = 75   // 启用或禁用仓位
    PassthroughCommandVDiffAlaram     PassthroughCommand = 78   // 电池压差告警阈值
    PassthroughCommandAmmeter         PassthroughCommand = 79   // 电表读数
    PassthroughCommandBatteryUsable   PassthroughCommand = 80   // 可用电池数量
    PassthroughCommandVersion         PassthroughCommand = 81   // 查询固件版本
    PassthroughCommandUpdateFW        PassthroughCommand = 82   // 升级电池柜固件
    PassthroughCommandGatingInfo      PassthroughCommand = 83   // 查询门控信息
    PassthroughCommandAPKLog          PassthroughCommand = 84   // 查询APK Log
    PassthroughCommandReboot          PassthroughCommand = 85   // 重启
    PassthroughCommand4GLog           PassthroughCommand = 86   // 获取4G Log
    PassthroughCommandFan             PassthroughCommand = 87   // 控制大风扇
    PassthroughCommandQRLogin         PassthroughCommand = 88   // 二维码登录
    PassthroughCommandScreenSaver     PassthroughCommand = 89   // 设置屏幕保护
    PassthroughCommandScreen          PassthroughCommand = 90   // 设置电柜显示图片
    PassthroughCommandFullSOC         PassthroughCommand = 1000 // 修改满电阈值
    PassthroughCommandSerialModify    PassthroughCommand = 1001 // 修改电柜编码
    PassthroughCommandExchangeTimeout PassthroughCommand = 1002 // 修改换电倒计时
    PassthroughCommandBinStatus       PassthroughCommand = 1003 // 获取当前仓位信息
    PassthroughCommandServerUrl       PassthroughCommand = 1005 // 获取当前服务器连接信息
)

// PassthroughCommand 透传指令
type PassthroughCommand uint16

func (c PassthroughCommand) Value() uint16 {
    return uint16(c)
}

type PassthroughCode int

const (
    PassthroughCodeSuccess          PassthroughCode = 0
    PassthroughCodeInternalError    PassthroughCode = 100
    PassthroughCodeNoContent        PassthroughCode = 102
    PassthroughCodeErrorContent     PassthroughCode = 103
    PassthroughCodeServerNoRsp      PassthroughCode = 106
    PassthroughCodeServerOff        PassthroughCode = 107
    PassthroughCodeDeviceNoRsp      PassthroughCode = 108
    PassthroughCodeDeviceOff        PassthroughCode = 109
    PassthroughCodeCmdNotSupport    PassthroughCode = 111
    PassthroughCodeParamInalid      PassthroughCode = 114
    PassthroughCodeParamMissing     PassthroughCode = 115
    PassthroughCodeOperationTimeout PassthroughCode = 135
    PassthroughCodeOperationError   PassthroughCode = 136
)

func NextPassthroughSN() (n byte) {
    n = passthroughSN.Load().(byte)
    if n >= math.MaxUint8-1 {
        n = defaultPassthroughSN
    }
    n += 1
    passthroughSN.Store(n)
    return
}

type PassthroughRequest struct {
    C     PassthroughCommand `json:"c"`
    TM    int64              `json:"tm"`
    Sign  []byte             `json:"sign,omitempty"`
    Param any                `json:"param,omitempty"`
}

type PassthroughResponse struct {
    TM     any             `json:"tm,omitempty"`
    Code   PassthroughCode `json:"code,omitempty"`
    Result string          `json:"result,omitempty"`
}

// PassthroughOpenDoorRequest 开仓请求
type PassthroughOpenDoorRequest struct {
    Action    int    `json:"action"`    // 1:开仓 2:关仓
    Index     int    `json:"opendoor"`  // 仓门
    TaskToken string `json:"tasktoken"` // 换电任务token
}

type PassthroughBinAction int

const (
    PassthroughBinActionEnable  PassthroughBinAction = 1 // 启用
    PassthroughBinActionDisable PassthroughBinAction = 2 // 禁用
)

// PassthroughBinAvailableRequest 仓位启用或禁用请求
type PassthroughBinAvailableRequest struct {
    Action    PassthroughBinAction `json:"action"`   // 1:启用 2:禁用
    Index     int                  `json:"opendoor"` // 仓位序号 (0-11)
    CabinetSn string               `json:"cabinetsn"`
}

// PassthroughAmmeterRequest 请求电表读数
type PassthroughAmmeterRequest struct {
    CabinetSn string `json:"cabinetsn"`
}

func (m *PassthroughRequest) Bytes() (b []byte) {
    // buf := adapter.NewBuffer()
    // defer adapter.ReleaseBuffer(buf)
    //
    // err = binary.Write(buf, binary.BigEndian, m.C)
    // if err != nil {
    //     return
    // }
    // buf.Write(adapter.ConvertString2Bytes(strconv.FormatInt(m.TM, 10)))
    //
    // m.Sign = EcbEncrypt(buf.Bytes(), adapter.ConvertString2Bytes(key))

    b, _ = jsoniter.Marshal(m)
    return
}

func unpackPassthroughResponse(sn byte, b []byte) {
    var (
        success bool
        result  string
    )

    defer func() {
        if v := recover(); v != nil {
            zap.L().Error("捕获错误", zap.Error(fmt.Errorf("%v", v)))
        }
    }()

    defer func() {
        // 加载需要处理的任务信息
        loaded, exists := passthroughTasks.Load(sn)
        if !exists {
            return
        }

        t, ok := loaded.(*PassthroughTask)
        // 如果任务非法直接返回
        if !ok {
            return
        }

        // 任务是否需要格式化返回
        switch v := t.result.(type) {
        case *AmmeterResponse:
            v.Unmarshal(result)
        }

        // 任务是否需要等待结果返回
        if t.success != nil {
            adapter.ChSafeSend(t.success, success)
        }
    }()

    data := new(PassthroughResponse)
    err := jsoniter.Unmarshal(b, data)
    if err != nil {
        zap.L().Error("控制响应数据解析错误", zap.Error(err))
    }

    success = data.Code == PassthroughCodeSuccess
    result = data.Result
    // TODO 后续处理
}
