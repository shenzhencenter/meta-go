package enums

type GroupfeedBackdatedTimeGranularityEnumParam string

const (
	GroupfeedBackdatedTimeGranularityEnumParamDay   GroupfeedBackdatedTimeGranularityEnumParam = "day"
	GroupfeedBackdatedTimeGranularityEnumParamHour  GroupfeedBackdatedTimeGranularityEnumParam = "hour"
	GroupfeedBackdatedTimeGranularityEnumParamMin   GroupfeedBackdatedTimeGranularityEnumParam = "min"
	GroupfeedBackdatedTimeGranularityEnumParamMonth GroupfeedBackdatedTimeGranularityEnumParam = "month"
	GroupfeedBackdatedTimeGranularityEnumParamNone  GroupfeedBackdatedTimeGranularityEnumParam = "none"
	GroupfeedBackdatedTimeGranularityEnumParamYear  GroupfeedBackdatedTimeGranularityEnumParam = "year"
)

func (value GroupfeedBackdatedTimeGranularityEnumParam) String() string {
	return string(value)
}
