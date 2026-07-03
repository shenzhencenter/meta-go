package enums

type PageabTestsOptimizationGoalEnumParam string

const (
	PageabTestsOptimizationGoalEnumParamAutoResolveToControl PageabTestsOptimizationGoalEnumParam = "AUTO_RESOLVE_TO_CONTROL"
	PageabTestsOptimizationGoalEnumParamAvgTimeWatched       PageabTestsOptimizationGoalEnumParam = "AVG_TIME_WATCHED"
	PageabTestsOptimizationGoalEnumParamComments             PageabTestsOptimizationGoalEnumParam = "COMMENTS"
	PageabTestsOptimizationGoalEnumParamImpressions          PageabTestsOptimizationGoalEnumParam = "IMPRESSIONS"
	PageabTestsOptimizationGoalEnumParamImpressionsUnique    PageabTestsOptimizationGoalEnumParam = "IMPRESSIONS_UNIQUE"
	PageabTestsOptimizationGoalEnumParamLinkClicks           PageabTestsOptimizationGoalEnumParam = "LINK_CLICKS"
	PageabTestsOptimizationGoalEnumParamOther                PageabTestsOptimizationGoalEnumParam = "OTHER"
	PageabTestsOptimizationGoalEnumParamReactions            PageabTestsOptimizationGoalEnumParam = "REACTIONS"
	PageabTestsOptimizationGoalEnumParamReelsPlays           PageabTestsOptimizationGoalEnumParam = "REELS_PLAYS"
	PageabTestsOptimizationGoalEnumParamShares               PageabTestsOptimizationGoalEnumParam = "SHARES"
	PageabTestsOptimizationGoalEnumParamVideoViewsX60S       PageabTestsOptimizationGoalEnumParam = "VIDEO_VIEWS_60S"
	PageabTestsOptimizationGoalEnumParamViewers              PageabTestsOptimizationGoalEnumParam = "VIEWERS"
	PageabTestsOptimizationGoalEnumParamViews                PageabTestsOptimizationGoalEnumParam = "VIEWS"
)

func (value PageabTestsOptimizationGoalEnumParam) String() string {
	return string(value)
}
