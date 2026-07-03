package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdExportPresetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdExportPresetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdExportPreset(ctx context.Context, client *core.Client, id string, params GetAdExportPresetParams) (*objects.AdExportPreset, error) {
	var out objects.AdExportPreset
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
