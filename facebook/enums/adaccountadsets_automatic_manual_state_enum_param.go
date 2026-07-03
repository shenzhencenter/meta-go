package enums

type AdaccountadsetsAutomaticManualStateEnumParam string

const (
	AdaccountadsetsAutomaticManualStateEnumParamAutomatic AdaccountadsetsAutomaticManualStateEnumParam = "AUTOMATIC"
	AdaccountadsetsAutomaticManualStateEnumParamManual    AdaccountadsetsAutomaticManualStateEnumParam = "MANUAL"
	AdaccountadsetsAutomaticManualStateEnumParamUnset     AdaccountadsetsAutomaticManualStateEnumParam = "UNSET"
)

func (value AdaccountadsetsAutomaticManualStateEnumParam) String() string {
	return string(value)
}
