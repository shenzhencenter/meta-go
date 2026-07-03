package enums

type BusinessvideosUnpublishedContentTypeEnumParam string

const (
	BusinessvideosUnpublishedContentTypeEnumParamAdsPost                  BusinessvideosUnpublishedContentTypeEnumParam = "ADS_POST"
	BusinessvideosUnpublishedContentTypeEnumParamDraft                    BusinessvideosUnpublishedContentTypeEnumParam = "DRAFT"
	BusinessvideosUnpublishedContentTypeEnumParamInlineCreated            BusinessvideosUnpublishedContentTypeEnumParam = "INLINE_CREATED"
	BusinessvideosUnpublishedContentTypeEnumParamPublished                BusinessvideosUnpublishedContentTypeEnumParam = "PUBLISHED"
	BusinessvideosUnpublishedContentTypeEnumParamPublishPending           BusinessvideosUnpublishedContentTypeEnumParam = "PUBLISH_PENDING"
	BusinessvideosUnpublishedContentTypeEnumParamReviewableBrandedContent BusinessvideosUnpublishedContentTypeEnumParam = "REVIEWABLE_BRANDED_CONTENT"
	BusinessvideosUnpublishedContentTypeEnumParamScheduled                BusinessvideosUnpublishedContentTypeEnumParam = "SCHEDULED"
	BusinessvideosUnpublishedContentTypeEnumParamScheduledRecurring       BusinessvideosUnpublishedContentTypeEnumParam = "SCHEDULED_RECURRING"
)

func (value BusinessvideosUnpublishedContentTypeEnumParam) String() string {
	return string(value)
}
