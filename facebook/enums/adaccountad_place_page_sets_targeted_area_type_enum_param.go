package enums

type AdaccountadPlacePageSetsTargetedAreaTypeEnumParam string

const (
	AdaccountadPlacePageSetsTargetedAreaTypeEnumParamCustomRadius  AdaccountadPlacePageSetsTargetedAreaTypeEnumParam = "CUSTOM_RADIUS"
	AdaccountadPlacePageSetsTargetedAreaTypeEnumParamMarketingArea AdaccountadPlacePageSetsTargetedAreaTypeEnumParam = "MARKETING_AREA"
	AdaccountadPlacePageSetsTargetedAreaTypeEnumParamNone          AdaccountadPlacePageSetsTargetedAreaTypeEnumParam = "NONE"
)

func (value AdaccountadPlacePageSetsTargetedAreaTypeEnumParam) String() string {
	return string(value)
}
