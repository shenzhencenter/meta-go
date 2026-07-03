package enums

type ProductgroupproductsAvailabilityEnumParam string

const (
	ProductgroupproductsAvailabilityEnumParamAvailableForOrder ProductgroupproductsAvailabilityEnumParam = "available for order"
	ProductgroupproductsAvailabilityEnumParamDiscontinued      ProductgroupproductsAvailabilityEnumParam = "discontinued"
	ProductgroupproductsAvailabilityEnumParamInStock           ProductgroupproductsAvailabilityEnumParam = "in stock"
	ProductgroupproductsAvailabilityEnumParamMarkAsExpired     ProductgroupproductsAvailabilityEnumParam = "mark_as_expired"
	ProductgroupproductsAvailabilityEnumParamMarkAsSold        ProductgroupproductsAvailabilityEnumParam = "mark_as_sold"
	ProductgroupproductsAvailabilityEnumParamOutOfStock        ProductgroupproductsAvailabilityEnumParam = "out of stock"
	ProductgroupproductsAvailabilityEnumParamPending           ProductgroupproductsAvailabilityEnumParam = "pending"
	ProductgroupproductsAvailabilityEnumParamPreorder          ProductgroupproductsAvailabilityEnumParam = "preorder"
)

func (value ProductgroupproductsAvailabilityEnumParam) String() string {
	return string(value)
}
