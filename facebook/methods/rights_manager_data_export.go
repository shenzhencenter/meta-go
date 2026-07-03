package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetRightsManagerDataExportParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetRightsManagerDataExportParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetRightsManagerDataExport(ctx context.Context, client *core.Client, id string, params GetRightsManagerDataExportParams) (*objects.RightsManagerDataExport, error) {
	var out objects.RightsManagerDataExport
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
