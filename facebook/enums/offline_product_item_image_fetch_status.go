package enums

type OfflineproductitemImageFetchStatus string

const (
	OfflineproductitemImageFetchStatusDirectUpload OfflineproductitemImageFetchStatus = "DIRECT_UPLOAD"
	OfflineproductitemImageFetchStatusFetched      OfflineproductitemImageFetchStatus = "FETCHED"
	OfflineproductitemImageFetchStatusFetchFailed  OfflineproductitemImageFetchStatus = "FETCH_FAILED"
	OfflineproductitemImageFetchStatusNoStatus     OfflineproductitemImageFetchStatus = "NO_STATUS"
	OfflineproductitemImageFetchStatusOutdated     OfflineproductitemImageFetchStatus = "OUTDATED"
	OfflineproductitemImageFetchStatusPartialFetch OfflineproductitemImageFetchStatus = "PARTIAL_FETCH"
)

func (value OfflineproductitemImageFetchStatus) String() string {
	return string(value)
}
