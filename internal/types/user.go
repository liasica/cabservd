// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-17
// Based on cabservd by liasica, magicrolan@qq.com.

package types

type UserType string

const (
    UserTypeCabinet  UserType = "cabinet"  // 电柜
    UserTypeManager  UserType = "manager"  // 后台操作
    UserTypeEmployee UserType = "employee" // 员工操作
    UserTypeRider    UserType = "rider"    // 骑手操作
)

type User struct {
    Type  UserType `json:"type" form:"type" binding:"required"` // 用户类别
    ID    uint64   `json:"id,omitempty" form:"id"`              // 用户ID
    Phone string   `json:"phone,omitempty" form:"phone"`        // 用户电话
}
