package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetPrivateLiftStudyInstance(ctx context.Context, client *core.Client, id string, params GetPrivateLiftStudyInstanceParams) (*objects.PrivateLiftStudyInstance, error) {
	var out objects.PrivateLiftStudyInstance
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func UpdatePrivateLiftStudyInstance(ctx context.Context, client *core.Client, id string, params UpdatePrivateLiftStudyInstanceParams) (*objects.PrivateLiftStudyInstance, error) {
	var out objects.PrivateLiftStudyInstance
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
