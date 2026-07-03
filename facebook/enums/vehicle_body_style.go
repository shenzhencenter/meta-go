package enums

type VehicleBodyStyle string

const (
	VehicleBodyStyleConvertible VehicleBodyStyle = "CONVERTIBLE"
	VehicleBodyStyleCoupe       VehicleBodyStyle = "COUPE"
	VehicleBodyStyleCrossover   VehicleBodyStyle = "CROSSOVER"
	VehicleBodyStyleEstate      VehicleBodyStyle = "ESTATE"
	VehicleBodyStyleGrandtourer VehicleBodyStyle = "GRANDTOURER"
	VehicleBodyStyleHatchback   VehicleBodyStyle = "HATCHBACK"
	VehicleBodyStyleMinibus     VehicleBodyStyle = "MINIBUS"
	VehicleBodyStyleMinivan     VehicleBodyStyle = "MINIVAN"
	VehicleBodyStyleMpv         VehicleBodyStyle = "MPV"
	VehicleBodyStyleNone        VehicleBodyStyle = "NONE"
	VehicleBodyStyleOther       VehicleBodyStyle = "OTHER"
	VehicleBodyStylePickup      VehicleBodyStyle = "PICKUP"
	VehicleBodyStyleRoadster    VehicleBodyStyle = "ROADSTER"
	VehicleBodyStyleSaloon      VehicleBodyStyle = "SALOON"
	VehicleBodyStyleSedan       VehicleBodyStyle = "SEDAN"
	VehicleBodyStyleSmallCar    VehicleBodyStyle = "SMALL_CAR"
	VehicleBodyStyleSportscar   VehicleBodyStyle = "SPORTSCAR"
	VehicleBodyStyleSupercar    VehicleBodyStyle = "SUPERCAR"
	VehicleBodyStyleSupermini   VehicleBodyStyle = "SUPERMINI"
	VehicleBodyStyleSuv         VehicleBodyStyle = "SUV"
	VehicleBodyStyleTruck       VehicleBodyStyle = "TRUCK"
	VehicleBodyStyleVan         VehicleBodyStyle = "VAN"
	VehicleBodyStyleWagon       VehicleBodyStyle = "WAGON"
)

func (value VehicleBodyStyle) String() string {
	return string(value)
}
