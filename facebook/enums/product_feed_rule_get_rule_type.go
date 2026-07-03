package enums

type ProductfeedrulegetRuleType string

const (
	ProductfeedrulegetRuleTypeFallbackRule     ProductfeedrulegetRuleType = "FALLBACK_RULE"
	ProductfeedrulegetRuleTypeLetterCaseRule   ProductfeedrulegetRuleType = "LETTER_CASE_RULE"
	ProductfeedrulegetRuleTypeMappingRule      ProductfeedrulegetRuleType = "MAPPING_RULE"
	ProductfeedrulegetRuleTypeRegexReplaceRule ProductfeedrulegetRuleType = "REGEX_REPLACE_RULE"
	ProductfeedrulegetRuleTypeValueMappingRule ProductfeedrulegetRuleType = "VALUE_MAPPING_RULE"
)

func (value ProductfeedrulegetRuleType) String() string {
	return string(value)
}
