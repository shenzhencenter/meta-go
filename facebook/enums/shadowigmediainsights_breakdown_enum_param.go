package enums

type ShadowigmediainsightsBreakdownEnumParam string

const (
	ShadowigmediainsightsBreakdownEnumParamActionType                ShadowigmediainsightsBreakdownEnumParam = "action_type"
	ShadowigmediainsightsBreakdownEnumParamFollowType                ShadowigmediainsightsBreakdownEnumParam = "follow_type"
	ShadowigmediainsightsBreakdownEnumParamStoryNavigationActionType ShadowigmediainsightsBreakdownEnumParam = "story_navigation_action_type"
	ShadowigmediainsightsBreakdownEnumParamSurfaceType               ShadowigmediainsightsBreakdownEnumParam = "surface_type"
)

func (value ShadowigmediainsightsBreakdownEnumParam) String() string {
	return string(value)
}
