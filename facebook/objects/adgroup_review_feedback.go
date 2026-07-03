package objects

type AdgroupReviewFeedback struct {
	Global            *map[string]string                      `json:"global,omitempty"`
	PlacementSpecific *AdgroupPlacementSpecificReviewFeedback `json:"placement_specific,omitempty"`
}
