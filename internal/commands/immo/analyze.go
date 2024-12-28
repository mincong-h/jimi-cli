package immo

import (
	"bytes"
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

type analysisInteraction struct {
	section     string
	interaction string
}

var analysisInteractions = []analysisInteraction{
	{
		section:     "1. Renovation",
		interaction: "mhuang-seloger:RenovationAnalysis",
	},
	{
		section:     "2. Location Intelligence",
		interaction: "mhuang-seloger:LocationIntelligenceAnalysis",
	},
	{
		section:     "3. Legal And Administrative",
		interaction: "mhuang-seloger:LegalAndAdministrativeAnalysis",
	},
	{
		section:     "4. Market Dynamics",
		interaction: "mhuang-seloger:MarketDynamicsAnalysis",
	},
	{
		section:     "5. Lifestyle and Fit",
		interaction: "mhuang-seloger:LifestyleAndFitAnalysis",
	},
	{
		section:     "6. Risks",
		interaction: "mhuang-seloger:RisksAnalysis",
	},
}

func runAnalyze(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("Please provide a real-estate listing to analyze.")
	}
	var (
		property = args[0]
		fileName = fmt.Sprintf(".jimi/analysis-%s.md", property)
	)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open output file: %w", err)
	}
	defer f.Close()

	fmt.Printf("Analyzing real-estate listing %q...\n", property)
	for _, a := range analysisInteractions {
		var (
			out     bytes.Buffer
			bashCmd = fmt.Sprintf(`composable run %s --data '{"real_estate_listing": "store:%s"}'`, a.interaction, property)
		)

		fmt.Println(bashCmd)
		interaction := exec.Command("bash", "-c", bashCmd)
		interaction.Stdout = &out
		interaction.Stderr = os.Stderr

		if err := interaction.Run(); err != nil {
			fmt.Printf("Failed to analyze: %v\n", err)
			return fmt.Errorf("failed to analyze: %w", err)
		}

		if _, err := f.WriteString(fmt.Sprintf("## %s\n\n", a.section)); err != nil {
			return fmt.Errorf("failed to write header: %w", err)
		}
		if _, err := f.Write(out.Bytes()); err != nil {
			return fmt.Errorf("failed to write content: %w", err)
		}
		if _, err := f.WriteString("\n\n"); err != nil {
			return fmt.Errorf("failed to write end: %w", err)
		}

		fmt.Println("---")
		fmt.Println()
	}
	fmt.Printf("Analysis completed. Visit file for details: %s\n", fileName)
	return nil
}
