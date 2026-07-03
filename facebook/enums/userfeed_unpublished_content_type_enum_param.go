package enums

type UserfeedUnpublishedContentTypeEnumParam string

const (
	UserfeedUnpublishedContentTypeEnumParamAdsPost                  UserfeedUnpublishedContentTypeEnumParam = "ADS_POST"
	UserfeedUnpublishedContentTypeEnumParamDraft                    UserfeedUnpublishedContentTypeEnumParam = "DRAFT"
	UserfeedUnpublishedContentTypeEnumParamInlineCreated            UserfeedUnpublishedContentTypeEnumParam = "INLINE_CREATED"
	UserfeedUnpublishedContentTypeEnumParamPublished                UserfeedUnpublishedContentTypeEnumParam = "PUBLISHED"
	UserfeedUnpublishedContentTypeEnumParamPublishPending           UserfeedUnpublishedContentTypeEnumParam = "PUBLISH_PENDING"
	UserfeedUnpublishedContentTypeEnumParamReviewableBrandedContent UserfeedUnpublishedContentTypeEnumParam = "REVIEWABLE_BRANDED_CONTENT"
	UserfeedUnpublishedContentTypeEnumParamScheduled                UserfeedUnpublishedContentTypeEnumParam = "SCHEDULED"
	UserfeedUnpublishedContentTypeEnumParamScheduledRecurring       UserfeedUnpublishedContentTypeEnumParam = "SCHEDULED_RECURRING"
)

func (value UserfeedUnpublishedContentTypeEnumParam) String() string {
	return string(value)
}
