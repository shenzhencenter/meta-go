package enums

type AdcampaigngroupinsightsActionReportTimeEnumParam string

const (
	AdcampaigngroupinsightsActionReportTimeEnumParamConversion AdcampaigngroupinsightsActionReportTimeEnumParam = "conversion"
	AdcampaigngroupinsightsActionReportTimeEnumParamImpression AdcampaigngroupinsightsActionReportTimeEnumParam = "impression"
	AdcampaigngroupinsightsActionReportTimeEnumParamLifetime   AdcampaigngroupinsightsActionReportTimeEnumParam = "lifetime"
	AdcampaigngroupinsightsActionReportTimeEnumParamMixed      AdcampaigngroupinsightsActionReportTimeEnumParam = "mixed"
)

func (value AdcampaigngroupinsightsActionReportTimeEnumParam) String() string {
	return string(value)
}
