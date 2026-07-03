package objects

type AdAccountReachEstimate struct {
	EstimateReady   *bool `json:"estimate_ready,omitempty"`
	UsersLowerBound *int  `json:"users_lower_bound,omitempty"`
	UsersUpperBound *int  `json:"users_upper_bound,omitempty"`
}
