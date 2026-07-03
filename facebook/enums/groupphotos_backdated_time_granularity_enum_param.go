package enums

type GroupphotosBackdatedTimeGranularityEnumParam string

const (
	GroupphotosBackdatedTimeGranularityEnumParamDay   GroupphotosBackdatedTimeGranularityEnumParam = "day"
	GroupphotosBackdatedTimeGranularityEnumParamHour  GroupphotosBackdatedTimeGranularityEnumParam = "hour"
	GroupphotosBackdatedTimeGranularityEnumParamMin   GroupphotosBackdatedTimeGranularityEnumParam = "min"
	GroupphotosBackdatedTimeGranularityEnumParamMonth GroupphotosBackdatedTimeGranularityEnumParam = "month"
	GroupphotosBackdatedTimeGranularityEnumParamNone  GroupphotosBackdatedTimeGranularityEnumParam = "none"
	GroupphotosBackdatedTimeGranularityEnumParamYear  GroupphotosBackdatedTimeGranularityEnumParam = "year"
)

func (value GroupphotosBackdatedTimeGranularityEnumParam) String() string {
	return string(value)
}
