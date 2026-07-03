package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetOffsiteSignalContainerBusinessObjectLinkedApplicationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOffsiteSignalContainerBusinessObjectLinkedApplicationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOffsiteSignalContainerBusinessObjectLinkedApplication(ctx context.Context, client *core.Client, id string, params GetOffsiteSignalContainerBusinessObjectLinkedApplicationParams) (*core.Cursor[objects.AdsDataset], error) {
	var out core.Cursor[objects.AdsDataset]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "linked_application"), params.ToParams(), &out)
	return &out, err
}

type GetOffsiteSignalContainerBusinessObjectLinkedPageParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOffsiteSignalContainerBusinessObjectLinkedPageParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOffsiteSignalContainerBusinessObjectLinkedPage(ctx context.Context, client *core.Client, id string, params GetOffsiteSignalContainerBusinessObjectLinkedPageParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "linked_page"), params.ToParams(), &out)
	return &out, err
}

type GetOffsiteSignalContainerBusinessObjectParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOffsiteSignalContainerBusinessObjectParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOffsiteSignalContainerBusinessObject(ctx context.Context, client *core.Client, id string, params GetOffsiteSignalContainerBusinessObjectParams) (*objects.OffsiteSignalContainerBusinessObject, error) {
	var out objects.OffsiteSignalContainerBusinessObject
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
