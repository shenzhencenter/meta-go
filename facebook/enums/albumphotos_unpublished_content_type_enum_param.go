package enums

type AlbumphotosUnpublishedContentTypeEnumParam string

const (
	AlbumphotosUnpublishedContentTypeEnumParamAdsPost                  AlbumphotosUnpublishedContentTypeEnumParam = "ADS_POST"
	AlbumphotosUnpublishedContentTypeEnumParamDraft                    AlbumphotosUnpublishedContentTypeEnumParam = "DRAFT"
	AlbumphotosUnpublishedContentTypeEnumParamInlineCreated            AlbumphotosUnpublishedContentTypeEnumParam = "INLINE_CREATED"
	AlbumphotosUnpublishedContentTypeEnumParamPublished                AlbumphotosUnpublishedContentTypeEnumParam = "PUBLISHED"
	AlbumphotosUnpublishedContentTypeEnumParamPublishPending           AlbumphotosUnpublishedContentTypeEnumParam = "PUBLISH_PENDING"
	AlbumphotosUnpublishedContentTypeEnumParamReviewableBrandedContent AlbumphotosUnpublishedContentTypeEnumParam = "REVIEWABLE_BRANDED_CONTENT"
	AlbumphotosUnpublishedContentTypeEnumParamScheduled                AlbumphotosUnpublishedContentTypeEnumParam = "SCHEDULED"
	AlbumphotosUnpublishedContentTypeEnumParamScheduledRecurring       AlbumphotosUnpublishedContentTypeEnumParam = "SCHEDULED_RECURRING"
)

func (value AlbumphotosUnpublishedContentTypeEnumParam) String() string {
	return string(value)
}
