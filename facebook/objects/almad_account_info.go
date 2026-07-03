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
