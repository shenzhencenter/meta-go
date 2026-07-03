package objects

type TargetingRelaxation struct {
	CustomAudience *uint64 `json:"custom_audience,omitempty"`
	Lookalike      *uint64 `json:"lookalike,omitempty"`
}

var TargetingRelaxationFields = struct {
	CustomAudience string
	Lookalike      string
}{
	CustomAudience: "custom_audience",
	Lookalike:      "lookalike",
}
