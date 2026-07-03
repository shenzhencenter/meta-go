package enums

type AdaccountadPlacePageSetsAsyncTargetedAreaTypeEnumParam string

const (
	AdaccountadPlacePageSetsAsyncTargetedAreaTypeEnumParamCustomRadius  AdaccountadPlacePageSetsAsyncTargetedAreaTypeEnumParam = "CUSTOM_RADIUS"
	AdaccountadPlacePageSetsAsyncTargetedAreaTypeEnumParamMarketingArea AdaccountadPlacePageSetsAsyncTargetedAreaTypeEnumParam = "MARKETING_AREA"
	AdaccountadPlacePageSetsAsyncTargetedAreaTypeEnumParamNone          AdaccountadPlacePageSetsAsyncTargetedAreaTypeEnumParam = "NONE"
)

func (value AdaccountadPlacePageSetsAsyncTargetedAreaTypeEnumParam) String() string {
	return string(value)
}
