package enums

type ProductcatalogvehiclesTransmissionEnumParam string

const (
	ProductcatalogvehiclesTransmissionEnumParamAutomatic ProductcatalogvehiclesTransmissionEnumParam = "AUTOMATIC"
	ProductcatalogvehiclesTransmissionEnumParamManual    ProductcatalogvehiclesTransmissionEnumParam = "MANUAL"
	ProductcatalogvehiclesTransmissionEnumParamNone      ProductcatalogvehiclesTransmissionEnumParam = "NONE"
	ProductcatalogvehiclesTransmissionEnumParamOther     ProductcatalogvehiclesTransmissionEnumParam = "OTHER"
)

func (value ProductcatalogvehiclesTransmissionEnumParam) String() string {
	return string(value)
}
