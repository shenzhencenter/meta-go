package objects

type TargetingProductAudienceSubSpec struct {
	MinRetentionSeconds *string `json:"min_retention_seconds,omitempty"`
	RetentionSeconds    *string `json:"retention_seconds,omitempty"`
	Rule                *string `json:"rule,omitempty"`
}

var TargetingProductAudienceSubSpecFields = struct {
	MinRetentionSeconds string
	RetentionSeconds    string
	Rule                string
}{
	MinRetentionSeconds: "min_retention_seconds",
	RetentionSeconds:    "retention_seconds",
	Rule:                "rule",
}
