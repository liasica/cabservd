// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

// 获取柜机业务属性
func fetchBusinessAttr(req *BusinessAttrRequest) (CabAttrs, error) {
	return doRequest[CabAttrs]("/openapi/cabinet/business/attr", req)
}

// 获取柜机物理属性
func fetchPhysicsAttr(req *SnListRequest) (CabAttrs, error) {
	return doRequest[CabAttrs]("/openapi/cabinet/physics/attr", req)
}

// 获取柜机APP设置属性
func fetchAppAttr(req *SnListRequest) (*AppAttr, error) {
	return doRequest[*AppAttr]("/openapi/cabinet/app/attr", req)
}

// 批量创建柜机
func fetchCabinetCreate(items []*CabinetMeta) (data any, err error) {
	return doRequest[any]("/openapi/cabinet/create", CabinetCreateRequest{CabinetList: items})
}

// 批量创建柜机型号
func fetchCabinetModelCreate(items []*CabinetModel) (data any, err error) {
	return doRequest[any]("/openapi/cabinet/model/create", CabinetModelCreateRequest{ModelList: items})
}

// 修改柜机属性
func fetchCrabinetModify(req *CabinetMeta) (data any, err error) {
	return doRequest[any]("/openapi/cabinet/modify", req)
}

// 批量删除柜机
func fetchCrabinetDelete(req *SnListRequest) (data any, err error) {
	return doRequest[any]("/openapi/cabinet/delete", req)
}

// 删除指定柜机型号
func fetchCabinetModelDelete(req string) (data any, err error) {
	return doRequest[any]("/openapi/cabinet/model/delete", req)
}

// 获取格挡属性
func fetchCellAttr(req *CellAttrRequest) (BinAttrs, error) {
	return doRequest[BinAttrs]("/openapi/cell/attr", req)
}

// FetchBatteryCreate 批量创建电池
func FetchBatteryCreate(req *BatteryCreateRequest) (data any, err error) {
	return doRequest[any]("/openapi/battery/create", req)
}

// 批量删除电池
func fetchBatteryDelete(req *BatteryDeleteRequest) (data any, err error) {
	return doRequest[any]("/openapi/battery/delete", req)
}

// 修改电池编号
func fetchBatteryModify(req *BatteryModifySnRequest) (data any, err error) {
	return doRequest[any]("/openapi/battery/modify", req)
}

// 获取电池属性
func fetchBatteryAttr(req *BatteryAttrRequest) (*BatteryAttr, error) {
	return doRequest[*BatteryAttr]("/openapi/battery/attr", req)
}

// 仓位控制
func fetchCellCommand(req *CellCommandRequest) (string, error) {
	result, err := request[ApiResponse[any]]("/openapi/cell/command", req)
	if err != nil {
		return "", err
	}
	return result.RequestID, nil
}

// 换电请求
func fetchExchange(req *bizExchangeRequest) (exchangeStatus, error) {
	result, err := doRequest[bizResponse]("/openapi/business/exchange/order", req)
	if err != nil {
		return exchangeStatusUnknown, err
	}
	status := exchangeStatus(result.Status)
	return status, status.error()
}

func fetchBizQuery(req *bizQueryRequest) (*bizQueryResult, error) {
	return doRequest[*bizQueryResult]("/openapi/business/order/query", req)
}

func fetchCoreCommand[T any](req *CabinetCommandRequest[T]) (any, error) {
	return doRequest[any]("/openapi/cabinet/core/command", req)
}
