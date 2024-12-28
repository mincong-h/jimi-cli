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

var analysisInteractions = map[string]string{
	"Renovation":               "mhuang-seloger:RenovationAnalysis",
	"Location Intelligence":    "mhuang-seloger:LocationIntelligenceAnalysis",
	"Legal And Administrative": "mhuang-seloger:LegalAndAdministrativeAnalysis",
	"Market Dynamics":          "mhuang-seloger:MarketDynamicsAnalysis",
	"Lifestyle and Fit":        "mhuang-seloger:LifestyleAndFitAnalysis",
	"Risks":                    "mhuang-seloger:RisksAnalysis",
}

func runAnalyze(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("Please provide a real-estate listing to analyze.")
	}
	var (
		property = args[0]
		fileName = fmt.Sprintf("analysis-%s.md", property)
	)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open output file: %w", err)
	}
	defer f.Close()

	fmt.Printf("Analyzing real-estate listing %q...\n", property)
	for header, interaction := range analysisInteractions {
		if _, err := f.WriteString(fmt.Sprintf("## %s\n\n", header)); err != nil {
			return fmt.Errorf("failed to write header: %w", err)
		}

		bashCmd := fmt.Sprintf(`composable run %s --data '{"real_estate_listing": "store:%s"}'`, interaction, property)
		fmt.Println(bashCmd)
		interaction := exec.Command("bash", "-c", bashCmd)
		interaction.Stdout = f
		interaction.Stderr = os.Stderr
		if err := interaction.Run(); err != nil {
			fmt.Printf("Failed to analyze: %v\n", err)
			return fmt.Errorf("failed to analyze: %w", err)
		}
		fmt.Println("---")
		fmt.Println()
	}
	fmt.Printf("Analysis completed. Visit file for details: %s\n", fileName)
	return nil
}
