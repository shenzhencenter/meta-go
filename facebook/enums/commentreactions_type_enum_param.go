package enums

type CommentreactionsTypeEnumParam string

const (
	CommentreactionsTypeEnumParamAngry    CommentreactionsTypeEnumParam = "ANGRY"
	CommentreactionsTypeEnumParamCare     CommentreactionsTypeEnumParam = "CARE"
	CommentreactionsTypeEnumParamFire     CommentreactionsTypeEnumParam = "FIRE"
	CommentreactionsTypeEnumParamHaha     CommentreactionsTypeEnumParam = "HAHA"
	CommentreactionsTypeEnumParamHundred  CommentreactionsTypeEnumParam = "HUNDRED"
	CommentreactionsTypeEnumParamLike     CommentreactionsTypeEnumParam = "LIKE"
	CommentreactionsTypeEnumParamLove     CommentreactionsTypeEnumParam = "LOVE"
	CommentreactionsTypeEnumParamNone     CommentreactionsTypeEnumParam = "NONE"
	CommentreactionsTypeEnumParamPride    CommentreactionsTypeEnumParam = "PRIDE"
	CommentreactionsTypeEnumParamSad      CommentreactionsTypeEnumParam = "SAD"
	CommentreactionsTypeEnumParamThankful CommentreactionsTypeEnumParam = "THANKFUL"
	CommentreactionsTypeEnumParamWow      CommentreactionsTypeEnumParam = "WOW"
)

func (value CommentreactionsTypeEnumParam) String() string {
	return string(value)
}
