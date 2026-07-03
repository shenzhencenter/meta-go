package enums

type PagevideosSwapModeEnumParam string

const (
	PagevideosSwapModeEnumParamReplace PagevideosSwapModeEnumParam = "replace"
)

func (value PagevideosSwapModeEnumParam) String() string {
	return string(value)
}
