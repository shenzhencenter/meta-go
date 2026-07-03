package enums

type UserbusinessesVerticalEnumParam string

const (
	UserbusinessesVerticalEnumParamAdvertising                  UserbusinessesVerticalEnumParam = "ADVERTISING"
	UserbusinessesVerticalEnumParamAutomotive                   UserbusinessesVerticalEnumParam = "AUTOMOTIVE"
	UserbusinessesVerticalEnumParamConsumerPackagedGoods        UserbusinessesVerticalEnumParam = "CONSUMER_PACKAGED_GOODS"
	UserbusinessesVerticalEnumParamEcommerce                    UserbusinessesVerticalEnumParam = "ECOMMERCE"
	UserbusinessesVerticalEnumParamEducation                    UserbusinessesVerticalEnumParam = "EDUCATION"
	UserbusinessesVerticalEnumParamEnergyAndUtilities           UserbusinessesVerticalEnumParam = "ENERGY_AND_UTILITIES"
	UserbusinessesVerticalEnumParamEntertainmentAndMedia        UserbusinessesVerticalEnumParam = "ENTERTAINMENT_AND_MEDIA"
	UserbusinessesVerticalEnumParamFinancialServices            UserbusinessesVerticalEnumParam = "FINANCIAL_SERVICES"
	UserbusinessesVerticalEnumParamGaming                       UserbusinessesVerticalEnumParam = "GAMING"
	UserbusinessesVerticalEnumParamGovernmentAndPolitics        UserbusinessesVerticalEnumParam = "GOVERNMENT_AND_POLITICS"
	UserbusinessesVerticalEnumParamHealth                       UserbusinessesVerticalEnumParam = "HEALTH"
	UserbusinessesVerticalEnumParamLuxury                       UserbusinessesVerticalEnumParam = "LUXURY"
	UserbusinessesVerticalEnumParamMarketing                    UserbusinessesVerticalEnumParam = "MARKETING"
	UserbusinessesVerticalEnumParamNonProfit                    UserbusinessesVerticalEnumParam = "NON_PROFIT"
	UserbusinessesVerticalEnumParamNotSet                       UserbusinessesVerticalEnumParam = "NOT_SET"
	UserbusinessesVerticalEnumParamOrganizationsAndAssociations UserbusinessesVerticalEnumParam = "ORGANIZATIONS_AND_ASSOCIATIONS"
	UserbusinessesVerticalEnumParamOther                        UserbusinessesVerticalEnumParam = "OTHER"
	UserbusinessesVerticalEnumParamProfessionalServices         UserbusinessesVerticalEnumParam = "PROFESSIONAL_SERVICES"
	UserbusinessesVerticalEnumParamRestaurant                   UserbusinessesVerticalEnumParam = "RESTAURANT"
	UserbusinessesVerticalEnumParamRetail                       UserbusinessesVerticalEnumParam = "RETAIL"
	UserbusinessesVerticalEnumParamTechnology                   UserbusinessesVerticalEnumParam = "TECHNOLOGY"
	UserbusinessesVerticalEnumParamTelecom                      UserbusinessesVerticalEnumParam = "TELECOM"
	UserbusinessesVerticalEnumParamTravel                       UserbusinessesVerticalEnumParam = "TRAVEL"
)

func (value UserbusinessesVerticalEnumParam) String() string {
	return string(value)
}
