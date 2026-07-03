package enums

type GroupphotosUnpublishedContentTypeEnumParam string

const (
	GroupphotosUnpublishedContentTypeEnumParamAdsPost                  GroupphotosUnpublishedContentTypeEnumParam = "ADS_POST"
	GroupphotosUnpublishedContentTypeEnumParamDraft                    GroupphotosUnpublishedContentTypeEnumParam = "DRAFT"
	GroupphotosUnpublishedContentTypeEnumParamInlineCreated            GroupphotosUnpublishedContentTypeEnumParam = "INLINE_CREATED"
	GroupphotosUnpublishedContentTypeEnumParamPublished                GroupphotosUnpublishedContentTypeEnumParam = "PUBLISHED"
	GroupphotosUnpublishedContentTypeEnumParamPublishPending           GroupphotosUnpublishedContentTypeEnumParam = "PUBLISH_PENDING"
	GroupphotosUnpublishedContentTypeEnumParamReviewableBrandedContent GroupphotosUnpublishedContentTypeEnumParam = "REVIEWABLE_BRANDED_CONTENT"
	GroupphotosUnpublishedContentTypeEnumParamScheduled                GroupphotosUnpublishedContentTypeEnumParam = "SCHEDULED"
	GroupphotosUnpublishedContentTypeEnumParamScheduledRecurring       GroupphotosUnpublishedContentTypeEnumParam = "SCHEDULED_RECURRING"
)

func (value GroupphotosUnpublishedContentTypeEnumParam) String() string {
	return string(value)
}
