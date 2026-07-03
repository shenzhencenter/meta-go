package objects

type PageMessageResponsivenessMetrics struct {
	IsVeryResponsive *bool    `json:"is_very_responsive,omitempty"`
	ResponseRate     *float64 `json:"response_rate,omitempty"`
	ResponseTime     *float64 `json:"response_time,omitempty"`
}
