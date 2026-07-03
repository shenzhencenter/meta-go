package enums

type AdaccountvideoAdsVideoStateEnumParam string

const (
	AdaccountvideoAdsVideoStateEnumParamDraft     AdaccountvideoAdsVideoStateEnumParam = "DRAFT"
	AdaccountvideoAdsVideoStateEnumParamPublished AdaccountvideoAdsVideoStateEnumParam = "PUBLISHED"
	AdaccountvideoAdsVideoStateEnumParamScheduled AdaccountvideoAdsVideoStateEnumParam = "SCHEDULED"
)

func (value AdaccountvideoAdsVideoStateEnumParam) String() string {
	return string(value)
}
