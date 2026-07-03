package objects

type AdAccountAmountSpentHistory struct {
	AmountSpent *int    `json:"amount_spent,omitempty"`
	SpendCap    *int    `json:"spend_cap,omitempty"`
	TimeStart   *string `json:"time_start,omitempty"`
	TimeStop    *string `json:"time_stop,omitempty"`
}

var AdAccountAmountSpentHistoryFields = struct {
	AmountSpent string
	SpendCap    string
	TimeStart   string
	TimeStop    string
}{
	AmountSpent: "amount_spent",
	SpendCap:    "spend_cap",
	TimeStart:   "time_start",
	TimeStop:    "time_stop",
}
