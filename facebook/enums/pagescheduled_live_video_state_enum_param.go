package enums

type PagescheduledLiveVideoStateEnumParam string

const (
	PagescheduledLiveVideoStateEnumParamDraft     PagescheduledLiveVideoStateEnumParam = "DRAFT"
	PagescheduledLiveVideoStateEnumParamPublished PagescheduledLiveVideoStateEnumParam = "PUBLISHED"
)

func (value PagescheduledLiveVideoStateEnumParam) String() string {
	return string(value)
}
