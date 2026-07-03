package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ShadowIGMediaBuilder struct {
	CopyrightCheckStatus *IGVideoCopyrightCheckStatus  `json:"copyright_check_status,omitempty"`
	ID                   *core.ID                      `json:"id,omitempty"`
	Status               *string                       `json:"status,omitempty"`
	StatusCode           *string                       `json:"status_code,omitempty"`
	VideoStatus          *IGResumableVideoUploadStatus `json:"video_status,omitempty"`
}
