package enums

type LocalservicebusinessVisibility string

const (
	LocalservicebusinessVisibilityPublished LocalservicebusinessVisibility = "PUBLISHED"
	LocalservicebusinessVisibilityStaging   LocalservicebusinessVisibility = "STAGING"
)

func (value LocalservicebusinessVisibility) String() string {
	return string(value)
}
