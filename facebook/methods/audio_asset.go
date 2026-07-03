package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAudioAssetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAudioAssetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAudioAsset(ctx context.Context, client *core.Client, id string, params GetAudioAssetParams) (*objects.AudioAsset, error) {
	var out objects.AudioAsset
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
