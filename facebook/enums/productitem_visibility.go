package enums

type ProductitemVisibility string

const (
	ProductitemVisibilityPublished ProductitemVisibility = "published"
	ProductitemVisibilityStaging   ProductitemVisibility = "staging"
)

func (value ProductitemVisibility) String() string {
	return string(value)
}
