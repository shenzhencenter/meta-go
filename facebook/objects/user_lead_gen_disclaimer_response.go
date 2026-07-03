package objects

type UserLeadGenDisclaimerResponse struct {
	CheckboxKey *string `json:"checkbox_key,omitempty"`
	IsChecked   *string `json:"is_checked,omitempty"`
}

var UserLeadGenDisclaimerResponseFields = struct {
	CheckboxKey string
	IsChecked   string
}{
	CheckboxKey: "checkbox_key",
	IsChecked:   "is_checked",
}
