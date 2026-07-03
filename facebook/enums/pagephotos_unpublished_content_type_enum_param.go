package enums

type PagephotosUnpublishedContentTypeEnumParam string

const (
	PagephotosUnpublishedContentTypeEnumParamAdsPost                  PagephotosUnpublishedContentTypeEnumParam = "ADS_POST"
	PagephotosUnpublishedContentTypeEnumParamDraft                    PagephotosUnpublishedContentTypeEnumParam = "DRAFT"
	PagephotosUnpublishedContentTypeEnumParamInlineCreated            PagephotosUnpublishedContentTypeEnumParam = "INLINE_CREATED"
	PagephotosUnpublishedContentTypeEnumParamPublished                PagephotosUnpublishedContentTypeEnumParam = "PUBLISHED"
	PagephotosUnpublishedContentTypeEnumParamPublishPending           PagephotosUnpublishedContentTypeEnumParam = "PUBLISH_PENDING"
	PagephotosUnpublishedContentTypeEnumParamReviewableBrandedContent PagephotosUnpublishedContentTypeEnumParam = "REVIEWABLE_BRANDED_CONTENT"
	PagephotosUnpublishedContentTypeEnumParamScheduled                PagephotosUnpublishedContentTypeEnumParam = "SCHEDULED"
	PagephotosUnpublishedContentTypeEnumParamScheduledRecurring       PagephotosUnpublishedContentTypeEnumParam = "SCHEDULED_RECURRING"
)

func (value PagephotosUnpublishedContentTypeEnumParam) String() string {
	return string(value)
}
