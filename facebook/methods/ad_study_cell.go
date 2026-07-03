package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdStudyCellAdaccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyCellAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyCellAdaccounts(ctx context.Context, client *core.Client, id string, params GetAdStudyCellAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adaccounts"), params.ToParams(), &out)
	return &out, err
}

type GetAdStudyCellAdsetsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyCellAdsetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyCellAdsets(ctx context.Context, client *core.Client, id string, params GetAdStudyCellAdsetsParams) (*core.Cursor[objects.AdSet], error) {
	var out core.Cursor[objects.AdSet]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adsets"), params.ToParams(), &out)
	return &out, err
}

type GetAdStudyCellCampaignsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyCellCampaignsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyCellCampaigns(ctx context.Context, client *core.Client, id string, params GetAdStudyCellCampaignsParams) (*core.Cursor[objects.Campaign], error) {
	var out core.Cursor[objects.Campaign]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "campaigns"), params.ToParams(), &out)
	return &out, err
}

type GetAdStudyCellParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyCellParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyCell(ctx context.Context, client *core.Client, id string, params GetAdStudyCellParams) (*objects.AdStudyCell, error) {
	var out objects.AdStudyCell
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateAdStudyCellParams struct {
	Adaccounts       *[]uint64                          `facebook:"adaccounts"`
	Ads              *[]string                          `facebook:"ads"`
	Adsets           *[]string                          `facebook:"adsets"`
	Campaigns        *[]string                          `facebook:"campaigns"`
	CreationTemplate *enums.AdstudycellCreationTemplate `facebook:"creation_template"`
	Description      *string                            `facebook:"description"`
	Name             *string                            `facebook:"name"`
	Extra            core.Params                        `facebook:"-"`
}

func (params UpdateAdStudyCellParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Adaccounts != nil {
		out["adaccounts"] = *params.Adaccounts
	}
	if params.Ads != nil {
		out["ads"] = *params.Ads
	}
	if params.Adsets != nil {
		out["adsets"] = *params.Adsets
	}
	if params.Campaigns != nil {
		out["campaigns"] = *params.Campaigns
	}
	if params.CreationTemplate != nil {
		out["creation_template"] = *params.CreationTemplate
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	return out
}

func UpdateAdStudyCell(ctx context.Context, client *core.Client, id string, params UpdateAdStudyCellParams) (*objects.AdStudyCell, error) {
	var out objects.AdStudyCell
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
