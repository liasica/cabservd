// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

type BusinessAttr struct {
	Name        string `json:"name"`
	Sn          string `json:"sn"`
	Online      int    `json:"online"`
	OnlineTime  int64  `json:"onlineTime"`
	ModelType   string `json:"modelType"`
	CellNums    int    `json:"cellNums"`
	CustomPhone string `json:"customPhone"`
}
