package objects

type ProductFeedUploadErrorReport struct {
	FileHandle   *string `json:"file_handle,omitempty"`
	ReportStatus *string `json:"report_status,omitempty"`
}

var ProductFeedUploadErrorReportFields = struct {
	FileHandle   string
	ReportStatus string
}{
	FileHandle:   "file_handle",
	ReportStatus: "report_status",
}
