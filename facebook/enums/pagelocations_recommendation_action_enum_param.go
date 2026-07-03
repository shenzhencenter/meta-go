package enums

type PagelocationsRecommendationActionEnumParam string

const (
	PagelocationsRecommendationActionEnumParamAcceptClosed PagelocationsRecommendationActionEnumParam = "ACCEPT_CLOSED"
	PagelocationsRecommendationActionEnumParamAcceptNew    PagelocationsRecommendationActionEnumParam = "ACCEPT_NEW"
	PagelocationsRecommendationActionEnumParamRejectClosed PagelocationsRecommendationActionEnumParam = "REJECT_CLOSED"
	PagelocationsRecommendationActionEnumParamRejectNew    PagelocationsRecommendationActionEnumParam = "REJECT_NEW"
)

func (value PagelocationsRecommendationActionEnumParam) String() string {
	return string(value)
}
