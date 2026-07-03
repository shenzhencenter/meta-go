package enums

type ProductfeeduploadgetInputMethod string

const (
	ProductfeeduploadgetInputMethodGoogleSheetsFetch        ProductfeeduploadgetInputMethod = "GOOGLE_SHEETS_FETCH"
	ProductfeeduploadgetInputMethodManualUpload             ProductfeeduploadgetInputMethod = "MANUAL_UPLOAD"
	ProductfeeduploadgetInputMethodReuploadExisting         ProductfeeduploadgetInputMethod = "REUPLOAD_EXISTING"
	ProductfeeduploadgetInputMethodReuploadLastFile         ProductfeeduploadgetInputMethod = "REUPLOAD_LAST_FILE"
	ProductfeeduploadgetInputMethodServerFetch              ProductfeeduploadgetInputMethod = "SERVER_FETCH"
	ProductfeeduploadgetInputMethodUserInitiatedServerFetch ProductfeeduploadgetInputMethod = "USER_INITIATED_SERVER_FETCH"
)

func (value ProductfeeduploadgetInputMethod) String() string {
	return string(value)
}
