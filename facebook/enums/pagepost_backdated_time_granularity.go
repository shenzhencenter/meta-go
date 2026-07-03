package enums

type PagepostBackdatedTimeGranularity string

const (
	PagepostBackdatedTimeGranularityDay   PagepostBackdatedTimeGranularity = "day"
	PagepostBackdatedTimeGranularityHour  PagepostBackdatedTimeGranularity = "hour"
	PagepostBackdatedTimeGranularityMin   PagepostBackdatedTimeGranularity = "min"
	PagepostBackdatedTimeGranularityMonth PagepostBackdatedTimeGranularity = "month"
	PagepostBackdatedTimeGranularityNone  PagepostBackdatedTimeGranularity = "none"
	PagepostBackdatedTimeGranularityYear  PagepostBackdatedTimeGranularity = "year"
)

func (value PagepostBackdatedTimeGranularity) String() string {
	return string(value)
}
