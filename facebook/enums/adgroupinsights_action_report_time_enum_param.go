package enums

type AdgroupinsightsActionReportTimeEnumParam string

const (
	AdgroupinsightsActionReportTimeEnumParamConversion AdgroupinsightsActionReportTimeEnumParam = "conversion"
	AdgroupinsightsActionReportTimeEnumParamImpression AdgroupinsightsActionReportTimeEnumParam = "impression"
	AdgroupinsightsActionReportTimeEnumParamLifetime   AdgroupinsightsActionReportTimeEnumParam = "lifetime"
	AdgroupinsightsActionReportTimeEnumParamMixed      AdgroupinsightsActionReportTimeEnumParam = "mixed"
)

func (value AdgroupinsightsActionReportTimeEnumParam) String() string {
	return string(value)
}
