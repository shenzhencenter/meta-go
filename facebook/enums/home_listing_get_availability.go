package enums

type HomelistinggetAvailability string

const (
	HomelistinggetAvailabilityAvailableSoon HomelistinggetAvailability = "AVAILABLE_SOON"
	HomelistinggetAvailabilityForRent       HomelistinggetAvailability = "FOR_RENT"
	HomelistinggetAvailabilityForSale       HomelistinggetAvailability = "FOR_SALE"
	HomelistinggetAvailabilityOffMarket     HomelistinggetAvailability = "OFF_MARKET"
	HomelistinggetAvailabilityRecentlySold  HomelistinggetAvailability = "RECENTLY_SOLD"
	HomelistinggetAvailabilitySalePending   HomelistinggetAvailability = "SALE_PENDING"
)

func (value HomelistinggetAvailability) String() string {
	return string(value)
}
