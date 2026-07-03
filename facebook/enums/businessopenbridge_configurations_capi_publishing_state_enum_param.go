package enums

type BusinessopenbridgeConfigurationsCapiPublishingStateEnumParam string

const (
	BusinessopenbridgeConfigurationsCapiPublishingStateEnumParamDisabled       BusinessopenbridgeConfigurationsCapiPublishingStateEnumParam = "DISABLED"
	BusinessopenbridgeConfigurationsCapiPublishingStateEnumParamEnabled        BusinessopenbridgeConfigurationsCapiPublishingStateEnumParam = "ENABLED"
	BusinessopenbridgeConfigurationsCapiPublishingStateEnumParamNotInitialized BusinessopenbridgeConfigurationsCapiPublishingStateEnumParam = "NOT_INITIALIZED"
)

func (value BusinessopenbridgeConfigurationsCapiPublishingStateEnumParam) String() string {
	return string(value)
}
