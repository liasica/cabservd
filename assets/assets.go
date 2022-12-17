// Copyright (C) liasica. 2022-present.
//
// Created at 2022-12-17
// Based on cabservd by liasica, magicrolan@qq.com.

package assets

import (
    "embed"
    _ "embed"
    log "github.com/sirupsen/logrus"
    "html/template"
    "io/fs"
    "net/http"
)

var (
    //go:embed config/config.yaml
    DefaultConfig []byte

    //go:embed templates/*
    f embed.FS
)

func LoadStatics(path string) http.FileSystem {
    sub, err := fs.Sub(f, path)
    if err != nil {
        log.Fatal("static load failed")
    }
    return http.FS(sub)
}

func LoadTemplates() (*template.Template, error) {
    return template.New("").ParseFS(f, "templates/**/*.html")
}
