package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetMediaCopyrightUpdateRecordParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMediaCopyrightUpdateRecordParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMediaCopyrightUpdateRecord(ctx context.Context, client *core.Client, id string, params GetMediaCopyrightUpdateRecordParams) (*objects.MediaCopyrightUpdateRecord, error) {
	var out objects.MediaCopyrightUpdateRecord
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
