// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-28
// Based on cabservd by liasica, magicrolan@qq.com.

package app

import (
    errs "github.com/auroraride/adapter/errors"
    "net/http"
)

func Panic(params ...any) {
    r := CreateResponse(params...)
    if r.Code == http.StatusOK {
        r.Code = http.StatusInternalServerError
    }

    if r.Message == "" {
        switch r.Code {
        case http.StatusBadRequest:
            r.Message = errs.BadRequest.Error()
        default:
            r.Message = errs.InternalServerError.Error()
        }
    }

    panic(r)
}
