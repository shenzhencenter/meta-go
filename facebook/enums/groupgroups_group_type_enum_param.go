package enums

type GroupgroupsGroupTypeEnumParam string

const (
	GroupgroupsGroupTypeEnumParamCasual           GroupgroupsGroupTypeEnumParam = "CASUAL"
	GroupgroupsGroupTypeEnumParamCoworkers        GroupgroupsGroupTypeEnumParam = "COWORKERS"
	GroupgroupsGroupTypeEnumParamCustom           GroupgroupsGroupTypeEnumParam = "CUSTOM"
	GroupgroupsGroupTypeEnumParamForSale          GroupgroupsGroupTypeEnumParam = "FOR_SALE"
	GroupgroupsGroupTypeEnumParamForWork          GroupgroupsGroupTypeEnumParam = "FOR_WORK"
	GroupgroupsGroupTypeEnumParamGame             GroupgroupsGroupTypeEnumParam = "GAME"
	GroupgroupsGroupTypeEnumParamHealthSupport    GroupgroupsGroupTypeEnumParam = "HEALTH_SUPPORT"
	GroupgroupsGroupTypeEnumParamJobs             GroupgroupsGroupTypeEnumParam = "JOBS"
	GroupgroupsGroupTypeEnumParamLearning         GroupgroupsGroupTypeEnumParam = "LEARNING"
	GroupgroupsGroupTypeEnumParamNone             GroupgroupsGroupTypeEnumParam = "NONE"
	GroupgroupsGroupTypeEnumParamParenting        GroupgroupsGroupTypeEnumParam = "PARENTING"
	GroupgroupsGroupTypeEnumParamStreamer         GroupgroupsGroupTypeEnumParam = "STREAMER"
	GroupgroupsGroupTypeEnumParamWorkAgentToAgent GroupgroupsGroupTypeEnumParam = "WORK_AGENT_TO_AGENT"
	GroupgroupsGroupTypeEnumParamWorkAnnouncement GroupgroupsGroupTypeEnumParam = "WORK_ANNOUNCEMENT"
	GroupgroupsGroupTypeEnumParamWorkDemoGroup    GroupgroupsGroupTypeEnumParam = "WORK_DEMO_GROUP"
	GroupgroupsGroupTypeEnumParamWorkDiscussion   GroupgroupsGroupTypeEnumParam = "WORK_DISCUSSION"
	GroupgroupsGroupTypeEnumParamWorkEphemeral    GroupgroupsGroupTypeEnumParam = "WORK_EPHEMERAL"
	GroupgroupsGroupTypeEnumParamWorkFeedback     GroupgroupsGroupTypeEnumParam = "WORK_FEEDBACK"
	GroupgroupsGroupTypeEnumParamWorkForSale      GroupgroupsGroupTypeEnumParam = "WORK_FOR_SALE"
	GroupgroupsGroupTypeEnumParamWorkGarden       GroupgroupsGroupTypeEnumParam = "WORK_GARDEN"
	GroupgroupsGroupTypeEnumParamWorkIntegrity    GroupgroupsGroupTypeEnumParam = "WORK_INTEGRITY"
	GroupgroupsGroupTypeEnumParamWorkLearning     GroupgroupsGroupTypeEnumParam = "WORK_LEARNING"
	GroupgroupsGroupTypeEnumParamWorkMentorship   GroupgroupsGroupTypeEnumParam = "WORK_MENTORSHIP"
	GroupgroupsGroupTypeEnumParamWorkMultiCompany GroupgroupsGroupTypeEnumParam = "WORK_MULTI_COMPANY"
	GroupgroupsGroupTypeEnumParamWorkRecruiting   GroupgroupsGroupTypeEnumParam = "WORK_RECRUITING"
	GroupgroupsGroupTypeEnumParamWorkSocial       GroupgroupsGroupTypeEnumParam = "WORK_SOCIAL"
	GroupgroupsGroupTypeEnumParamWorkStages       GroupgroupsGroupTypeEnumParam = "WORK_STAGES"
	GroupgroupsGroupTypeEnumParamWorkTeam         GroupgroupsGroupTypeEnumParam = "WORK_TEAM"
	GroupgroupsGroupTypeEnumParamWorkTeamwork     GroupgroupsGroupTypeEnumParam = "WORK_TEAMWORK"
)

func (value GroupgroupsGroupTypeEnumParam) String() string {
	return string(value)
}
