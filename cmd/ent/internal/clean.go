// Copyright (C) liasica. 2021-present.
//
// Created at 2021/12/10
// Based on aurservd by liasica, magicrolan@qq.com.

package internal

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	defaultEntPath = "./internal/ent"
)

type FileInfoType uint8

const (
	FileInfoTypeFile = iota
	FileInfoTypeDir
)

var (
	removeNotable = map[string]FileInfoType{
		"connect.go":          FileInfoTypeFile,
		"bin_attached.go":     FileInfoTypeFile,
		"updater.go":          FileInfoTypeFile,
		"cabinet_attached.go": FileInfoTypeFile,
		"console_attached.go": FileInfoTypeFile,
		"internal":            FileInfoTypeDir,
		"schema":              FileInfoTypeDir,
	}

	// keep = regexp.MustCompile(`^db.go|^connect.go|^slslog.go|^cabinet_task.go|^internal|^schema`)
)

func Clean(path ...string) {
	p := defaultEntPath
	if len(path) > 0 {
		p = path[0]
	}

	// _ = filepath.WalkDir(p, func(path string, d fs.DirEntry, err error) error {
	//     name := strings.Replace(path, p, "", 1)
	//     if name == "" || d.IsDir() {
	//         return nil
	//     }
	//     if !keep.MatchString(name) {
	//         _ = os.RemoveAll(path)
	//     }
	//     return nil
	// })

	fs, _ := ioutil.ReadDir(p)
	for _, f := range fs {
		name := f.Name()
		if _, ok := removeNotable[name]; !ok {
			_ = os.RemoveAll(filepath.Join(p, name))
		}
	}
}

func CleanCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "clean [path]",
		Short: "clean generated go code for the ent directory",
		Run: func(cmd *cobra.Command, path []string) {
			Clean(path...)
		},
	}
}
