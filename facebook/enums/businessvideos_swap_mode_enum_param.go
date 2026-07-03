package enums

type BusinessvideosSwapModeEnumParam string

const (
	BusinessvideosSwapModeEnumParamReplace BusinessvideosSwapModeEnumParam = "replace"
)

func (value BusinessvideosSwapModeEnumParam) String() string {
	return string(value)
}
