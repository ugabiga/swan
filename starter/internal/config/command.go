package config

import (
	"github.com/ugabiga/swan/core"
)

func ProvideCommand() *core.Command {
	return core.NewCommand()
}
