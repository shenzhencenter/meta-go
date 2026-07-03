package enums

type AdaccountadPlacePageSetsAsyncLocationTypesEnumParam string

const (
	AdaccountadPlacePageSetsAsyncLocationTypesEnumParamHome   AdaccountadPlacePageSetsAsyncLocationTypesEnumParam = "home"
	AdaccountadPlacePageSetsAsyncLocationTypesEnumParamRecent AdaccountadPlacePageSetsAsyncLocationTypesEnumParam = "recent"
)

func (value AdaccountadPlacePageSetsAsyncLocationTypesEnumParam) String() string {
	return string(value)
}
