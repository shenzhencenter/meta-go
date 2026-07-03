package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdCreationPackageConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdCreationPackageConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdCreationPackageConfigBatchCall(id string, params GetAdCreationPackageConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdCreationPackageConfigBatchRequest(id string, params GetAdCreationPackageConfigParams, options ...core.BatchOption) *core.BatchRequest[objects.AdCreationPackageConfig] {
	return core.NewBatchRequest[objects.AdCreationPackageConfig](GetAdCreationPackageConfigBatchCall(id, params, options...))
}

func DecodeGetAdCreationPackageConfigBatchResponse(response *core.BatchResponse) (*objects.AdCreationPackageConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdCreationPackageConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdCreationPackageConfig(ctx context.Context, client *core.Client, id string, params GetAdCreationPackageConfigParams) (*objects.AdCreationPackageConfig, error) {
	var out objects.AdCreationPackageConfig
	call := GetAdCreationPackageConfigBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
