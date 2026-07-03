package enums

type ProductitemImageFetchStatus string

const (
	ProductitemImageFetchStatusDirectUpload ProductitemImageFetchStatus = "DIRECT_UPLOAD"
	ProductitemImageFetchStatusFetched      ProductitemImageFetchStatus = "FETCHED"
	ProductitemImageFetchStatusFetchFailed  ProductitemImageFetchStatus = "FETCH_FAILED"
	ProductitemImageFetchStatusNoStatus     ProductitemImageFetchStatus = "NO_STATUS"
	ProductitemImageFetchStatusOutdated     ProductitemImageFetchStatus = "OUTDATED"
	ProductitemImageFetchStatusPartialFetch ProductitemImageFetchStatus = "PARTIAL_FETCH"
)

func (value ProductitemImageFetchStatus) String() string {
	return string(value)
}
