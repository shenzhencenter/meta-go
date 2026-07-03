package enums

type PagevisitorPostsWithEnumParam string

const (
	PagevisitorPostsWithEnumParamLocation PagevisitorPostsWithEnumParam = "LOCATION"
)

func (value PagevisitorPostsWithEnumParam) String() string {
	return string(value)
}
