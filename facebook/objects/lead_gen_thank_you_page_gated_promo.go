package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type LeadGenThankYouPageGatedPromo struct {
	ID              *core.ID `json:"id,omitempty"`
	OnlineOfferURL  *string  `json:"online_offer_url,omitempty"`
	OnlinePromoCode *string  `json:"online_promo_code,omitempty"`
}
