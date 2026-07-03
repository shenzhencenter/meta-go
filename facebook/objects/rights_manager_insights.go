package objects

type RightsManagerInsights struct {
	Error          *string                                `json:"error,omitempty"`
	ErrorMessage   *string                                `json:"error_message,omitempty"`
	Metadata       *[]map[string]map[string]interface{}   `json:"metadata,omitempty"`
	Totals         *[]map[string]map[string]interface{}   `json:"totals,omitempty"`
	XAxisBreakdown *[][]map[string]map[string]interface{} `json:"x_axis_breakdown,omitempty"`
}
