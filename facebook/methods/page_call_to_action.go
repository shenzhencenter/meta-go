package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeletePageCallToActionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeletePageCallToActionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeletePageCallToAction(ctx context.Context, client *core.Client, id string, params DeletePageCallToActionParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetPageCallToActionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageCallToActionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageCallToAction(ctx context.Context, client *core.Client, id string, params GetPageCallToActionParams) (*objects.PageCallToAction, error) {
	var out objects.PageCallToAction
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdatePageCallToActionParams struct {
	AndroidAppID           *core.ID                                      `facebook:"android_app_id"`
	AndroidDestinationType *enums.PagecalltoactionAndroidDestinationType `facebook:"android_destination_type"`
	AndroidPackageName     *string                                       `facebook:"android_package_name"`
	AndroidURL             *string                                       `facebook:"android_url"`
	EmailAddress           *string                                       `facebook:"email_address"`
	IntlNumberWithPlus     *string                                       `facebook:"intl_number_with_plus"`
	IphoneAppID            *core.ID                                      `facebook:"iphone_app_id"`
	IphoneDestinationType  *enums.PagecalltoactionIphoneDestinationType  `facebook:"iphone_destination_type"`
	IphoneURL              *string                                       `facebook:"iphone_url"`
	Type                   *enums.PagecalltoactionType                   `facebook:"type"`
	WebDestinationType     *enums.PagecalltoactionWebDestinationType     `facebook:"web_destination_type"`
	WebURL                 *string                                       `facebook:"web_url"`
	Extra                  core.Params                                   `facebook:"-"`
}

func (params UpdatePageCallToActionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AndroidAppID != nil {
		out["android_app_id"] = *params.AndroidAppID
	}
	if params.AndroidDestinationType != nil {
		out["android_destination_type"] = *params.AndroidDestinationType
	}
	if params.AndroidPackageName != nil {
		out["android_package_name"] = *params.AndroidPackageName
	}
	if params.AndroidURL != nil {
		out["android_url"] = *params.AndroidURL
	}
	if params.EmailAddress != nil {
		out["email_address"] = *params.EmailAddress
	}
	if params.IntlNumberWithPlus != nil {
		out["intl_number_with_plus"] = *params.IntlNumberWithPlus
	}
	if params.IphoneAppID != nil {
		out["iphone_app_id"] = *params.IphoneAppID
	}
	if params.IphoneDestinationType != nil {
		out["iphone_destination_type"] = *params.IphoneDestinationType
	}
	if params.IphoneURL != nil {
		out["iphone_url"] = *params.IphoneURL
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	if params.WebDestinationType != nil {
		out["web_destination_type"] = *params.WebDestinationType
	}
	if params.WebURL != nil {
		out["web_url"] = *params.WebURL
	}
	return out
}

func UpdatePageCallToAction(ctx context.Context, client *core.Client, id string, params UpdatePageCallToActionParams) (*objects.PageCallToAction, error) {
	var out objects.PageCallToAction
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
