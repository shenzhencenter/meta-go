package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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
	PlayableAdUploadTime   *core.Time `json:"playable_ad_upload_time,omitempty"`
}

var CloudGameFields = struct {
	ID                     string
	Name                   string
	Owner                  string
	PlayableAdFileSize     string
	PlayableAdOrientation  string
	PlayableAdPackageName  string
	PlayableAdRejectReason string
	PlayableAdStatus       string
	PlayableAdUploadTime   string
}{
	ID:                     "id",
	Name:                   "name",
	Owner:                  "owner",
	PlayableAdFileSize:     "playable_ad_file_size",
	PlayableAdOrientation:  "playable_ad_orientation",
	PlayableAdPackageName:  "playable_ad_package_name",
	PlayableAdRejectReason: "playable_ad_reject_reason",
	PlayableAdStatus:       "playable_ad_status",
	PlayableAdUploadTime:   "playable_ad_upload_time",
}
