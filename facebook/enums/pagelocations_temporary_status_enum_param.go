package enums

type PagelocationsTemporaryStatusEnumParam string

const (
	PagelocationsTemporaryStatusEnumParamDifferentlyOpen   PagelocationsTemporaryStatusEnumParam = "DIFFERENTLY_OPEN"
	PagelocationsTemporaryStatusEnumParamNoData            PagelocationsTemporaryStatusEnumParam = "NO_DATA"
	PagelocationsTemporaryStatusEnumParamOperatingAsUsual  PagelocationsTemporaryStatusEnumParam = "OPERATING_AS_USUAL"
	PagelocationsTemporaryStatusEnumParamTemporarilyClosed PagelocationsTemporaryStatusEnumParam = "TEMPORARILY_CLOSED"
)

func (value PagelocationsTemporaryStatusEnumParam) String() string {
	return string(value)
}
