// Copyright (C) liasica. 2023-present.
//
// Created at 2023-03-01
// Based on cabservd by liasica, magicrolan@qq.com.

package rpc

import (
	"net"

	"github.com/auroraride/adapter/rpc/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/auroraride/cabservd/internal/g"
)

func Run() {
	lis, err := net.Listen("tcp", g.Config.Rpc.Bind)
	if err != nil {
		zap.L().Fatal("RPC启动失败", zap.Error(err))
		return
	}

	s := grpc.NewServer(
	// TODO DUMP MIDDLEWARE
	)
	pb.RegisterCabinetServer(
		s,
		&cabinetServer{},
	)
	err = s.Serve(lis)
	if err != nil {
		zap.L().Fatal("RPC启动失败", zap.Error(err))
	}
}
