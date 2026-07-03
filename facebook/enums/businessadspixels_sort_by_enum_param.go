package enums

type BusinessadspixelsSortByEnumParam string

const (
	BusinessadspixelsSortByEnumParamLastFiredTime BusinessadspixelsSortByEnumParam = "LAST_FIRED_TIME"
	BusinessadspixelsSortByEnumParamName          BusinessadspixelsSortByEnumParam = "NAME"
)

func (value BusinessadspixelsSortByEnumParam) String() string {
	return string(value)
}
