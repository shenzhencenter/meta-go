package enums

type AdsvalueadjustmentrulecollectionEntryPoint string

const (
	AdsvalueadjustmentrulecollectionEntryPointAdvertisingSettings  AdsvalueadjustmentrulecollectionEntryPoint = "ADVERTISING_SETTINGS"
	AdsvalueadjustmentrulecollectionEntryPointL2ConversionLocation AdsvalueadjustmentrulecollectionEntryPoint = "L2_CONVERSION_LOCATION"
	AdsvalueadjustmentrulecollectionEntryPointL2Global             AdsvalueadjustmentrulecollectionEntryPoint = "L2_GLOBAL"
	AdsvalueadjustmentrulecollectionEntryPointL2NcaGoal            AdsvalueadjustmentrulecollectionEntryPoint = "L2_NCA_GOAL"
	AdsvalueadjustmentrulecollectionEntryPointL2Placement          AdsvalueadjustmentrulecollectionEntryPoint = "L2_PLACEMENT"
)

func (value AdsvalueadjustmentrulecollectionEntryPoint) String() string {
	return string(value)
}
