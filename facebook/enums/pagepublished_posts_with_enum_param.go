package enums

type PagepublishedPostsWithEnumParam string

const (
	PagepublishedPostsWithEnumParamLocation PagepublishedPostsWithEnumParam = "LOCATION"
)

func (value PagepublishedPostsWithEnumParam) String() string {
	return string(value)
}
