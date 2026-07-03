package enums

type PostreactionsTypeEnumParam string

const (
	PostreactionsTypeEnumParamAngry    PostreactionsTypeEnumParam = "ANGRY"
	PostreactionsTypeEnumParamCare     PostreactionsTypeEnumParam = "CARE"
	PostreactionsTypeEnumParamFire     PostreactionsTypeEnumParam = "FIRE"
	PostreactionsTypeEnumParamHaha     PostreactionsTypeEnumParam = "HAHA"
	PostreactionsTypeEnumParamHundred  PostreactionsTypeEnumParam = "HUNDRED"
	PostreactionsTypeEnumParamLike     PostreactionsTypeEnumParam = "LIKE"
	PostreactionsTypeEnumParamLove     PostreactionsTypeEnumParam = "LOVE"
	PostreactionsTypeEnumParamNone     PostreactionsTypeEnumParam = "NONE"
	PostreactionsTypeEnumParamPride    PostreactionsTypeEnumParam = "PRIDE"
	PostreactionsTypeEnumParamSad      PostreactionsTypeEnumParam = "SAD"
	PostreactionsTypeEnumParamThankful PostreactionsTypeEnumParam = "THANKFUL"
	PostreactionsTypeEnumParamWow      PostreactionsTypeEnumParam = "WOW"
)

func (value PostreactionsTypeEnumParam) String() string {
	return string(value)
}
