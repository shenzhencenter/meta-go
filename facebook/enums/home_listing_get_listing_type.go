package enums

type HomelistinggetListingType string

const (
	HomelistinggetListingTypeEmptyValue      HomelistinggetListingType = "EMPTY_VALUE"
	HomelistinggetListingTypeForeclosed      HomelistinggetListingType = "FORECLOSED"
	HomelistinggetListingTypeForRentByAgent  HomelistinggetListingType = "FOR_RENT_BY_AGENT"
	HomelistinggetListingTypeForRentByOwner  HomelistinggetListingType = "FOR_RENT_BY_OWNER"
	HomelistinggetListingTypeForSaleByAgent  HomelistinggetListingType = "FOR_SALE_BY_AGENT"
	HomelistinggetListingTypeForSaleByOwner  HomelistinggetListingType = "FOR_SALE_BY_OWNER"
	HomelistinggetListingTypeNewConstruction HomelistinggetListingType = "NEW_CONSTRUCTION"
	HomelistinggetListingTypeNewListing      HomelistinggetListingType = "NEW_LISTING"
	HomelistinggetListingTypeOther           HomelistinggetListingType = "OTHER"
)

func (value HomelistinggetListingType) String() string {
	return string(value)
}
