package objects

type IGResumableVideoUploadStatus struct {
	ProcessingPhase *VideoStatusProcessingPhase `json:"processing_phase,omitempty"`
	UploadingPhase  *VideoStatusUploadingPhase  `json:"uploading_phase,omitempty"`
}

var IGResumableVideoUploadStatusFields = struct {
	ProcessingPhase string
	UploadingPhase  string
}{
	ProcessingPhase: "processing_phase",
	UploadingPhase:  "uploading_phase",
}
