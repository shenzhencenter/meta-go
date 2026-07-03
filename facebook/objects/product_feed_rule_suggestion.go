package objects

type ProductFeedRuleSuggestion struct {
	Attribute *string              `json:"attribute,omitempty"`
	Params    *[]map[string]string `json:"params,omitempty"`
	Type      *string              `json:"type,omitempty"`
}

var ProductFeedRuleSuggestionFields = struct {
	Attribute string
	Params    string
	Type      string
}{
	Attribute: "attribute",
	Params:    "params",
	Type:      "type",
}
