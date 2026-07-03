package objects

type CustomAudienceGroup struct {
	AudienceTypeParamName *string `json:"audience_type_param_name,omitempty"`
	ExistingCustomerTag   *string `json:"existing_customer_tag,omitempty"`
	NewCustomerTag        *string `json:"new_customer_tag,omitempty"`
}

var CustomAudienceGroupFields = struct {
	AudienceTypeParamName string
	ExistingCustomerTag   string
	NewCustomerTag        string
}{
	AudienceTypeParamName: "audience_type_param_name",
	ExistingCustomerTag:   "existing_customer_tag",
	NewCustomerTag:        "new_customer_tag",
}
