package enums

type AdaccountasyncadrequestsetsNotificationModeEnumParam string

const (
	AdaccountasyncadrequestsetsNotificationModeEnumParamOff        AdaccountasyncadrequestsetsNotificationModeEnumParam = "OFF"
	AdaccountasyncadrequestsetsNotificationModeEnumParamOnComplete AdaccountasyncadrequestsetsNotificationModeEnumParam = "ON_COMPLETE"
)

func (value AdaccountasyncadrequestsetsNotificationModeEnumParam) String() string {
	return string(value)
}
