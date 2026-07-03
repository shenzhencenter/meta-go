package enums

type ProductcatalogvehiclesAvailabilityEnumParam string

const (
	ProductcatalogvehiclesAvailabilityEnumParamAvailable    ProductcatalogvehiclesAvailabilityEnumParam = "AVAILABLE"
	ProductcatalogvehiclesAvailabilityEnumParamNotAvailable ProductcatalogvehiclesAvailabilityEnumParam = "NOT_AVAILABLE"
	ProductcatalogvehiclesAvailabilityEnumParamPending      ProductcatalogvehiclesAvailabilityEnumParam = "PENDING"
	ProductcatalogvehiclesAvailabilityEnumParamUnknown      ProductcatalogvehiclesAvailabilityEnumParam = "UNKNOWN"
)

func (value ProductcatalogvehiclesAvailabilityEnumParam) String() string {
	return string(value)
}
