package enums

type ApplicationdaChecksConnectionMethodEnumParam string

const (
	ApplicationdaChecksConnectionMethodEnumParamAll     ApplicationdaChecksConnectionMethodEnumParam = "ALL"
	ApplicationdaChecksConnectionMethodEnumParamApp     ApplicationdaChecksConnectionMethodEnumParam = "APP"
	ApplicationdaChecksConnectionMethodEnumParamBrowser ApplicationdaChecksConnectionMethodEnumParam = "BROWSER"
	ApplicationdaChecksConnectionMethodEnumParamServer  ApplicationdaChecksConnectionMethodEnumParam = "SERVER"
)

func (value ApplicationdaChecksConnectionMethodEnumParam) String() string {
	return string(value)
}
