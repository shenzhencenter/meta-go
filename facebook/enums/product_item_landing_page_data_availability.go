package enums

type ProductitemlandingpagedataAvailability string

const (
	ProductitemlandingpagedataAvailabilityAvailableForOrder ProductitemlandingpagedataAvailability = "available for order"
	ProductitemlandingpagedataAvailabilityDiscontinued      ProductitemlandingpagedataAvailability = "discontinued"
	ProductitemlandingpagedataAvailabilityInStock           ProductitemlandingpagedataAvailability = "in stock"
	ProductitemlandingpagedataAvailabilityMarkAsExpired     ProductitemlandingpagedataAvailability = "mark_as_expired"
	ProductitemlandingpagedataAvailabilityMarkAsSold        ProductitemlandingpagedataAvailability = "mark_as_sold"
	ProductitemlandingpagedataAvailabilityOutOfStock        ProductitemlandingpagedataAvailability = "out of stock"
	ProductitemlandingpagedataAvailabilityPending           ProductitemlandingpagedataAvailability = "pending"
	ProductitemlandingpagedataAvailabilityPreorder          ProductitemlandingpagedataAvailability = "preorder"
)

func (value ProductitemlandingpagedataAvailability) String() string {
	return string(value)
}
