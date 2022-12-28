// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package app

import (
    "github.com/go-playground/validator/v10"
    "net/http"
)

type Validator struct {
    validator *validator.Validate
}

func NewValidator() *Validator {
    return &Validator{validator: validator.New()}
}

func (v *Validator) Validate(i interface{}) error {
    if err := v.validator.Struct(i); err != nil {
        // Optionally, you could return the error to give each route more control over the status code
        Panic(http.StatusBadRequest, err)
    }
    return nil
}
