package enums

type ProfileProfileType string

const (
	ProfileProfileTypeApplication ProfileProfileType = "application"
	ProfileProfileTypeEvent       ProfileProfileType = "event"
	ProfileProfileTypeGroup       ProfileProfileType = "group"
	ProfileProfileTypePage        ProfileProfileType = "page"
	ProfileProfileTypeUser        ProfileProfileType = "user"
)

func (value ProfileProfileType) String() string {
	return string(value)
}
