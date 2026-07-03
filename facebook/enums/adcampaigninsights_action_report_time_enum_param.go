package enums

type AdcampaigninsightsActionReportTimeEnumParam string

const (
	AdcampaigninsightsActionReportTimeEnumParamConversion AdcampaigninsightsActionReportTimeEnumParam = "conversion"
	AdcampaigninsightsActionReportTimeEnumParamImpression AdcampaigninsightsActionReportTimeEnumParam = "impression"
	AdcampaigninsightsActionReportTimeEnumParamLifetime   AdcampaigninsightsActionReportTimeEnumParam = "lifetime"
	AdcampaigninsightsActionReportTimeEnumParamMixed      AdcampaigninsightsActionReportTimeEnumParam = "mixed"
)

func (value AdcampaigninsightsActionReportTimeEnumParam) String() string {
	return string(value)
}
