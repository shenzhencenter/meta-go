package enums

type GroupliveVideosSpatialAudioFormatEnumParam string

const (
	GroupliveVideosSpatialAudioFormatEnumParamAmbixX4 GroupliveVideosSpatialAudioFormatEnumParam = "ambiX_4"
)

func (value GroupliveVideosSpatialAudioFormatEnumParam) String() string {
	return string(value)
}
