package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type CloudGame struct {
	ID                     *core.ID   `json:"id,omitempty"`
	Name                   *string    `json:"name,omitempty"`
	Owner                  *Profile   `json:"owner,omitempty"`
	PlayableAdFileSize     *uint64    `json:"playable_ad_file_size,omitempty"`
	PlayableAdOrientation  *string    `json:"playable_ad_orientation,omitempty"`
	PlayableAdPackageName  *string    `json:"playable_ad_package_name,omitempty"`
	PlayableAdRejectReason *string    `json:"playable_ad_reject_reason,omitempty"`
	PlayableAdStatus       *string    `json:"playable_ad_status,omitempty"`
	PlayableAdUploadTime   *time.Time `json:"playable_ad_upload_time,omitempty"`
}
