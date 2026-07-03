package objects

type VideoStatusUploadingPhase struct {
	BytesTransferred *uint64             `json:"bytes_transferred,omitempty"`
	Errors           *[]VideoStatusError `json:"errors,omitempty"`
	SourceFileSize   *uint64             `json:"source_file_size,omitempty"`
	Status           *string             `json:"status,omitempty"`
}
