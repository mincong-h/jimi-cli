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
		fmt.Printf("City %q (%s)\n", city.Name, city.ZipCode)
		cityStats[city.ZipCode] = city
	}

	for i, good := range cfg.Goods {
		fmt.Printf("%d. Mortgages for %q (%.0fK)\n", i+1, good.Name, math.Round(good.Price/1000))
		fmt.Println(good.OfferUrl)
		fmt.Println("==========")
		for j, mortgage := range cfg.EstimatedMortgages {
			fmt.Printf("%d.%d. Mortgage %.0fK\n", i+1, j+1, math.Round(mortgage.Amount/1000))
			fmt.Println("----------")

			result := evaluate(EvaluationContext{
				Family:          cfg.Family,
				CurrentProperty: cfg.CurrentProperty,
				Mortgage:        mortgage,
				CityStats:       cityStats,
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

func evaluate(ctx EvaluationContext, good Property) EvaluationResult {
	var alerts []string

	// Assume the agent fees are included in the price of the good.
	purchaseCost := good.Price * (1 + notaryFeesRate) // house + fees

	contribution := purchaseCost - ctx.Mortgage.Amount

	reminingAssets := ctx.Family.TotalAssets - contribution

	if contribution > ctx.Family.ContributionThreshold {
		alerts = append(alerts, fmt.Sprintf("Contribution is above threshold (%.0fK > %.0fK)",
			contribution/1000,
			ctx.Family.ContributionThreshold/1000),
		)
	}

	// performances
	performance := GoodPerformance{}
	if good.Type == "house" {
		if stats, exists := ctx.CityStats[good.ZipCode]; exists {
			avg := stats.HouseAveragePricePerM2
			performance.PricePerM2 = math.Round(good.PricePerM2())
			performance.AveragePricePerM2 = math.Round(avg)

			if good.PricePerM2() > avg {
				performance.Comment = fmt.Sprintf("House price is %.0f%% above the average. (%.0f > %.0f)",
					(good.PricePerM2()-avg)/avg*100,
					good.PricePerM2(),
					avg,
				)
			} else {
				performance.Comment = fmt.Sprintf("House price is %.0f%% below the average. (%.0f < %.0f)",
					(avg-good.PricePerM2())/avg*100,
					good.PricePerM2(),
					avg,
				)
			}
		} else {
			alerts = append(alerts, "City stats not found")
		}
	} else {
		if stats, exists := ctx.CityStats[good.ZipCode]; exists {
			avg := stats.ApartmentAveragePricePerM2
			performance.PricePerM2 = math.Round(good.PricePerM2())
			performance.AveragePricePerM2 = math.Round(avg)

			if good.PricePerM2() > avg {
				performance.Comment = fmt.Sprintf("Flat price is %.0f%% above the average. (%.0f > %.0f)",
					(good.PricePerM2()-avg)/avg*100,
					good.PricePerM2(),
					avg,
				)
			} else {
				performance.Comment = fmt.Sprintf("Flat price is %.0f%% below the average. (%.0f < %.0f)",
					(avg-good.PricePerM2())/avg*100,
					good.PricePerM2(),
					avg,
				)
			}
		} else {
			alerts = append(alerts, "City stats not found")
		}
	}

	// ----------
	// Maintenance costs: start
	//
	// If the good is bigger than the current home, the monthly housing charges will increase proportionally.
	monthlyHousingCharges := ctx.CurrentProperty.MonthlyCharges * (good.TotalLivingSpaceM2 / ctx.CurrentProperty.SurfaceM2)
	monthlyExpenses := ctx.Family.MonthlyExpenses + monthlyHousingCharges + ctx.Mortgage.MonthlyCost // note: we cannot remove existing charges until we rent the current flat
	if good.HasGarage {
		monthlyExpenses -= ctx.Family.MonthlyParkingFee
	}
	monthlyExpensesDiff := fmt.Sprintf("%.0f (%.0f%%)",
		monthlyExpenses-ctx.Family.MonthlyExpenses,
		(monthlyExpenses-ctx.Family.MonthlyExpenses)/ctx.Family.MonthlyExpenses*100,
	)
	annualHousingCost := (monthlyHousingCharges+ctx.Mortgage.MonthlyCost)*12 + good.AnnualPropertyTax
	// Maintenance costs: end
	// ----------

	// ----------
	// Renting: start
	cp := ctx.CurrentProperty
	// e.g. (1200-385)*(1-0.08) - 1378/12 - 920 = 635
	rentingGain := (cp.MonthlyIncome-cp.MonthlyCharges)*(1-cp.GestionFeesRate) - cp.AnnualPropertyTax/12 - cp.MonthlyMortgage
	renting := RentingPerformance{
		NetMonthlyGain:    math.Round(rentingGain),
		MonthlyMortgage:   cp.MonthlyMortgage,
		SurfaceM2:         cp.SurfaceM2,
		MonthlyIncome:     cp.MonthlyIncome,
		MonthlyCharges:    cp.MonthlyCharges,
		GestionFeesRate:   cp.GestionFeesRate,
		AnnualPropertyTax: cp.AnnualPropertyTax,
	}
	// Renting: end
	// ----------

	return EvaluationResult{
		PurchaseCost: PurchaseCost{
			MortgageAmount:    math.Round(ctx.Mortgage.Amount),
			Contribution:      math.Round(contribution),
			TotalPurchaseCost: math.Round(purchaseCost),
			RemainingAssets:   math.Round(reminingAssets),
		},
		OperationalCost: OperationalCost{
			MonthlyMortgageCost:    math.Round(ctx.Mortgage.MonthlyCost + ctx.Mortgage.Insurance),
			MonthlyHousingCharges:  math.Round(monthlyHousingCharges),
			MonthlyExpenses:        math.Round(monthlyExpenses),
			MonthlyExpensesDiff:    monthlyExpensesDiff,
			AnnualPropertyTax:      math.Round(good.AnnualPropertyTax),
			TotalAnnualHousingCost: math.Round(annualHousingCost),
		},
		Performance: performance,
		Renting:     renting,
		Alerts:      alerts,
	}
}
