package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCPASBusinessSetupConfigAdAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCPASBusinessSetupConfigAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCPASBusinessSetupConfigAdAccounts(ctx context.Context, client *core.Client, id string, params GetCPASBusinessSetupConfigAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ad_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetCPASBusinessSetupConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCPASBusinessSetupConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCPASBusinessSetupConfig(ctx context.Context, client *core.Client, id string, params GetCPASBusinessSetupConfigParams) (*objects.CPASBusinessSetupConfig, error) {
	var out objects.CPASBusinessSetupConfig
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
