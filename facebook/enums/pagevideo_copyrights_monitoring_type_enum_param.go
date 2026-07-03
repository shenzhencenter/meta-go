package enums

type PagevideoCopyrightsMonitoringTypeEnumParam string

const (
	PagevideoCopyrightsMonitoringTypeEnumParamAudioOnly     PagevideoCopyrightsMonitoringTypeEnumParam = "AUDIO_ONLY"
	PagevideoCopyrightsMonitoringTypeEnumParamVideoAndAudio PagevideoCopyrightsMonitoringTypeEnumParam = "VIDEO_AND_AUDIO"
	PagevideoCopyrightsMonitoringTypeEnumParamVideoOnly     PagevideoCopyrightsMonitoringTypeEnumParam = "VIDEO_ONLY"
)

func (value PagevideoCopyrightsMonitoringTypeEnumParam) String() string {
	return string(value)
}
