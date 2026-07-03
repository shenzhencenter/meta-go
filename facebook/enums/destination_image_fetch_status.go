package enums

type DestinationImageFetchStatus string

const (
	DestinationImageFetchStatusDirectUpload DestinationImageFetchStatus = "DIRECT_UPLOAD"
	DestinationImageFetchStatusFetched      DestinationImageFetchStatus = "FETCHED"
	DestinationImageFetchStatusFetchFailed  DestinationImageFetchStatus = "FETCH_FAILED"
	DestinationImageFetchStatusNoStatus     DestinationImageFetchStatus = "NO_STATUS"
	DestinationImageFetchStatusOutdated     DestinationImageFetchStatus = "OUTDATED"
	DestinationImageFetchStatusPartialFetch DestinationImageFetchStatus = "PARTIAL_FETCH"
)

func (value DestinationImageFetchStatus) String() string {
	return string(value)
}
