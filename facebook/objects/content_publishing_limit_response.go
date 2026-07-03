package objects

type ContentPublishingLimitResponse struct {
	Config     *map[string]interface{} `json:"config,omitempty"`
	QuotaUsage *int                    `json:"quota_usage,omitempty"`
}

var ContentPublishingLimitResponseFields = struct {
	Config     string
	QuotaUsage string
}{
	Config:     "config",
	QuotaUsage: "quota_usage",
}
