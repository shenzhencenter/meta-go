package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CustomAudiencesharedAccountInfo struct {
	AccountID     *core.ID `json:"account_id,omitempty"`
	AccountName   *string  `json:"account_name,omitempty"`
	BusinessID    *core.ID `json:"business_id,omitempty"`
	BusinessName  *string  `json:"business_name,omitempty"`
	SharingStatus *string  `json:"sharing_status,omitempty"`
}

var CustomAudiencesharedAccountInfoFields = struct {
	AccountID     string
	AccountName   string
	BusinessID    string
	BusinessName  string
	SharingStatus string
}{
	AccountID:     "account_id",
	AccountName:   "account_name",
	BusinessID:    "business_id",
	BusinessName:  "business_name",
	SharingStatus: "sharing_status",
}
