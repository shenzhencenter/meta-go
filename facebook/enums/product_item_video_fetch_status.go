package enums

type ProductitemVideoFetchStatus string

const (
	ProductitemVideoFetchStatusDirectUpload ProductitemVideoFetchStatus = "DIRECT_UPLOAD"
	ProductitemVideoFetchStatusFetched      ProductitemVideoFetchStatus = "FETCHED"
	ProductitemVideoFetchStatusFetchFailed  ProductitemVideoFetchStatus = "FETCH_FAILED"
	ProductitemVideoFetchStatusNoStatus     ProductitemVideoFetchStatus = "NO_STATUS"
	ProductitemVideoFetchStatusOutdated     ProductitemVideoFetchStatus = "OUTDATED"
	ProductitemVideoFetchStatusPartialFetch ProductitemVideoFetchStatus = "PARTIAL_FETCH"
)

func (value ProductitemVideoFetchStatus) String() string {
	return string(value)
}
