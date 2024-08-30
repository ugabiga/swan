package core

import (
	"github.com/spf13/cobra"
)

type Command struct {
	mainCommand *cobra.Command
	commands    []*cobra.Command
}

func NewCommand() *Command {
	return &Command{
		mainCommand: &cobra.Command{},
	}
}

func (c *Command) RegisterCommand(commands ...*cobra.Command) {
	c.commands = append(c.commands, commands...)
}

func (c *Command) registerMainCommand(command *cobra.Command) {
	c.mainCommand = command
}

func (c *Command) Run() error {
	for _, command := range c.commands {
		c.mainCommand.AddCommand(command)
	}

	if err := c.mainCommand.Execute(); err != nil {
		return err
	}

	return nil
}
