package enums

type AdaccountadspixelsSortByEnumParam string

const (
	AdaccountadspixelsSortByEnumParamLastFiredTime AdaccountadspixelsSortByEnumParam = "LAST_FIRED_TIME"
	AdaccountadspixelsSortByEnumParamName          AdaccountadspixelsSortByEnumParam = "NAME"
)

func (value AdaccountadspixelsSortByEnumParam) String() string {
	return string(value)
}
