package objects

type ColumnSuggestions struct {
	Explanations      *map[string]interface{} `json:"explanations,omitempty"`
	Format            *[]string               `json:"format,omitempty"`
	Objective         *[]string               `json:"objective,omitempty"`
	OptimizationGoals *[]string               `json:"optimization_goals,omitempty"`
}
