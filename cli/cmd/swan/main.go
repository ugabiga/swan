package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "swan",
}

func main() {
	_ = rootCmd.Execute()
}
