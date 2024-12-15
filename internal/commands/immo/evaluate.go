package immo

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

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
	for _, mortgage := range cfg.EstimatedMortgages {
		println("Estimating mortgage amount for", mortgage.Amount)
		ctx := EvaluationContext{
			EstimatedMortgageAmount: mortgage.Amount,
		}
		result := evaluate(ctx)
		printResult(result)
	}
}

func loadConfig() (ImmoConfig, error) {
	var (
		configPath = os.Getenv("JIMI_CONFIG")
		config     ImmoConfig
	)
	if configPath == "" {
		return config, errors.New("JIMI_CONFIG is not set")
	}

	println("Loading config from")
	bytes, err := os.ReadFile(os.Getenv("JIMI_CONFIG") + "/immo.yaml")
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

func evaluate(ctx EvaluationContext) EvaludationResult {
	result := EvaludationResult{
		MortgageAmount: ctx.EstimatedMortgageAmount,
		// MonthlyMortgagePayment: calculateMonthlyMortgagePayment(ctx),
		// TotalPurchaseCost:      calculateTotalPurchaseCost(ctx),
		// TotalMonthHousingCose:  calculateTotalMonthHousingCost(ctx),
		// TotalAnnualHousingCost: calculateTotalAnnualHousingCost(ctx),
		// TotalCostOfLoan:        calculateTotalCostOfLoan(ctx),
		// RemainingAssets:        calculateRemainingAssets(ctx),
		// MonthlyNetBalance:      calculateMonthlyNetBalance(ctx),
	}
	return result
}
