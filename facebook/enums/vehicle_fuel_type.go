package enums

type VehicleFuelType string

const (
	VehicleFuelTypeDiesel       VehicleFuelType = "DIESEL"
	VehicleFuelTypeElectric     VehicleFuelType = "ELECTRIC"
	VehicleFuelTypeFlex         VehicleFuelType = "FLEX"
	VehicleFuelTypeGasoline     VehicleFuelType = "GASOLINE"
	VehicleFuelTypeHybrid       VehicleFuelType = "HYBRID"
	VehicleFuelTypeNone         VehicleFuelType = "NONE"
	VehicleFuelTypeOther        VehicleFuelType = "OTHER"
	VehicleFuelTypePetrol       VehicleFuelType = "PETROL"
	VehicleFuelTypePluginHybrid VehicleFuelType = "PLUGIN_HYBRID"
)

func (value VehicleFuelType) String() string {
	return string(value)
}
