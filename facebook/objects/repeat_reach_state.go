package objects

type RepeatReachState struct {
	CurrentSaturationLevel    *float64 `json:"current_saturation_level,omitempty"`
	ForecastedSaturationLevel *float64 `json:"forecasted_saturation_level,omitempty"`
	HighSaturationThreshold   *float64 `json:"high_saturation_threshold,omitempty"`
	ShouldDisplayCpr          *bool    `json:"should_display_cpr,omitempty"`
}

var RepeatReachStateFields = struct {
	CurrentSaturationLevel    string
	ForecastedSaturationLevel string
	HighSaturationThreshold   string
	ShouldDisplayCpr          string
}{
	CurrentSaturationLevel:    "current_saturation_level",
	ForecastedSaturationLevel: "forecasted_saturation_level",
	HighSaturationThreshold:   "high_saturation_threshold",
	ShouldDisplayCpr:          "should_display_cpr",
}
