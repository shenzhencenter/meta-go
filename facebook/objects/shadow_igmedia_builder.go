package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ShadowIGMediaBuilder struct {
	CopyrightCheckStatus *IGVideoCopyrightCheckStatus  `json:"copyright_check_status,omitempty"`
	ID                   *core.ID                      `json:"id,omitempty"`
	Status               *string                       `json:"status,omitempty"`
	StatusCode           *string                       `json:"status_code,omitempty"`
	VideoStatus          *IGResumableVideoUploadStatus `json:"video_status,omitempty"`
}

var ShadowIGMediaBuilderFields = struct {
	CopyrightCheckStatus string
	ID                   string
	Status               string
	StatusCode           string
	VideoStatus          string
}{
	CopyrightCheckStatus: "copyright_check_status",
	ID:                   "id",
	Status:               "status",
	StatusCode:           "status_code",
	VideoStatus:          "video_status",
}
