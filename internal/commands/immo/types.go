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

	ContributionThreshold float64 `yaml:"contribution_threshold"`
}

// EvaluationContext represents the context of an evaluation.
type EvaluationContext struct {
	TotalFamilyAssets      float64
	TotalFamilyLiabilities float64
	ContributionThreshold  float64
	MortgageAmount         float64
	MortgageInterestRate   float64
	MortgageYears          int
	MortgageMonthlyCost    float64
}

// EvaludationResult represents the result of an evaluation.
// The fields are grouped into two categories: one-time costs and continuous costs.
// For each category, the fields are ordered by importance.
type EvaludationResult struct {
	// One-time costs

	TotalPurchaseCost float64 `yaml:"total_purchase_cost"` // house + fees
	Contribution      float64 `yaml:"contribution"`
	MortgageAmount    float64 `yaml:"mortgage_amount"`
	RemainingAssets   float64 `yaml:"remaining_assets"` // after initial contribution

	// Continuous costs

	MortgageMonthlyCost    float64 `yaml:"mortgage_monthly_cost"`
	AnnualPropertyTax      float64 `yaml:"annual_property_tax"` // annual
	TotalAnnualHousingCost float64 `yaml:"total_annual_housing_cost"`

	// Metadata
	Alerts []string `yaml:"alerts"`
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
