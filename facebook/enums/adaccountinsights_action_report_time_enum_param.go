package enums

type AdaccountinsightsActionReportTimeEnumParam string

const (
	AdaccountinsightsActionReportTimeEnumParamConversion AdaccountinsightsActionReportTimeEnumParam = "conversion"
	AdaccountinsightsActionReportTimeEnumParamImpression AdaccountinsightsActionReportTimeEnumParam = "impression"
	AdaccountinsightsActionReportTimeEnumParamLifetime   AdaccountinsightsActionReportTimeEnumParam = "lifetime"
	AdaccountinsightsActionReportTimeEnumParamMixed      AdaccountinsightsActionReportTimeEnumParam = "mixed"
)

func (value AdaccountinsightsActionReportTimeEnumParam) String() string {
	return string(value)
}
