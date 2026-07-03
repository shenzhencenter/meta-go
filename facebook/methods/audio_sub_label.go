package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAudioSubLabelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAudioSubLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAudioSubLabelBatchCall(id string, params GetAudioSubLabelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAudioSubLabelBatchRequest(id string, params GetAudioSubLabelParams, options ...core.BatchOption) *core.BatchRequest[objects.AudioSubLabel] {
	return core.NewBatchRequest[objects.AudioSubLabel](GetAudioSubLabelBatchCall(id, params, options...))
}

func DecodeGetAudioSubLabelBatchResponse(response *core.BatchResponse) (*objects.AudioSubLabel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AudioSubLabel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAudioSubLabelWithResponse(ctx context.Context, client *core.Client, id string, params GetAudioSubLabelParams) (*objects.AudioSubLabel, *core.Response, error) {
	var out objects.AudioSubLabel
	call := GetAudioSubLabelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAudioSubLabel(ctx context.Context, client *core.Client, id string, params GetAudioSubLabelParams) (*objects.AudioSubLabel, error) {
	out, _, err := GetAudioSubLabelWithResponse(ctx, client, id, params)
	return out, err
}
