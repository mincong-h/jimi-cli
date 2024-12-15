package immo

type ImmoConfig struct {
	Family FamilyContext `yaml:"family"`

	// EstimatedMortgages is the estimated mortgages for different scenarios.
	EstimatedMortgages []Mortgate `yaml:"estimated_mortgages"`

	// Goods is the list of goods to evaluate.
	Goods []Good `yaml:"goods"`
}

// FamilyContext represents the family situation. It contains the common information
// shared by all evaluations.
type FamilyContext struct {
	// TotalAssets is the total assets of the family, including cash, stocks, commodities, etc.
	TotalAssets float64 `yaml:"total_assets"`

	// TotalLiabilities is the total liabilities of the family, including mortgage, car loan, etc.
	TotalLiabilities float64 `yaml:"total_liabilities"`
}

// EvaluationContext represents the context of an evaluation.
type EvaluationContext struct {
	TotalFamilyAssets      float64
	TotalFamilyLiabilities float64
	MortgageAmount         float64
	MortgageInterestRate   float64
	MortgageYears          int
	MortgageMonthlyCost    float64
}

// EvaludationResult represents the result of an evaluation.
type EvaludationResult struct {
	MortgageAmount         float64 `yaml:"mortgage_amount"`
	MortgageMonthlyCost    float64 `yaml:"mortgage_monthly_cost"`
	TotalPurchaseCost      float64 `yaml:"total_purchase_cost"`      // house + fees
	TotalMonthHousingCost  float64 `yaml:"total_month_housing_host"` // mortgage + fees
	TotalAnnualHousingCost float64 `yaml:"total_annual_housing_cost"`
	TotalCostOfLoan        float64 `yaml:"total_cost_of_loan"`
	Contribution           float64 `yaml:"contribution"`
	RemainingAssets        float64 `yaml:"remaining_assets"` // after initial contribution
	MonthlyNetBalance      float64 `yaml:"monthly_net_balance"`
}

type Mortgate struct {
	Bank         string  `yaml:"bank"`
	Amount       float64 `yaml:"amount"`
	InterestRate float64 `yaml:"interest_rate"`
	Years        int     `yaml:"years"`
	MonthlyCost  float64 `yaml:"monthly_cost"`
	Comment      string  `yaml:"comment"`
}

type Good struct {
	Name          string  `yaml:"name"`
	Price         float64 `yaml:"price"`
	Address       string  `yaml:"address"`
	Link          string  `yaml:"link"`
	LivingSpaceM2 float64 `yaml:"living_space_m2"` // loi carrez
	LandSurfaceM2 float64 `yaml:"land_surface_m2"`
	Pieces        int     `yaml:"pieces"`
	Rooms         int     `yaml:"rooms"`
	PropertyTax   float64 `yaml:"property_tax"` // annual
	Comment       string  `yaml:"comment"`
}
