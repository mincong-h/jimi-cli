package immo

import (
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var evaluateCmd = &cobra.Command{
	Use:   "evaluate",
	Short: "Evaluate different scenarios of a real-estate purchase.",
	Run:   runEvaluate,
}

func runEvaluate(cmd *cobra.Command, args []string) {
	family := loadFamilyContext()
	for i := range family.EstimatedMortgageAmounts {
		ctx := EvaluationContext{
			TotalFamilyAssets:             family.TotalFamilyAssets,
			TotalFamilyLiabilities:        family.TotalFamilyLiabilities,
			EstimatedMortgageAmount:       family.EstimatedMortgageAmounts[i],
			EstimatedMortgageInterestRate: family.EstimatedMortgageInterestRates[i],
			EstimatedMortgageDuration:     family.EstimatedMortgageDurations[i],
			EstimatedAgencyFees:           family.EstimatedAdditionalFees[i],
			EstimatedContribution:         family.EstimatedContributions[i],
		}
		result := evaluate(ctx)
		printResult(result)
	}
}

func loadFamilyContext() FamilyContext {
	return FamilyContext{
		TotalFamilyAssets:              100000,
		TotalFamilyLiabilities:         50000,
		EstimatedMortgageAmounts:       []float64{200000, 300000, 400000},
		EstimatedMortgageInterestRates: []float64{0.02, 0.025, 0.03},
		EstimatedMortgageDurations:     []int{15, 20, 30},
		EstimatedAdditionalFees:        []float64{10000, 15000, 20000},
		EstimatedContributions:         []float64{50000, 60000, 70000},
	}
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
