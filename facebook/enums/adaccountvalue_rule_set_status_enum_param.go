package enums

type AdaccountvalueRuleSetStatusEnumParam string

const (
	AdaccountvalueRuleSetStatusEnumParamActive  AdaccountvalueRuleSetStatusEnumParam = "ACTIVE"
	AdaccountvalueRuleSetStatusEnumParamDeleted AdaccountvalueRuleSetStatusEnumParam = "DELETED"
	AdaccountvalueRuleSetStatusEnumParamDraft   AdaccountvalueRuleSetStatusEnumParam = "DRAFT"
)

func (value AdaccountvalueRuleSetStatusEnumParam) String() string {
	return string(value)
}
