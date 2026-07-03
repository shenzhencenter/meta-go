package enums

type AdaccountasyncadcreativesNotificationModeEnumParam string

const (
	AdaccountasyncadcreativesNotificationModeEnumParamOff        AdaccountasyncadcreativesNotificationModeEnumParam = "OFF"
	AdaccountasyncadcreativesNotificationModeEnumParamOnComplete AdaccountasyncadcreativesNotificationModeEnumParam = "ON_COMPLETE"
)

func (value AdaccountasyncadcreativesNotificationModeEnumParam) String() string {
	return string(value)
}
