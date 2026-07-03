package enums

type PageliveVideosStreamTypeEnumParam string

const (
	PageliveVideosStreamTypeEnumParamAmbient PageliveVideosStreamTypeEnumParam = "AMBIENT"
	PageliveVideosStreamTypeEnumParamRegular PageliveVideosStreamTypeEnumParam = "REGULAR"
)

func (value PageliveVideosStreamTypeEnumParam) String() string {
	return string(value)
}
