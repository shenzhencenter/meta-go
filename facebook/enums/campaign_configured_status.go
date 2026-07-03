package enums

type CampaignConfiguredStatus string

const (
	CampaignConfiguredStatusActive   CampaignConfiguredStatus = "ACTIVE"
	CampaignConfiguredStatusArchived CampaignConfiguredStatus = "ARCHIVED"
	CampaignConfiguredStatusDeleted  CampaignConfiguredStatus = "DELETED"
	CampaignConfiguredStatusPaused   CampaignConfiguredStatus = "PAUSED"
)

func (value CampaignConfiguredStatus) String() string {
	return string(value)
}
