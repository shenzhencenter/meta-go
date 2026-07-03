package enums

type InstagrambusinessassetagenciesPermittedTasksEnumParam string

const (
	InstagrambusinessassetagenciesPermittedTasksEnumParamAdvertise         InstagrambusinessassetagenciesPermittedTasksEnumParam = "ADVERTISE"
	InstagrambusinessassetagenciesPermittedTasksEnumParamAnalyze           InstagrambusinessassetagenciesPermittedTasksEnumParam = "ANALYZE"
	InstagrambusinessassetagenciesPermittedTasksEnumParamCommunityActivity InstagrambusinessassetagenciesPermittedTasksEnumParam = "COMMUNITY_ACTIVITY"
	InstagrambusinessassetagenciesPermittedTasksEnumParamContent           InstagrambusinessassetagenciesPermittedTasksEnumParam = "CONTENT"
	InstagrambusinessassetagenciesPermittedTasksEnumParamMessages          InstagrambusinessassetagenciesPermittedTasksEnumParam = "MESSAGES"
)

func (value InstagrambusinessassetagenciesPermittedTasksEnumParam) String() string {
	return string(value)
}
