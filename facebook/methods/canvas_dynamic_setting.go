package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCanvasDynamicSettingParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCanvasDynamicSettingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCanvasDynamicSetting(ctx context.Context, client *core.Client, id string, params GetCanvasDynamicSettingParams) (*objects.CanvasDynamicSetting, error) {
	var out objects.CanvasDynamicSetting
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
