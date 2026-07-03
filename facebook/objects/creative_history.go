package objects

type CreativeHistory struct {
	CreativeFingerprint *int                      `json:"creative_fingerprint,omitempty"`
	TimeRanges          *[]map[string]interface{} `json:"time_ranges,omitempty"`
}
