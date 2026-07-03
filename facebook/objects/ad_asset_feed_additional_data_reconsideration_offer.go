package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdAssetFeedAdditionalDataReconsiderationOffer struct {
	OfferID   *core.ID `json:"offer_id,omitempty"`
	OfferName *string  `json:"offer_name,omitempty"`
}
