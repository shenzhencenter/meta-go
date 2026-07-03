package enums

type CampaignEffectiveStatus string

const (
	CampaignEffectiveStatusActive     CampaignEffectiveStatus = "ACTIVE"
	CampaignEffectiveStatusArchived   CampaignEffectiveStatus = "ARCHIVED"
	CampaignEffectiveStatusDeleted    CampaignEffectiveStatus = "DELETED"
	CampaignEffectiveStatusInProcess  CampaignEffectiveStatus = "IN_PROCESS"
	CampaignEffectiveStatusPaused     CampaignEffectiveStatus = "PAUSED"
	CampaignEffectiveStatusWithIssues CampaignEffectiveStatus = "WITH_ISSUES"
)

func (value CampaignEffectiveStatus) String() string {
	return string(value)
}
