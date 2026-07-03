package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteAdgroupGendeleteParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdgroupGendeleteParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdgroupGendeleteBatchCall(id string, params DeleteAdgroupGendeleteParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, ""), params.ToParams(), options...)
}

func NewDeleteAdgroupGendeleteBatchRequest(id string, params DeleteAdgroupGendeleteParams, options ...core.BatchOption) *core.BatchRequest[objects.AdgroupDelete] {
	return core.NewBatchRequest[objects.AdgroupDelete](DeleteAdgroupGendeleteBatchCall(id, params, options...))
}

func DecodeDeleteAdgroupGendeleteBatchResponse(response *core.BatchResponse) (*objects.AdgroupDelete, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdgroupDelete
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteAdgroupGendelete(ctx context.Context, client *core.Client, id string, params DeleteAdgroupGendeleteParams) (*objects.AdgroupDelete, error) {
	var out objects.AdgroupDelete
	call := DeleteAdgroupGendeleteBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
