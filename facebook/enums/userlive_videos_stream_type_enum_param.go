package enums

type UserliveVideosStreamTypeEnumParam string

const (
	UserliveVideosStreamTypeEnumParamAmbient UserliveVideosStreamTypeEnumParam = "AMBIENT"
	UserliveVideosStreamTypeEnumParamRegular UserliveVideosStreamTypeEnumParam = "REGULAR"
)

func (value UserliveVideosStreamTypeEnumParam) String() string {
	return string(value)
}
