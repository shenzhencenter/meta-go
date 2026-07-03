package enums

type CampaignStatus string

const (
	CampaignStatusActive   CampaignStatus = "ACTIVE"
	CampaignStatusArchived CampaignStatus = "ARCHIVED"
	CampaignStatusDeleted  CampaignStatus = "DELETED"
	CampaignStatusPaused   CampaignStatus = "PAUSED"
)

func (value CampaignStatus) String() string {
	return string(value)
}
