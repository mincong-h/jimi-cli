package commands

import (
	"github.com/mincong-h/jimi-cli/internal/commands/immo"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "jimi",
	Short: "jimi - A CLI tool for all personal projects of Jimi.",
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(immo.ImmoCmd)
}
