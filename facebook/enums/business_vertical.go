package enums

type BusinessVertical string

const (
	BusinessVerticalAdvertising                  BusinessVertical = "ADVERTISING"
	BusinessVerticalAutomotive                   BusinessVertical = "AUTOMOTIVE"
	BusinessVerticalConsumerPackagedGoods        BusinessVertical = "CONSUMER_PACKAGED_GOODS"
	BusinessVerticalEcommerce                    BusinessVertical = "ECOMMERCE"
	BusinessVerticalEducation                    BusinessVertical = "EDUCATION"
	BusinessVerticalEnergyAndUtilities           BusinessVertical = "ENERGY_AND_UTILITIES"
	BusinessVerticalEntertainmentAndMedia        BusinessVertical = "ENTERTAINMENT_AND_MEDIA"
	BusinessVerticalFinancialServices            BusinessVertical = "FINANCIAL_SERVICES"
	BusinessVerticalGaming                       BusinessVertical = "GAMING"
	BusinessVerticalGovernmentAndPolitics        BusinessVertical = "GOVERNMENT_AND_POLITICS"
	BusinessVerticalHealth                       BusinessVertical = "HEALTH"
	BusinessVerticalLuxury                       BusinessVertical = "LUXURY"
	BusinessVerticalMarketing                    BusinessVertical = "MARKETING"
	BusinessVerticalNonProfit                    BusinessVertical = "NON_PROFIT"
	BusinessVerticalNotSet                       BusinessVertical = "NOT_SET"
	BusinessVerticalOrganizationsAndAssociations BusinessVertical = "ORGANIZATIONS_AND_ASSOCIATIONS"
	BusinessVerticalOther                        BusinessVertical = "OTHER"
	BusinessVerticalProfessionalServices         BusinessVertical = "PROFESSIONAL_SERVICES"
	BusinessVerticalRestaurant                   BusinessVertical = "RESTAURANT"
	BusinessVerticalRetail                       BusinessVertical = "RETAIL"
	BusinessVerticalTechnology                   BusinessVertical = "TECHNOLOGY"
	BusinessVerticalTelecom                      BusinessVertical = "TELECOM"
	BusinessVerticalTravel                       BusinessVertical = "TRAVEL"
)

func (value BusinessVertical) String() string {
	return string(value)
}
