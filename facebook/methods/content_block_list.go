package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetContentBlockListAppliedAdAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetContentBlockListAppliedAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetContentBlockListAppliedAdAccounts(ctx context.Context, client *core.Client, id string, params GetContentBlockListAppliedAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "applied_ad_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetContentBlockListFacebookContentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetContentBlockListFacebookContentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetContentBlockListFacebookContent(ctx context.Context, client *core.Client, id string, params GetContentBlockListFacebookContentParams) (*core.Cursor[objects.Postish], error) {
	var out core.Cursor[objects.Postish]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "facebook_content"), params.ToParams(), &out)
	return &out, err
}

type GetContentBlockListInstagramContentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetContentBlockListInstagramContentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetContentBlockListInstagramContent(ctx context.Context, client *core.Client, id string, params GetContentBlockListInstagramContentParams) (*core.Cursor[objects.IGMedia], error) {
	var out core.Cursor[objects.IGMedia]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "instagram_content"), params.ToParams(), &out)
	return &out, err
}

type GetContentBlockListParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetContentBlockListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetContentBlockList(ctx context.Context, client *core.Client, id string, params GetContentBlockListParams) (*objects.ContentBlockList, error) {
	var out objects.ContentBlockList
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
