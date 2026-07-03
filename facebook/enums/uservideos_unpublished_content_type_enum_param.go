package enums

type UservideosUnpublishedContentTypeEnumParam string

const (
	UservideosUnpublishedContentTypeEnumParamAdsPost                  UservideosUnpublishedContentTypeEnumParam = "ADS_POST"
	UservideosUnpublishedContentTypeEnumParamDraft                    UservideosUnpublishedContentTypeEnumParam = "DRAFT"
	UservideosUnpublishedContentTypeEnumParamInlineCreated            UservideosUnpublishedContentTypeEnumParam = "INLINE_CREATED"
	UservideosUnpublishedContentTypeEnumParamPublished                UservideosUnpublishedContentTypeEnumParam = "PUBLISHED"
	UservideosUnpublishedContentTypeEnumParamPublishPending           UservideosUnpublishedContentTypeEnumParam = "PUBLISH_PENDING"
	UservideosUnpublishedContentTypeEnumParamReviewableBrandedContent UservideosUnpublishedContentTypeEnumParam = "REVIEWABLE_BRANDED_CONTENT"
	UservideosUnpublishedContentTypeEnumParamScheduled                UservideosUnpublishedContentTypeEnumParam = "SCHEDULED"
	UservideosUnpublishedContentTypeEnumParamScheduledRecurring       UservideosUnpublishedContentTypeEnumParam = "SCHEDULED_RECURRING"
)

func (value UservideosUnpublishedContentTypeEnumParam) String() string {
	return string(value)
}
