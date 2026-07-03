package objects

type MinimumBudget struct {
	Currency                 *string `json:"currency,omitempty"`
	MinDailyBudgetHighFreq   *int    `json:"min_daily_budget_high_freq,omitempty"`
	MinDailyBudgetImp        *int    `json:"min_daily_budget_imp,omitempty"`
	MinDailyBudgetLowFreq    *int    `json:"min_daily_budget_low_freq,omitempty"`
	MinDailyBudgetVideoViews *int    `json:"min_daily_budget_video_views,omitempty"`
}

var MinimumBudgetFields = struct {
	Currency                 string
	MinDailyBudgetHighFreq   string
	MinDailyBudgetImp        string
	MinDailyBudgetLowFreq    string
	MinDailyBudgetVideoViews string
}{
	Currency:                 "currency",
	MinDailyBudgetHighFreq:   "min_daily_budget_high_freq",
	MinDailyBudgetImp:        "min_daily_budget_imp",
	MinDailyBudgetLowFreq:    "min_daily_budget_low_freq",
	MinDailyBudgetVideoViews: "min_daily_budget_video_views",
}
