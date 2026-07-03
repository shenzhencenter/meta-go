package enums

type ProductcatalogvehiclesStateOfVehicleEnumParam string

const (
	ProductcatalogvehiclesStateOfVehicleEnumParamCpo  ProductcatalogvehiclesStateOfVehicleEnumParam = "CPO"
	ProductcatalogvehiclesStateOfVehicleEnumParamNew  ProductcatalogvehiclesStateOfVehicleEnumParam = "NEW"
	ProductcatalogvehiclesStateOfVehicleEnumParamUsed ProductcatalogvehiclesStateOfVehicleEnumParam = "USED"
)

func (value ProductcatalogvehiclesStateOfVehicleEnumParam) String() string {
	return string(value)
}
