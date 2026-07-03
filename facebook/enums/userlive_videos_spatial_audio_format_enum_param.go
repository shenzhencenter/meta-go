package enums

type UserliveVideosSpatialAudioFormatEnumParam string

const (
	UserliveVideosSpatialAudioFormatEnumParamAmbixX4 UserliveVideosSpatialAudioFormatEnumParam = "ambiX_4"
)

func (value UserliveVideosSpatialAudioFormatEnumParam) String() string {
	return string(value)
}
