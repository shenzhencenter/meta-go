package enums

type AutomotivemodelVisibility string

const (
	AutomotivemodelVisibilityPublished AutomotivemodelVisibility = "PUBLISHED"
	AutomotivemodelVisibilityStaging   AutomotivemodelVisibility = "STAGING"
)

func (value AutomotivemodelVisibility) String() string {
	return string(value)
}
