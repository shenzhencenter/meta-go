package enums

type CustomaudienceAudienceLabels string

const (
	CustomaudienceAudienceLabelsAppInstallers      CustomaudienceAudienceLabels = "APP_INSTALLERS"
	CustomaudienceAudienceLabelsAppUsers           CustomaudienceAudienceLabels = "APP_USERS"
	CustomaudienceAudienceLabelsAtRisk             CustomaudienceAudienceLabels = "AT_RISK"
	CustomaudienceAudienceLabelsCartAbandoners     CustomaudienceAudienceLabels = "CART_ABANDONERS"
	CustomaudienceAudienceLabelsCustomerLeads      CustomaudienceAudienceLabels = "CUSTOMER_LEADS"
	CustomaudienceAudienceLabelsDisengaged         CustomaudienceAudienceLabels = "DISENGAGED"
	CustomaudienceAudienceLabelsDisqualifiedLeads  CustomaudienceAudienceLabels = "DISQUALIFIED_LEADS"
	CustomaudienceAudienceLabelsEngagedUsers       CustomaudienceAudienceLabels = "ENGAGED_USERS"
	CustomaudienceAudienceLabelsHighValueCustomers CustomaudienceAudienceLabels = "HIGH_VALUE_CUSTOMERS"
	CustomaudienceAudienceLabelsLowValueCustomers  CustomaudienceAudienceLabels = "LOW_VALUE_CUSTOMERS"
	CustomaudienceAudienceLabelsOtherX1            CustomaudienceAudienceLabels = "OTHER_1"
	CustomaudienceAudienceLabelsOtherX2            CustomaudienceAudienceLabels = "OTHER_2"
	CustomaudienceAudienceLabelsOtherX3            CustomaudienceAudienceLabels = "OTHER_3"
	CustomaudienceAudienceLabelsPersonas           CustomaudienceAudienceLabels = "PERSONAS"
	CustomaudienceAudienceLabelsQualifiedLeads     CustomaudienceAudienceLabels = "QUALIFIED_LEADS"
	CustomaudienceAudienceLabelsRecentPurchasers   CustomaudienceAudienceLabels = "RECENT_PURCHASERS"
	CustomaudienceAudienceLabelsRestrictedUsers    CustomaudienceAudienceLabels = "RESTRICTED_USERS"
	CustomaudienceAudienceLabelsTrialUsers         CustomaudienceAudienceLabels = "TRIAL_USERS"
	CustomaudienceAudienceLabelsUnwantedCustomers  CustomaudienceAudienceLabels = "UNWANTED_CUSTOMERS"
)

func (value CustomaudienceAudienceLabels) String() string {
	return string(value)
}
