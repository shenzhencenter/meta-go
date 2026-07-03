package enums

type ProductfeeduploadInputMethod string

const (
	ProductfeeduploadInputMethodGoogleSheetsFetch        ProductfeeduploadInputMethod = "Google Sheets Fetch"
	ProductfeeduploadInputMethodManualUpload             ProductfeeduploadInputMethod = "Manual Upload"
	ProductfeeduploadInputMethodReuploadLastFile         ProductfeeduploadInputMethod = "Reupload Last File"
	ProductfeeduploadInputMethodServerFetch              ProductfeeduploadInputMethod = "Server Fetch"
	ProductfeeduploadInputMethodUserInitiatedServerFetch ProductfeeduploadInputMethod = "User initiated server fetch"
)

func (value ProductfeeduploadInputMethod) String() string {
	return string(value)
}
