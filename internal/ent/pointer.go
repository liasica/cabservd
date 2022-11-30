// Code generated by ent, DO NOT EDIT.

package ent

import "github.com/auroraride/cabservd/internal/ent/cabinet"

type BinPointers []*BinPointer
type BinPointer struct {
	UUID      *string  `json:"uuid,omitempty"`
	Brand     *string  `json:"brand,omitempty"`
	Serial    *string  `json:"serial,omitempty"`
	Lock      *bool    `json:"lock,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Index     *int     `json:"index,omitempty"`
	Open      *bool    `json:"open,omitempty"`
	Enable    *bool    `json:"enable,omitempty"`
	Health    *bool    `json:"health,omitempty"`
	BatterySn *string  `json:"battery_sn,omitempty"`
	Voltage   *float64 `json:"voltage,omitempty"`
	Current   *float64 `json:"current,omitempty"`
	Soc       *float64 `json:"soc,omitempty"`
	Soh       *float64 `json:"soh,omitempty"`
}

type CabinetPointers []*CabinetPointer
type CabinetPointer struct {
	Brand       *string         `json:"brand,omitempty"`
	Serial      *string         `json:"serial,omitempty"`
	Status      *cabinet.Status `json:"status,omitempty"`
	Enable      *bool           `json:"enable,omitempty"`
	Lng         *float64        `json:"lng,omitempty"`
	Lat         *float64        `json:"lat,omitempty"`
	Gsm         *float64        `json:"gsm,omitempty"`
	Voltage     *float64        `json:"voltage,omitempty"`
	Current     *float64        `json:"current,omitempty"`
	Temperature *float64        `json:"temperature,omitempty"`
	Electricity *float64        `json:"electricity,omitempty"`
}
