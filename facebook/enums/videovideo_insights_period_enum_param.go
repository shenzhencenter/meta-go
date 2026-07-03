package enums

type VideovideoInsightsPeriodEnumParam string

const (
	VideovideoInsightsPeriodEnumParamDay            VideovideoInsightsPeriodEnumParam = "day"
	VideovideoInsightsPeriodEnumParamDaysX28        VideovideoInsightsPeriodEnumParam = "days_28"
	VideovideoInsightsPeriodEnumParamLifetime       VideovideoInsightsPeriodEnumParam = "lifetime"
	VideovideoInsightsPeriodEnumParamMonth          VideovideoInsightsPeriodEnumParam = "month"
	VideovideoInsightsPeriodEnumParamTotalOverRange VideovideoInsightsPeriodEnumParam = "total_over_range"
	VideovideoInsightsPeriodEnumParamWeek           VideovideoInsightsPeriodEnumParam = "week"
)

func (value VideovideoInsightsPeriodEnumParam) String() string {
	return string(value)
}
