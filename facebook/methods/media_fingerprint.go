package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetMediaFingerprintParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMediaFingerprintParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMediaFingerprint(ctx context.Context, client *core.Client, id string, params GetMediaFingerprintParams) (*objects.MediaFingerprint, error) {
	var out objects.MediaFingerprint
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateMediaFingerprintParams struct {
	Metadata           *[]interface{}  `facebook:"metadata"`
	Source             *core.FileParam `facebook:"source"`
	Title              *string         `facebook:"title"`
	UniversalContentID *core.ID        `facebook:"universal_content_id"`
	Extra              core.Params     `facebook:"-"`
}

func (params UpdateMediaFingerprintParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Metadata != nil {
		out["metadata"] = *params.Metadata
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	if params.UniversalContentID != nil {
		out["universal_content_id"] = *params.UniversalContentID
	}
	return out
}

func UpdateMediaFingerprint(ctx context.Context, client *core.Client, id string, params UpdateMediaFingerprintParams) (*objects.MediaFingerprint, error) {
	var out objects.MediaFingerprint
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
