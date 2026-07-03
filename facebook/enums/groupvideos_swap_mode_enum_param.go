package enums

type GroupvideosSwapModeEnumParam string

const (
	GroupvideosSwapModeEnumParamReplace GroupvideosSwapModeEnumParam = "replace"
)

func (value GroupvideosSwapModeEnumParam) String() string {
	return string(value)
}
