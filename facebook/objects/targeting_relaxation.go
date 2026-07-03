package objects

type TargetingRelaxation struct {
	CustomAudience *uint64 `json:"custom_audience,omitempty"`
	Lookalike      *uint64 `json:"lookalike,omitempty"`
}
