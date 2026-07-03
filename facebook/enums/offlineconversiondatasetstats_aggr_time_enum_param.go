package enums

type OfflineconversiondatasetstatsAggrTimeEnumParam string

const (
	OfflineconversiondatasetstatsAggrTimeEnumParamEventTime  OfflineconversiondatasetstatsAggrTimeEnumParam = "event_time"
	OfflineconversiondatasetstatsAggrTimeEnumParamUploadTime OfflineconversiondatasetstatsAggrTimeEnumParam = "upload_time"
)

func (value OfflineconversiondatasetstatsAggrTimeEnumParam) String() string {
	return string(value)
}
