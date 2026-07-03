package enums

type LocalservicebusinessAvailability string

const (
	LocalservicebusinessAvailabilityAvailableForOrder LocalservicebusinessAvailability = "AVAILABLE_FOR_ORDER"
	LocalservicebusinessAvailabilityDiscontinued      LocalservicebusinessAvailability = "DISCONTINUED"
	LocalservicebusinessAvailabilityInStock           LocalservicebusinessAvailability = "IN_STOCK"
	LocalservicebusinessAvailabilityMarkAsExpired     LocalservicebusinessAvailability = "MARK_AS_EXPIRED"
	LocalservicebusinessAvailabilityMarkAsSold        LocalservicebusinessAvailability = "MARK_AS_SOLD"
	LocalservicebusinessAvailabilityOutOfStock        LocalservicebusinessAvailability = "OUT_OF_STOCK"
	LocalservicebusinessAvailabilityPending           LocalservicebusinessAvailability = "PENDING"
	LocalservicebusinessAvailabilityPreorder          LocalservicebusinessAvailability = "PREORDER"
)

func (value LocalservicebusinessAvailability) String() string {
	return string(value)
}
