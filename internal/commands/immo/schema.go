package immo

import (
	"github.com/spf13/cobra"
)

var showSchemaCmd = &cobra.Command{
	Use:   "evaluate",
	Short: "Evaluate different scenarios of a real-estate purchase.",
	Run:   runShowSchema,
}

func runShowSchema(cmd *cobra.Command, args []string) {
}
