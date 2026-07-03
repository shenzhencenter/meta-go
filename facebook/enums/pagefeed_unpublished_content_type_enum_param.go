package enums

type PagefeedUnpublishedContentTypeEnumParam string

const (
	PagefeedUnpublishedContentTypeEnumParamAdsPost                  PagefeedUnpublishedContentTypeEnumParam = "ADS_POST"
	PagefeedUnpublishedContentTypeEnumParamDraft                    PagefeedUnpublishedContentTypeEnumParam = "DRAFT"
	PagefeedUnpublishedContentTypeEnumParamInlineCreated            PagefeedUnpublishedContentTypeEnumParam = "INLINE_CREATED"
	PagefeedUnpublishedContentTypeEnumParamPublished                PagefeedUnpublishedContentTypeEnumParam = "PUBLISHED"
	PagefeedUnpublishedContentTypeEnumParamPublishPending           PagefeedUnpublishedContentTypeEnumParam = "PUBLISH_PENDING"
	PagefeedUnpublishedContentTypeEnumParamReviewableBrandedContent PagefeedUnpublishedContentTypeEnumParam = "REVIEWABLE_BRANDED_CONTENT"
	PagefeedUnpublishedContentTypeEnumParamScheduled                PagefeedUnpublishedContentTypeEnumParam = "SCHEDULED"
	PagefeedUnpublishedContentTypeEnumParamScheduledRecurring       PagefeedUnpublishedContentTypeEnumParam = "SCHEDULED_RECURRING"
)

func (value PagefeedUnpublishedContentTypeEnumParam) String() string {
	return string(value)
}
