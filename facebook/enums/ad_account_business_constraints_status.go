package enums

type AdaccountbusinessconstraintsStatus string

const (
	AdaccountbusinessconstraintsStatusActive                AdaccountbusinessconstraintsStatus = "ACTIVE"
	AdaccountbusinessconstraintsStatusApplicationInProgress AdaccountbusinessconstraintsStatus = "APPLICATION_IN_PROGRESS"
	AdaccountbusinessconstraintsStatusWithCampaignError     AdaccountbusinessconstraintsStatus = "WITH_CAMPAIGN_ERROR"
)

func (value AdaccountbusinessconstraintsStatus) String() string {
	return string(value)
}
