package enums

type UserphotosBackdatedTimeGranularityEnumParam string

const (
	UserphotosBackdatedTimeGranularityEnumParamDay   UserphotosBackdatedTimeGranularityEnumParam = "day"
	UserphotosBackdatedTimeGranularityEnumParamHour  UserphotosBackdatedTimeGranularityEnumParam = "hour"
	UserphotosBackdatedTimeGranularityEnumParamMin   UserphotosBackdatedTimeGranularityEnumParam = "min"
	UserphotosBackdatedTimeGranularityEnumParamMonth UserphotosBackdatedTimeGranularityEnumParam = "month"
	UserphotosBackdatedTimeGranularityEnumParamNone  UserphotosBackdatedTimeGranularityEnumParam = "none"
	UserphotosBackdatedTimeGranularityEnumParamYear  UserphotosBackdatedTimeGranularityEnumParam = "year"
)

func (value UserphotosBackdatedTimeGranularityEnumParam) String() string {
	return string(value)
}
