//go:build !embed

package web

import (
	"io/fs"
)

func GetDistFS() *fs.FS {
	return nil
}
