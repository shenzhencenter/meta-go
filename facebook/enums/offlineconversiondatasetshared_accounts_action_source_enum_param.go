package enums

type OfflineconversiondatasetsharedAccountsActionSourceEnumParam string

const (
	OfflineconversiondatasetsharedAccountsActionSourceEnumParamPhysicalStore OfflineconversiondatasetsharedAccountsActionSourceEnumParam = "PHYSICAL_STORE"
	OfflineconversiondatasetsharedAccountsActionSourceEnumParamWebsite       OfflineconversiondatasetsharedAccountsActionSourceEnumParam = "WEBSITE"
)

func (value OfflineconversiondatasetsharedAccountsActionSourceEnumParam) String() string {
	return string(value)
}
