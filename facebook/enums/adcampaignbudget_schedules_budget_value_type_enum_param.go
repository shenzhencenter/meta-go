package enums

type AdcampaignbudgetSchedulesBudgetValueTypeEnumParam string

const (
	AdcampaignbudgetSchedulesBudgetValueTypeEnumParamAbsolute   AdcampaignbudgetSchedulesBudgetValueTypeEnumParam = "ABSOLUTE"
	AdcampaignbudgetSchedulesBudgetValueTypeEnumParamMultiplier AdcampaignbudgetSchedulesBudgetValueTypeEnumParam = "MULTIPLIER"
)

func (value AdcampaignbudgetSchedulesBudgetValueTypeEnumParam) String() string {
	return string(value)
}
