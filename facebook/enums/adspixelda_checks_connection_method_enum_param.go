package enums

type AdspixeldaChecksConnectionMethodEnumParam string

const (
	AdspixeldaChecksConnectionMethodEnumParamAll     AdspixeldaChecksConnectionMethodEnumParam = "ALL"
	AdspixeldaChecksConnectionMethodEnumParamApp     AdspixeldaChecksConnectionMethodEnumParam = "APP"
	AdspixeldaChecksConnectionMethodEnumParamBrowser AdspixeldaChecksConnectionMethodEnumParam = "BROWSER"
	AdspixeldaChecksConnectionMethodEnumParamServer  AdspixeldaChecksConnectionMethodEnumParam = "SERVER"
)

func (value AdspixeldaChecksConnectionMethodEnumParam) String() string {
	return string(value)
}
