package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ALMAdAccountInfo struct {
	AdAccountID        *core.ID   `json:"ad_account_id,omitempty"`
	ID                 *core.ID   `json:"id,omitempty"`
	ManagedBy          *string    `json:"managed_by,omitempty"`
	OwnedBy            *string    `json:"owned_by,omitempty"`
	ParentAdvertiserID *core.ID   `json:"parent_advertiser_id,omitempty"`
	SubVertical        *string    `json:"sub_vertical,omitempty"`
	Tag                *[]string  `json:"tag,omitempty"`
	UserIds            *[]core.ID `json:"user_ids,omitempty"`
	Vertical           *string    `json:"vertical,omitempty"`
}

var ALMAdAccountInfoFields = struct {
	AdAccountID        string
	ID                 string
	ManagedBy          string
	OwnedBy            string
	ParentAdvertiserID string
	SubVertical        string
	Tag                string
	UserIds            string
	Vertical           string
}{
	AdAccountID:        "ad_account_id",
	ID:                 "id",
	ManagedBy:          "managed_by",
	OwnedBy:            "owned_by",
	ParentAdvertiserID: "parent_advertiser_id",
	SubVertical:        "sub_vertical",
	Tag:                "tag",
	UserIds:            "user_ids",
	Vertical:           "vertical",
}
