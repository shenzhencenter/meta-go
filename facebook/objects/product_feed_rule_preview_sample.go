package objects

type ProductFeedRulePreviewSample struct {
	PropertiesAfter  *[]map[string]string `json:"properties_after,omitempty"`
	PropertiesBefore *[]map[string]string `json:"properties_before,omitempty"`
}
