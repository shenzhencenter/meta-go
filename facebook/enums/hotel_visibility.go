package enums

type HotelVisibility string

const (
	HotelVisibilityPublished HotelVisibility = "PUBLISHED"
	HotelVisibilityStaging   HotelVisibility = "STAGING"
)

func (value HotelVisibility) String() string {
	return string(value)
}
