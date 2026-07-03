package objects

type ReportingAudience struct {
	CustomAudiences             *[]RawCustomAudience `json:"custom_audiences,omitempty"`
	CustomAudiencesURLParamName *string              `json:"custom_audiences_url_param_name,omitempty"`
	CustomAudiencesURLParamType *string              `json:"custom_audiences_url_param_type,omitempty"`
}
