package config

import "github.com/ugabiga/swan/starter/internal/example"

var Commands = map[string]any{
	"crawl": example.InvokeCommand,
}
