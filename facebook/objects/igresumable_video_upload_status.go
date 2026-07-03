package objects

type IGResumableVideoUploadStatus struct {
	ProcessingPhase *VideoStatusProcessingPhase `json:"processing_phase,omitempty"`
	UploadingPhase  *VideoStatusUploadingPhase  `json:"uploading_phase,omitempty"`
}
