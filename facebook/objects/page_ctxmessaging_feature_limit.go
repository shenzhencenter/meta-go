package objects

type PageCTXMessagingFeatureLimit struct {
	MessagingFeatureLimitDuration *int    `json:"messaging_feature_limit_duration,omitempty"`
	MessagingFeatureLimitType     *string `json:"messaging_feature_limit_type,omitempty"`
	MessagingViolationType        *string `json:"messaging_violation_type,omitempty"`
}

var PageCTXMessagingFeatureLimitFields = struct {
	MessagingFeatureLimitDuration string
	MessagingFeatureLimitType     string
	MessagingViolationType        string
}{
	MessagingFeatureLimitDuration: "messaging_feature_limit_duration",
	MessagingFeatureLimitType:     "messaging_feature_limit_type",
	MessagingViolationType:        "messaging_violation_type",
}
