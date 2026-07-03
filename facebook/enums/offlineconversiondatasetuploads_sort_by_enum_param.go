package enums

type OfflineconversiondatasetuploadsSortByEnumParam string

const (
	OfflineconversiondatasetuploadsSortByEnumParamAPICalls          OfflineconversiondatasetuploadsSortByEnumParam = "API_CALLS"
	OfflineconversiondatasetuploadsSortByEnumParamCreationTime      OfflineconversiondatasetuploadsSortByEnumParam = "CREATION_TIME"
	OfflineconversiondatasetuploadsSortByEnumParamEventTimeMax      OfflineconversiondatasetuploadsSortByEnumParam = "EVENT_TIME_MAX"
	OfflineconversiondatasetuploadsSortByEnumParamEventTimeMin      OfflineconversiondatasetuploadsSortByEnumParam = "EVENT_TIME_MIN"
	OfflineconversiondatasetuploadsSortByEnumParamFirstUploadTime   OfflineconversiondatasetuploadsSortByEnumParam = "FIRST_UPLOAD_TIME"
	OfflineconversiondatasetuploadsSortByEnumParamIsExcludedForLift OfflineconversiondatasetuploadsSortByEnumParam = "IS_EXCLUDED_FOR_LIFT"
	OfflineconversiondatasetuploadsSortByEnumParamLastUploadTime    OfflineconversiondatasetuploadsSortByEnumParam = "LAST_UPLOAD_TIME"
)

func (value OfflineconversiondatasetuploadsSortByEnumParam) String() string {
	return string(value)
}
