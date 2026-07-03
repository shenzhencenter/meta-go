package enums

type ProductcatalogvehiclesDrivetrainEnumParam string

const (
	ProductcatalogvehiclesDrivetrainEnumParamAwd    ProductcatalogvehiclesDrivetrainEnumParam = "AWD"
	ProductcatalogvehiclesDrivetrainEnumParamFourWd ProductcatalogvehiclesDrivetrainEnumParam = "FOUR_WD"
	ProductcatalogvehiclesDrivetrainEnumParamFwd    ProductcatalogvehiclesDrivetrainEnumParam = "FWD"
	ProductcatalogvehiclesDrivetrainEnumParamNone   ProductcatalogvehiclesDrivetrainEnumParam = "NONE"
	ProductcatalogvehiclesDrivetrainEnumParamOther  ProductcatalogvehiclesDrivetrainEnumParam = "OTHER"
	ProductcatalogvehiclesDrivetrainEnumParamRwd    ProductcatalogvehiclesDrivetrainEnumParam = "RWD"
	ProductcatalogvehiclesDrivetrainEnumParamTwoWd  ProductcatalogvehiclesDrivetrainEnumParam = "TWO_WD"
)

func (value ProductcatalogvehiclesDrivetrainEnumParam) String() string {
	return string(value)
}
