package enums

type LivevideoreactionsTypeEnumParam string

const (
	LivevideoreactionsTypeEnumParamAngry    LivevideoreactionsTypeEnumParam = "ANGRY"
	LivevideoreactionsTypeEnumParamCare     LivevideoreactionsTypeEnumParam = "CARE"
	LivevideoreactionsTypeEnumParamFire     LivevideoreactionsTypeEnumParam = "FIRE"
	LivevideoreactionsTypeEnumParamHaha     LivevideoreactionsTypeEnumParam = "HAHA"
	LivevideoreactionsTypeEnumParamHundred  LivevideoreactionsTypeEnumParam = "HUNDRED"
	LivevideoreactionsTypeEnumParamLike     LivevideoreactionsTypeEnumParam = "LIKE"
	LivevideoreactionsTypeEnumParamLove     LivevideoreactionsTypeEnumParam = "LOVE"
	LivevideoreactionsTypeEnumParamNone     LivevideoreactionsTypeEnumParam = "NONE"
	LivevideoreactionsTypeEnumParamPride    LivevideoreactionsTypeEnumParam = "PRIDE"
	LivevideoreactionsTypeEnumParamSad      LivevideoreactionsTypeEnumParam = "SAD"
	LivevideoreactionsTypeEnumParamThankful LivevideoreactionsTypeEnumParam = "THANKFUL"
	LivevideoreactionsTypeEnumParamWow      LivevideoreactionsTypeEnumParam = "WOW"
)

func (value LivevideoreactionsTypeEnumParam) String() string {
	return string(value)
}
