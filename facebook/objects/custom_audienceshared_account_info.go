package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CustomAudiencesharedAccountInfo struct {
	AccountID     *core.ID `json:"account_id,omitempty"`
	AccountName   *string  `json:"account_name,omitempty"`
	BusinessID    *core.ID `json:"business_id,omitempty"`
	BusinessName  *string  `json:"business_name,omitempty"`
	SharingStatus *string  `json:"sharing_status,omitempty"`
}
