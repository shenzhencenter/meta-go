package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetProfilePicture(ctx context.Context, client *core.Client, id string, params GetProfilePictureParams) (*core.Cursor[objects.ProfilePictureSource], error) {
	var out core.Cursor[objects.ProfilePictureSource]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "picture"), params.ToParams(), &out)
	return &out, err
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

func GetProfile(ctx context.Context, client *core.Client, id string, params GetProfileParams) (*objects.Profile, error) {
	var out objects.Profile
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
