package enums

type AdaccountadPlacePageSetsLocationTypesEnumParam string

const (
	AdaccountadPlacePageSetsLocationTypesEnumParamHome   AdaccountadPlacePageSetsLocationTypesEnumParam = "home"
	AdaccountadPlacePageSetsLocationTypesEnumParamRecent AdaccountadPlacePageSetsLocationTypesEnumParam = "recent"
)

func (value AdaccountadPlacePageSetsLocationTypesEnumParam) String() string {
	return string(value)
}
