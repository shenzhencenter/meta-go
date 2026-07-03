package objects

type MessagingFeatureReview struct {
	Feature *string `json:"feature,omitempty"`
	Status  *string `json:"status,omitempty"`
}
