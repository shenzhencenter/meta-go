package enums

type ProductfeedrulesRuleTypeEnumParam string

const (
	ProductfeedrulesRuleTypeEnumParamFallbackRule     ProductfeedrulesRuleTypeEnumParam = "fallback_rule"
	ProductfeedrulesRuleTypeEnumParamLetterCaseRule   ProductfeedrulesRuleTypeEnumParam = "letter_case_rule"
	ProductfeedrulesRuleTypeEnumParamMappingRule      ProductfeedrulesRuleTypeEnumParam = "mapping_rule"
	ProductfeedrulesRuleTypeEnumParamRegexReplaceRule ProductfeedrulesRuleTypeEnumParam = "regex_replace_rule"
	ProductfeedrulesRuleTypeEnumParamValueMappingRule ProductfeedrulesRuleTypeEnumParam = "value_mapping_rule"
)

func (value ProductfeedrulesRuleTypeEnumParam) String() string {
	return string(value)
}
