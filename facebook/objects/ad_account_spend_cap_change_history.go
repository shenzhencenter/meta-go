package objects

type AdAccountSpendCapChangeHistory struct {
	Action    *string `json:"action,omitempty"`
	SpendCap  *int    `json:"spend_cap,omitempty"`
	TimeStart *string `json:"time_start,omitempty"`
	TimeStop  *string `json:"time_stop,omitempty"`
}
