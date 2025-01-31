package commands_v2

import (
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/swctl/internal/commands_v2/handler"
	"github.com/ugabiga/swan/swctl/internal/commands_v2/new"
)

func Commands() []*cobra.Command {
	return []*cobra.Command{
		new.Cmd,
		handler.Cmd,
	}
}
