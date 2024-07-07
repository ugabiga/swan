package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "swan",
	//Run: func(cmd *cobra.Command, args []string) {
	//
	//},
}

func main() {
	_ = rootCmd.Execute()
}
