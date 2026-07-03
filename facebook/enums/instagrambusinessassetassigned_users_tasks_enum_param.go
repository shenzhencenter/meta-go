package enums

type InstagrambusinessassetassignedUsersTasksEnumParam string

const (
	InstagrambusinessassetassignedUsersTasksEnumParamAdvertise         InstagrambusinessassetassignedUsersTasksEnumParam = "ADVERTISE"
	InstagrambusinessassetassignedUsersTasksEnumParamAnalyze           InstagrambusinessassetassignedUsersTasksEnumParam = "ANALYZE"
	InstagrambusinessassetassignedUsersTasksEnumParamCommunityActivity InstagrambusinessassetassignedUsersTasksEnumParam = "COMMUNITY_ACTIVITY"
	InstagrambusinessassetassignedUsersTasksEnumParamContent           InstagrambusinessassetassignedUsersTasksEnumParam = "CONTENT"
	InstagrambusinessassetassignedUsersTasksEnumParamMessages          InstagrambusinessassetassignedUsersTasksEnumParam = "MESSAGES"
)

func (value InstagrambusinessassetassignedUsersTasksEnumParam) String() string {
	return string(value)
}
