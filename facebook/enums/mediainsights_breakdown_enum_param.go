package enums

type MediainsightsBreakdownEnumParam string

const (
	MediainsightsBreakdownEnumParamActionType                MediainsightsBreakdownEnumParam = "action_type"
	MediainsightsBreakdownEnumParamFollowType                MediainsightsBreakdownEnumParam = "follow_type"
	MediainsightsBreakdownEnumParamStoryNavigationActionType MediainsightsBreakdownEnumParam = "story_navigation_action_type"
	MediainsightsBreakdownEnumParamSurfaceType               MediainsightsBreakdownEnumParam = "surface_type"
)

func (value MediainsightsBreakdownEnumParam) String() string {
	return string(value)
}
