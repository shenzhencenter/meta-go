package objects

type AdBidAdjustments struct {
	AgeRange   *map[string]float64     `json:"age_range,omitempty"`
	PageTypes  *map[string]interface{} `json:"page_types,omitempty"`
	UserGroups *string                 `json:"user_groups,omitempty"`
}

var AdBidAdjustmentsFields = struct {
	AgeRange   string
	PageTypes  string
	UserGroups string
}{
	AgeRange:   "age_range",
	PageTypes:  "page_types",
	UserGroups: "user_groups",
}
