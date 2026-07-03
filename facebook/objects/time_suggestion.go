package objects

type TimeSuggestion struct {
	HighDemandPeriods *[]map[string]interface{} `json:"high_demand_periods,omitempty"`
	IsEnabled         *bool                     `json:"is_enabled,omitempty"`
}
