package objects

type CTXDefaultOptimizationConfig struct {
	DefaultingSource *string `json:"defaulting_source,omitempty"`
	DestinationType  *string `json:"destination_type,omitempty"`
	Objective        *string `json:"objective,omitempty"`
	OptimizationGoal *string `json:"optimization_goal,omitempty"`
}

var CTXDefaultOptimizationConfigFields = struct {
	DefaultingSource string
	DestinationType  string
	Objective        string
	OptimizationGoal string
}{
	DefaultingSource: "defaulting_source",
	DestinationType:  "destination_type",
	Objective:        "objective",
	OptimizationGoal: "optimization_goal",
}
