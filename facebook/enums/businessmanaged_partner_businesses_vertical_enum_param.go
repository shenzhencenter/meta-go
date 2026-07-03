package enums

type BusinessmanagedPartnerBusinessesVerticalEnumParam string

const (
	BusinessmanagedPartnerBusinessesVerticalEnumParamAdvertising                  BusinessmanagedPartnerBusinessesVerticalEnumParam = "ADVERTISING"
	BusinessmanagedPartnerBusinessesVerticalEnumParamAutomotive                   BusinessmanagedPartnerBusinessesVerticalEnumParam = "AUTOMOTIVE"
	BusinessmanagedPartnerBusinessesVerticalEnumParamConsumerPackagedGoods        BusinessmanagedPartnerBusinessesVerticalEnumParam = "CONSUMER_PACKAGED_GOODS"
	BusinessmanagedPartnerBusinessesVerticalEnumParamEcommerce                    BusinessmanagedPartnerBusinessesVerticalEnumParam = "ECOMMERCE"
	BusinessmanagedPartnerBusinessesVerticalEnumParamEducation                    BusinessmanagedPartnerBusinessesVerticalEnumParam = "EDUCATION"
	BusinessmanagedPartnerBusinessesVerticalEnumParamEnergyAndUtilities           BusinessmanagedPartnerBusinessesVerticalEnumParam = "ENERGY_AND_UTILITIES"
	BusinessmanagedPartnerBusinessesVerticalEnumParamEntertainmentAndMedia        BusinessmanagedPartnerBusinessesVerticalEnumParam = "ENTERTAINMENT_AND_MEDIA"
	BusinessmanagedPartnerBusinessesVerticalEnumParamFinancialServices            BusinessmanagedPartnerBusinessesVerticalEnumParam = "FINANCIAL_SERVICES"
	BusinessmanagedPartnerBusinessesVerticalEnumParamGaming                       BusinessmanagedPartnerBusinessesVerticalEnumParam = "GAMING"
	BusinessmanagedPartnerBusinessesVerticalEnumParamGovernmentAndPolitics        BusinessmanagedPartnerBusinessesVerticalEnumParam = "GOVERNMENT_AND_POLITICS"
	BusinessmanagedPartnerBusinessesVerticalEnumParamHealth                       BusinessmanagedPartnerBusinessesVerticalEnumParam = "HEALTH"
	BusinessmanagedPartnerBusinessesVerticalEnumParamLuxury                       BusinessmanagedPartnerBusinessesVerticalEnumParam = "LUXURY"
	BusinessmanagedPartnerBusinessesVerticalEnumParamMarketing                    BusinessmanagedPartnerBusinessesVerticalEnumParam = "MARKETING"
	BusinessmanagedPartnerBusinessesVerticalEnumParamNonProfit                    BusinessmanagedPartnerBusinessesVerticalEnumParam = "NON_PROFIT"
	BusinessmanagedPartnerBusinessesVerticalEnumParamNotSet                       BusinessmanagedPartnerBusinessesVerticalEnumParam = "NOT_SET"
	BusinessmanagedPartnerBusinessesVerticalEnumParamOrganizationsAndAssociations BusinessmanagedPartnerBusinessesVerticalEnumParam = "ORGANIZATIONS_AND_ASSOCIATIONS"
	BusinessmanagedPartnerBusinessesVerticalEnumParamOther                        BusinessmanagedPartnerBusinessesVerticalEnumParam = "OTHER"
	BusinessmanagedPartnerBusinessesVerticalEnumParamProfessionalServices         BusinessmanagedPartnerBusinessesVerticalEnumParam = "PROFESSIONAL_SERVICES"
	BusinessmanagedPartnerBusinessesVerticalEnumParamRestaurant                   BusinessmanagedPartnerBusinessesVerticalEnumParam = "RESTAURANT"
	BusinessmanagedPartnerBusinessesVerticalEnumParamRetail                       BusinessmanagedPartnerBusinessesVerticalEnumParam = "RETAIL"
	BusinessmanagedPartnerBusinessesVerticalEnumParamTechnology                   BusinessmanagedPartnerBusinessesVerticalEnumParam = "TECHNOLOGY"
	BusinessmanagedPartnerBusinessesVerticalEnumParamTelecom                      BusinessmanagedPartnerBusinessesVerticalEnumParam = "TELECOM"
	BusinessmanagedPartnerBusinessesVerticalEnumParamTravel                       BusinessmanagedPartnerBusinessesVerticalEnumParam = "TRAVEL"
)

func (value BusinessmanagedPartnerBusinessesVerticalEnumParam) String() string {
	return string(value)
}
