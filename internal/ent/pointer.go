// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/auroraride/adapter"
	"github.com/auroraride/cabservd/internal/ent/cabinet"
	"github.com/auroraride/cabservd/internal/ent/console"
	"github.com/google/uuid"
)

type BinPointers []*BinPointer
type BinPointer struct {
	UUID          *string        `json:"uuid,omitempty"`
	CabinetID     *uint64        `json:"cabinet_id,omitempty"`
	Brand         *adapter.Brand `json:"brand,omitempty"`
	Serial        *string        `json:"serial,omitempty"`
	Name          *string        `json:"name,omitempty"`
	Ordinal       *int           `json:"ordinal,omitempty"`
	Open          *bool          `json:"open,omitempty"`
	Enable        *bool          `json:"enable,omitempty"`
	Health        *bool          `json:"health,omitempty"`
	BatteryExists *bool          `json:"battery_exists,omitempty"`
	BatterySn     *string        `json:"battery_sn,omitempty"`
	Voltage       *float64       `json:"voltage,omitempty"`
	Current       *float64       `json:"current,omitempty"`
	Soc           *float64       `json:"soc,omitempty"`
	Soh           *float64       `json:"soh,omitempty"`
	Remark        *string        `json:"remark,omitempty"`
}

type CabinetPointers []*CabinetPointer
type CabinetPointer struct {
	Online      *bool           `json:"online,omitempty"`
	Brand       *adapter.Brand  `json:"brand,omitempty"`
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

type ConsolePointers []*ConsolePointer
type ConsolePointer struct {
	CabinetID *uint64               `json:"cabinet_id,omitempty"`
	BinID     *uint64               `json:"bin_id,omitempty"`
	Operate   *adapter.Operate      `json:"operate,omitempty"`
	Serial    *string               `json:"serial,omitempty"`
	UUID      *uuid.UUID            `json:"uuid,omitempty"`
	Type      *console.Type         `json:"type,omitempty"`
	UserID    *string               `json:"user_id,omitempty"`
	UserType  *adapter.UserType     `json:"user_type,omitempty"`
	Step      *adapter.ExchangeStep `json:"step,omitempty"`
	Status    *console.Status       `json:"status,omitempty"`
	BeforeBin **adapter.BinInfo     `json:"before_bin,omitempty"`
	AfterBin  **adapter.BinInfo     `json:"after_bin,omitempty"`
	Message   *string               `json:"message,omitempty"`
	StartAt   *time.Time            `json:"startAt,omitempty"`
	StopAt    *time.Time            `json:"stopAt,omitempty"`
	Duration  *float64              `json:"duration,omitempty"`
}

type ScanPointers []*ScanPointer
type ScanPointer struct {
	CabinetID *uint64                          `json:"cabinet_id,omitempty"`
	UUID      *uuid.UUID                       `json:"uuid,omitempty"`
	Efficient *bool                            `json:"efficient,omitempty"`
	UserID    *string                          `json:"user_id,omitempty"`
	UserType  *adapter.UserType                `json:"user_type,omitempty"`
	Serial    *string                          `json:"serial,omitempty"`
	Data      **adapter.ExchangeUsableResponse `json:"data,omitempty"`
}
