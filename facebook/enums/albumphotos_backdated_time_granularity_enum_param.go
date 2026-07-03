package enums

type AlbumphotosBackdatedTimeGranularityEnumParam string

const (
	AlbumphotosBackdatedTimeGranularityEnumParamDay   AlbumphotosBackdatedTimeGranularityEnumParam = "day"
	AlbumphotosBackdatedTimeGranularityEnumParamHour  AlbumphotosBackdatedTimeGranularityEnumParam = "hour"
	AlbumphotosBackdatedTimeGranularityEnumParamMin   AlbumphotosBackdatedTimeGranularityEnumParam = "min"
	AlbumphotosBackdatedTimeGranularityEnumParamMonth AlbumphotosBackdatedTimeGranularityEnumParam = "month"
	AlbumphotosBackdatedTimeGranularityEnumParamNone  AlbumphotosBackdatedTimeGranularityEnumParam = "none"
	AlbumphotosBackdatedTimeGranularityEnumParamYear  AlbumphotosBackdatedTimeGranularityEnumParam = "year"
)

func (value AlbumphotosBackdatedTimeGranularityEnumParam) String() string {
	return string(value)
}
