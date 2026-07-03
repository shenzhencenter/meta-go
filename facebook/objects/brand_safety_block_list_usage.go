package objects

type BrandSafetyBlockListUsage struct {
	CurrentUsage *int    `json:"current_usage,omitempty"`
	NewUsage     *int    `json:"new_usage,omitempty"`
	Platform     *string `json:"platform,omitempty"`
	Position     *string `json:"position,omitempty"`
	Threshold    *int    `json:"threshold,omitempty"`
}
