package enums

type PagecallsActionEnumParam string

const (
	PagecallsActionEnumParamAccept      PagecallsActionEnumParam = "ACCEPT"
	PagecallsActionEnumParamConnect     PagecallsActionEnumParam = "CONNECT"
	PagecallsActionEnumParamMediaUpdate PagecallsActionEnumParam = "MEDIA_UPDATE"
	PagecallsActionEnumParamReject      PagecallsActionEnumParam = "REJECT"
	PagecallsActionEnumParamTerminate   PagecallsActionEnumParam = "TERMINATE"
)

func (value PagecallsActionEnumParam) String() string {
	return string(value)
}
