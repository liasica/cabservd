// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-03
// Based on cabservd by liasica, magicrolan@qq.com.

package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/app"
	"github.com/auroraride/adapter/defs/cabdef"
	"go.uber.org/zap"

	"github.com/auroraride/cabservd/internal/biz"
	"github.com/auroraride/cabservd/internal/brands/xlls"
	"github.com/auroraride/cabservd/internal/core"
	"github.com/auroraride/cabservd/internal/ent"
	"github.com/auroraride/cabservd/internal/ent/bin"
	"github.com/auroraride/cabservd/internal/ent/console"
	"github.com/auroraride/cabservd/internal/g"
	"github.com/auroraride/cabservd/internal/sync"
	"github.com/auroraride/cabservd/internal/types"
)

type binService struct {
	*app.BaseService
	orm *ent.BinClient
}

func NewBin(params ...any) *binService {
	return &binService{
		BaseService: app.NewService(params...),
		orm:         ent.Database.Bin,
	}
}

func (s *binService) QueryAllBin() ent.Bins {
	items, _ := s.orm.Query().All(s.GetContext())
	return items
}

func (s *binService) Query(id uint64) (*ent.Bin, error) {
	return s.orm.Query().Where(bin.ID(id)).First(s.GetContext())
}

func (s *binService) QuerySerialOrdinal(serial string, ordinal int) (*ent.Bin, error) {
	return s.orm.Query().Where(bin.Serial(serial), bin.Ordinal(ordinal)).First(s.GetContext())
}

// Operate 按步骤操作某个仓位
func (s *binService) Operate(bo *types.Bin) (err error) {
	if bo.StepCallback == nil {
		return adapter.ErrorBadRequest
	}

	// 查询仓位
	eb, _ := NewBin(s.GetUser()).QuerySerialOrdinal(bo.Serial, bo.Ordinal)
	if eb == nil {
		return adapter.ErrorBinNotFound
	}

	// TODO 是否有必要操作之前验证当前状态, 操作值 等于 当前状态时直接返回成功
	// TODO 其他详细日志

	// 操作超时时间
	timeout := time.After(time.Duration(bo.Timeout) * time.Second)

	// 监听数据库变动
	ch := make(chan *ent.Bin)
	sync.Bin.SetListener(eb, ch)

	// 步骤结果监听器
	stepper := make(chan *types.BinResult)

	// 创建可中断任务
	biztask := biz.Create(bo.Serial, bo.Ordinal, bo.Business, bo.MainOperate, s.GetUser())
	defer biztask.Del()

	defer func() {
		// 退出时删除监听
		sync.Bin.RemoveListener(ch)
		close(stepper)

		// 判定是否成功以更新备注
		if err == nil && bo.BinRemark != nil {
			_ = s.orm.UpdateOneID(eb.ID).
				SetNillableRemark(bo.BinRemark).
				Exec(s.GetContext())
		}
	}()

	// 开启监听
	go func() {
		for {
			select {
			case msg := <-biztask.Interrupter:
				// TODO 中断任务
				err = errors.New("强制中断: " + msg)
				stepper <- types.NewBinResult(nil, err)
				return
			case <-timeout:
				err = adapter.ErrorOperateTimeout
				stepper <- types.NewBinResult(nil, err)
				return
			case x := <-ch:
				// 如果通道关闭直接返回
				if x == nil {
					return
				}

				// 更新仓位信息
				*eb = *x

				var doorOk, batteryOk, binOk bool

				step := bo.Current()

				switch step.Door {
				case cabdef.DetectDoorIgnore:
					// 忽略仓门检测
					doorOk = true
				case cabdef.DetectDoorOpen:
					// 检测仓门是否开启
					doorOk = x.Open
				case cabdef.DetectDoorClose:
					// 检测仓门是否关闭
					doorOk = !x.Open
				}

				switch step.Battery {
				case cabdef.DetectBatteryIgnore:
					// 忽略电池检测
					batteryOk = true
				case cabdef.DetectBatteryPutin:
					// 严格检测电池是否放入
					batteryOk = x.IsStrictHasBattery(g.Fakevoltage)
				case cabdef.DetectBatteryPutout:
					// 检测电池是否取出
					batteryOk = x.IsLooseNoBattery(g.Fakevoltage)
				}

				switch step.Bin {
				case cabdef.DetectBinIgnore:
					// 忽略仓位检测
					binOk = true
				case cabdef.DetectBinUsable:
					binOk = x.IsUsable()
					// 如果需要仓位可用但是仓位不可用, 直接发送任务失败并返回
					if !binOk {
						stepper <- types.NewBinResult(eb, adapter.ErrorBinNotUsable)
						return
					}
				case cabdef.DetectBinEnable:
					binOk = x.Enable
				case cabdef.DetectBinDisable:
					binOk = !x.Enable
				}

				if batteryOk && doorOk && binOk {
					// 若有bms通讯, 需检查放入电池编号是否匹配
					if !g.Config.NonBms && step.Battery == cabdef.DetectBatteryPutin && bo.Battery != "" && eb.BatterySn != bo.Battery {
						err = adapter.ErrorBatteryPutin
					}

					adapter.ChSafeSend(stepper, types.NewBinResult(eb, err))

					// 如果有错误, 终止
					if err != nil {
						return
					}

					// 尝试开启下次任务, 如果没有下次任务, 终止
					if !bo.Next() {
						return
					}

					// 将本次仓位信息结果传递到下次
					// TODO 取消数据库监听, 使用程序判定每一步的仓位变动
					// ch <- x
				}
			}
		}
	}()

	for _, step := range bo.Steps {
		err = s.doOperateStep(bo, eb, step, stepper)

		// 遇到错误, 直接返回
		if err != nil {
			return
		}
	}

	return
}

// IsExchangeFirstStep 是否换电第一步
func (s *binService) IsExchangeFirstStep(business adapter.Business, step *types.BinStep) bool {
	return business == adapter.BusinessExchange && step.Step == 1
}

// IsExchangeThirdStep 是否换电第三步
func (s *binService) IsExchangeThirdStep(business adapter.Business, step *types.BinStep) bool {
	return business == adapter.BusinessExchange && step.Step == 3
}

// 按步骤操作
func (s *binService) doOperateStep(bo *types.Bin, eb *ent.Bin, step *types.BinStep, stepper chan *types.BinResult) (err error) {
	// 创建记录
	var co *ent.Console
	co, err = ent.Database.Console.Create().
		SetOperate(step.Operate).
		SetCabinetID(eb.CabinetID).
		SetBinID(eb.ID).
		SetSerial(eb.Serial).
		SetUserID(s.GetUser().ID).
		SetUserType(s.GetUser().Type).
		SetStatus(console.StatusRunning).
		SetStartAt(time.Now()).
		SetBeforeBin(eb.Info()).
		SetStep(step.Step).
		SetBusiness(bo.Business).
		SetUUID(bo.UUID).
		SetRemark(bo.Remark).
		Save(s.GetContext())
	if err != nil {
		return
	}

	buf := adapter.NewBuffer()
	defer adapter.ReleaseBuffer(buf)

	buf.WriteString("处理结果: <")
	buf.WriteString(s.GetUser().String())
	buf.WriteString("> [电柜: ")
	buf.WriteString(eb.Serial)
	buf.WriteString(", 仓门:")
	buf.WriteString(strconv.Itoa(eb.Ordinal))
	buf.WriteString("] { ")
	buf.WriteString(bo.Business.Text())
	buf.WriteString("业务")
	buf.WriteString(step.String())
	buf.WriteString(" }")

	times := 0
	defer func() {
		res := NewConsole(s.GetUser()).Update(co, eb, times, err).OperateResult()

		if err != nil {
			zap.L().Error(buf.String(), zap.Error(err))
		} else {
			zap.L().Info(buf.String())
		}

		// 同步回调结果
		bo.StepCallback(res)
	}()

	// 电柜控制重试监听器
	var ticker *time.Timer

	// 电柜控制
	// TODO: 抽离接口
	switch g.Config.Brand {
	default:
		if step.Operate.IsCommand() {
			// 重复检测器
			// 初始设置为0立即执行指令
			ticker = time.AfterFunc(0, func() {
				// 「换电第一步」如果超过指定次数, 终止重复指令
				// TODO 这部分代码太丑了, 需要进行优化
				if s.IsExchangeFirstStep(bo.Business, step) && times >= g.ExchangeFirstStepRetryTimes ||
					s.IsExchangeThirdStep(bo.Business, step) && times >= g.ExchangeThirdStepRetryTimes {
					ticker.Stop()
					return
				}
				// 次数+1
				times += 1

				// 电柜控制
				err = core.Hub.Bean.SendOperate(eb.Serial, step.Operate, eb.Ordinal, times)

				// 如果电柜控制失败, 直接返回错误
				if err != nil {
					stepper <- types.NewBinResult(nil, err)
					return
				}

				// 「换电第一步」如果需要重复开仓, 则重置为每隔3s检测一次是否响应, 若指令无响应则重复开仓
				if g.ExchangeThirdStepRetryTimes > 1 && (s.IsExchangeFirstStep(bo.Business, step) || s.IsExchangeThirdStep(bo.Business, step)) && ticker != nil {
					ticker.Reset(4 * time.Second)
				}
			})

			// // 电柜控制
			// err = core.Hub.Bean.SendOperate(eb.Serial, step.Operate, eb.Ordinal)
			//
			// // TODO: 开仓失败后是否重复弹开逻辑???
			// // TODO: 详细失败日志???
			// if err != nil {
			// 	return
			// }
		}
	case adapter.CabinetBrandXiliulouServer:
		err = xlls.BinOperate(bo, step)
	}

	// 监听步骤结果
	r := <-stepper

	// 终止重复指令
	if ticker != nil {
		ticker.Stop()
	}

	_, err = r.Result()
	return
}

func (s *binService) BinInfo(req *cabdef.BinInfoRequest) (info *cabdef.BinInfo, err error) {
	b, _ := s.QuerySerialOrdinal(req.Serial, *req.Ordinal)
	if b == nil {
		err = adapter.ErrorBinNotFound
		return
	}
	info = b.Info()
	return
}

// Deactivate 禁用或启用仓位 (逻辑禁用)
func (s *binService) Deactivate(req *cabdef.BinDeactivateRequest) error {
	// 查找电柜
	_, err := NewCabinet(s.GetUser()).QuerySerial(req.Serial)
	if err != nil {
		return adapter.ErrorCabinetNotFound
	}

	updater := ent.Database.Bin.Update().Where(bin.Serial(req.Serial), bin.Ordinal(req.Ordinal)).SetNillableDeactivate(req.Deactivate)
	if *req.Deactivate {
		updater.SetNillableDeactivateReason(req.Reason)
	} else {
		updater.ClearDeactivateReason()
	}

	return updater.Exec(s.GetContext())
}
