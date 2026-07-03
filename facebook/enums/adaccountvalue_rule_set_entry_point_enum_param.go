package enums

type AdaccountvalueRuleSetEntryPointEnumParam string

const (
	AdaccountvalueRuleSetEntryPointEnumParamAdvertisingSettings  AdaccountvalueRuleSetEntryPointEnumParam = "ADVERTISING_SETTINGS"
	AdaccountvalueRuleSetEntryPointEnumParamL2ConversionLocation AdaccountvalueRuleSetEntryPointEnumParam = "L2_CONVERSION_LOCATION"
	AdaccountvalueRuleSetEntryPointEnumParamL2Global             AdaccountvalueRuleSetEntryPointEnumParam = "L2_GLOBAL"
	AdaccountvalueRuleSetEntryPointEnumParamL2NcaGoal            AdaccountvalueRuleSetEntryPointEnumParam = "L2_NCA_GOAL"
	AdaccountvalueRuleSetEntryPointEnumParamL2Placement          AdaccountvalueRuleSetEntryPointEnumParam = "L2_PLACEMENT"
)

func (value AdaccountvalueRuleSetEntryPointEnumParam) String() string {
	return string(value)
}
