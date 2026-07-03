package objects

type AdAccountLiveVideoAdvertiser struct {
	IgLvaDefaultDurationS          *int  `json:"ig_lva_default_duration_s,omitempty"`
	IsLvaToggleOn                  *bool `json:"is_lva_toggle_on,omitempty"`
	LvaDefaultBudget               *int  `json:"lva_default_budget,omitempty"`
	LvaDefaultDurationS            *int  `json:"lva_default_duration_s,omitempty"`
	ShouldDefaultCurrentLive       *bool `json:"should_default_current_live,omitempty"`
	ShouldDefaultScheduledLive     *bool `json:"should_default_scheduled_live,omitempty"`
	ShouldDefaultToggleOnFromModel *bool `json:"should_default_toggle_on_from_model,omitempty"`
	ShouldShowLvaToggle            *bool `json:"should_show_lva_toggle,omitempty"`
}

var AdAccountLiveVideoAdvertiserFields = struct {
	IgLvaDefaultDurationS          string
	IsLvaToggleOn                  string
	LvaDefaultBudget               string
	LvaDefaultDurationS            string
	ShouldDefaultCurrentLive       string
	ShouldDefaultScheduledLive     string
	ShouldDefaultToggleOnFromModel string
	ShouldShowLvaToggle            string
}{
	IgLvaDefaultDurationS:          "ig_lva_default_duration_s",
	IsLvaToggleOn:                  "is_lva_toggle_on",
	LvaDefaultBudget:               "lva_default_budget",
	LvaDefaultDurationS:            "lva_default_duration_s",
	ShouldDefaultCurrentLive:       "should_default_current_live",
	ShouldDefaultScheduledLive:     "should_default_scheduled_live",
	ShouldDefaultToggleOnFromModel: "should_default_toggle_on_from_model",
	ShouldShowLvaToggle:            "should_show_lva_toggle",
}
