package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PageCallToAction struct {
	AndroidApp             *Application `json:"android_app,omitempty"`
	AndroidDeeplink        *string      `json:"android_deeplink,omitempty"`
	AndroidDestinationType *string      `json:"android_destination_type,omitempty"`
	AndroidPackageName     *string      `json:"android_package_name,omitempty"`
	AndroidURL             *string      `json:"android_url,omitempty"`
	CreatedTime            *core.Time   `json:"created_time,omitempty"`
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
	UpdatedTime            *core.Time   `json:"updated_time,omitempty"`
	WebDestinationType     *string      `json:"web_destination_type,omitempty"`
	WebURL                 *string      `json:"web_url,omitempty"`
}

var PageCallToActionFields = struct {
	AndroidApp             string
	AndroidDeeplink        string
	AndroidDestinationType string
	AndroidPackageName     string
	AndroidURL             string
	CreatedTime            string
	EmailAddress           string
	From                   string
	ID                     string
	IntlNumberWithPlus     string
	IphoneApp              string
	IphoneDeeplink         string
	IphoneDestinationType  string
	IphoneURL              string
	Status                 string
	Type                   string
	UpdatedTime            string
	WebDestinationType     string
	WebURL                 string
}{
	AndroidApp:             "android_app",
	AndroidDeeplink:        "android_deeplink",
	AndroidDestinationType: "android_destination_type",
	AndroidPackageName:     "android_package_name",
	AndroidURL:             "android_url",
	CreatedTime:            "created_time",
	EmailAddress:           "email_address",
	From:                   "from",
	ID:                     "id",
	IntlNumberWithPlus:     "intl_number_with_plus",
	IphoneApp:              "iphone_app",
	IphoneDeeplink:         "iphone_deeplink",
	IphoneDestinationType:  "iphone_destination_type",
	IphoneURL:              "iphone_url",
	Status:                 "status",
	Type:                   "type",
	UpdatedTime:            "updated_time",
	WebDestinationType:     "web_destination_type",
	WebURL:                 "web_url",
}
