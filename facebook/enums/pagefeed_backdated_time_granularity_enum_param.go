package enums

type PagefeedBackdatedTimeGranularityEnumParam string

const (
	PagefeedBackdatedTimeGranularityEnumParamDay   PagefeedBackdatedTimeGranularityEnumParam = "day"
	PagefeedBackdatedTimeGranularityEnumParamHour  PagefeedBackdatedTimeGranularityEnumParam = "hour"
	PagefeedBackdatedTimeGranularityEnumParamMin   PagefeedBackdatedTimeGranularityEnumParam = "min"
	PagefeedBackdatedTimeGranularityEnumParamMonth PagefeedBackdatedTimeGranularityEnumParam = "month"
	PagefeedBackdatedTimeGranularityEnumParamNone  PagefeedBackdatedTimeGranularityEnumParam = "none"
	PagefeedBackdatedTimeGranularityEnumParamYear  PagefeedBackdatedTimeGranularityEnumParam = "year"
)

func (value PagefeedBackdatedTimeGranularityEnumParam) String() string {
	return string(value)
}
