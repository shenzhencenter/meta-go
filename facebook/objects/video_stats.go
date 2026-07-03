package objects

type VideoStats struct {
	Aggregate      *[]map[string]map[string]interface{}   `json:"aggregate,omitempty"`
	Error          *string                                `json:"error,omitempty"`
	Metadata       *[]map[string]map[string]interface{}   `json:"metadata,omitempty"`
	TimeSeries     *[]map[string][]map[string]interface{} `json:"time_series,omitempty"`
	Totals         *[]map[string]map[string]interface{}   `json:"totals,omitempty"`
	XAxisBreakdown *[][]map[string]map[string]interface{} `json:"x_axis_breakdown,omitempty"`
}
