// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-16
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

type BatterySnData struct {
	BatterySn string `json:"batterySn"`
}

type BatteryAttrs []*BatteryAttr

// BatteryAttr 电池属性
type BatteryAttr struct {
	BatteryV      *float64  `json:"batteryV"`
	BatteryA      *float64  `json:"batteryA"`
	Soc           *int      `json:"soc"`
	Capacity      *float64  `json:"capacity"`
	CoreNum       *int      `json:"coreNum"`
	EnvTemp       *float64  `json:"envTemp"`
	CoreTemp      *float64  `json:"coreTemp"`
	BoardTemp     *float64  `json:"boardTemp"`
	CoreV         []float64 `json:"coreV"`
	BatteryHealth []string  `json:"batteryHealth"`
	Longitude     *string   `json:"longitude"`
	Latitude      *string   `json:"latitude"`
	ModelType     *string   `json:"modelType"`
}
