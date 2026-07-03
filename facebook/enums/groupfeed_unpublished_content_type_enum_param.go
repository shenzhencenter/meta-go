package enums

type GroupfeedUnpublishedContentTypeEnumParam string

const (
	GroupfeedUnpublishedContentTypeEnumParamAdsPost                  GroupfeedUnpublishedContentTypeEnumParam = "ADS_POST"
	GroupfeedUnpublishedContentTypeEnumParamDraft                    GroupfeedUnpublishedContentTypeEnumParam = "DRAFT"
	GroupfeedUnpublishedContentTypeEnumParamInlineCreated            GroupfeedUnpublishedContentTypeEnumParam = "INLINE_CREATED"
	GroupfeedUnpublishedContentTypeEnumParamPublished                GroupfeedUnpublishedContentTypeEnumParam = "PUBLISHED"
	GroupfeedUnpublishedContentTypeEnumParamPublishPending           GroupfeedUnpublishedContentTypeEnumParam = "PUBLISH_PENDING"
	GroupfeedUnpublishedContentTypeEnumParamReviewableBrandedContent GroupfeedUnpublishedContentTypeEnumParam = "REVIEWABLE_BRANDED_CONTENT"
	GroupfeedUnpublishedContentTypeEnumParamScheduled                GroupfeedUnpublishedContentTypeEnumParam = "SCHEDULED"
	GroupfeedUnpublishedContentTypeEnumParamScheduledRecurring       GroupfeedUnpublishedContentTypeEnumParam = "SCHEDULED_RECURRING"
)

func (value GroupfeedUnpublishedContentTypeEnumParam) String() string {
	return string(value)
}
