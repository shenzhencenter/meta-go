package enums

type PagecommerceOrdersFiltersEnumParam string

const (
	PagecommerceOrdersFiltersEnumParamHasCancellations PagecommerceOrdersFiltersEnumParam = "HAS_CANCELLATIONS"
	PagecommerceOrdersFiltersEnumParamHasFulfillments  PagecommerceOrdersFiltersEnumParam = "HAS_FULFILLMENTS"
	PagecommerceOrdersFiltersEnumParamHasRefunds       PagecommerceOrdersFiltersEnumParam = "HAS_REFUNDS"
	PagecommerceOrdersFiltersEnumParamNoCancellations  PagecommerceOrdersFiltersEnumParam = "NO_CANCELLATIONS"
	PagecommerceOrdersFiltersEnumParamNoRefunds        PagecommerceOrdersFiltersEnumParam = "NO_REFUNDS"
	PagecommerceOrdersFiltersEnumParamNoShipments      PagecommerceOrdersFiltersEnumParam = "NO_SHIPMENTS"
)

func (value PagecommerceOrdersFiltersEnumParam) String() string {
	return string(value)
}
