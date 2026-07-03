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
