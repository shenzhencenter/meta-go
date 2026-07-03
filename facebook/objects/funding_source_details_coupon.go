package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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
	Expiration            *core.Time                         `json:"expiration,omitempty"`
	OriginalAmount        *int                               `json:"original_amount,omitempty"`
	OriginalDisplayAmount *string                            `json:"original_display_amount,omitempty"`
	StartDate             *core.Time                         `json:"start_date,omitempty"`
	VendorID              *core.ID                           `json:"vendor_id,omitempty"`
}

var FundingSourceDetailsCouponFields = struct {
	Amount                string
	CampaignIds           string
	ChildAdAccountID      string
	ChildBmID             string
	CouponID              string
	CouponTiering         string
	Currency              string
	DisplayAmount         string
	Expiration            string
	OriginalAmount        string
	OriginalDisplayAmount string
	StartDate             string
	VendorID              string
}{
	Amount:                "amount",
	CampaignIds:           "campaign_ids",
	ChildAdAccountID:      "child_ad_account_id",
	ChildBmID:             "child_bm_id",
	CouponID:              "coupon_id",
	CouponTiering:         "coupon_tiering",
	Currency:              "currency",
	DisplayAmount:         "display_amount",
	Expiration:            "expiration",
	OriginalAmount:        "original_amount",
	OriginalDisplayAmount: "original_display_amount",
	StartDate:             "start_date",
	VendorID:              "vendor_id",
}
