package objects

type MessengerProfile struct {
	AccountLinkingURL          *string                   `json:"account_linking_url,omitempty"`
	Commands                   *[]map[string]interface{} `json:"commands,omitempty"`
	GetStarted                 *map[string]interface{}   `json:"get_started,omitempty"`
	IceBreakers                *[]map[string]interface{} `json:"ice_breakers,omitempty"`
	PersistentMenu             *[]map[string]interface{} `json:"persistent_menu,omitempty"`
	SubjectToNewEuPrivacyRules *bool                     `json:"subject_to_new_eu_privacy_rules,omitempty"`
	WhitelistedDomains         *[]string                 `json:"whitelisted_domains,omitempty"`
}

var MessengerProfileFields = struct {
	AccountLinkingURL          string
	Commands                   string
	GetStarted                 string
	IceBreakers                string
	PersistentMenu             string
	SubjectToNewEuPrivacyRules string
	WhitelistedDomains         string
}{
	AccountLinkingURL:          "account_linking_url",
	Commands:                   "commands",
	GetStarted:                 "get_started",
	IceBreakers:                "ice_breakers",
	PersistentMenu:             "persistent_menu",
	SubjectToNewEuPrivacyRules: "subject_to_new_eu_privacy_rules",
	WhitelistedDomains:         "whitelisted_domains",
}
