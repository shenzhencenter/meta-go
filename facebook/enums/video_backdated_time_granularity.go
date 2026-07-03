package enums

type VideoBackdatedTimeGranularity string

const (
	VideoBackdatedTimeGranularityDay   VideoBackdatedTimeGranularity = "day"
	VideoBackdatedTimeGranularityHour  VideoBackdatedTimeGranularity = "hour"
	VideoBackdatedTimeGranularityMin   VideoBackdatedTimeGranularity = "min"
	VideoBackdatedTimeGranularityMonth VideoBackdatedTimeGranularity = "month"
	VideoBackdatedTimeGranularityNone  VideoBackdatedTimeGranularity = "none"
	VideoBackdatedTimeGranularityYear  VideoBackdatedTimeGranularity = "year"
)

func (value VideoBackdatedTimeGranularity) String() string {
	return string(value)
}
