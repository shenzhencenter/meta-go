package enums

type ShadowiguserinsightsBreakdownEnumParam string

const (
	ShadowiguserinsightsBreakdownEnumParamAge               ShadowiguserinsightsBreakdownEnumParam = "age"
	ShadowiguserinsightsBreakdownEnumParamCity              ShadowiguserinsightsBreakdownEnumParam = "city"
	ShadowiguserinsightsBreakdownEnumParamContactButtonType ShadowiguserinsightsBreakdownEnumParam = "contact_button_type"
	ShadowiguserinsightsBreakdownEnumParamCountry           ShadowiguserinsightsBreakdownEnumParam = "country"
	ShadowiguserinsightsBreakdownEnumParamFollowType        ShadowiguserinsightsBreakdownEnumParam = "follow_type"
	ShadowiguserinsightsBreakdownEnumParamGender            ShadowiguserinsightsBreakdownEnumParam = "gender"
	ShadowiguserinsightsBreakdownEnumParamMediaProductType  ShadowiguserinsightsBreakdownEnumParam = "media_product_type"
)

func (value ShadowiguserinsightsBreakdownEnumParam) String() string {
	return string(value)
}
