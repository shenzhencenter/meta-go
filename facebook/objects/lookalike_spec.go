package objects

type LookalikeSpec struct {
	Country                    *string                   `json:"country,omitempty"`
	IsCreatedByRecommendedDfca *bool                     `json:"is_created_by_recommended_dfca,omitempty"`
	IsFinancialService         *bool                     `json:"is_financial_service,omitempty"`
	IsParentLal                *bool                     `json:"is_parent_lal,omitempty"`
	Origin                     *[]map[string]interface{} `json:"origin,omitempty"`
	OriginEventName            *string                   `json:"origin_event_name,omitempty"`
	OriginEventSourceName      *string                   `json:"origin_event_source_name,omitempty"`
	OriginEventSourceType      *string                   `json:"origin_event_source_type,omitempty"`
	ProductSetName             *string                   `json:"product_set_name,omitempty"`
	Ratio                      *float64                  `json:"ratio,omitempty"`
	StartingRatio              *float64                  `json:"starting_ratio,omitempty"`
	TargetCountries            *[]string                 `json:"target_countries,omitempty"`
	TargetCountryNames         *[]interface{}            `json:"target_country_names,omitempty"`
	Type                       *string                   `json:"type,omitempty"`
}
