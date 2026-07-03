package enums

type AdcampaigngroupbudgetSchedulesBudgetValueTypeEnumParam string

const (
	AdcampaigngroupbudgetSchedulesBudgetValueTypeEnumParamAbsolute   AdcampaigngroupbudgetSchedulesBudgetValueTypeEnumParam = "ABSOLUTE"
	AdcampaigngroupbudgetSchedulesBudgetValueTypeEnumParamMultiplier AdcampaigngroupbudgetSchedulesBudgetValueTypeEnumParam = "MULTIPLIER"
)

func (value AdcampaigngroupbudgetSchedulesBudgetValueTypeEnumParam) String() string {
	return string(value)
}
