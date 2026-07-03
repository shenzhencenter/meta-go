package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type DeleteExtendedCreditAllocationConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteExtendedCreditAllocationConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteExtendedCreditAllocationConfig(ctx context.Context, client *core.Client, id string, params DeleteExtendedCreditAllocationConfigParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetExtendedCreditAllocationConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetExtendedCreditAllocationConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetExtendedCreditAllocationConfig(ctx context.Context, client *core.Client, id string, params GetExtendedCreditAllocationConfigParams) (*objects.ExtendedCreditAllocationConfig, error) {
	var out objects.ExtendedCreditAllocationConfig
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateExtendedCreditAllocationConfigParams struct {
	Amount *map[string]interface{} `facebook:"amount"`
	Extra  core.Params             `facebook:"-"`
}

func (params UpdateExtendedCreditAllocationConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Amount != nil {
		out["amount"] = *params.Amount
	}
	return out
}

func UpdateExtendedCreditAllocationConfig(ctx context.Context, client *core.Client, id string, params UpdateExtendedCreditAllocationConfigParams) (*objects.ExtendedCreditAllocationConfig, error) {
	var out objects.ExtendedCreditAllocationConfig
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
