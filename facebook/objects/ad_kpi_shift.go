package objects

type AdKpiShift struct {
	AdSet               *AdSet   `json:"ad_set,omitempty"`
	CostPerResultShift  *float64 `json:"cost_per_result_shift,omitempty"`
	EnoughEffectiveDays *bool    `json:"enough_effective_days,omitempty"`
	ResultIndicator     *string  `json:"result_indicator,omitempty"`
	ResultShift         *float64 `json:"result_shift,omitempty"`
	SpendShift          *float64 `json:"spend_shift,omitempty"`
}

var AdKpiShiftFields = struct {
	AdSet               string
	CostPerResultShift  string
	EnoughEffectiveDays string
	ResultIndicator     string
	ResultShift         string
	SpendShift          string
}{
	AdSet:               "ad_set",
	CostPerResultShift:  "cost_per_result_shift",
	EnoughEffectiveDays: "enough_effective_days",
	ResultIndicator:     "result_indicator",
	ResultShift:         "result_shift",
	SpendShift:          "spend_shift",
}
