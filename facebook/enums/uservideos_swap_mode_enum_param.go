package enums

type UservideosSwapModeEnumParam string

const (
	UservideosSwapModeEnumParamReplace UservideosSwapModeEnumParam = "replace"
)

func (value UservideosSwapModeEnumParam) String() string {
	return string(value)
}
