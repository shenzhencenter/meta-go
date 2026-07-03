package enums

type AdaccountadvideosUnpublishedContentTypeEnumParam string

const (
	AdaccountadvideosUnpublishedContentTypeEnumParamAdsPost                  AdaccountadvideosUnpublishedContentTypeEnumParam = "ADS_POST"
	AdaccountadvideosUnpublishedContentTypeEnumParamDraft                    AdaccountadvideosUnpublishedContentTypeEnumParam = "DRAFT"
	AdaccountadvideosUnpublishedContentTypeEnumParamInlineCreated            AdaccountadvideosUnpublishedContentTypeEnumParam = "INLINE_CREATED"
	AdaccountadvideosUnpublishedContentTypeEnumParamPublished                AdaccountadvideosUnpublishedContentTypeEnumParam = "PUBLISHED"
	AdaccountadvideosUnpublishedContentTypeEnumParamPublishPending           AdaccountadvideosUnpublishedContentTypeEnumParam = "PUBLISH_PENDING"
	AdaccountadvideosUnpublishedContentTypeEnumParamReviewableBrandedContent AdaccountadvideosUnpublishedContentTypeEnumParam = "REVIEWABLE_BRANDED_CONTENT"
	AdaccountadvideosUnpublishedContentTypeEnumParamScheduled                AdaccountadvideosUnpublishedContentTypeEnumParam = "SCHEDULED"
	AdaccountadvideosUnpublishedContentTypeEnumParamScheduledRecurring       AdaccountadvideosUnpublishedContentTypeEnumParam = "SCHEDULED_RECURRING"
)

func (value AdaccountadvideosUnpublishedContentTypeEnumParam) String() string {
	return string(value)
}
