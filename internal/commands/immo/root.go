package immo

import "github.com/spf13/cobra"

var ImmoCmd = &cobra.Command{
	Use:   "immo",
	Short: `Commands for a real-estate  project.`,
	RunE:  runImmo,
}

func runImmo(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

func init() {
	ImmoCmd.AddCommand(evaluateCmd)
	ImmoCmd.AddCommand(showSchemaCmd)
}
