package enums

type BusinessopenbridgeConfigurationsEventEnrichmentStateEnumParam string

const (
	BusinessopenbridgeConfigurationsEventEnrichmentStateEnumParamNo             BusinessopenbridgeConfigurationsEventEnrichmentStateEnumParam = "NO"
	BusinessopenbridgeConfigurationsEventEnrichmentStateEnumParamNotInitialized BusinessopenbridgeConfigurationsEventEnrichmentStateEnumParam = "NOT_INITIALIZED"
	BusinessopenbridgeConfigurationsEventEnrichmentStateEnumParamYes            BusinessopenbridgeConfigurationsEventEnrichmentStateEnumParam = "YES"
)

func (value BusinessopenbridgeConfigurationsEventEnrichmentStateEnumParam) String() string {
	return string(value)
}
