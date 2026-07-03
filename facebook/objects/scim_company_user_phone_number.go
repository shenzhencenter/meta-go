package objects

type ScimCompanyUserPhoneNumber struct {
	Number  *string `json:"number,omitempty"`
	Primary *bool   `json:"primary,omitempty"`
	Type    *string `json:"type,omitempty"`
}

var ScimCompanyUserPhoneNumberFields = struct {
	Number  string
	Primary string
	Type    string
}{
	Number:  "number",
	Primary: "primary",
	Type:    "type",
}
