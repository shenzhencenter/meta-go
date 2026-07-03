package enums

type PostBackdatedTimeGranularity string

const (
	PostBackdatedTimeGranularityDay   PostBackdatedTimeGranularity = "day"
	PostBackdatedTimeGranularityHour  PostBackdatedTimeGranularity = "hour"
	PostBackdatedTimeGranularityMin   PostBackdatedTimeGranularity = "min"
	PostBackdatedTimeGranularityMonth PostBackdatedTimeGranularity = "month"
	PostBackdatedTimeGranularityNone  PostBackdatedTimeGranularity = "none"
	PostBackdatedTimeGranularityYear  PostBackdatedTimeGranularity = "year"
)

func (value PostBackdatedTimeGranularity) String() string {
	return string(value)
}
