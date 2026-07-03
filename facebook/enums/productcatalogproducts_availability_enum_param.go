package enums

type ProductcatalogproductsAvailabilityEnumParam string

const (
	ProductcatalogproductsAvailabilityEnumParamAvailableForOrder ProductcatalogproductsAvailabilityEnumParam = "available for order"
	ProductcatalogproductsAvailabilityEnumParamDiscontinued      ProductcatalogproductsAvailabilityEnumParam = "discontinued"
	ProductcatalogproductsAvailabilityEnumParamInStock           ProductcatalogproductsAvailabilityEnumParam = "in stock"
	ProductcatalogproductsAvailabilityEnumParamMarkAsExpired     ProductcatalogproductsAvailabilityEnumParam = "mark_as_expired"
	ProductcatalogproductsAvailabilityEnumParamMarkAsSold        ProductcatalogproductsAvailabilityEnumParam = "mark_as_sold"
	ProductcatalogproductsAvailabilityEnumParamOutOfStock        ProductcatalogproductsAvailabilityEnumParam = "out of stock"
	ProductcatalogproductsAvailabilityEnumParamPending           ProductcatalogproductsAvailabilityEnumParam = "pending"
	ProductcatalogproductsAvailabilityEnumParamPreorder          ProductcatalogproductsAvailabilityEnumParam = "preorder"
)

func (value ProductcatalogproductsAvailabilityEnumParam) String() string {
	return string(value)
}
