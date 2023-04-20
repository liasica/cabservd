// Copyright (C) liasica. 2021-present.
//
// Created at 2021/12/10
// Based on aurservd by liasica, magicrolan@qq.com.

package main

import (
	"github.com/auroraride/cabservd/cmd/ent/internal"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{Use: "ent"}
	cmd.AddCommand(
		internal.InitCmd(),
		internal.GenerateCmd(),
		internal.CleanCmd(),
	)
	_ = cmd.Execute()
}
