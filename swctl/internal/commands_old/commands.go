package commands

import (
	"github.com/spf13/cobra"
)

func Commands() []*cobra.Command {
	return []*cobra.Command{
		// MakeStruct,
		MakeDBClient,
		MakeCommandCommand,
		// MakeEventCommand,
	}
}
