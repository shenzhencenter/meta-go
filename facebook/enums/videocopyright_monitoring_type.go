package enums

type VideocopyrightMonitoringType string

const (
	VideocopyrightMonitoringTypeAudioOnly     VideocopyrightMonitoringType = "AUDIO_ONLY"
	VideocopyrightMonitoringTypeVideoAndAudio VideocopyrightMonitoringType = "VIDEO_AND_AUDIO"
	VideocopyrightMonitoringTypeVideoOnly     VideocopyrightMonitoringType = "VIDEO_ONLY"
)

func (value VideocopyrightMonitoringType) String() string {
	return string(value)
}
