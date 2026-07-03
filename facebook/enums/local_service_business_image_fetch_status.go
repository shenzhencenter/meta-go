package enums

type LocalservicebusinessImageFetchStatus string

const (
	LocalservicebusinessImageFetchStatusDirectUpload LocalservicebusinessImageFetchStatus = "DIRECT_UPLOAD"
	LocalservicebusinessImageFetchStatusFetched      LocalservicebusinessImageFetchStatus = "FETCHED"
	LocalservicebusinessImageFetchStatusFetchFailed  LocalservicebusinessImageFetchStatus = "FETCH_FAILED"
	LocalservicebusinessImageFetchStatusNoStatus     LocalservicebusinessImageFetchStatus = "NO_STATUS"
	LocalservicebusinessImageFetchStatusOutdated     LocalservicebusinessImageFetchStatus = "OUTDATED"
	LocalservicebusinessImageFetchStatusPartialFetch LocalservicebusinessImageFetchStatus = "PARTIAL_FETCH"
)

func (value LocalservicebusinessImageFetchStatus) String() string {
	return string(value)
}
