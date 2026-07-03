package enums

type UserfeedBackdatedTimeGranularityEnumParam string

const (
	UserfeedBackdatedTimeGranularityEnumParamDay   UserfeedBackdatedTimeGranularityEnumParam = "day"
	UserfeedBackdatedTimeGranularityEnumParamHour  UserfeedBackdatedTimeGranularityEnumParam = "hour"
	UserfeedBackdatedTimeGranularityEnumParamMin   UserfeedBackdatedTimeGranularityEnumParam = "min"
	UserfeedBackdatedTimeGranularityEnumParamMonth UserfeedBackdatedTimeGranularityEnumParam = "month"
	UserfeedBackdatedTimeGranularityEnumParamNone  UserfeedBackdatedTimeGranularityEnumParam = "none"
	UserfeedBackdatedTimeGranularityEnumParamYear  UserfeedBackdatedTimeGranularityEnumParam = "year"
)

func (value UserfeedBackdatedTimeGranularityEnumParam) String() string {
	return string(value)
}
