package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type FundingSourceDetailsCoupon struct {
	Amount                *int                               `json:"amount,omitempty"`
	CampaignIds           *[]core.ID                         `json:"campaign_ids,omitempty"`
	ChildAdAccountID      *core.ID                           `json:"child_ad_account_id,omitempty"`
	ChildBmID             *core.ID                           `json:"child_bm_id,omitempty"`
	CouponID              *core.ID                           `json:"coupon_id,omitempty"`
	CouponTiering         *FundingSourceDetailsCouponTiering `json:"coupon_tiering,omitempty"`
	Currency              *string                            `json:"currency,omitempty"`
	DisplayAmount         *string                            `json:"display_amount,omitempty"`
	Expiration            *time.Time                         `json:"expiration,omitempty"`
	OriginalAmount        *int                               `json:"original_amount,omitempty"`
	OriginalDisplayAmount *string                            `json:"original_display_amount,omitempty"`
	StartDate             *time.Time                         `json:"start_date,omitempty"`
	VendorID              *core.ID                           `json:"vendor_id,omitempty"`
}
