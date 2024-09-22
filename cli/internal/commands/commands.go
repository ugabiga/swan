package commands

import "github.com/spf13/cobra"

func Commands() []*cobra.Command {
	return []*cobra.Command{
		NewCmd,
		MakeStruct,
		MakeDomainCmd,
		MakeDBClient,
	}
}
