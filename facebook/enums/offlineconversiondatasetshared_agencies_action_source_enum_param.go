package enums

type OfflineconversiondatasetsharedAgenciesActionSourceEnumParam string

const (
	OfflineconversiondatasetsharedAgenciesActionSourceEnumParamPhysicalStore OfflineconversiondatasetsharedAgenciesActionSourceEnumParam = "PHYSICAL_STORE"
	OfflineconversiondatasetsharedAgenciesActionSourceEnumParamWebsite       OfflineconversiondatasetsharedAgenciesActionSourceEnumParam = "WEBSITE"
)

func (value OfflineconversiondatasetsharedAgenciesActionSourceEnumParam) String() string {
	return string(value)
}
