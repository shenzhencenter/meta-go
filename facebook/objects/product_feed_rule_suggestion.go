package objects

type ProductFeedRuleSuggestion struct {
	Attribute *string              `json:"attribute,omitempty"`
	Params    *[]map[string]string `json:"params,omitempty"`
	Type      *string              `json:"type,omitempty"`
}
