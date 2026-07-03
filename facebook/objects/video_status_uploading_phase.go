package objects

type VideoStatusUploadingPhase struct {
	BytesTransferred *uint64             `json:"bytes_transferred,omitempty"`
	Errors           *[]VideoStatusError `json:"errors,omitempty"`
	SourceFileSize   *uint64             `json:"source_file_size,omitempty"`
	Status           *string             `json:"status,omitempty"`
}

var VideoStatusUploadingPhaseFields = struct {
	BytesTransferred string
	Errors           string
	SourceFileSize   string
	Status           string
}{
	BytesTransferred: "bytes_transferred",
	Errors:           "errors",
	SourceFileSize:   "source_file_size",
	Status:           "status",
}
