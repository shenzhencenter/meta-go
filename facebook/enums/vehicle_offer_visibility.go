package enums

type VehicleofferVisibility string

const (
	VehicleofferVisibilityPublished VehicleofferVisibility = "PUBLISHED"
	VehicleofferVisibilityStaging   VehicleofferVisibility = "STAGING"
)

func (value VehicleofferVisibility) String() string {
	return string(value)
}
