package enums

type AutomotivemodelImageFetchStatus string

const (
	AutomotivemodelImageFetchStatusDirectUpload AutomotivemodelImageFetchStatus = "DIRECT_UPLOAD"
	AutomotivemodelImageFetchStatusFetched      AutomotivemodelImageFetchStatus = "FETCHED"
	AutomotivemodelImageFetchStatusFetchFailed  AutomotivemodelImageFetchStatus = "FETCH_FAILED"
	AutomotivemodelImageFetchStatusNoStatus     AutomotivemodelImageFetchStatus = "NO_STATUS"
	AutomotivemodelImageFetchStatusOutdated     AutomotivemodelImageFetchStatus = "OUTDATED"
	AutomotivemodelImageFetchStatusPartialFetch AutomotivemodelImageFetchStatus = "PARTIAL_FETCH"
)

func (value AutomotivemodelImageFetchStatus) String() string {
	return string(value)
}
