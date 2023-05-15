// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

type SnList struct {
	SnList []string `json:"snList"`
}

type BusinessAttrs []*BusinessAttr
type BusinessAttr struct {
	Name        string `json:"name"`
	Sn          string `json:"sn"`
	Online      int    `json:"online"`
	OnlineTime  int64  `json:"onlineTime"`
	ModelType   string `json:"modelType"`
	CellNums    int    `json:"cellNums"`
	CustomPhone string `json:"customPhone"`
}

type PhysicsAttrs []*PhysicsAttr
type PhysicsAttr struct {
	Sn                      string      `json:"sn"`
	Version                 string      `json:"version"`
	Voltage                 float64     `json:"voltage"`
	Current                 float64     `json:"current"`
	Temp                    float64     `json:"temp"`
	Iccid                   string      `json:"iccid"`
	PowerConsumption        float64     `json:"powerConsumption"`
	FanStatus               int         `json:"fanStatus"`
	LightStatus             int         `json:"lightStatus"`
	Power                   float64     `json:"power"`
	PowerFactor             float64     `json:"powerFactor"`
	ActiveElectricityEnergy float64     `json:"activeElectricityEnergy"`
	WaterPumpStatus         int         `json:"waterPumpStatus"`
	WaterLevelWarningStatus interface{} `json:"waterLevelWarningStatus"`
	WaterLeachingWarning    int         `json:"waterLeachingWarning"`
	Humidity                float64     `json:"humidity"`
	DoorStatus              int         `json:"doorStatus"`
	SmokeSensorStatus       int         `json:"smokeSensorStatus"`
	ExtinguisherStatus      int         `json:"extinguisherStatus"`
	CVersion                string      `json:"cVersion"`
}
