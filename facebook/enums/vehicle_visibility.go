package enums

type VehicleVisibility string

const (
	VehicleVisibilityPublished VehicleVisibility = "PUBLISHED"
	VehicleVisibilityStaging   VehicleVisibility = "STAGING"
)

func (value VehicleVisibility) String() string {
	return string(value)
}
