package enums

type PagevideoReelsVideoStateEnumParam string

const (
	PagevideoReelsVideoStateEnumParamDraft     PagevideoReelsVideoStateEnumParam = "DRAFT"
	PagevideoReelsVideoStateEnumParamPublished PagevideoReelsVideoStateEnumParam = "PUBLISHED"
	PagevideoReelsVideoStateEnumParamScheduled PagevideoReelsVideoStateEnumParam = "SCHEDULED"
)

func (value PagevideoReelsVideoStateEnumParam) String() string {
	return string(value)
}
