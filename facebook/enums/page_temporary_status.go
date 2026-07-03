package enums

type PageTemporaryStatus string

const (
	PageTemporaryStatusDifferentlyOpen   PageTemporaryStatus = "DIFFERENTLY_OPEN"
	PageTemporaryStatusNoData            PageTemporaryStatus = "NO_DATA"
	PageTemporaryStatusOperatingAsUsual  PageTemporaryStatus = "OPERATING_AS_USUAL"
	PageTemporaryStatusTemporarilyClosed PageTemporaryStatus = "TEMPORARILY_CLOSED"
)

func (value PageTemporaryStatus) String() string {
	return string(value)
}
