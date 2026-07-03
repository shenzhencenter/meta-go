package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CallAdsPhoneData struct {
	CallAdsPhoneDataUseCase *string  `json:"call_ads_phone_data_use_case,omitempty"`
	CallbackVariant         *string  `json:"callback_variant,omitempty"`
	DestinationWebsiteURL   *string  `json:"destination_website_url,omitempty"`
	ID                      *core.ID `json:"id,omitempty"`
	Page                    *Page    `json:"page,omitempty"`
	PhoneNumber             *string  `json:"phone_number,omitempty"`
}

var CallAdsPhoneDataFields = struct {
	CallAdsPhoneDataUseCase string
	CallbackVariant         string
	DestinationWebsiteURL   string
	ID                      string
	Page                    string
	PhoneNumber             string
}{
	CallAdsPhoneDataUseCase: "call_ads_phone_data_use_case",
	CallbackVariant:         "callback_variant",
	DestinationWebsiteURL:   "destination_website_url",
	ID:                      "id",
	Page:                    "page",
	PhoneNumber:             "phone_number",
}
