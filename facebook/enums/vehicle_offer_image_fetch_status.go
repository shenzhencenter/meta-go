package enums

type VehicleofferImageFetchStatus string

const (
	VehicleofferImageFetchStatusDirectUpload VehicleofferImageFetchStatus = "DIRECT_UPLOAD"
	VehicleofferImageFetchStatusFetched      VehicleofferImageFetchStatus = "FETCHED"
	VehicleofferImageFetchStatusFetchFailed  VehicleofferImageFetchStatus = "FETCH_FAILED"
	VehicleofferImageFetchStatusNoStatus     VehicleofferImageFetchStatus = "NO_STATUS"
	VehicleofferImageFetchStatusOutdated     VehicleofferImageFetchStatus = "OUTDATED"
	VehicleofferImageFetchStatusPartialFetch VehicleofferImageFetchStatus = "PARTIAL_FETCH"
)

func (value VehicleofferImageFetchStatus) String() string {
	return string(value)
}
