package enums

type VehicleDrivetrain string

const (
	VehicleDrivetrainAwd    VehicleDrivetrain = "AWD"
	VehicleDrivetrainFourWd VehicleDrivetrain = "FOUR_WD"
	VehicleDrivetrainFwd    VehicleDrivetrain = "FWD"
	VehicleDrivetrainNone   VehicleDrivetrain = "NONE"
	VehicleDrivetrainOther  VehicleDrivetrain = "OTHER"
	VehicleDrivetrainRwd    VehicleDrivetrain = "RWD"
	VehicleDrivetrainTwoWd  VehicleDrivetrain = "TWO_WD"
)

func (value VehicleDrivetrain) String() string {
	return string(value)
}
