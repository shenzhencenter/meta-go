package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAssetFeedAdditionalDataReconsiderationOffer struct {
	OfferID   *core.ID `json:"offer_id,omitempty"`
	OfferName *string  `json:"offer_name,omitempty"`
}

var AdAssetFeedAdditionalDataReconsiderationOfferFields = struct {
	OfferID   string
	OfferName string
}{
	OfferID:   "offer_id",
	OfferName: "offer_name",
}
