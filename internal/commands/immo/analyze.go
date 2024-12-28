package immo

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze a real-estate listing using Vertesia AI.",
	RunE:  runAnalyze,
}

func runAnalyze(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("Please provide a real-estate listing to analyze.")
	}
	property := args[0]
	fmt.Printf("Analyzing real-estate listing %s...\n", property)
	interaction := exec.Command("bash", "-c", "composable profiles")
	interaction.Stdout = os.Stdout
	interaction.Stderr = os.Stderr
	if err := interaction.Run(); err != nil {
		fmt.Printf("Failed to analyze: %v\n", err)
		return fmt.Errorf("failed to analyze: %w", err)
	}
	return nil
}
