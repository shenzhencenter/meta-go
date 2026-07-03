package objects

type VideoStatus struct {
	CopyrightCheckStatus *VideoCopyrightCheckStatus  `json:"copyright_check_status,omitempty"`
	ProcessingPhase      *VideoStatusProcessingPhase `json:"processing_phase,omitempty"`
	ProcessingProgress   *uint64                     `json:"processing_progress,omitempty"`
	PublishingPhase      *VideoStatusPublishingPhase `json:"publishing_phase,omitempty"`
	UploadingPhase       *VideoStatusUploadingPhase  `json:"uploading_phase,omitempty"`
	VideoStatus          *string                     `json:"video_status,omitempty"`
}

var VideoStatusFields = struct {
	CopyrightCheckStatus string
	ProcessingPhase      string
	ProcessingProgress   string
	PublishingPhase      string
	UploadingPhase       string
	VideoStatus          string
}{
	CopyrightCheckStatus: "copyright_check_status",
	ProcessingPhase:      "processing_phase",
	ProcessingProgress:   "processing_progress",
	PublishingPhase:      "publishing_phase",
	UploadingPhase:       "uploading_phase",
	VideoStatus:          "video_status",
}
