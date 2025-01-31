package tpl

import "embed"

//go:embed create/*.tmpl
var CreateTemplateFS embed.FS
