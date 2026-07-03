package enums

type FlightImageFetchStatus string

const (
	FlightImageFetchStatusDirectUpload FlightImageFetchStatus = "DIRECT_UPLOAD"
	FlightImageFetchStatusFetched      FlightImageFetchStatus = "FETCHED"
	FlightImageFetchStatusFetchFailed  FlightImageFetchStatus = "FETCH_FAILED"
	FlightImageFetchStatusNoStatus     FlightImageFetchStatus = "NO_STATUS"
	FlightImageFetchStatusOutdated     FlightImageFetchStatus = "OUTDATED"
	FlightImageFetchStatusPartialFetch FlightImageFetchStatus = "PARTIAL_FETCH"
)

func (value FlightImageFetchStatus) String() string {
	return string(value)
}
