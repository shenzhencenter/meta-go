package enums

type AdcampaignmessageDeliveryEstimatePacingTypeEnumParam string

const (
	AdcampaignmessageDeliveryEstimatePacingTypeEnumParamDayParting            AdcampaignmessageDeliveryEstimatePacingTypeEnumParam = "DAY_PARTING"
	AdcampaignmessageDeliveryEstimatePacingTypeEnumParamDisabled              AdcampaignmessageDeliveryEstimatePacingTypeEnumParam = "DISABLED"
	AdcampaignmessageDeliveryEstimatePacingTypeEnumParamNoPacing              AdcampaignmessageDeliveryEstimatePacingTypeEnumParam = "NO_PACING"
	AdcampaignmessageDeliveryEstimatePacingTypeEnumParamProbabilisticPacing   AdcampaignmessageDeliveryEstimatePacingTypeEnumParam = "PROBABILISTIC_PACING"
	AdcampaignmessageDeliveryEstimatePacingTypeEnumParamProbabilisticPacingV2 AdcampaignmessageDeliveryEstimatePacingTypeEnumParam = "PROBABILISTIC_PACING_V2"
	AdcampaignmessageDeliveryEstimatePacingTypeEnumParamStandard              AdcampaignmessageDeliveryEstimatePacingTypeEnumParam = "STANDARD"
)

func (value AdcampaignmessageDeliveryEstimatePacingTypeEnumParam) String() string {
	return string(value)
}
