package enums

type ProductcatalogvehiclesFuelTypeEnumParam string

const (
	ProductcatalogvehiclesFuelTypeEnumParamDiesel       ProductcatalogvehiclesFuelTypeEnumParam = "DIESEL"
	ProductcatalogvehiclesFuelTypeEnumParamElectric     ProductcatalogvehiclesFuelTypeEnumParam = "ELECTRIC"
	ProductcatalogvehiclesFuelTypeEnumParamFlex         ProductcatalogvehiclesFuelTypeEnumParam = "FLEX"
	ProductcatalogvehiclesFuelTypeEnumParamGasoline     ProductcatalogvehiclesFuelTypeEnumParam = "GASOLINE"
	ProductcatalogvehiclesFuelTypeEnumParamHybrid       ProductcatalogvehiclesFuelTypeEnumParam = "HYBRID"
	ProductcatalogvehiclesFuelTypeEnumParamNone         ProductcatalogvehiclesFuelTypeEnumParam = "NONE"
	ProductcatalogvehiclesFuelTypeEnumParamOther        ProductcatalogvehiclesFuelTypeEnumParam = "OTHER"
	ProductcatalogvehiclesFuelTypeEnumParamPetrol       ProductcatalogvehiclesFuelTypeEnumParam = "PETROL"
	ProductcatalogvehiclesFuelTypeEnumParamPluginHybrid ProductcatalogvehiclesFuelTypeEnumParam = "PLUGIN_HYBRID"
)

func (value ProductcatalogvehiclesFuelTypeEnumParam) String() string {
	return string(value)
}
