package enums

type GroupliveVideosStreamTypeEnumParam string

const (
	GroupliveVideosStreamTypeEnumParamAmbient GroupliveVideosStreamTypeEnumParam = "AMBIENT"
	GroupliveVideosStreamTypeEnumParamRegular GroupliveVideosStreamTypeEnumParam = "REGULAR"
)

func (value GroupliveVideosStreamTypeEnumParam) String() string {
	return string(value)
}
