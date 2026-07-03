package enums

type UserphotosUnpublishedContentTypeEnumParam string

const (
	UserphotosUnpublishedContentTypeEnumParamAdsPost                  UserphotosUnpublishedContentTypeEnumParam = "ADS_POST"
	UserphotosUnpublishedContentTypeEnumParamDraft                    UserphotosUnpublishedContentTypeEnumParam = "DRAFT"
	UserphotosUnpublishedContentTypeEnumParamInlineCreated            UserphotosUnpublishedContentTypeEnumParam = "INLINE_CREATED"
	UserphotosUnpublishedContentTypeEnumParamPublished                UserphotosUnpublishedContentTypeEnumParam = "PUBLISHED"
	UserphotosUnpublishedContentTypeEnumParamPublishPending           UserphotosUnpublishedContentTypeEnumParam = "PUBLISH_PENDING"
	UserphotosUnpublishedContentTypeEnumParamReviewableBrandedContent UserphotosUnpublishedContentTypeEnumParam = "REVIEWABLE_BRANDED_CONTENT"
	UserphotosUnpublishedContentTypeEnumParamScheduled                UserphotosUnpublishedContentTypeEnumParam = "SCHEDULED"
	UserphotosUnpublishedContentTypeEnumParamScheduledRecurring       UserphotosUnpublishedContentTypeEnumParam = "SCHEDULED_RECURRING"
)

func (value UserphotosUnpublishedContentTypeEnumParam) String() string {
	return string(value)
}
