// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-29
// Based on cabservd by liasica, magicrolan@qq.com.

package core

import (
	"context"
	"time"

	"github.com/auroraride/adapter/log"
	"github.com/auroraride/cabservd/internal/ent"
	"github.com/auroraride/cabservd/internal/ent/bin"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/g"
	"go.uber.org/zap"
)

type CabinetUpdater interface {
	GetSerial() (string, bool)
	GetCabinet() (*ent.CabinetPointer, bool)
	GetBins() ent.BinPointers
}

func UpdateCabinet(h Hook, p CabinetUpdater) {
	ctx := context.Background()

	serial, ok := p.GetSerial()
	if !ok {
		return
	}

	cab := LoadOrStoreCabinet(ctx, serial)
	if cab == nil {
		zap.L().Error("仓位保存失败: 未找到电柜信息")
		return
	}

	cp, exists := p.GetCabinet()
	if exists {
		saveCabinet(h, cab, cp)
	}

	bp := p.GetBins()
	saveBins(cab, bp)
}

func LoadOrStoreCabinet(ctx context.Context, serial string) (cab *ent.Cabinet) {
	orm := ent.Database.Cabinet
	cab, _ = orm.Query().Where(cabinet.Serial(serial)).First(ctx)
	if cab != nil {
		return
	}
	var err error
	cab, err = orm.Create().SetSerial(serial).Save(ctx)
	if err != nil {
		zap.L().Error("电柜保存失败", zap.Error(err))
	}
	return
}

func saveCabinet(_ Hook, cab *ent.Cabinet, item *ent.CabinetPointer) {
	ctx := context.Background()
	u := ent.Database.Cabinet.UpdateOne(cab).SetUpdatedAt(time.Now())
	// 在线
	if item.Online != nil {
		u.SetOnline(*item.Online)
	}

	// 状态
	if item.Status != nil {
		u.SetStatus(*item.Status)
	}

	// 经度
	if item.Lng != nil {
		u.SetLng(*item.Lng)
	}

	// 纬度
	if item.Lat != nil {
		u.SetLat(*item.Lat)
	}

	// GSM
	if item.Gsm != nil {
		u.SetGsm(*item.Gsm)
	}

	// 电压
	if item.Voltage != nil {
		u.SetVoltage(*item.Voltage)
	}

	// 电流
	if item.Current != nil {
		u.SetCurrent(*item.Current)
	}

	// 温度
	if item.Temperature != nil {
		u.SetTemperature(*item.Temperature)
	}

	// 启用
	if item.Enable != nil {
		u.SetEnable(*item.Enable)
	}

	// 总用电
	if item.Electricity != nil {
		u.SetElectricity(*item.Electricity)
	}

	err := u.Exec(ctx)
	if err != nil {
		zap.L().Error("电柜保存失败", zap.Error(err), log.Payload(item))
	}
}

func binSaver(cab *ent.Cabinet, ordinal int, setter func(u *ent.BinMutation, b *ent.Bin)) error {
	ctx := context.Background()
	b, _ := ent.Database.Bin.Query().Where(bin.Serial(cab.Serial), bin.Ordinal(ordinal)).Only(ctx)
	var (
		mu      *ent.BinMutation
		creator *ent.BinCreate
		updater *ent.BinUpdateOne
	)
	if b == nil {
		creator = ent.Database.Bin.Create()
		mu = creator.Mutation()
	} else {
		updater = ent.Database.Bin.UpdateOne(b)
		mu = updater.Mutation()
	}

	mu.SetSerial(cab.Serial)
	mu.SetOrdinal(ordinal)
	mu.SetCabinetID(cab.ID)
	setter(mu, b)

	if creator != nil {
		return creator.Exec(ctx)
	}

	return updater.Exec(ctx)
}

func saveBins(cab *ent.Cabinet, items ent.BinPointers) {
	if len(items) == 0 {
		return
	}

	for _, item := range items {
		// fmt.Println(item.String())
		// TODO 删除DEBUG
		err := binSaver(cab, *item.Ordinal, func(u *ent.BinMutation, old *ent.Bin) {
			// u, old := binSaver(tx, cab, *item.Ordinal)
			u.SetName(*item.Name)
			u.SetUpdatedAt(time.Now())

			// 健康状态
			if item.Health != nil {
				u.SetHealth(*item.Health)
			}

			// 仓门状态
			if item.Open != nil {
				u.SetOpen(*item.Open)
			}

			// 仓位启用状态
			if item.Enable != nil {
				u.SetEnable(*item.Enable)
			}

			// 电压
			if item.Voltage != nil {
				u.SetVoltage(*item.Voltage)
			}

			// 电流
			if item.Current != nil {
				u.SetCurrent(*item.Current)
			}

			// 电量
			if item.Soc != nil {
				u.SetSoc(*item.Soc)
			}

			// 健康
			if item.Soh != nil {
				u.SetSoh(*item.Soh)
			}

			// 电池编号
			if item.BatterySn != nil {
				u.SetBatterySn(*item.BatterySn)
				// 如果需要自动清除电池数据
				if *item.BatterySn == "" && g.AutoResetWithoutBatterySN {
					u.ResetBattery()
				}
				// 如果无在位检测, 需要处理电池在位标记
				if !g.BatteryReign {
					u.SetBatteryExists(*item.BatterySn != "")
				}
			}

			// 电池在位
			if item.BatteryExists != nil {
				u.SetBatteryExists(*item.BatteryExists)
			}

			// // TODO 需要完善发送锁仓指令
			// // 当启用中的旧仓位中有电池时, 若非开门操作中电池编号丢失或在位丢失, 直接锁仓
			// if !mem.BinInOperation(cab.Serial, *item.Ordinal).IsOpen() &&
			//     old != nil && (old.BatteryExists || old.BatterySn != "") && old.Enable &&
			//     ((item.BatterySn != nil && *item.BatterySn == "") || (item.BatteryExists != nil && !*item.BatteryExists)) {
			//     u.SetEnable(false)
			//     u.SetRemark("未开门状态电池丢失")
			//     // TODO 如何发送锁仓指令
			//
			//     be := "True"
			//     if !old.BatteryExists {
			//         be = "False"
			//     }
			//     zap.L().Info(cab.Serial + " " + *item.Name + ", 未开门状态电池丢失, 旧电池信息: sn = " + old.BatterySn + ", battery_exists = " + be)
			// }
		})
		if err != nil {
			zap.L().Error("仓位保存失败", zap.Error(err), log.Payload(item))
		}
	}
}

// ResetBins 重置电柜仓位信息
func ResetBins(sn string) error {
	return ent.Database.Bin.Update().
		Where(bin.Serial(sn)).
		SetBatterySn("").
		SetSoc(0).
		SetSoh(0).
		SetVoltage(0).
		SetCurrent(0).
		// SetEnable(true). // TODO 是否单独设置LOCK
		SetOpen(false).
		Exec(context.Background())
}
