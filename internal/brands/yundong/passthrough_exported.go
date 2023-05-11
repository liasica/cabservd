// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-17
// Based on yundong by liasica, magicrolan@qq.com.

package yundong

import (
	"github.com/google/uuid"
)

func commandIsSuccess(task *PassthroughTask, err error) bool {
	if err != nil {
		return false
	}

	return task.Wait()
}

// CommandReboot 重启
func CommandReboot(serial string) bool {
	_, err := sendPassthrough(serial, PassthroughCommandReboot, 1)
	return err == nil
}

// CommandDisable 禁用仓位, 从0开始
func CommandDisable(serial string, index int, times int) bool {
	return commandIsSuccess(sendPassthrough(serial, &PassthroughBinAvailableRequest{
		Action:    PassthroughBinActionDisable,
		Index:     index,
		CabinetSn: serial,
	}, times))
}

// CommandEnable 启用仓位, 从0开始
func CommandEnable(serial string, index int, times int) bool {
	return commandIsSuccess(sendPassthrough(serial, &PassthroughBinAvailableRequest{
		Action:    PassthroughBinActionEnable,
		Index:     index,
		CabinetSn: serial,
	}, times))
}

// CommandOpen 开启仓位
func CommandOpen(serial string, index int, times int) bool {
	return commandIsSuccess(sendPassthrough(serial, &PassthroughOpenDoorRequest{
		Action:    1,
		Index:     index,
		TaskToken: uuid.New().String(),
	}, times))
}
