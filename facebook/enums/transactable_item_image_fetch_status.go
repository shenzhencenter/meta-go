package enums

type TransactableitemImageFetchStatus string

const (
	TransactableitemImageFetchStatusDirectUpload TransactableitemImageFetchStatus = "DIRECT_UPLOAD"
	TransactableitemImageFetchStatusFetched      TransactableitemImageFetchStatus = "FETCHED"
	TransactableitemImageFetchStatusFetchFailed  TransactableitemImageFetchStatus = "FETCH_FAILED"
	TransactableitemImageFetchStatusNoStatus     TransactableitemImageFetchStatus = "NO_STATUS"
	TransactableitemImageFetchStatusOutdated     TransactableitemImageFetchStatus = "OUTDATED"
	TransactableitemImageFetchStatusPartialFetch TransactableitemImageFetchStatus = "PARTIAL_FETCH"
)

func (value TransactableitemImageFetchStatus) String() string {
	return string(value)
}
