package objects

type ColumnSuggestions struct {
	Explanations      *map[string]interface{} `json:"explanations,omitempty"`
	Format            *[]string               `json:"format,omitempty"`
	Objective         *[]string               `json:"objective,omitempty"`
	OptimizationGoals *[]string               `json:"optimization_goals,omitempty"`
}

var ColumnSuggestionsFields = struct {
	Explanations      string
	Format            string
	Objective         string
	OptimizationGoals string
}{
	Explanations:      "explanations",
	Format:            "format",
	Objective:         "objective",
	OptimizationGoals: "optimization_goals",
}
