package enums

type FlightVisibility string

const (
	FlightVisibilityPublished FlightVisibility = "PUBLISHED"
	FlightVisibilityStaging   FlightVisibility = "STAGING"
)

func (value FlightVisibility) String() string {
	return string(value)
}
