// Copyright (C) liasica. 2022-present.
//
// Created at 2022-08-01
// Based on aurservd by liasica, magicrolan@qq.com.

package assets

import (
    "embed"
    "github.com/labstack/echo/v4"
    "html/template"
    "io"
    "io/fs"
    "strings"
)

const (
    templatesDir = "templates"
)

var (
    //go:embed templates/*
    views embed.FS

    Templates *htmlTemplate
)

type htmlTemplate struct {
    Templates map[string]*template.Template
}

func (t *htmlTemplate) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
    return t.Templates[name].ExecuteTemplate(w, name, data)
}

var (
    funcMaps = template.FuncMap{
        "add": func(x, y int) int {
            return x + y
        },
    }
)

func LoadTemplates() {
    Templates = &htmlTemplate{Templates: make(map[string]*template.Template)}

    _ = fs.WalkDir(views, templatesDir, func(path string, d fs.DirEntry, _ error) (err error) {
        if d.IsDir() {
            return
        }

        name := strings.Replace(path, templatesDir+"/", "", 1)
        pt := template.New(name)
        b, _ := views.ReadFile(path)
        _, _ = pt.Funcs(funcMaps).Parse(string(b))

        Templates.Templates[name] = pt
        return
    })
}
