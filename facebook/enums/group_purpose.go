package enums

type GroupPurpose string

const (
	GroupPurposeCasual           GroupPurpose = "CASUAL"
	GroupPurposeCoworkers        GroupPurpose = "COWORKERS"
	GroupPurposeCustom           GroupPurpose = "CUSTOM"
	GroupPurposeForSale          GroupPurpose = "FOR_SALE"
	GroupPurposeForWork          GroupPurpose = "FOR_WORK"
	GroupPurposeGame             GroupPurpose = "GAME"
	GroupPurposeHealthSupport    GroupPurpose = "HEALTH_SUPPORT"
	GroupPurposeJobs             GroupPurpose = "JOBS"
	GroupPurposeLearning         GroupPurpose = "LEARNING"
	GroupPurposeNone             GroupPurpose = "NONE"
	GroupPurposeParenting        GroupPurpose = "PARENTING"
	GroupPurposeStreamer         GroupPurpose = "STREAMER"
	GroupPurposeWorkAgentToAgent GroupPurpose = "WORK_AGENT_TO_AGENT"
	GroupPurposeWorkAnnouncement GroupPurpose = "WORK_ANNOUNCEMENT"
	GroupPurposeWorkDemoGroup    GroupPurpose = "WORK_DEMO_GROUP"
	GroupPurposeWorkDiscussion   GroupPurpose = "WORK_DISCUSSION"
	GroupPurposeWorkEphemeral    GroupPurpose = "WORK_EPHEMERAL"
	GroupPurposeWorkFeedback     GroupPurpose = "WORK_FEEDBACK"
	GroupPurposeWorkForSale      GroupPurpose = "WORK_FOR_SALE"
	GroupPurposeWorkGarden       GroupPurpose = "WORK_GARDEN"
	GroupPurposeWorkIntegrity    GroupPurpose = "WORK_INTEGRITY"
	GroupPurposeWorkLearning     GroupPurpose = "WORK_LEARNING"
	GroupPurposeWorkMentorship   GroupPurpose = "WORK_MENTORSHIP"
	GroupPurposeWorkMultiCompany GroupPurpose = "WORK_MULTI_COMPANY"
	GroupPurposeWorkRecruiting   GroupPurpose = "WORK_RECRUITING"
	GroupPurposeWorkSocial       GroupPurpose = "WORK_SOCIAL"
	GroupPurposeWorkStages       GroupPurpose = "WORK_STAGES"
	GroupPurposeWorkTeam         GroupPurpose = "WORK_TEAM"
	GroupPurposeWorkTeamwork     GroupPurpose = "WORK_TEAMWORK"
)

func (value GroupPurpose) String() string {
	return string(value)
}
