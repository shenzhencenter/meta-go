package enums

type ProductcatalogvehiclesBodyStyleEnumParam string

const (
	ProductcatalogvehiclesBodyStyleEnumParamConvertible ProductcatalogvehiclesBodyStyleEnumParam = "CONVERTIBLE"
	ProductcatalogvehiclesBodyStyleEnumParamCoupe       ProductcatalogvehiclesBodyStyleEnumParam = "COUPE"
	ProductcatalogvehiclesBodyStyleEnumParamCrossover   ProductcatalogvehiclesBodyStyleEnumParam = "CROSSOVER"
	ProductcatalogvehiclesBodyStyleEnumParamEstate      ProductcatalogvehiclesBodyStyleEnumParam = "ESTATE"
	ProductcatalogvehiclesBodyStyleEnumParamGrandtourer ProductcatalogvehiclesBodyStyleEnumParam = "GRANDTOURER"
	ProductcatalogvehiclesBodyStyleEnumParamHatchback   ProductcatalogvehiclesBodyStyleEnumParam = "HATCHBACK"
	ProductcatalogvehiclesBodyStyleEnumParamMinibus     ProductcatalogvehiclesBodyStyleEnumParam = "MINIBUS"
	ProductcatalogvehiclesBodyStyleEnumParamMinivan     ProductcatalogvehiclesBodyStyleEnumParam = "MINIVAN"
	ProductcatalogvehiclesBodyStyleEnumParamMpv         ProductcatalogvehiclesBodyStyleEnumParam = "MPV"
	ProductcatalogvehiclesBodyStyleEnumParamNone        ProductcatalogvehiclesBodyStyleEnumParam = "NONE"
	ProductcatalogvehiclesBodyStyleEnumParamOther       ProductcatalogvehiclesBodyStyleEnumParam = "OTHER"
	ProductcatalogvehiclesBodyStyleEnumParamPickup      ProductcatalogvehiclesBodyStyleEnumParam = "PICKUP"
	ProductcatalogvehiclesBodyStyleEnumParamRoadster    ProductcatalogvehiclesBodyStyleEnumParam = "ROADSTER"
	ProductcatalogvehiclesBodyStyleEnumParamSaloon      ProductcatalogvehiclesBodyStyleEnumParam = "SALOON"
	ProductcatalogvehiclesBodyStyleEnumParamSedan       ProductcatalogvehiclesBodyStyleEnumParam = "SEDAN"
	ProductcatalogvehiclesBodyStyleEnumParamSmallCar    ProductcatalogvehiclesBodyStyleEnumParam = "SMALL_CAR"
	ProductcatalogvehiclesBodyStyleEnumParamSportscar   ProductcatalogvehiclesBodyStyleEnumParam = "SPORTSCAR"
	ProductcatalogvehiclesBodyStyleEnumParamSupercar    ProductcatalogvehiclesBodyStyleEnumParam = "SUPERCAR"
	ProductcatalogvehiclesBodyStyleEnumParamSupermini   ProductcatalogvehiclesBodyStyleEnumParam = "SUPERMINI"
	ProductcatalogvehiclesBodyStyleEnumParamSuv         ProductcatalogvehiclesBodyStyleEnumParam = "SUV"
	ProductcatalogvehiclesBodyStyleEnumParamTruck       ProductcatalogvehiclesBodyStyleEnumParam = "TRUCK"
	ProductcatalogvehiclesBodyStyleEnumParamVan         ProductcatalogvehiclesBodyStyleEnumParam = "VAN"
	ProductcatalogvehiclesBodyStyleEnumParamWagon       ProductcatalogvehiclesBodyStyleEnumParam = "WAGON"
)

func (value ProductcatalogvehiclesBodyStyleEnumParam) String() string {
	return string(value)
}
