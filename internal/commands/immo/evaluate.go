package immo

import "github.com/spf13/cobra"

var evaluateCmd = &cobra.Command{
	Use:   "evaluate",
	Short: "Evaluate different scenarios of a real-estate purchase.",
	Run:   runEvaluate,
}

func runEvaluate(cmd *cobra.Command, args []string) {
}
