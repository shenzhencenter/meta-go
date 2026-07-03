package enums

type UserinsightsBreakdownEnumParam string

const (
	UserinsightsBreakdownEnumParamAge               UserinsightsBreakdownEnumParam = "age"
	UserinsightsBreakdownEnumParamCity              UserinsightsBreakdownEnumParam = "city"
	UserinsightsBreakdownEnumParamContactButtonType UserinsightsBreakdownEnumParam = "contact_button_type"
	UserinsightsBreakdownEnumParamCountry           UserinsightsBreakdownEnumParam = "country"
	UserinsightsBreakdownEnumParamFollowType        UserinsightsBreakdownEnumParam = "follow_type"
	UserinsightsBreakdownEnumParamGender            UserinsightsBreakdownEnumParam = "gender"
	UserinsightsBreakdownEnumParamMediaProductType  UserinsightsBreakdownEnumParam = "media_product_type"
)

func (value UserinsightsBreakdownEnumParam) String() string {
	return string(value)
}
