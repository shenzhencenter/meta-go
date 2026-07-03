package objects

type ProductFeedRulePreviewSample struct {
	PropertiesAfter  *[]map[string]string `json:"properties_after,omitempty"`
	PropertiesBefore *[]map[string]string `json:"properties_before,omitempty"`
}

var ProductFeedRulePreviewSampleFields = struct {
	PropertiesAfter  string
	PropertiesBefore string
}{
	PropertiesAfter:  "properties_after",
	PropertiesBefore: "properties_before",
}
