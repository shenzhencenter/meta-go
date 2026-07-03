package enums

type AdaccountvalueRuleSetProductTypeEnumParam string

const (
	AdaccountvalueRuleSetProductTypeEnumParamAudience    AdaccountvalueRuleSetProductTypeEnumParam = "AUDIENCE"
	AdaccountvalueRuleSetProductTypeEnumParamLeadgenAds  AdaccountvalueRuleSetProductTypeEnumParam = "LEADGEN_ADS"
	AdaccountvalueRuleSetProductTypeEnumParamOmniChannel AdaccountvalueRuleSetProductTypeEnumParam = "OMNI_CHANNEL"
)

func (value AdaccountvalueRuleSetProductTypeEnumParam) String() string {
	return string(value)
}
