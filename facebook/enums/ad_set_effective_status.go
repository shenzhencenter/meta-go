package enums

type AdsetEffectiveStatus string

const (
	AdsetEffectiveStatusActive         AdsetEffectiveStatus = "ACTIVE"
	AdsetEffectiveStatusArchived       AdsetEffectiveStatus = "ARCHIVED"
	AdsetEffectiveStatusCampaignPaused AdsetEffectiveStatus = "CAMPAIGN_PAUSED"
	AdsetEffectiveStatusDeleted        AdsetEffectiveStatus = "DELETED"
	AdsetEffectiveStatusInProcess      AdsetEffectiveStatus = "IN_PROCESS"
	AdsetEffectiveStatusPaused         AdsetEffectiveStatus = "PAUSED"
	AdsetEffectiveStatusWithIssues     AdsetEffectiveStatus = "WITH_ISSUES"
)

func (value AdsetEffectiveStatus) String() string {
	return string(value)
}
