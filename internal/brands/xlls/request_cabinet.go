// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

// FetchBusinessAttr 获取柜机业务属性
func FetchBusinessAttr(sns []string) (BusinessAttrs, error) {
	res, err := doRequest[Response[BusinessAttrs]]("/openapi/cabinet/business/attr", SnList{SnList: sns})
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

// FetchPhysicsAttr 获取柜机物理属性
func FetchPhysicsAttr(sns []string) (PhysicsAttrs, error) {
	res, err := doRequest[Response[PhysicsAttrs]]("/openapi/cabinet/physics/attr", SnList{SnList: sns})
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}
