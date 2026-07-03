package objects

type AudienceFunnel struct {
	AudienceTypeParamName    *string                `json:"audience_type_param_name,omitempty"`
	AudienceTypeParamTags    *[]map[string]string   `json:"audience_type_param_tags,omitempty"`
	CustomAudienceGroupsInfo *[]map[string][]string `json:"custom_audience_groups_info,omitempty"`
}
