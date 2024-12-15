package immo

// FamilyContext represents the family situation. It contains the common information
// shared by all evaluations.
type FamilyContext struct {
	// TotalFamilyAssets is the total assets of the family, including cash, stocks, commodities, etc.
	TotalFamilyAssets float64

	// TotalFamilyLiabilities is the total liabilities of the family, including mortgage, car loan, etc.
	TotalFamilyLiabilities float64

	// EstimatedMortgageAmounts is the estimated mortgage amounts for different scenarios.
	EstimatedMortgageAmounts []float64

	// EstimatedMortgageInterestRates is the estimated interest rates for different scenarios.
	EstimatedMortgageInterestRates []float64

	// EstimatedMortgageDurations is the estimated mortgage durations for different scenarios.
	EstimatedMortgageDurations []int

	// EstimatedAgencyFees is the estimated additional fees for each scenario.
	EstimatedAdditionalFees []float64

	// EstimatedContributions is the estimated contributions for each scenario.
	EstimatedContributions []float64
}

// EvaluationContext represents the context of an evaluation.
type EvaluationContext struct {
	TotalFamilyAssets             float64
	TotalFamilyLiabilities        float64
	EstimatedMortgageAmount       float64
	EstimatedMortgageInterestRate float64
	EstimatedMortgageDuration     int
	EstimatedAgencyFees           float64
	EstimatedContribution         float64
}

// EvaludationResult represents the result of an evaluation.
type EvaludationResult struct {
	MortgageAmount         float64 `yaml:"mortgage_amount"`
	MonthlyMortgagePayment float64 `yaml:"monthly_mortgage_payment"`
	TotalPurchaseCost      float64 `yaml:"total_purchase_cost"`      // house + fees
	TotalMonthHousingCost  float64 `yaml:"total_month_housing_host"` // mortgage + fees
	TotalAnnualHousingCost float64 `yaml:"total_annual_housing_cost"`
	TotalCostOfLoan        float64 `yaml:"total_cost_of_loan"`
	RemainingAssets        float64 `yaml:"remaining_assets"` // after initial contribution
	MonthlyNetBalance      float64 `yaml:"monthly_net_balance"`
}
