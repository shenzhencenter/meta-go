package objects

type CTXOptimizationEligibility struct {
	Ctm *map[string]interface{} `json:"ctm,omitempty"`
}

var CTXOptimizationEligibilityFields = struct {
	Ctm string
}{
	Ctm: "ctm",
}
