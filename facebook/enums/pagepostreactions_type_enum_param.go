package enums

type PagepostreactionsTypeEnumParam string

const (
	PagepostreactionsTypeEnumParamAngry    PagepostreactionsTypeEnumParam = "ANGRY"
	PagepostreactionsTypeEnumParamCare     PagepostreactionsTypeEnumParam = "CARE"
	PagepostreactionsTypeEnumParamFire     PagepostreactionsTypeEnumParam = "FIRE"
	PagepostreactionsTypeEnumParamHaha     PagepostreactionsTypeEnumParam = "HAHA"
	PagepostreactionsTypeEnumParamHundred  PagepostreactionsTypeEnumParam = "HUNDRED"
	PagepostreactionsTypeEnumParamLike     PagepostreactionsTypeEnumParam = "LIKE"
	PagepostreactionsTypeEnumParamLove     PagepostreactionsTypeEnumParam = "LOVE"
	PagepostreactionsTypeEnumParamNone     PagepostreactionsTypeEnumParam = "NONE"
	PagepostreactionsTypeEnumParamPride    PagepostreactionsTypeEnumParam = "PRIDE"
	PagepostreactionsTypeEnumParamSad      PagepostreactionsTypeEnumParam = "SAD"
	PagepostreactionsTypeEnumParamThankful PagepostreactionsTypeEnumParam = "THANKFUL"
	PagepostreactionsTypeEnumParamWow      PagepostreactionsTypeEnumParam = "WOW"
)

func (value PagepostreactionsTypeEnumParam) String() string {
	return string(value)
}
