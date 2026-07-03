package enums

type HomelistingImageFetchStatus string

const (
	HomelistingImageFetchStatusDirectUpload HomelistingImageFetchStatus = "DIRECT_UPLOAD"
	HomelistingImageFetchStatusFetched      HomelistingImageFetchStatus = "FETCHED"
	HomelistingImageFetchStatusFetchFailed  HomelistingImageFetchStatus = "FETCH_FAILED"
	HomelistingImageFetchStatusNoStatus     HomelistingImageFetchStatus = "NO_STATUS"
	HomelistingImageFetchStatusOutdated     HomelistingImageFetchStatus = "OUTDATED"
	HomelistingImageFetchStatusPartialFetch HomelistingImageFetchStatus = "PARTIAL_FETCH"
)

func (value HomelistingImageFetchStatus) String() string {
	return string(value)
}
