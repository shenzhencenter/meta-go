package objects

type AdsPixelDeliveryRecommendations struct {
	CustomEventType  *string `json:"custom_event_type,omitempty"`
	OptimizationGoal *string `json:"optimization_goal,omitempty"`
}

var AdsPixelDeliveryRecommendationsFields = struct {
	CustomEventType  string
	OptimizationGoal string
}{
	CustomEventType:  "custom_event_type",
	OptimizationGoal: "optimization_goal",
}
