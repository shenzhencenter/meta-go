package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdLabelAdcreativesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLabelAdcreativesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLabelAdcreatives(ctx context.Context, client *core.Client, id string, params GetAdLabelAdcreativesParams) (*core.Cursor[objects.AdCreative], error) {
	var out core.Cursor[objects.AdCreative]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adcreatives"), params.ToParams(), &out)
	return &out, err
}

type GetAdLabelAdsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLabelAdsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLabelAds(ctx context.Context, client *core.Client, id string, params GetAdLabelAdsParams) (*core.Cursor[objects.Ad], error) {
	var out core.Cursor[objects.Ad]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ads"), params.ToParams(), &out)
	return &out, err
}

type GetAdLabelAdsetsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLabelAdsetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLabelAdsets(ctx context.Context, client *core.Client, id string, params GetAdLabelAdsetsParams) (*core.Cursor[objects.AdSet], error) {
	var out core.Cursor[objects.AdSet]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adsets"), params.ToParams(), &out)
	return &out, err
}

type GetAdLabelCampaignsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLabelCampaignsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLabelCampaigns(ctx context.Context, client *core.Client, id string, params GetAdLabelCampaignsParams) (*core.Cursor[objects.Campaign], error) {
	var out core.Cursor[objects.Campaign]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "campaigns"), params.ToParams(), &out)
	return &out, err
}

type DeleteAdLabelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdLabel(ctx context.Context, client *core.Client, id string, params DeleteAdLabelParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetAdLabelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLabel(ctx context.Context, client *core.Client, id string, params GetAdLabelParams) (*objects.AdLabel, error) {
	var out objects.AdLabel
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateAdLabelParams struct {
	Name  string      `facebook:"name"`
	Extra core.Params `facebook:"-"`
}

func (params UpdateAdLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["name"] = params.Name
	return out
}

func UpdateAdLabel(ctx context.Context, client *core.Client, id string, params UpdateAdLabelParams) (*objects.AdLabel, error) {
	var out objects.AdLabel
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
