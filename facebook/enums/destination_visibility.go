package enums

type DestinationVisibility string

const (
	DestinationVisibilityPublished DestinationVisibility = "PUBLISHED"
	DestinationVisibilityStaging   DestinationVisibility = "STAGING"
)

func (value DestinationVisibility) String() string {
	return string(value)
}
