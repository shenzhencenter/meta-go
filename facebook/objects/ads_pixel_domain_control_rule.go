package objects

type AdsPixelDomainControlRule struct {
	DomainList *[]map[string]interface{} `json:"domain_list,omitempty"`
	Type       *string                   `json:"type,omitempty"`
}

var AdsPixelDomainControlRuleFields = struct {
	DomainList string
	Type       string
}{
	DomainList: "domain_list",
	Type:       "type",
}
