package enums

type PagephotosBackdatedTimeGranularityEnumParam string

const (
	PagephotosBackdatedTimeGranularityEnumParamDay   PagephotosBackdatedTimeGranularityEnumParam = "day"
	PagephotosBackdatedTimeGranularityEnumParamHour  PagephotosBackdatedTimeGranularityEnumParam = "hour"
	PagephotosBackdatedTimeGranularityEnumParamMin   PagephotosBackdatedTimeGranularityEnumParam = "min"
	PagephotosBackdatedTimeGranularityEnumParamMonth PagephotosBackdatedTimeGranularityEnumParam = "month"
	PagephotosBackdatedTimeGranularityEnumParamNone  PagephotosBackdatedTimeGranularityEnumParam = "none"
	PagephotosBackdatedTimeGranularityEnumParamYear  PagephotosBackdatedTimeGranularityEnumParam = "year"
)

func (value PagephotosBackdatedTimeGranularityEnumParam) String() string {
	return string(value)
}
