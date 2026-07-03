package enums

type PagevideosUnpublishedContentTypeEnumParam string

const (
	PagevideosUnpublishedContentTypeEnumParamAdsPost                  PagevideosUnpublishedContentTypeEnumParam = "ADS_POST"
	PagevideosUnpublishedContentTypeEnumParamDraft                    PagevideosUnpublishedContentTypeEnumParam = "DRAFT"
	PagevideosUnpublishedContentTypeEnumParamInlineCreated            PagevideosUnpublishedContentTypeEnumParam = "INLINE_CREATED"
	PagevideosUnpublishedContentTypeEnumParamPublished                PagevideosUnpublishedContentTypeEnumParam = "PUBLISHED"
	PagevideosUnpublishedContentTypeEnumParamPublishPending           PagevideosUnpublishedContentTypeEnumParam = "PUBLISH_PENDING"
	PagevideosUnpublishedContentTypeEnumParamReviewableBrandedContent PagevideosUnpublishedContentTypeEnumParam = "REVIEWABLE_BRANDED_CONTENT"
	PagevideosUnpublishedContentTypeEnumParamScheduled                PagevideosUnpublishedContentTypeEnumParam = "SCHEDULED"
	PagevideosUnpublishedContentTypeEnumParamScheduledRecurring       PagevideosUnpublishedContentTypeEnumParam = "SCHEDULED_RECURRING"
)

func (value PagevideosUnpublishedContentTypeEnumParam) String() string {
	return string(value)
}
