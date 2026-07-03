package enums

type HotelImageFetchStatus string

const (
	HotelImageFetchStatusDirectUpload HotelImageFetchStatus = "DIRECT_UPLOAD"
	HotelImageFetchStatusFetched      HotelImageFetchStatus = "FETCHED"
	HotelImageFetchStatusFetchFailed  HotelImageFetchStatus = "FETCH_FAILED"
	HotelImageFetchStatusNoStatus     HotelImageFetchStatus = "NO_STATUS"
	HotelImageFetchStatusOutdated     HotelImageFetchStatus = "OUTDATED"
	HotelImageFetchStatusPartialFetch HotelImageFetchStatus = "PARTIAL_FETCH"
)

func (value HotelImageFetchStatus) String() string {
	return string(value)
}
