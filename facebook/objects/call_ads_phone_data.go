package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CallAdsPhoneData struct {
	CallAdsPhoneDataUseCase *string  `json:"call_ads_phone_data_use_case,omitempty"`
	CallbackVariant         *string  `json:"callback_variant,omitempty"`
	DestinationWebsiteURL   *string  `json:"destination_website_url,omitempty"`
	ID                      *core.ID `json:"id,omitempty"`
	Page                    *Page    `json:"page,omitempty"`
	PhoneNumber             *string  `json:"phone_number,omitempty"`
}
