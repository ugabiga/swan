package commands

import (
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/swctl/internal/commands/event"
	"github.com/ugabiga/swan/swctl/internal/commands/general"
	"github.com/ugabiga/swan/swctl/internal/commands/handler"
	"github.com/ugabiga/swan/swctl/internal/commands/new"
)

func Commands() []*cobra.Command {
	return []*cobra.Command{
		new.Cmd,
		handler.Cmd,
		event.Cmd,
		general.Cmd,
	}
}
