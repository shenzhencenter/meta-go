package objects

type ProductFeedUploadErrorReport struct {
	FileHandle   *string `json:"file_handle,omitempty"`
	ReportStatus *string `json:"report_status,omitempty"`
}
