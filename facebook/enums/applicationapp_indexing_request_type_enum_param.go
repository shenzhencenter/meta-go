package enums

type ApplicationappIndexingRequestTypeEnumParam string

const (
	ApplicationappIndexingRequestTypeEnumParamAppIndexing    ApplicationappIndexingRequestTypeEnumParam = "APP_INDEXING"
	ApplicationappIndexingRequestTypeEnumParamButtonSampling ApplicationappIndexingRequestTypeEnumParam = "BUTTON_SAMPLING"
	ApplicationappIndexingRequestTypeEnumParamPlugin         ApplicationappIndexingRequestTypeEnumParam = "PLUGIN"
)

func (value ApplicationappIndexingRequestTypeEnumParam) String() string {
	return string(value)
}
