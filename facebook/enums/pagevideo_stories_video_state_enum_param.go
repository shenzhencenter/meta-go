package enums

type PagevideoStoriesVideoStateEnumParam string

const (
	PagevideoStoriesVideoStateEnumParamDraft     PagevideoStoriesVideoStateEnumParam = "DRAFT"
	PagevideoStoriesVideoStateEnumParamPublished PagevideoStoriesVideoStateEnumParam = "PUBLISHED"
	PagevideoStoriesVideoStateEnumParamScheduled PagevideoStoriesVideoStateEnumParam = "SCHEDULED"
)

func (value PagevideoStoriesVideoStateEnumParam) String() string {
	return string(value)
}
