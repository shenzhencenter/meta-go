package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AudiencePermission struct {
	Audience         *CustomAudience `json:"audience,omitempty"`
	ShareAccountID   *core.ID        `json:"share_account_id,omitempty"`
	ShareAccountName *string         `json:"share_account_name,omitempty"`
}

var AudiencePermissionFields = struct {
	Audience         string
	ShareAccountID   string
	ShareAccountName string
}{
	Audience:         "audience",
	ShareAccountID:   "share_account_id",
	ShareAccountName: "share_account_name",
}
