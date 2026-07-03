package enums

type ProductcatalogvehiclesVehicleTypeEnumParam string

const (
	ProductcatalogvehiclesVehicleTypeEnumParamBoat       ProductcatalogvehiclesVehicleTypeEnumParam = "BOAT"
	ProductcatalogvehiclesVehicleTypeEnumParamCarTruck   ProductcatalogvehiclesVehicleTypeEnumParam = "CAR_TRUCK"
	ProductcatalogvehiclesVehicleTypeEnumParamCommercial ProductcatalogvehiclesVehicleTypeEnumParam = "COMMERCIAL"
	ProductcatalogvehiclesVehicleTypeEnumParamMotorcycle ProductcatalogvehiclesVehicleTypeEnumParam = "MOTORCYCLE"
	ProductcatalogvehiclesVehicleTypeEnumParamOther      ProductcatalogvehiclesVehicleTypeEnumParam = "OTHER"
	ProductcatalogvehiclesVehicleTypeEnumParamPowersport ProductcatalogvehiclesVehicleTypeEnumParam = "POWERSPORT"
	ProductcatalogvehiclesVehicleTypeEnumParamRvCamper   ProductcatalogvehiclesVehicleTypeEnumParam = "RV_CAMPER"
	ProductcatalogvehiclesVehicleTypeEnumParamTrailer    ProductcatalogvehiclesVehicleTypeEnumParam = "TRAILER"
)

func (value ProductcatalogvehiclesVehicleTypeEnumParam) String() string {
	return string(value)
}
