package enums

type EventliveVideosSpatialAudioFormatEnumParam string

const (
	EventliveVideosSpatialAudioFormatEnumParamAmbixX4 EventliveVideosSpatialAudioFormatEnumParam = "ambiX_4"
)

func (value EventliveVideosSpatialAudioFormatEnumParam) String() string {
	return string(value)
}
