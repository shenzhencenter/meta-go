package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type PageCallToAction struct {
	AndroidApp             *Application `json:"android_app,omitempty"`
	AndroidDeeplink        *string      `json:"android_deeplink,omitempty"`
	AndroidDestinationType *string      `json:"android_destination_type,omitempty"`
	AndroidPackageName     *string      `json:"android_package_name,omitempty"`
	AndroidURL             *string      `json:"android_url,omitempty"`
	CreatedTime            *time.Time   `json:"created_time,omitempty"`
	EmailAddress           *string      `json:"email_address,omitempty"`
	From                   *Page        `json:"from,omitempty"`
	ID                     *core.ID     `json:"id,omitempty"`
	IntlNumberWithPlus     *string      `json:"intl_number_with_plus,omitempty"`
	IphoneApp              *Application `json:"iphone_app,omitempty"`
	IphoneDeeplink         *string      `json:"iphone_deeplink,omitempty"`
	IphoneDestinationType  *string      `json:"iphone_destination_type,omitempty"`
	IphoneURL              *string      `json:"iphone_url,omitempty"`
	Status                 *string      `json:"status,omitempty"`
	Type                   *string      `json:"type,omitempty"`
	UpdatedTime            *time.Time   `json:"updated_time,omitempty"`
	WebDestinationType     *string      `json:"web_destination_type,omitempty"`
	WebURL                 *string      `json:"web_url,omitempty"`
}
