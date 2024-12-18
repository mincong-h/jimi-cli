package immo

import (
	"errors"
	"fmt"
	"math"
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
	fmt.Printf("Found %d goods and %d mortgages to evaluate\n\n", len(cfg.Goods), len(cfg.EstimatedMortgages))

	var cityStats = make(map[string]CityStats)
	for _, city := range cfg.CityStats {
		cityStats[city.ZipCode] = city
	}

	for i, good := range cfg.Goods {
		fmt.Printf("%d. Mortgages for %q (%.0fK)\n", i+1, good.Name, math.Round(good.Price/1000))
		fmt.Println(good.Link)
		fmt.Println("==========")
		for j, mortgage := range cfg.EstimatedMortgages {
			fmt.Printf("%d.%d. Mortgage %.0fK\n", i+1, j+1, math.Round(mortgage.Amount/1000))
			fmt.Println("----------")

			result := evaluate(EvaluationContext{
				TotalFamilyAssets:      cfg.Family.TotalAssets,
				TotalFamilyLiabilities: cfg.Family.TotalLiabilities,
				ContributionThreshold:  cfg.Family.ContributionThreshold,
				Mortgage:               mortgage,
				CityStats:              cityStats,
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

func printResult(result EvaluationResult) {
	data, _ := yaml.Marshal(result)
	println(string(data))
}

func evaluate(ctx EvaluationContext, good Good) EvaluationResult {
	var alerts []string

	// Assume the agent fees are included in the price of the good.
	purchaseCost := good.Price * (1 + notaryFeesRate) // house + fees

	contribution := purchaseCost - ctx.Mortgage.Amount

	// TODO: add more expenses here
	reminingAssets := ctx.TotalFamilyAssets - contribution

	annualHousingCost := ctx.Mortgage.MonthlyCost*12 + good.PropertyTax

	if contribution > ctx.ContributionThreshold {
		alerts = append(alerts, "Contribution is above threshold")
	}

	// performances
	performance := GoodPerformance{}
	if good.Type == "house" {
		if stats, exists := ctx.CityStats[good.ZipCode]; exists {
			avg := stats.HouseAveragePricePerM2
			if good.Price > avg {
				performance.Comment = fmt.Sprintf("House price is %.0f%% above the average. (%.0f > %.0f)",
					(good.Price-avg)/avg*100,
					good.Price,
					avg,
				)
			} else {
				performance.Comment = fmt.Sprintf("House price is %.0f%% below the average. (%.0f < %.0f)",
					(avg-good.Price)/avg*100,
					good.Price,
					avg,
				)
			}
		} else {
			alerts = append(alerts, "City stats not found")
		}
	} else {
		if stats, exists := ctx.CityStats[good.ZipCode]; exists {
			avg := stats.ApartmentAveragePricePerM2
			if good.Price > avg {
				performance.Comment = fmt.Sprintf("Flat price is %.0f%% above the average. (%.0f > %.0f)",
					(good.Price-avg)/avg*100,
					good.Price,
					avg,
				)
			} else {
				performance.Comment = fmt.Sprintf("Flat price is %.0f%% below the average. (%.0f < %.0f)",
					(avg-good.Price)/avg*100,
					good.Price,
					avg,
				)
			}
		} else {
			alerts = append(alerts, "City stats not found")
		}
	}

	return EvaluationResult{
		PurchaseCost: PurchaseCost{
			MortgageAmount:    math.Round(ctx.Mortgage.Amount),
			Contribution:      math.Round(contribution),
			TotalPurchaseCost: math.Round(purchaseCost),
			RemainingAssets:   math.Round(reminingAssets),
		},
		MaintenanceCost: MaintenanceCost{
			MortgageMonthlyCost:    math.Round(ctx.Mortgage.MonthlyCost + ctx.Mortgage.Insurance),
			AnnualPropertyTax:      math.Round(good.PropertyTax),
			TotalAnnualHousingCost: math.Round(annualHousingCost),
		},
		Performance: performance,
		Alerts:      alerts,
	}
}
