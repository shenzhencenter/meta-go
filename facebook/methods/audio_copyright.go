package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAudioCopyrightUpdateRecordsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAudioCopyrightUpdateRecordsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAudioCopyrightUpdateRecords(ctx context.Context, client *core.Client, id string, params GetAudioCopyrightUpdateRecordsParams) (*core.Cursor[objects.MediaCopyrightUpdateRecord], error) {
	var out core.Cursor[objects.MediaCopyrightUpdateRecord]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "update_records"), params.ToParams(), &out)
	return &out, err
}

type GetAudioCopyrightParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAudioCopyrightParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAudioCopyright(ctx context.Context, client *core.Client, id string, params GetAudioCopyrightParams) (*objects.AudioCopyright, error) {
	var out objects.AudioCopyright
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
