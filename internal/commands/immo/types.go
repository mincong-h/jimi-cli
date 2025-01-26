package immo

type ImmoConfig struct {
	Family FamilyContext `yaml:"family"`

	// EstimatedMortgages is the estimated mortgages for different scenarios.
	EstimatedMortgages []Mortgage `yaml:"estimated_mortgages"`

	// Goods is the list of goods to evaluate.
	Goods []Property `yaml:"goods"`

	// CityStats are statistics of the city where the goods are located.
	CityStats []CityStats `yaml:"cities"`

	// CurrentProperty is the context of the current property.
	CurrentProperty CurrentPropertyContext `yaml:"current_property"`
}

type CityStats struct {
	Name                       string  `yaml:"name"`
	ZipCode                    string  `yaml:"zip_code"`
	HouseAveragePricePerM2     float64 `yaml:"house_average_price_per_m2"`
	ApartmentAveragePricePerM2 float64 `yaml:"apartment_average_price_per_m2"`
}

type CurrentPropertyContext struct {
	MonthlyMortgage   float64 `yaml:"monthly_mortgage"`
	SurfaceM2         float64 `yaml:"surface_m2"`
	MonthlyIncome     float64 `yaml:"monthly_income"`
	MonthlyCharges    float64 `yaml:"monthly_charges"`
	GestionFeesRate   float64 `yaml:"gestion_fees_rate"`
	AnnualPropertyTax float64 `yaml:"annual_property_tax"`
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

	MonthlyParkingFee float64 `yaml:"monthly_parking_fee"`

	// MonthlySecondaryResidenceCost is the monthly cost of the secondary residence, including the
	// rent/mortgage, insurance, eletricity, etc.
	MonthlySecondaryResidenceCost float64 `yaml:"monthly_secondary_residence_cost"`
}

// EvaluationContext represents the context of an evaluation.
type EvaluationContext struct {
	Family          FamilyContext
	CurrentProperty CurrentPropertyContext
	Mortgage        Mortgage
	CityStats       map[string]CityStats // key: zip code
}

// EvaluationResult represents the result of an evaluation.
type EvaluationResult struct {
	CostSummary                CostSummary        `yaml:"cost_summary"`
	NewPropertyPurchaseCost    PurchaseCost       `yaml:"new_property_purchase"`
	NewPropertyOperationalCost OperationalCost    `yaml:"new_property_operational_cost"`
	NewPropertyPerformance     GoodPerformance    `yaml:"new_property_performance"`
	Renting                    RentingPerformance `yaml:"renting"`
	Alerts                     []string           `yaml:"alerts"`
}

type PurchaseCost struct {
	TotalPurchaseCost     float64 `yaml:"total_purchase_cost"` // house + fees
	Contribution          float64 `yaml:"contribution"`
	MortgageAmount        float64 `yaml:"mortgage_amount"`
	RemainingAssets       float64 `yaml:"remaining_assets"` // after initial contribution
	RenovationCost        float64 `yaml:"renovation_cost"`
	RenovationDescription string  `yaml:"renovation_description"`
}

type OperationalCost struct {
	MonthlyMortgageCost    float64 `yaml:"monthly_mortgage_cost"`
	MonthlyHousingCharges  float64 `yaml:"monthly_housing_charges"`
	MonthlyExpenses        float64 `yaml:"monthly_expenses"`
	MonthlyExpensesDiff    string  `yaml:"monthly_expenses_diff"`
	AnnualPropertyTax      float64 `yaml:"annual_property_tax"`
	TotalAnnualHousingCost float64 `yaml:"total_annual_housing_cost"`
}

type CostSummary struct {
	AnnualPropertyTax float64 `yaml:"annual_property_tax"`
}

type RentingPerformance struct {
	NetMonthlyGain    float64 `yaml:"net_monthly_gain"`
	MonthlyMortgage   float64 `yaml:"monthly_mortgage"`
	SurfaceM2         float64 `yaml:"surface_m2"`
	MonthlyIncome     float64 `yaml:"monthly_income"`
	MonthlyCharges    float64 `yaml:"monthly_charges"`
	GestionFeesRate   float64 `yaml:"gestion_fees_rate"`
	GestionFees       float64 `yaml:"gestion_fees"`
	AnnualPropertyTax float64 `yaml:"annual_property_tax"`
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

type Property struct {
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
	// Renovation
	// ----------
	RenovationCost                         float64 `yaml:"renovation_cost,omitempty" json:"renovation_cost,omitempty"`
	RenovationDescription                  string  `yaml:"renovation_description,omitempty" json:"renovation_description,omitempty"`
	EnergyPerformanceRatingAfterRenovation string  `yaml:"energy_performance_rating_after_renovation,omitempty" json:"energy_performance_rating_after_renovation,omitempty"`
	EnergyConsumptionAfterRenovation       float64 `yaml:"energy_consumption_after_renovation,omitempty" json:"energy_consumption_after_renovation,omitempty"`

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

	// WcCount is the number of bathrooms. Optional.
	WcCount int `yaml:"wc_count,omitempty" json:"wc_count,omitempty"`

	// Type is the type of the good. Required.
	Type string `yaml:"type" json:"type" jsonschema:"enum=house,enum=apartment"` // house or apartment

	// HasBalcony indicates if the good has a balcony. Optional.
	HasBalcony bool `yaml:"has_balcony,omitempty" json:"has_balcony,omitempty"`

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

	// RecentRenovationsOrUpdates is a list of recent renovations or updates. Optional.
	RecentRenovationsOrUpdates []string `yaml:"recent_renovations_or_updates,omitempty" json:"recent_renovations_or_updates,omitempty"`

	// ----------
	// Location Intelligence
	// ----------

	// PropertyNeighborhood is the neighborhood of the good. Optional.
	PropertyNeighborhood string `yaml:"property_neighborhood,omitempty" json:"property_neighborhood,omitempty"`

	// PropertyAddress is the address of the property. Optional.
	//
	// This field is optional because most of the real estate websites do not provide the address.
	// However, we can fill it manually in the configuration file after the first visit.
	PropertyAddress string `yaml:"property_address,omitempty" json:"property_address,omitempty"`

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

	// EnergyConsumptionAnnualCost is the annual energy consumption of the good. Optional.
	EnergyConsumptionAnnualCost float64 `yaml:"energy_consumption_annual_cost,omitempty" json:"energy_consumption_annual_cost,omitempty"`

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

func (p Property) PricePerM2() float64 {
	return p.Price / p.LivingSpaceLoiCarrezM2
}
