package objects

type RightsManagerInsights struct {
	Error          *string                                `json:"error,omitempty"`
	ErrorMessage   *string                                `json:"error_message,omitempty"`
	Metadata       *[]map[string]map[string]interface{}   `json:"metadata,omitempty"`
	Totals         *[]map[string]map[string]interface{}   `json:"totals,omitempty"`
	XAxisBreakdown *[][]map[string]map[string]interface{} `json:"x_axis_breakdown,omitempty"`
}

var RightsManagerInsightsFields = struct {
	Error          string
	ErrorMessage   string
	Metadata       string
	Totals         string
	XAxisBreakdown string
}{
	Error:          "error",
	ErrorMessage:   "error_message",
	Metadata:       "metadata",
	Totals:         "totals",
	XAxisBreakdown: "x_axis_breakdown",
}
