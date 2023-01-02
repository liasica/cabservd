// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-02
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
    "github.com/auroraride/adapter"
    "github.com/auroraride/cabservd/internal/app"
    "github.com/auroraride/cabservd/internal/ent"
    "github.com/auroraride/cabservd/internal/types"
    "github.com/jinzhu/copier"
    "golang.org/x/exp/slices"
    "net/http"
)

type businessService struct {
    *BaseService
}

func NewBusiness(params ...any) *businessService {
    return &businessService{
        BaseService: newService(params...),
    }
}

func (s *businessService) RiderBusinessVerify(b adapter.Business) bool {
    return slices.Contains(adapter.RiderBusiness, b)
}

func (s *businessService) RiderBusinessVerifyX(b adapter.Business) {
    if !slices.Contains(adapter.RiderBusiness, b) {
        app.Panic(http.StatusBadRequest, adapter.ErrorBusiness)
    }
}

// Usable 获取业务可用仓位
func (s *businessService) Usable(req *adapter.BusinuessUsableRequest) (res *adapter.CabinetBinUsableResponse) {
    s.RiderBusinessVerifyX(req.Business)

    // 判定最小空仓位和最小电池数量
    var minemptybins, minbatteries int
    var useEmpty bool

    switch req.Business {
    case adapter.BusinessActive, adapter.BusinessContinue:
        minbatteries = 2
    case adapter.BusinessPause, adapter.BusinessUnsubscribe:
        minemptybins = 2
        useEmpty = true
    default:
        return
    }

    cs := NewCabinet(s.User)
    // 查找电柜和仓位
    cab, _ := cs.QuerySerialWithBin(req.Serial)

    // 验证电柜是否可以办理业务
    err := cs.DetectCabinet(cab)
    if err != nil {
        app.Panic(http.StatusBadRequest, err)
    }

    // 获取可办理业务的仓位
    var fully, empty, target *ent.Bin
    fully, empty, err = cs.BusinessInfo(req.Model, cab, req.Minsoc, minbatteries, minemptybins)
    if err != nil {
        app.Panic(http.StatusBadRequest, err)
    }

    if useEmpty {
        target = empty
    } else {
        target = fully
    }

    // 存储扫码记录
    res = &adapter.CabinetBinUsableResponse{
        Cabinet:     new(adapter.Cabinet),
        BusinessBin: new(adapter.Bin),
    }

    // 拷贝属性
    _ = copier.Copy(res.Cabinet, cab)
    _ = copier.Copy(res.BusinessBin, target)

    sm := NewScan(s.User).Create(req.Business, req.Serial, cab, res)
    res.UUID = sm.UUID.String()

    return
}

// Do 执行业务
func (s *businessService) Do(req *adapter.BusinessRequest) (res adapter.BusinessResponse) {
    s.RiderBusinessVerifyX(req.Business)

    sc := NewScan(s.User).CensorX(req.UUID, req.Timeout, 0)

    cb := func(r *adapter.OperateStepResult) {
        res.Results = append(res.Results, r)
    }

    var conf types.BinSteps

    switch req.Business {
    case adapter.BusinessActive, adapter.BusinessContinue:
        conf = types.PutoutConfigure
    case adapter.BusinessPause, adapter.BusinessUnsubscribe:
        conf = types.PutinConfigure
    default:
        app.Panic(http.StatusBadRequest, adapter.ErrorBusiness)
    }

    // 开始操作
    err := NewBin(s.User).Operate(&types.Bin{
        Timeout:      req.Timeout,
        Serial:       req.Serial,
        UUID:         req.UUID,
        Ordinal:      sc.Data.BusinessBin.Ordinal,
        Business:     req.Business,
        Steps:        conf,
        Battery:      req.Battery,
        StepCallback: cb,
    })

    if err != nil {
        res.Error = err.Error()
    }

    return
}
