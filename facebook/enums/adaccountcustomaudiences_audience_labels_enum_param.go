package enums

type AdaccountcustomaudiencesAudienceLabelsEnumParam string

const (
	AdaccountcustomaudiencesAudienceLabelsEnumParamAppInstallers      AdaccountcustomaudiencesAudienceLabelsEnumParam = "APP_INSTALLERS"
	AdaccountcustomaudiencesAudienceLabelsEnumParamAppUsers           AdaccountcustomaudiencesAudienceLabelsEnumParam = "APP_USERS"
	AdaccountcustomaudiencesAudienceLabelsEnumParamAtRisk             AdaccountcustomaudiencesAudienceLabelsEnumParam = "AT_RISK"
	AdaccountcustomaudiencesAudienceLabelsEnumParamCartAbandoners     AdaccountcustomaudiencesAudienceLabelsEnumParam = "CART_ABANDONERS"
	AdaccountcustomaudiencesAudienceLabelsEnumParamCustomerLeads      AdaccountcustomaudiencesAudienceLabelsEnumParam = "CUSTOMER_LEADS"
	AdaccountcustomaudiencesAudienceLabelsEnumParamDisengaged         AdaccountcustomaudiencesAudienceLabelsEnumParam = "DISENGAGED"
	AdaccountcustomaudiencesAudienceLabelsEnumParamDisqualifiedLeads  AdaccountcustomaudiencesAudienceLabelsEnumParam = "DISQUALIFIED_LEADS"
	AdaccountcustomaudiencesAudienceLabelsEnumParamEngagedUsers       AdaccountcustomaudiencesAudienceLabelsEnumParam = "ENGAGED_USERS"
	AdaccountcustomaudiencesAudienceLabelsEnumParamHighValueCustomers AdaccountcustomaudiencesAudienceLabelsEnumParam = "HIGH_VALUE_CUSTOMERS"
	AdaccountcustomaudiencesAudienceLabelsEnumParamLowValueCustomers  AdaccountcustomaudiencesAudienceLabelsEnumParam = "LOW_VALUE_CUSTOMERS"
	AdaccountcustomaudiencesAudienceLabelsEnumParamOtherX1            AdaccountcustomaudiencesAudienceLabelsEnumParam = "OTHER_1"
	AdaccountcustomaudiencesAudienceLabelsEnumParamOtherX2            AdaccountcustomaudiencesAudienceLabelsEnumParam = "OTHER_2"
	AdaccountcustomaudiencesAudienceLabelsEnumParamOtherX3            AdaccountcustomaudiencesAudienceLabelsEnumParam = "OTHER_3"
	AdaccountcustomaudiencesAudienceLabelsEnumParamPersonas           AdaccountcustomaudiencesAudienceLabelsEnumParam = "PERSONAS"
	AdaccountcustomaudiencesAudienceLabelsEnumParamQualifiedLeads     AdaccountcustomaudiencesAudienceLabelsEnumParam = "QUALIFIED_LEADS"
	AdaccountcustomaudiencesAudienceLabelsEnumParamRecentPurchasers   AdaccountcustomaudiencesAudienceLabelsEnumParam = "RECENT_PURCHASERS"
	AdaccountcustomaudiencesAudienceLabelsEnumParamRestrictedUsers    AdaccountcustomaudiencesAudienceLabelsEnumParam = "RESTRICTED_USERS"
	AdaccountcustomaudiencesAudienceLabelsEnumParamTrialUsers         AdaccountcustomaudiencesAudienceLabelsEnumParam = "TRIAL_USERS"
	AdaccountcustomaudiencesAudienceLabelsEnumParamUnwantedCustomers  AdaccountcustomaudiencesAudienceLabelsEnumParam = "UNWANTED_CUSTOMERS"
)

func (value AdaccountcustomaudiencesAudienceLabelsEnumParam) String() string {
	return string(value)
}
