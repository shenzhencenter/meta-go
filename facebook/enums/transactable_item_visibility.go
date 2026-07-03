package enums

type TransactableitemVisibility string

const (
	TransactableitemVisibilityPublished TransactableitemVisibility = "PUBLISHED"
	TransactableitemVisibilityStaging   TransactableitemVisibility = "STAGING"
)

func (value TransactableitemVisibility) String() string {
	return string(value)
}
