package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProfilePictureParams struct {
	Height   *int                               `facebook:"height"`
	Redirect *bool                              `facebook:"redirect"`
	Type     *enums.ProfilepictureTypeEnumParam `facebook:"type"`
	Width    *int                               `facebook:"width"`
	Extra    core.Params                        `facebook:"-"`
}

func (params GetProfilePictureParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Height != nil {
		out["height"] = *params.Height
	}
	if params.Redirect != nil {
		out["redirect"] = *params.Redirect
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	if params.Width != nil {
		out["width"] = *params.Width
	}
	return out
}

func GetProfilePictureBatchCall(id string, params GetProfilePictureParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "picture"), params.ToParams(), options...)
}

func NewGetProfilePictureBatchRequest(id string, params GetProfilePictureParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProfilePictureSource]] {
	return core.NewBatchRequest[core.Cursor[objects.ProfilePictureSource]](GetProfilePictureBatchCall(id, params, options...))
}

func DecodeGetProfilePictureBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProfilePictureSource], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProfilePictureSource]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProfilePictureWithResponse(ctx context.Context, client *core.Client, id string, params GetProfilePictureParams) (*core.Cursor[objects.ProfilePictureSource], *core.Response, error) {
	var out core.Cursor[objects.ProfilePictureSource]
	call := GetProfilePictureBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetProfilePicture(ctx context.Context, client *core.Client, id string, params GetProfilePictureParams) (*core.Cursor[objects.ProfilePictureSource], error) {
	out, _, err := GetProfilePictureWithResponse(ctx, client, id, params)
	return out, err
}

type GetProfileParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProfileParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProfileBatchCall(id string, params GetProfileParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetProfileBatchRequest(id string, params GetProfileParams, options ...core.BatchOption) *core.BatchRequest[objects.Profile] {
	return core.NewBatchRequest[objects.Profile](GetProfileBatchCall(id, params, options...))
}

func DecodeGetProfileBatchResponse(response *core.BatchResponse) (*objects.Profile, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Profile
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProfileWithResponse(ctx context.Context, client *core.Client, id string, params GetProfileParams) (*objects.Profile, *core.Response, error) {
	var out objects.Profile
	call := GetProfileBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetProfile(ctx context.Context, client *core.Client, id string, params GetProfileParams) (*objects.Profile, error) {
	out, _, err := GetProfileWithResponse(ctx, client, id, params)
	return out, err
}
