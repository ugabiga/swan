package commands

import (
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/swctl/internal/commands/command"
	"github.com/ugabiga/swan/swctl/internal/commands/entdb"
	"github.com/ugabiga/swan/swctl/internal/commands/event"
	"github.com/ugabiga/swan/swctl/internal/commands/general_struct"
	"github.com/ugabiga/swan/swctl/internal/commands/handler"
	"github.com/ugabiga/swan/swctl/internal/commands/migration"
	"github.com/ugabiga/swan/swctl/internal/commands/new"
)

func Commands() []*cobra.Command {
	return []*cobra.Command{
		new.Cmd,
		handler.Cmd,
		event.Cmd,
		command.Cmd,
		general_struct.Cmd,
		entdb.NewCmd,
		entdb.GenerateCmd,
		migration.CreateCmd,
		migration.HashCmd,
		migration.DropCmd,
		migration.UpCmd,
		migration.DownCmd,
		migration.ForceCmd,
	}
}
