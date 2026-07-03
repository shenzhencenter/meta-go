package enums

type BusinessmanagedBusinessesVerticalEnumParam string

const (
	BusinessmanagedBusinessesVerticalEnumParamAdvertising                  BusinessmanagedBusinessesVerticalEnumParam = "ADVERTISING"
	BusinessmanagedBusinessesVerticalEnumParamAutomotive                   BusinessmanagedBusinessesVerticalEnumParam = "AUTOMOTIVE"
	BusinessmanagedBusinessesVerticalEnumParamConsumerPackagedGoods        BusinessmanagedBusinessesVerticalEnumParam = "CONSUMER_PACKAGED_GOODS"
	BusinessmanagedBusinessesVerticalEnumParamEcommerce                    BusinessmanagedBusinessesVerticalEnumParam = "ECOMMERCE"
	BusinessmanagedBusinessesVerticalEnumParamEducation                    BusinessmanagedBusinessesVerticalEnumParam = "EDUCATION"
	BusinessmanagedBusinessesVerticalEnumParamEnergyAndUtilities           BusinessmanagedBusinessesVerticalEnumParam = "ENERGY_AND_UTILITIES"
	BusinessmanagedBusinessesVerticalEnumParamEntertainmentAndMedia        BusinessmanagedBusinessesVerticalEnumParam = "ENTERTAINMENT_AND_MEDIA"
	BusinessmanagedBusinessesVerticalEnumParamFinancialServices            BusinessmanagedBusinessesVerticalEnumParam = "FINANCIAL_SERVICES"
	BusinessmanagedBusinessesVerticalEnumParamGaming                       BusinessmanagedBusinessesVerticalEnumParam = "GAMING"
	BusinessmanagedBusinessesVerticalEnumParamGovernmentAndPolitics        BusinessmanagedBusinessesVerticalEnumParam = "GOVERNMENT_AND_POLITICS"
	BusinessmanagedBusinessesVerticalEnumParamHealth                       BusinessmanagedBusinessesVerticalEnumParam = "HEALTH"
	BusinessmanagedBusinessesVerticalEnumParamLuxury                       BusinessmanagedBusinessesVerticalEnumParam = "LUXURY"
	BusinessmanagedBusinessesVerticalEnumParamMarketing                    BusinessmanagedBusinessesVerticalEnumParam = "MARKETING"
	BusinessmanagedBusinessesVerticalEnumParamNonProfit                    BusinessmanagedBusinessesVerticalEnumParam = "NON_PROFIT"
	BusinessmanagedBusinessesVerticalEnumParamNotSet                       BusinessmanagedBusinessesVerticalEnumParam = "NOT_SET"
	BusinessmanagedBusinessesVerticalEnumParamOrganizationsAndAssociations BusinessmanagedBusinessesVerticalEnumParam = "ORGANIZATIONS_AND_ASSOCIATIONS"
	BusinessmanagedBusinessesVerticalEnumParamOther                        BusinessmanagedBusinessesVerticalEnumParam = "OTHER"
	BusinessmanagedBusinessesVerticalEnumParamProfessionalServices         BusinessmanagedBusinessesVerticalEnumParam = "PROFESSIONAL_SERVICES"
	BusinessmanagedBusinessesVerticalEnumParamRestaurant                   BusinessmanagedBusinessesVerticalEnumParam = "RESTAURANT"
	BusinessmanagedBusinessesVerticalEnumParamRetail                       BusinessmanagedBusinessesVerticalEnumParam = "RETAIL"
	BusinessmanagedBusinessesVerticalEnumParamTechnology                   BusinessmanagedBusinessesVerticalEnumParam = "TECHNOLOGY"
	BusinessmanagedBusinessesVerticalEnumParamTelecom                      BusinessmanagedBusinessesVerticalEnumParam = "TELECOM"
	BusinessmanagedBusinessesVerticalEnumParamTravel                       BusinessmanagedBusinessesVerticalEnumParam = "TRAVEL"
)

func (value BusinessmanagedBusinessesVerticalEnumParam) String() string {
	return string(value)
}
