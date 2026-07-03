package objects

type ScimCompanyUserPhoneNumber struct {
	Number  *string `json:"number,omitempty"`
	Primary *bool   `json:"primary,omitempty"`
	Type    *string `json:"type,omitempty"`
}
