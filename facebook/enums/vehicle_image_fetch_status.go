package enums

type VehicleImageFetchStatus string

const (
	VehicleImageFetchStatusDirectUpload VehicleImageFetchStatus = "DIRECT_UPLOAD"
	VehicleImageFetchStatusFetched      VehicleImageFetchStatus = "FETCHED"
	VehicleImageFetchStatusFetchFailed  VehicleImageFetchStatus = "FETCH_FAILED"
	VehicleImageFetchStatusNoStatus     VehicleImageFetchStatus = "NO_STATUS"
	VehicleImageFetchStatusOutdated     VehicleImageFetchStatus = "OUTDATED"
	VehicleImageFetchStatusPartialFetch VehicleImageFetchStatus = "PARTIAL_FETCH"
)

func (value VehicleImageFetchStatus) String() string {
	return string(value)
}
