package enums

type ProductitemAvailabilityEnum string

const (
	ProductitemAvailabilityEnumAvailableForOrder ProductitemAvailabilityEnum = "available for order"
	ProductitemAvailabilityEnumDiscontinued      ProductitemAvailabilityEnum = "discontinued"
	ProductitemAvailabilityEnumInStock           ProductitemAvailabilityEnum = "in stock"
	ProductitemAvailabilityEnumMarkAsExpired     ProductitemAvailabilityEnum = "mark_as_expired"
	ProductitemAvailabilityEnumMarkAsSold        ProductitemAvailabilityEnum = "mark_as_sold"
	ProductitemAvailabilityEnumOutOfStock        ProductitemAvailabilityEnum = "out of stock"
	ProductitemAvailabilityEnumPending           ProductitemAvailabilityEnum = "pending"
	ProductitemAvailabilityEnumPreorder          ProductitemAvailabilityEnum = "preorder"
)

func (value ProductitemAvailabilityEnum) String() string {
	return string(value)
}
