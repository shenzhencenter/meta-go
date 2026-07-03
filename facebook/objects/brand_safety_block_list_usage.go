package objects

type BrandSafetyBlockListUsage struct {
	CurrentUsage *int    `json:"current_usage,omitempty"`
	NewUsage     *int    `json:"new_usage,omitempty"`
	Platform     *string `json:"platform,omitempty"`
	Position     *string `json:"position,omitempty"`
	Threshold    *int    `json:"threshold,omitempty"`
}

var BrandSafetyBlockListUsageFields = struct {
	CurrentUsage string
	NewUsage     string
	Platform     string
	Position     string
	Threshold    string
}{
	CurrentUsage: "current_usage",
	NewUsage:     "new_usage",
	Platform:     "platform",
	Position:     "position",
	Threshold:    "threshold",
}
