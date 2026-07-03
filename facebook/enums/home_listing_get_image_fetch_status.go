package enums

type HomelistinggetImageFetchStatus string

const (
	HomelistinggetImageFetchStatusDirectUpload HomelistinggetImageFetchStatus = "DIRECT_UPLOAD"
	HomelistinggetImageFetchStatusFetched      HomelistinggetImageFetchStatus = "FETCHED"
	HomelistinggetImageFetchStatusFetchFailed  HomelistinggetImageFetchStatus = "FETCH_FAILED"
	HomelistinggetImageFetchStatusNoStatus     HomelistinggetImageFetchStatus = "NO_STATUS"
	HomelistinggetImageFetchStatusOutdated     HomelistinggetImageFetchStatus = "OUTDATED"
	HomelistinggetImageFetchStatusPartialFetch HomelistinggetImageFetchStatus = "PARTIAL_FETCH"
)

func (value HomelistinggetImageFetchStatus) String() string {
	return string(value)
}
