package objects

type VideoStatus struct {
	CopyrightCheckStatus *VideoCopyrightCheckStatus  `json:"copyright_check_status,omitempty"`
	ProcessingPhase      *VideoStatusProcessingPhase `json:"processing_phase,omitempty"`
	ProcessingProgress   *uint64                     `json:"processing_progress,omitempty"`
	PublishingPhase      *VideoStatusPublishingPhase `json:"publishing_phase,omitempty"`
	UploadingPhase       *VideoStatusUploadingPhase  `json:"uploading_phase,omitempty"`
	VideoStatus          *string                     `json:"video_status,omitempty"`
}
