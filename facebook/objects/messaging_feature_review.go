package objects

type MessagingFeatureReview struct {
	Feature *string `json:"feature,omitempty"`
	Status  *string `json:"status,omitempty"`
}

var MessagingFeatureReviewFields = struct {
	Feature string
	Status  string
}{
	Feature: "feature",
	Status:  "status",
}
