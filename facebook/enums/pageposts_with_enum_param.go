package enums

type PagepostsWithEnumParam string

const (
	PagepostsWithEnumParamLocation PagepostsWithEnumParam = "LOCATION"
)

func (value PagepostsWithEnumParam) String() string {
	return string(value)
}
