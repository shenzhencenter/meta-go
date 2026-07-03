package objects

type TargetingAutomation struct {
	AdvantageAudience       *uint64                 `json:"advantage_audience,omitempty"`
	CreativeAudiencePairing *uint64                 `json:"creative_audience_pairing,omitempty"`
	IndividualSetting       *map[string]interface{} `json:"individual_setting,omitempty"`
	SharedAudiences         *uint64                 `json:"shared_audiences,omitempty"`
	ValueExpression         *uint64                 `json:"value_expression,omitempty"`
}

var TargetingAutomationFields = struct {
	AdvantageAudience       string
	CreativeAudiencePairing string
	IndividualSetting       string
	SharedAudiences         string
	ValueExpression         string
}{
	AdvantageAudience:       "advantage_audience",
	CreativeAudiencePairing: "creative_audience_pairing",
	IndividualSetting:       "individual_setting",
	SharedAudiences:         "shared_audiences",
	ValueExpression:         "value_expression",
}
