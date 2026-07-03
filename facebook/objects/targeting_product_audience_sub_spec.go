package objects

type TargetingProductAudienceSubSpec struct {
	MinRetentionSeconds *string `json:"min_retention_seconds,omitempty"`
	RetentionSeconds    *string `json:"retention_seconds,omitempty"`
	Rule                *string `json:"rule,omitempty"`
}
