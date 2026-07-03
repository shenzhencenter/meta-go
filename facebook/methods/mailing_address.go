package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetMailingAddressParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMailingAddressParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMailingAddressBatchCall(id string, params GetMailingAddressParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetMailingAddressBatchRequest(id string, params GetMailingAddressParams, options ...core.BatchOption) *core.BatchRequest[objects.MailingAddress] {
	return core.NewBatchRequest[objects.MailingAddress](GetMailingAddressBatchCall(id, params, options...))
}

func DecodeGetMailingAddressBatchResponse(response *core.BatchResponse) (*objects.MailingAddress, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.MailingAddress
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetMailingAddressWithResponse(ctx context.Context, client *core.Client, id string, params GetMailingAddressParams) (*objects.MailingAddress, *core.Response, error) {
	var out objects.MailingAddress
	call := GetMailingAddressBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetMailingAddress(ctx context.Context, client *core.Client, id string, params GetMailingAddressParams) (*objects.MailingAddress, error) {
	out, _, err := GetMailingAddressWithResponse(ctx, client, id, params)
	return out, err
}
