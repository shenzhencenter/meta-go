package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LeadGenThankYouPageGatedPromo struct {
	ID              *core.ID `json:"id,omitempty"`
	OnlineOfferURL  *string  `json:"online_offer_url,omitempty"`
	OnlinePromoCode *string  `json:"online_promo_code,omitempty"`
}

var LeadGenThankYouPageGatedPromoFields = struct {
	ID              string
	OnlineOfferURL  string
	OnlinePromoCode string
}{
	ID:              "id",
	OnlineOfferURL:  "online_offer_url",
	OnlinePromoCode: "online_promo_code",
}
