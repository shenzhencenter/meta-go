package objects

type AdAccountAdRulesCountByType struct {
	Count          *int    `json:"count,omitempty"`
	EvaluationType *string `json:"evaluation_type,omitempty"`
}

var AdAccountAdRulesCountByTypeFields = struct {
	Count          string
	EvaluationType string
}{
	Count:          "count",
	EvaluationType: "evaluation_type",
}
