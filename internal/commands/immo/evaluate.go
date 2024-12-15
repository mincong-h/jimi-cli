package immo

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// note: we can improve this later
// https://www.anil.org/outils/outils-de-calcul/frais-dacquisition-dits-frais-de-notaire/
const notaryFeesRate = 0.08

var evaluateCmd = &cobra.Command{
	Use:   "evaluate",
	Short: "Evaluate different scenarios of a real-estate purchase.",
	Run:   runEvaluate,
}

func runEvaluate(cmd *cobra.Command, args []string) {
	cfg, err := loadConfig()
	if err != nil {
		println(err)
		os.Exit(1)
	}
	for _, good := range cfg.Goods {
		println("Estimating mortgages for", good.Name)
		for _, mortgage := range cfg.EstimatedMortgages {
			fmt.Printf("Estimating mortgage %.0f\n", mortgage.Amount)
			result := evaluate(EvaluationContext{
				TotalFamilyAssets:      cfg.Family.TotalAssets,
				TotalFamilyLiabilities: cfg.Family.TotalLiabilities,
				MortgageAmount:         mortgage.Amount,
				MortgageInterestRate:   mortgage.InterestRate,
				MortgageYears:          mortgage.Years,
				MortgageMonthlyCost:    mortgage.MonthlyCost,
			}, good)
			printResult(result)
		}
	}
}

func loadConfig() (ImmoConfig, error) {
	var (
		rootConfigPath = os.Getenv("JIMI_CONFIG")
		configPath     = rootConfigPath + "/immo.yaml"
		config         ImmoConfig
	)
	if rootConfigPath == "" {
		return config, errors.New("JIMI_CONFIG is not set")
	}

	println("Loading config from", configPath)
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return config, err
	}
	return config, nil
}

func printResult(result EvaludationResult) {
	data, _ := yaml.Marshal(result)
	println(string(data))
}

func evaluate(ctx EvaluationContext, good Good) EvaludationResult {
	// Assume the agent fees are included in the price of the good.
	purchaseCost := good.Price * (1 + notaryFeesRate) // house + fees

	contribution := purchaseCost - ctx.MortgageAmount

	// TODO: add more expenses here
	reminingAssets := ctx.TotalFamilyAssets - contribution

	return EvaludationResult{
		MortgageAmount:      ctx.MortgageAmount,
		MortgageMonthlyCost: ctx.MortgageMonthlyCost,
		Contribution:        contribution,
		TotalPurchaseCost:   purchaseCost,
		RemainingAssets:     reminingAssets,
	}
}
