package enums

type PagestoriesStatusEnumParam string

const (
	PagestoriesStatusEnumParamArchived  PagestoriesStatusEnumParam = "ARCHIVED"
	PagestoriesStatusEnumParamPublished PagestoriesStatusEnumParam = "PUBLISHED"
)

func (value PagestoriesStatusEnumParam) String() string {
	return string(value)
}
