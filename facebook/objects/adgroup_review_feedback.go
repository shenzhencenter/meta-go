package objects

type AdgroupReviewFeedback struct {
	Global            *map[string]string                      `json:"global,omitempty"`
	PlacementSpecific *AdgroupPlacementSpecificReviewFeedback `json:"placement_specific,omitempty"`
}

var AdgroupReviewFeedbackFields = struct {
	Global            string
	PlacementSpecific string
}{
	Global:            "global",
	PlacementSpecific: "placement_specific",
}
