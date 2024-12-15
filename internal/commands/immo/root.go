package immo

import "github.com/spf13/cobra"

var ImmoCmd = &cobra.Command{
	Use:   "immo",
	Short: `Commands for a real-estate  project.`,
	Run:   runImmo,
}

func runImmo(cmd *cobra.Command, args []string) {
}
