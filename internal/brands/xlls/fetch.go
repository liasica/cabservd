// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

// FetchBusinessAttr 获取柜机业务属性
func FetchBusinessAttr(req *BusinessAttrRequest) (CabAttrs, error) {
	return doRequest[CabAttrs]("/openapi/cabinet/business/attr", req)
}

// FetchPhysicsAttr 获取柜机物理属性
func FetchPhysicsAttr(req *SnListRequest) (CabAttrs, error) {
	return doRequest[CabAttrs]("/openapi/cabinet/physics/attr", req)
}

// FetchAppAttr 获取柜机APP设置属性
func FetchAppAttr(req *SnListRequest) (*AppAttr, error) {
	return doRequest[*AppAttr]("/openapi/cabinet/app/attr", req)
}

// FetchCabinetCreate 批量创建柜机
func FetchCabinetCreate(items []*CabinetMeta) (data any, err error) {
	return doRequest[any]("/openapi/cabinet/create", CabinetCreateRequest{CabinetList: items})
}

// FetchCabinetModelCreate 批量创建柜机型号
func FetchCabinetModelCreate(items []*CabinetModel) (data any, err error) {
	return doRequest[any]("/openapi/cabinet/model/create", CabinetModelCreateRequest{ModelList: items})
}

// FetchCrabinetModify 修改柜机属性
func FetchCrabinetModify(req *CabinetMeta) (data any, err error) {
	return doRequest[any]("/openapi/cabinet/modify", req)
}

// FetchCrabinetDelete 批量删除柜机
func FetchCrabinetDelete(req *SnListRequest) (data any, err error) {
	return doRequest[any]("/openapi/cabinet/delete", req)
}

// FetchCabinetModelDelete 删除指定柜机型号
func FetchCabinetModelDelete(req string) (data any, err error) {
	return doRequest[any]("/openapi/cabinet/model/delete", req)
}

// FetchCellAttr 获取格挡属性
func FetchCellAttr(req *CellAttrRequest) (BinAttrs, error) {
	return doRequest[BinAttrs]("/openapi/cell/attr", req)
}

// FetchBatteryCreate 批量创建电池
func FetchBatteryCreate(req *BatteryCreateRequest) (data any, err error) {
	return doRequest[any]("/openapi/battery/create", req)
}

// FetchBatteryDelete 批量删除电池
func FetchBatteryDelete(req *BatteryDeleteRequest) (data any, err error) {
	return doRequest[any]("/openapi/battery/delete", req)
}

// FetchBatteryModify 修改电池编号
func FetchBatteryModify(req *BatteryModifySnRequest) (data any, err error) {
	return doRequest[any]("/openapi/battery/modify", req)
}

// FetchBatteryAttr 获取电池属性
func FetchBatteryAttr(req *BatteryAttrRequest) (*BatteryAttr, error) {
	return doRequest[*BatteryAttr]("/openapi/battery/attr", req)
}
