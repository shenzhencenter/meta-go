package enums

type OfflineconversiondatasetuploadsOrderEnumParam string

const (
	OfflineconversiondatasetuploadsOrderEnumParamAscending  OfflineconversiondatasetuploadsOrderEnumParam = "ASCENDING"
	OfflineconversiondatasetuploadsOrderEnumParamDescending OfflineconversiondatasetuploadsOrderEnumParam = "DESCENDING"
)

func (value OfflineconversiondatasetuploadsOrderEnumParam) String() string {
	return string(value)
}
