package enums

type PagefeedWithEnumParam string

const (
	PagefeedWithEnumParamLocation PagefeedWithEnumParam = "LOCATION"
)

func (value PagefeedWithEnumParam) String() string {
	return string(value)
}
