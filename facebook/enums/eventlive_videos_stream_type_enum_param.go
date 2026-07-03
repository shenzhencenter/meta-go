package enums

type EventliveVideosStreamTypeEnumParam string

const (
	EventliveVideosStreamTypeEnumParamAmbient EventliveVideosStreamTypeEnumParam = "AMBIENT"
	EventliveVideosStreamTypeEnumParamRegular EventliveVideosStreamTypeEnumParam = "REGULAR"
)

func (value EventliveVideosStreamTypeEnumParam) String() string {
	return string(value)
}
