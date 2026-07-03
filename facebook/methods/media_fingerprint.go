package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
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

func GetMediaFingerprintBatchCall(id string, params GetMediaFingerprintParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetMediaFingerprintBatchRequest(id string, params GetMediaFingerprintParams, options ...core.BatchOption) *core.BatchRequest[objects.MediaFingerprint] {
	return core.NewBatchRequest[objects.MediaFingerprint](GetMediaFingerprintBatchCall(id, params, options...))
}

func DecodeGetMediaFingerprintBatchResponse(response *core.BatchResponse) (*objects.MediaFingerprint, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.MediaFingerprint
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetMediaFingerprint(ctx context.Context, client *core.Client, id string, params GetMediaFingerprintParams) (*objects.MediaFingerprint, error) {
	var out objects.MediaFingerprint
	call := GetMediaFingerprintBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func UpdateMediaFingerprintBatchCall(id string, params UpdateMediaFingerprintParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateMediaFingerprintBatchRequest(id string, params UpdateMediaFingerprintParams, options ...core.BatchOption) *core.BatchRequest[objects.MediaFingerprint] {
	return core.NewBatchRequest[objects.MediaFingerprint](UpdateMediaFingerprintBatchCall(id, params, options...))
}

func DecodeUpdateMediaFingerprintBatchResponse(response *core.BatchResponse) (*objects.MediaFingerprint, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.MediaFingerprint
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateMediaFingerprint(ctx context.Context, client *core.Client, id string, params UpdateMediaFingerprintParams) (*objects.MediaFingerprint, error) {
	var out objects.MediaFingerprint
	call := UpdateMediaFingerprintBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
