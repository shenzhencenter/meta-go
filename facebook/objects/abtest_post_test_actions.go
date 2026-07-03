package objects

type ABTestPostTestActions struct {
	AutoChangeBudgetToWinner *int    `json:"auto_change_budget_to_winner,omitempty"`
	WinnerBudget             *string `json:"winner_budget,omitempty"`
}

var ABTestPostTestActionsFields = struct {
	AutoChangeBudgetToWinner string
	WinnerBudget             string
}{
	AutoChangeBudgetToWinner: "auto_change_budget_to_winner",
	WinnerBudget:             "winner_budget",
}
