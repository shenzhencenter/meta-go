package enums

type ProductitemVisibilityEnum string

const (
	ProductitemVisibilityEnumPublished ProductitemVisibilityEnum = "published"
	ProductitemVisibilityEnumStaging   ProductitemVisibilityEnum = "staging"
)

func (value ProductitemVisibilityEnum) String() string {
	return string(value)
}
