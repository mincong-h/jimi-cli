package immo

type ImmoConfig struct {
	Family FamilyContext `yaml:"family"`

	// EstimatedMortgages is the estimated mortgages for different scenarios.
	EstimatedMortgages []Mortgage `yaml:"estimated_mortgages"`

	// Goods is the list of goods to evaluate.
	Goods []Good `yaml:"goods"`

	// CityStats are statistics of the city where the goods are located.
	CityStats []CityStats `yaml:"cities"`
}

type CityStats struct {
	Name                       string  `yaml:"name"`
	ZipCode                    string  `yaml:"zip_code"`
	HouseAveragePricePerM2     float64 `yaml:"house_average_price_per_m2"`
	ApartmentAveragePricePerM2 float64 `yaml:"apartment_average_price_per_m2"`
}

// FamilyContext represents the family situation. It contains the common information
// shared by all evaluations.
type FamilyContext struct {
	// TotalAssets is the total assets of the family, including cash, stocks, commodities, etc.
	TotalAssets float64 `yaml:"total_assets"`

	// TotalLiabilities is the total liabilities of the family, including mortgage, car loan, etc.
	TotalLiabilities float64 `yaml:"total_liabilities"`

	ContributionThreshold float64 `yaml:"contribution_threshold"`

	// MonthlyExpenses is the total monthly expenses of the family, based on the average of the last 6 months.
	MonthlyExpenses float64 `yaml:"monthly_expenses"`

	// MonthlyHousingCharges is the total monthly housing charges of the family.
	MonthlyHousingCharges float64 `yaml:"monthly_housing_charges"`

	HomeSurfaceM2 float64 `yaml:"home_surface_m2"`
}

// EvaluationContext represents the context of an evaluation.
type EvaluationContext struct {
	Family    FamilyContext
	Mortgage  Mortgage
	CityStats map[string]CityStats // key: zip code
}

// EvaluationResult represents the result of an evaluation.
type EvaluationResult struct {
	PurchaseCost    PurchaseCost    `yaml:"purchase"`
	OperationalCost OperationalCost `yaml:"maintenance"`
	Performance     GoodPerformance `yaml:"performance"`
	Alerts          []string        `yaml:"alerts"`
}

type PurchaseCost struct {
	TotalPurchaseCost float64 `yaml:"total_purchase_cost"` // house + fees
	Contribution      float64 `yaml:"contribution"`
	MortgageAmount    float64 `yaml:"mortgage_amount"`
	RemainingAssets   float64 `yaml:"remaining_assets"` // after initial contribution
}

type OperationalCost struct {
	MonthlyMortgageCost    float64 `yaml:"monthly_mortgage_cost"`
	MonthlyHousingCharges  float64 `yaml:"monthly_housing_charges"`
	MonthlyExpenses        float64 `yaml:"monthly_expenses"`
	MonthlyExpensesDiff    string  `yaml:"monthly_expenses_diff"`
	AnnualPropertyTax      float64 `yaml:"annual_property_tax"`
	TotalAnnualHousingCost float64 `yaml:"total_annual_housing_cost"`
}

type GoodPerformance struct {
	PricePerM2        float64 `yaml:"price_per_m2"`
	AveragePricePerM2 float64 `yaml:"avg_price_per_m2"`
	Comment           string  `yaml:"comment"`
}

type Mortgage struct {
	Bank         string  `yaml:"bank"`
	Amount       float64 `yaml:"amount"`
	InterestRate float64 `yaml:"interest_rate"`
	Years        int     `yaml:"years"`
	MonthlyCost  float64 `yaml:"monthly_cost"` // without insurance
	Insurance    float64 `yaml:"insurance"`    // monthly
	Comment      string  `yaml:"comment"`
}

type Good struct {
	// ----------
	// General Information
	// ----------

	// Name is the name of the good. Required.
	Name string `yaml:"name" json:"name"`

	// OfferUrl is the link to the offer. Required.
	OfferUrl string `yaml:"offer_url" json:"offer_url"`

	// OfferDescription is the description of the offer. Required.
	OfferDescription string `yaml:"offer_description" json:"offer_description"`

	// ----------
	// Financials
	// ----------

	// Price is the price of the good shown in the offer. Required.
	Price float64 `yaml:"price" json:"price"`

	// BathroomCount is the number of bathrooms. Optional.
	AnnualPropertyTax float64 `yaml:"annual_property_tax,omitempty" json:"annual_property_tax,omitempty"`

	// ----------
	// Property Characteristics
	// ----------

	// TotalLivingSpaceM2 is the total living space in square meters. Required.
	//
	// It includes areas with a ceiling height below 1.8 meters
	TotalLivingSpaceM2 float64 `yaml:"total_living_space_m2" json:"total_living_space_m2"`

	// LivingSpaceLoiCarrezM2 is the living space in square meters respecting loi-carrez. Required.
	//
	// It excludes areas with a ceiling height below 1.8 meters.
	LivingSpaceLoiCarrezM2 float64 `yaml:"living_space_loi_carrez_m2" json:"living_space_loi_carrez_m2"`

	// LandSurfaceM2 is the land surface in square meters. Required.
	LandSurfaceM2 float64 `yaml:"land_surface_m2" json:"land_surface_m2"`

	// RoomCount is the number of rooms. Required.
	RoomCount int `yaml:"room_count" json:"room_count"`

	// BedroomCount is the number of bedrooms. Required.
	BedroomCount int `yaml:"bedroom_count" json:"bedroom_count"`

	// Type is the type of the good. Required.
	Type string `yaml:"type" json:"type" jsonschema:"enum=house,enum=apartment"` // house or apartment

	// HasGarden indicates if the good has a garden. Optional.
	HasGarden bool `yaml:"has_garden,omitempty" json:"has_garden,omitempty"`

	// HasTerrace indicates if the good has a terrace. Optional.
	HasTerrace bool `yaml:"has_terrace,omitempty" json:"has_terrace,omitempty"`

	// HasBox indicates if the good has a box. Optional.
	HasBox bool `yaml:"has_box,omitempty" json:"has_box,omitempty"`

	// HasGarage indicates if the good has a garage. Optional.
	HasGarage bool `yaml:"has_garage,omitempty" json:"has_garage,omitempty"`

	// ConstructionYear indicates the year of construction of the good. Optional.
	ConstructionYear int `yaml:"construction_year,omitempty" json:"construction_year,omitempty"`

	// ----------
	// Location Intelligence
	// ----------

	// GoodAddress is the address of the good. Optional.
	//
	// This field is optional because most of the real estate websites do not provide the address.
	// However, we can fill it manually in the configuration file after the first visit.
	GoodAddress string `yaml:"good_address,omitempty" json:"good_address,omitempty"`

	// ZipCode is the zip code of the good. Required.
	ZipCode string `yaml:"zip_code" json:"zip_code"`

	// DistanceByWalkToRer is the distance to the nearest RER station. Optional.
	DistanceByWalkToRer string `yaml:"distance_by_walk_to_rer,omitempty" json:"distance_by_walk_to_rer,omitempty"`

	// DistanceByWalkToBus is the distance to the nearest bus station. Optional.
	DistanceByWalkToBus string `yaml:"distance_by_walk_to_bus,omitempty" json:"distance_by_walk_to_bus,omitempty"`

	// ----------
	// Energy And Diagnosis
	// ----------

	// HeatingSystem describes the type of heating system (e.g., "individual gas heating: radiator")
	HeatingSystem string `yaml:"heating_system" json:"heating_system"`

	// HeatingType specifies the energy source used for heating (e.g., "gas", "electric")
	HeatingType string `yaml:"heating_type" json:"heating_type"`

	// HeatingMethod indicates the heating distribution method (e.g., "radiator")
	HeatingMethod string `yaml:"heating_method" json:"heating_method"`

	// EnergyPerformanceRating represents the DPE (Diagnostic de performance énergétique) rating
	// Rating from A to G, where A is the most efficient
	EnergyPerformanceRating string `yaml:"energy_performance_rating" json:"energy_performance_rating"`

	// GreenhouseGasRating represents the GES (Gaz à effet de serre) rating
	// Rating from A to G, where A has the lowest emissions
	EnergyGreenhouseGasRating string `yaml:"energy_greenhouse_gas_rating" json:"energy_greenhouse_gas_rating"`

	// EnergyConsumption in kWh/m²/year (optional)
	EnergyConsumption float64 `yaml:"energy_consumption,omitempty" json:"energy_consumption,omitempty"`

	// GHGEmissions in kgCO2/m²/year (optional)
	EnergyGHGEmissions float64 `yaml:"energy_ghg_emissions,omitempty" json:"energy_ghg_emissions,omitempty"`

	// EnergyEstimatedAnnualConsumption is the estimated annual energy consumption of the good. Optional.
	EnergyEstimatedAnnualConsumption string `yaml:"energy_estimated_annual_consumption,omitempty" json:"energy_estimated_annual_consumption,omitempty"`

	// ----------
	// Agency Information
	// ----------

	// AgencyName is the name of the agency. Optional.
	AgencyName string `yaml:"agency_name,omitempty" json:"agency_name,omitempty"`

	// AgencyEmail is the email of the agency. Optional.
	AgencyEmail string `yaml:"agency_email,omitempty" json:"agency_email,omitempty"`

	// AgencyTel is the phone number of the agency. Optional.
	AgencyTel string `yaml:"agency_tel,omitempty" json:"agency_tel,omitempty"`

	// Comment is a comment about the good. Optional.
	//
	// This is only used in the configuration file.
	Comment string `yaml:"comment" json:"-"`
}

func (g Good) PricePerM2() float64 {
	return g.Price / g.LivingSpaceLoiCarrezM2
}
