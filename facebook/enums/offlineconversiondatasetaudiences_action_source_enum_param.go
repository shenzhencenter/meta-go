package enums

type OfflineconversiondatasetaudiencesActionSourceEnumParam string

const (
	OfflineconversiondatasetaudiencesActionSourceEnumParamPhysicalStore OfflineconversiondatasetaudiencesActionSourceEnumParam = "PHYSICAL_STORE"
	OfflineconversiondatasetaudiencesActionSourceEnumParamWebsite       OfflineconversiondatasetaudiencesActionSourceEnumParam = "WEBSITE"
)

func (value OfflineconversiondatasetaudiencesActionSourceEnumParam) String() string {
	return string(value)
}
