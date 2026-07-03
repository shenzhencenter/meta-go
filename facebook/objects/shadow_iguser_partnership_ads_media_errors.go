package objects

type ShadowIGUserPartnershipAdsMediaErrors struct {
	AdCode     *string   `json:"ad_code,omitempty"`
	ErrorCodes *string   `json:"error_codes,omitempty"`
	Errors     *[]string `json:"errors,omitempty"`
	Permalink  *string   `json:"permalink,omitempty"`
}
