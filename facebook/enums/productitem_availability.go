package enums

type ProductitemAvailability string

const (
	ProductitemAvailabilityAvailableForOrder ProductitemAvailability = "available for order"
	ProductitemAvailabilityDiscontinued      ProductitemAvailability = "discontinued"
	ProductitemAvailabilityInStock           ProductitemAvailability = "in stock"
	ProductitemAvailabilityMarkAsExpired     ProductitemAvailability = "mark_as_expired"
	ProductitemAvailabilityMarkAsSold        ProductitemAvailability = "mark_as_sold"
	ProductitemAvailabilityOutOfStock        ProductitemAvailability = "out of stock"
	ProductitemAvailabilityPending           ProductitemAvailability = "pending"
	ProductitemAvailabilityPreorder          ProductitemAvailability = "preorder"
)

func (value ProductitemAvailability) String() string {
	return string(value)
}
