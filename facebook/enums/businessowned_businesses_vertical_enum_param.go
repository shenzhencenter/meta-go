package enums

type BusinessownedBusinessesVerticalEnumParam string

const (
	BusinessownedBusinessesVerticalEnumParamAdvertising                  BusinessownedBusinessesVerticalEnumParam = "ADVERTISING"
	BusinessownedBusinessesVerticalEnumParamAutomotive                   BusinessownedBusinessesVerticalEnumParam = "AUTOMOTIVE"
	BusinessownedBusinessesVerticalEnumParamConsumerPackagedGoods        BusinessownedBusinessesVerticalEnumParam = "CONSUMER_PACKAGED_GOODS"
	BusinessownedBusinessesVerticalEnumParamEcommerce                    BusinessownedBusinessesVerticalEnumParam = "ECOMMERCE"
	BusinessownedBusinessesVerticalEnumParamEducation                    BusinessownedBusinessesVerticalEnumParam = "EDUCATION"
	BusinessownedBusinessesVerticalEnumParamEnergyAndUtilities           BusinessownedBusinessesVerticalEnumParam = "ENERGY_AND_UTILITIES"
	BusinessownedBusinessesVerticalEnumParamEntertainmentAndMedia        BusinessownedBusinessesVerticalEnumParam = "ENTERTAINMENT_AND_MEDIA"
	BusinessownedBusinessesVerticalEnumParamFinancialServices            BusinessownedBusinessesVerticalEnumParam = "FINANCIAL_SERVICES"
	BusinessownedBusinessesVerticalEnumParamGaming                       BusinessownedBusinessesVerticalEnumParam = "GAMING"
	BusinessownedBusinessesVerticalEnumParamGovernmentAndPolitics        BusinessownedBusinessesVerticalEnumParam = "GOVERNMENT_AND_POLITICS"
	BusinessownedBusinessesVerticalEnumParamHealth                       BusinessownedBusinessesVerticalEnumParam = "HEALTH"
	BusinessownedBusinessesVerticalEnumParamLuxury                       BusinessownedBusinessesVerticalEnumParam = "LUXURY"
	BusinessownedBusinessesVerticalEnumParamMarketing                    BusinessownedBusinessesVerticalEnumParam = "MARKETING"
	BusinessownedBusinessesVerticalEnumParamNonProfit                    BusinessownedBusinessesVerticalEnumParam = "NON_PROFIT"
	BusinessownedBusinessesVerticalEnumParamNotSet                       BusinessownedBusinessesVerticalEnumParam = "NOT_SET"
	BusinessownedBusinessesVerticalEnumParamOrganizationsAndAssociations BusinessownedBusinessesVerticalEnumParam = "ORGANIZATIONS_AND_ASSOCIATIONS"
	BusinessownedBusinessesVerticalEnumParamOther                        BusinessownedBusinessesVerticalEnumParam = "OTHER"
	BusinessownedBusinessesVerticalEnumParamProfessionalServices         BusinessownedBusinessesVerticalEnumParam = "PROFESSIONAL_SERVICES"
	BusinessownedBusinessesVerticalEnumParamRestaurant                   BusinessownedBusinessesVerticalEnumParam = "RESTAURANT"
	BusinessownedBusinessesVerticalEnumParamRetail                       BusinessownedBusinessesVerticalEnumParam = "RETAIL"
	BusinessownedBusinessesVerticalEnumParamTechnology                   BusinessownedBusinessesVerticalEnumParam = "TECHNOLOGY"
	BusinessownedBusinessesVerticalEnumParamTelecom                      BusinessownedBusinessesVerticalEnumParam = "TELECOM"
	BusinessownedBusinessesVerticalEnumParamTravel                       BusinessownedBusinessesVerticalEnumParam = "TRAVEL"
)

func (value BusinessownedBusinessesVerticalEnumParam) String() string {
	return string(value)
}
