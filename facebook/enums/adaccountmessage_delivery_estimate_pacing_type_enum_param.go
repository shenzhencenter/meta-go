package enums

type AdaccountmessageDeliveryEstimatePacingTypeEnumParam string

const (
	AdaccountmessageDeliveryEstimatePacingTypeEnumParamDayParting            AdaccountmessageDeliveryEstimatePacingTypeEnumParam = "DAY_PARTING"
	AdaccountmessageDeliveryEstimatePacingTypeEnumParamDisabled              AdaccountmessageDeliveryEstimatePacingTypeEnumParam = "DISABLED"
	AdaccountmessageDeliveryEstimatePacingTypeEnumParamNoPacing              AdaccountmessageDeliveryEstimatePacingTypeEnumParam = "NO_PACING"
	AdaccountmessageDeliveryEstimatePacingTypeEnumParamProbabilisticPacing   AdaccountmessageDeliveryEstimatePacingTypeEnumParam = "PROBABILISTIC_PACING"
	AdaccountmessageDeliveryEstimatePacingTypeEnumParamProbabilisticPacingV2 AdaccountmessageDeliveryEstimatePacingTypeEnumParam = "PROBABILISTIC_PACING_V2"
	AdaccountmessageDeliveryEstimatePacingTypeEnumParamStandard              AdaccountmessageDeliveryEstimatePacingTypeEnumParam = "STANDARD"
)

func (value AdaccountmessageDeliveryEstimatePacingTypeEnumParam) String() string {
	return string(value)
}
