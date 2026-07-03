package enums

type IguserexportforcaminsightsMetricsEnumParam string

const (
	IguserexportforcaminsightsMetricsEnumParamCreatorEngagedAccounts IguserexportforcaminsightsMetricsEnumParam = "CREATOR_ENGAGED_ACCOUNTS"
	IguserexportforcaminsightsMetricsEnumParamCreatorReach           IguserexportforcaminsightsMetricsEnumParam = "CREATOR_REACH"
	IguserexportforcaminsightsMetricsEnumParamReelsHookRate          IguserexportforcaminsightsMetricsEnumParam = "REELS_HOOK_RATE"
	IguserexportforcaminsightsMetricsEnumParamReelsInteractionRate   IguserexportforcaminsightsMetricsEnumParam = "REELS_INTERACTION_RATE"
	IguserexportforcaminsightsMetricsEnumParamTotalFollowers         IguserexportforcaminsightsMetricsEnumParam = "TOTAL_FOLLOWERS"
)

func (value IguserexportforcaminsightsMetricsEnumParam) String() string {
	return string(value)
}
