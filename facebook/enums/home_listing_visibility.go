package enums

type HomelistingVisibility string

const (
	HomelistingVisibilityPublished HomelistingVisibility = "PUBLISHED"
	HomelistingVisibilityStaging   HomelistingVisibility = "STAGING"
)

func (value HomelistingVisibility) String() string {
	return string(value)
}
