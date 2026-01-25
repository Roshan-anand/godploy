package ui

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var embedded embed.FS

var DistDirFS, _ = fs.Sub(embedded, "dist")
