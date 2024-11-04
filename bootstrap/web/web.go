//go:build embed
// +build embed

package web

import (
	"embed"
	"io/fs"
)

//go:embed all:dist/*
var dist embed.FS

func GetDistFS() *fs.FS {
	var f fs.FS = dist
	return &f
}
