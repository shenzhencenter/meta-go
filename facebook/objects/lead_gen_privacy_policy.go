package objects

type LeadGenPrivacyPolicy struct {
	LinkText *string `json:"link_text,omitempty"`
	URL      *string `json:"url,omitempty"`
}

var LeadGenPrivacyPolicyFields = struct {
	LinkText string
	URL      string
}{
	LinkText: "link_text",
	URL:      "url",
}
