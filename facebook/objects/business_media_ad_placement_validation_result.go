package objects

type BusinessMediaAdPlacementValidationResult struct {
	AdPlacement      *string   `json:"ad_placement,omitempty"`
	AdPlacementLabel *string   `json:"ad_placement_label,omitempty"`
	ErrorMessages    *[]string `json:"error_messages,omitempty"`
	IsValid          *bool     `json:"is_valid,omitempty"`
}
