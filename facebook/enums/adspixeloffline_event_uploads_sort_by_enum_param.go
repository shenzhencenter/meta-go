package enums

type AdspixelofflineEventUploadsSortByEnumParam string

const (
	AdspixelofflineEventUploadsSortByEnumParamAPICalls          AdspixelofflineEventUploadsSortByEnumParam = "API_CALLS"
	AdspixelofflineEventUploadsSortByEnumParamCreationTime      AdspixelofflineEventUploadsSortByEnumParam = "CREATION_TIME"
	AdspixelofflineEventUploadsSortByEnumParamEventTimeMax      AdspixelofflineEventUploadsSortByEnumParam = "EVENT_TIME_MAX"
	AdspixelofflineEventUploadsSortByEnumParamEventTimeMin      AdspixelofflineEventUploadsSortByEnumParam = "EVENT_TIME_MIN"
	AdspixelofflineEventUploadsSortByEnumParamFirstUploadTime   AdspixelofflineEventUploadsSortByEnumParam = "FIRST_UPLOAD_TIME"
	AdspixelofflineEventUploadsSortByEnumParamIsExcludedForLift AdspixelofflineEventUploadsSortByEnumParam = "IS_EXCLUDED_FOR_LIFT"
	AdspixelofflineEventUploadsSortByEnumParamLastUploadTime    AdspixelofflineEventUploadsSortByEnumParam = "LAST_UPLOAD_TIME"
)

func (value AdspixelofflineEventUploadsSortByEnumParam) String() string {
	return string(value)
}
