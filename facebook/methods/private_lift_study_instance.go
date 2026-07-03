package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPrivateLiftStudyInstanceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPrivateLiftStudyInstanceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPrivateLiftStudyInstanceBatchCall(id string, params GetPrivateLiftStudyInstanceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPrivateLiftStudyInstanceBatchRequest(id string, params GetPrivateLiftStudyInstanceParams, options ...core.BatchOption) *core.BatchRequest[objects.PrivateLiftStudyInstance] {
	return core.NewBatchRequest[objects.PrivateLiftStudyInstance](GetPrivateLiftStudyInstanceBatchCall(id, params, options...))
}

func DecodeGetPrivateLiftStudyInstanceBatchResponse(response *core.BatchResponse) (*objects.PrivateLiftStudyInstance, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PrivateLiftStudyInstance
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPrivateLiftStudyInstanceWithResponse(ctx context.Context, client *core.Client, id string, params GetPrivateLiftStudyInstanceParams) (*objects.PrivateLiftStudyInstance, *core.Response, error) {
	var out objects.PrivateLiftStudyInstance
	call := GetPrivateLiftStudyInstanceBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPrivateLiftStudyInstance(ctx context.Context, client *core.Client, id string, params GetPrivateLiftStudyInstanceParams) (*objects.PrivateLiftStudyInstance, error) {
	out, _, err := GetPrivateLiftStudyInstanceWithResponse(ctx, client, id, params)
	return out, err
}

type UpdatePrivateLiftStudyInstanceParams struct {
	Operation *enums.PrivateliftstudyinstanceOperation `facebook:"operation"`
	RunID     *core.ID                                 `facebook:"run_id"`
	Extra     core.Params                              `facebook:"-"`
}

func (params UpdatePrivateLiftStudyInstanceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Operation != nil {
		out["operation"] = *params.Operation
	}
	if params.RunID != nil {
		out["run_id"] = *params.RunID
	}
	return out
}

func UpdatePrivateLiftStudyInstanceBatchCall(id string, params UpdatePrivateLiftStudyInstanceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdatePrivateLiftStudyInstanceBatchRequest(id string, params UpdatePrivateLiftStudyInstanceParams, options ...core.BatchOption) *core.BatchRequest[objects.PrivateLiftStudyInstance] {
	return core.NewBatchRequest[objects.PrivateLiftStudyInstance](UpdatePrivateLiftStudyInstanceBatchCall(id, params, options...))
}

func DecodeUpdatePrivateLiftStudyInstanceBatchResponse(response *core.BatchResponse) (*objects.PrivateLiftStudyInstance, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PrivateLiftStudyInstance
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdatePrivateLiftStudyInstanceWithResponse(ctx context.Context, client *core.Client, id string, params UpdatePrivateLiftStudyInstanceParams) (*objects.PrivateLiftStudyInstance, *core.Response, error) {
	var out objects.PrivateLiftStudyInstance
	call := UpdatePrivateLiftStudyInstanceBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdatePrivateLiftStudyInstance(ctx context.Context, client *core.Client, id string, params UpdatePrivateLiftStudyInstanceParams) (*objects.PrivateLiftStudyInstance, error) {
	out, _, err := UpdatePrivateLiftStudyInstanceWithResponse(ctx, client, id, params)
	return out, err
}
