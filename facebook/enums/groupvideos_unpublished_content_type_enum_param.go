package enums

type GroupvideosUnpublishedContentTypeEnumParam string

const (
	GroupvideosUnpublishedContentTypeEnumParamAdsPost                  GroupvideosUnpublishedContentTypeEnumParam = "ADS_POST"
	GroupvideosUnpublishedContentTypeEnumParamDraft                    GroupvideosUnpublishedContentTypeEnumParam = "DRAFT"
	GroupvideosUnpublishedContentTypeEnumParamInlineCreated            GroupvideosUnpublishedContentTypeEnumParam = "INLINE_CREATED"
	GroupvideosUnpublishedContentTypeEnumParamPublished                GroupvideosUnpublishedContentTypeEnumParam = "PUBLISHED"
	GroupvideosUnpublishedContentTypeEnumParamPublishPending           GroupvideosUnpublishedContentTypeEnumParam = "PUBLISH_PENDING"
	GroupvideosUnpublishedContentTypeEnumParamReviewableBrandedContent GroupvideosUnpublishedContentTypeEnumParam = "REVIEWABLE_BRANDED_CONTENT"
	GroupvideosUnpublishedContentTypeEnumParamScheduled                GroupvideosUnpublishedContentTypeEnumParam = "SCHEDULED"
	GroupvideosUnpublishedContentTypeEnumParamScheduledRecurring       GroupvideosUnpublishedContentTypeEnumParam = "SCHEDULED_RECURRING"
)

func (value GroupvideosUnpublishedContentTypeEnumParam) String() string {
	return string(value)
}
