package objects

type AdAccountReachEstimate struct {
	EstimateReady   *bool `json:"estimate_ready,omitempty"`
	UsersLowerBound *int  `json:"users_lower_bound,omitempty"`
	UsersUpperBound *int  `json:"users_upper_bound,omitempty"`
}

var AdAccountReachEstimateFields = struct {
	EstimateReady   string
	UsersLowerBound string
	UsersUpperBound string
}{
	EstimateReady:   "estimate_ready",
	UsersLowerBound: "users_lower_bound",
	UsersUpperBound: "users_upper_bound",
}
